package renderer

import (
	"fmt"
	"log"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/kudobuilder/kudo/pkg/engine"
	"github.com/kudobuilder/kudo/pkg/engine/resource"
	"github.com/kudobuilder/kudo/pkg/util/kudo"
)

// Enhancer takes your kubernetes template and kudo related Metadata and applies them to all resources in form of labels
// and annotations
// it also takes care of setting an owner of all the resources to the provided object
type Enhancer interface {
	Apply(templates map[string]string, metadata Metadata) ([]runtime.Object, error)
}

// DefaultEnhancer is implementation of Enhancer that applies the defined conventions by directly editing runtime.Objects (Unstructured).
type DefaultEnhancer struct {
	Scheme    *runtime.Scheme
	Client    client.Client
	Discovery discovery.CachedDiscoveryInterface
}

// Apply accepts templates to be rendered in kubernetes and enhances them with our own KUDO conventions
// These include the way we name our objects and what labels we apply to them
func (de *DefaultEnhancer) Apply(templates map[string]string, metadata Metadata) (objsToAdd []runtime.Object, err error) {
	unstructuredObjs := make([]*unstructured.Unstructured, 0, len(templates))

	for name, v := range templates {
		parsed, err := YamlToObject(v)
		if err != nil {
			return nil, fmt.Errorf("%wparsing YAML from %s: %v", engine.ErrFatalExecution, name, err)
		}
		for _, obj := range parsed {
			unstructMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
			if err != nil {
				return nil, fmt.Errorf("%wconverting to unstructured failed: %v", engine.ErrFatalExecution, err)
			}

			if err = addLabels(unstructMap, metadata); err != nil {
				return nil, fmt.Errorf("%wadding labels on parsed object: %v", engine.ErrFatalExecution, err)
			}
			if err = addAnnotations(unstructMap, metadata); err != nil {
				return nil, fmt.Errorf("%wadding annotations on parsed object %s: %v", engine.ErrFatalExecution, obj.GetObjectKind(), err)
			}

			objUnstructured := &unstructured.Unstructured{Object: unstructMap}

			isNamespaced, err := resource.IsNamespacedObject(obj, de.Discovery)
			if err != nil {
				return nil, fmt.Errorf("failed to determine if object %s is namespaced: %v", obj.GetObjectKind(), err)
			}

			// Note: Cross-namespace owner references are disallowed by design. This means:
			// 1) Namespace-scoped dependents can only specify owners in the same namespace, and owners that are cluster-scoped.
			// 2) Cluster-scoped dependents can only specify cluster-scoped owners, but not namespace-scoped owners.
			// More: https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/
			if isNamespaced {
				objUnstructured.SetNamespace(metadata.InstanceNamespace)
				if err = setControllerReference(metadata.ResourcesOwner, objUnstructured, de.Scheme); err != nil {
					return nil, fmt.Errorf("%wsetting controller reference on parsed object %s: %v", engine.ErrFatalExecution, obj.GetObjectKind(), err)
				}
			}

			// We convert to the typed version here and back to clean additional labels and annotations that were
			// added. This needs to be done as otherwise the dependency calculator calculates hashes based on the added
			// fields which are later missing if the resource is fetched from the api-server
			err = runtime.DefaultUnstructuredConverter.FromUnstructured(objUnstructured.UnstructuredContent(), obj)
			if err != nil {
				return nil, fmt.Errorf("%wconverting from unstructured failed: %v", engine.ErrFatalExecution, err)
			}
			unstructMap, err = runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
			if err != nil {
				return nil, fmt.Errorf("%wconverting to unstructured failed: %v", engine.ErrFatalExecution, err)
			}
			objUnstructured = &unstructured.Unstructured{Object: unstructMap}

			unstructuredObjs = append(unstructuredObjs, objUnstructured)
		}
	}

	dc := newDependencyCalculator(de.Client, unstructuredObjs)
	for _, uo := range unstructuredObjs {
		deps, err := calculateResourceDependencies(uo)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate resource dependencies for %s/%s: %v", uo.GetNamespace(), uo.GetName(), err)
		}
		if deps != nil && !deps.empty() {
			err = dc.calculateAndSetHash(uo, deps)
			if err != nil {
				return nil, fmt.Errorf("failed to add dependency hash to %s/%s: %v", uo.GetNamespace(), uo.GetName(), err)
			}
		}
	}

	// This is pretty important, if we don't convert it to the actual type everything will be Unstructured.
	// We depend on types later on in the processing e.g. when evaluating health.
	// Additionally, as we add annotations and labels to all possible paths, this step gets rid of anything
	// that doesn't belong to the specific object type.
	objs := make([]runtime.Object, 0, len(unstructuredObjs))
	for _, uo := range unstructuredObjs {
		obj, err := de.Scheme.New(uo.GroupVersionKind())
		if err != nil {
			objs = append(objs, uo)
			continue
		}

		err = runtime.DefaultUnstructuredConverter.FromUnstructured(uo.UnstructuredContent(), obj)
		if err != nil {
			return nil, fmt.Errorf("%wconverting from unstructured failed: %v", engine.ErrFatalExecution, err)
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func addLabels(obj map[string]interface{}, metadata Metadata) error {
	// List of paths for labels from here:
	// https://github.com/kubernetes-sigs/kustomize/blob/master/api/konfig/builtinpluginconsts/commonlabels.go

	// Labels are added to top-level resources as well as pod template specs. All labels here do not change
	// over the livetime of the instance and can not trigger unwanted restarts of the pods
	labelPaths := [][]string{
		{"metadata", "labels"},
		{"spec", "template", "metadata", "labels"},
		{"spec", "volumeClaimTemplates[]", "metadata", "labels"},
		{"spec", "jobTemplate", "metadata", "labels"},
		{"spec", "jobTemplate", "spec", "template", "metadata", "labels"},
	}

	fieldsToAdd := map[string]string{
		kudo.HeritageLabel: "kudo",
		kudo.OperatorLabel: metadata.OperatorName,
		kudo.InstanceLabel: metadata.InstanceName,
	}

	for _, path := range labelPaths {
		if err := addMapValues(obj, fieldsToAdd, path...); err != nil {
			return err
		}
	}

	return nil
}

func addAnnotations(obj map[string]interface{}, metadata Metadata) error {
	// List of paths for annotations from here:
	// https://github.com/kubernetes-sigs/kustomize/blob/master/api/konfig/builtinpluginconsts/commonannotations.go

	// For all pod template specs, we only add the operator version. It is pretty stable
	// and shouldn't change often, therefore not trigger an unwanted restart of the created pod
	annotationPaths := [][]string{
		{"metadata", "annotations"},
		{"spec", "template", "metadata", "annotations"},
		{"spec", "jobTemplate", "metadata", "annotations"},
		{"spec", "jobTemplate", "spec", "template", "metadata", "annotations"},
	}

	fieldsToAdd := map[string]string{
		kudo.OperatorVersionAnnotation: metadata.OperatorVersion,
	}

	for _, path := range annotationPaths {
		if err := addMapValues(obj, fieldsToAdd, path...); err != nil {
			return err
		}
	}

	// The plan, phase and step annotations are only added to the top level resources, not and pod template specs, as
	// that may lead to unwanted restarts of the pods
	topLevelFieldsToAdd := map[string]string{
		kudo.PlanAnnotation:    metadata.PlanName,
		kudo.PhaseAnnotation:   metadata.PhaseName,
		kudo.StepAnnotation:    metadata.StepName,
		kudo.PlanUIDAnnotation: string(metadata.PlanUID),
	}
	if err := addMapValues(obj, topLevelFieldsToAdd, "metadata", "annotations"); err != nil {
		return err
	}

	return nil
}

func addMapValues(obj map[string]interface{}, fieldsToAdd map[string]string, path ...string) error {
	for i, p := range path {
		// If we have an element with a slice in the path, apply the fields to all elements of the
		// slice with the remaining path
		if strings.HasSuffix(p, "[]") {
			sliceField := strings.TrimSuffix(p, "[]")

			subPath := append(path[0:i], sliceField)
			remainingPath := path[i+1:]

			unstructuredSlice, found, err := unstructured.NestedSlice(obj, subPath...)
			if !found || err != nil {
				// We don't return err here, as it just means that path is invalid for this object.
				// This is ok and does not indicate an error
				return nil
			}
			for _, s := range unstructuredSlice {
				if sliceMap, ok := s.(map[string]interface{}); ok {
					if err = addMapValues(sliceMap, fieldsToAdd, remainingPath...); err != nil {
						return err
					}
				}
			}
			if err = unstructured.SetNestedSlice(obj, unstructuredSlice, subPath...); err != nil {
				return err
			}

			return nil
		}
	}

	// Merge added fields to map at specified path
	stringMap, _, err := unstructured.NestedStringMap(obj, path...)
	if err != nil {
		// We don't return err here, as it just means that path is invalid for this object.
		// This is ok and does not indicate an error
		return nil
	}
	if stringMap == nil {
		stringMap = make(map[string]string)
	}
	for k, v := range fieldsToAdd {
		stringMap[k] = v
	}
	return unstructured.SetNestedStringMap(obj, stringMap, path...)
}

func setControllerReference(owner metav1.Object, object metav1.Object, scheme *runtime.Scheme) error {
	ownerNs := owner.GetNamespace()
	if ownerNs != "" {
		objNs := object.GetNamespace()
		if objNs == "" {
			// we're trying to create cluster-scoped resource from and bind Instance as owner of that
			// that is disallowed by design, see https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/#owners-and-dependents
			// for now solve by not adding the owner
			log.Printf("Not adding owner to resource %s because it's cluster-scoped and cannot be owned by namespace-scoped instance %s/%s", object.GetName(), owner.GetNamespace(), owner.GetName())
			return nil
		}
		if ownerNs != objNs {
			// we're trying to create resource in another namespace as is Instance's namespace, Instance cannot be owner of such resource
			// that is disallowed by design, see https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/#owners-and-dependents
			// for now solve by not adding the owner
			log.Printf("Not adding owner to resource %s/%s because it's in different namespace than instance %s/%s and thus cannot be owned by that instance", object.GetNamespace(), object.GetName(), owner.GetNamespace(), owner.GetName())
			return nil
		}
	}
	if err := controllerutil.SetControllerReference(owner, object, scheme); err != nil {
		return err
	}
	return nil
}

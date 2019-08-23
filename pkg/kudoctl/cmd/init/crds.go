package init

import (
	"strings"

	"github.com/ghodss/yaml"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	group   = "kudo.dev"
	version = "v1alpha1"
)

// OperatorCrd provides the Operator CRD manifest for printing
func OperatorCrd() *apiextv1beta1.CustomResourceDefinition {
	crd := generateOperator()
	crd.TypeMeta = metav1.TypeMeta{
		Kind:       "CustomResourceDefinition",
		APIVersion: "apiextensions.k8s.io/v1beta1",
	}
	return crd
}

func generateOperator() *apiextv1beta1.CustomResourceDefinition {
	crd := generateCrd("Operator", "operators")
	specProps := map[string]apiextv1beta1.JSONSchemaProps{
		"description":       apiextv1beta1.JSONSchemaProps{Type: "string"},
		"kubernetesVersion": apiextv1beta1.JSONSchemaProps{Type: "string"},
		"kudoVersion":       apiextv1beta1.JSONSchemaProps{Type: "string"},
		"maintainers": apiextv1beta1.JSONSchemaProps{
			Type:  "array",
			Items: &apiextv1beta1.JSONSchemaPropsOrArray{Schema: &apiextv1beta1.JSONSchemaProps{Type: "string"}, JSONSchemas: []apiextv1beta1.JSONSchemaProps{}},
		},
		"url": apiextv1beta1.JSONSchemaProps{Type: "string"},
	}

	validationProps := map[string]apiextv1beta1.JSONSchemaProps{
		"apiVersion": apiextv1beta1.JSONSchemaProps{Type: "string"},
		"kind":       apiextv1beta1.JSONSchemaProps{Type: "string"},
		"meta":       apiextv1beta1.JSONSchemaProps{Type: "object"},
		"spec":       apiextv1beta1.JSONSchemaProps{Properties: specProps, Type: "object"},
		"status":     apiextv1beta1.JSONSchemaProps{Type: "object"},
	}

	crd.Spec.Validation = &apiextv1beta1.CustomResourceValidation{
		OpenAPIV3Schema: &apiextv1beta1.JSONSchemaProps{
			Properties: validationProps,
		},
	}
	return crd
}

// OperatorVersionCrd provides the OperatorVersion CRD manifest for printing
func OperatorVersionCrd() *apiextv1beta1.CustomResourceDefinition {
	crd := generateOperatorVersion()
	crd.TypeMeta = metav1.TypeMeta{
		Kind:       "CustomResourceDefinition",
		APIVersion: "apiextensions.k8s.io/v1beta1",
	}
	return crd
}

func generateOperatorVersion() *apiextv1beta1.CustomResourceDefinition {
	crd := generateCrd("OperatorVersion", "operatorversions")
	dependProps := map[string]apiextv1beta1.JSONSchemaProps{
		"referenceName": apiextv1beta1.JSONSchemaProps{Type: "string", Description: "Name specifies the name of the dependency.  Referenced via this in defaults.config"},
		"version":       apiextv1beta1.JSONSchemaProps{Type: "string", Description: "Version captures the requirements for what versions of the above object are allowed Example: ^3.1.4"},
	}
	paramProps := map[string]apiextv1beta1.JSONSchemaProps{
		"default":     apiextv1beta1.JSONSchemaProps{Type: "string", Description: "Default is a default value if no paramter is provided by the instance"},
		"description": apiextv1beta1.JSONSchemaProps{Type: "string", Description: "Description captures a longer description of how the variable will be used"},
		"displayName": apiextv1beta1.JSONSchemaProps{Type: "string", Description: "Human friendly version of the parameter name"},
		"name":        apiextv1beta1.JSONSchemaProps{Type: "string", Description: "Name is the string that should be used in the template file for example, if `name: COUNT` then using the variable in a spec like:  spec:   replicas:  {{ .Params.COUNT }}"},
		"required":    apiextv1beta1.JSONSchemaProps{Type: "boolean", Description: "Required specifies if the parameter is required to be provided by all instances, or whether a default can suffice"},
		"trigger":     apiextv1beta1.JSONSchemaProps{Type: "string", Description: "Trigger identifies the plan that gets executed when this parameter changes in the Instance object. Default is `update` if present, or `deploy` if not present"},
	}
	specProps := map[string]apiextv1beta1.JSONSchemaProps{
		"connectionString": apiextv1beta1.JSONSchemaProps{Type: "string", Description: "ConnectionString defines a mustached string that can be used to connect to an instance of the Operator"},
		"dependencies": apiextv1beta1.JSONSchemaProps{
			Type: "array",
			Items: &apiextv1beta1.JSONSchemaPropsOrArray{Schema: &apiextv1beta1.JSONSchemaProps{
				Type:       "object",
				Required:   []string{"referenceName", "version"},
				Properties: dependProps,
			}, JSONSchemas: []apiextv1beta1.JSONSchemaProps{}},
		},
		"operator": apiextv1beta1.JSONSchemaProps{Type: "object"},
		"parameters": apiextv1beta1.JSONSchemaProps{
			Type: "array",
			Items: &apiextv1beta1.JSONSchemaPropsOrArray{Schema: &apiextv1beta1.JSONSchemaProps{
				Type:       "object",
				Properties: paramProps,
			}, JSONSchemas: []apiextv1beta1.JSONSchemaProps{}},
		},
		"plans":     apiextv1beta1.JSONSchemaProps{Type: "object", Description: "Plans specify a map a plans that specify how to"},
		"tasks":     apiextv1beta1.JSONSchemaProps{Type: "object"},
		"templates": apiextv1beta1.JSONSchemaProps{Type: "object", Description: "List of go templates YAML files that define the application operator instance"},
		"upgradableFrom": apiextv1beta1.JSONSchemaProps{
			Type:        "array",
			Description: "UpgradableFrom lists all OperatorVersions that can upgrade to this OperatorVersion",
			Items:       &apiextv1beta1.JSONSchemaPropsOrArray{Schema: &apiextv1beta1.JSONSchemaProps{Type: "object"}, JSONSchemas: []apiextv1beta1.JSONSchemaProps{}},
		},
		"version": apiextv1beta1.JSONSchemaProps{Type: "string"},
	}

	validationProps := map[string]apiextv1beta1.JSONSchemaProps{
		"apiVersion": apiextv1beta1.JSONSchemaProps{Type: "string"},
		"kind":       apiextv1beta1.JSONSchemaProps{Type: "string"},
		"meta":       apiextv1beta1.JSONSchemaProps{Type: "object"},
		"spec":       apiextv1beta1.JSONSchemaProps{Properties: specProps, Type: "object"},
		"status":     apiextv1beta1.JSONSchemaProps{Type: "object"},
	}

	crd.Spec.Validation = &apiextv1beta1.CustomResourceValidation{
		OpenAPIV3Schema: &apiextv1beta1.JSONSchemaProps{
			Properties: validationProps,
		},
	}
	return crd
}

// InstanceCrd provides the Instance CRD manifest for printing
func InstanceCrd() *apiextv1beta1.CustomResourceDefinition {
	crd := generateInstance()
	crd.TypeMeta = metav1.TypeMeta{
		Kind:       "CustomResourceDefinition",
		APIVersion: "apiextensions.k8s.io/v1beta1",
	}
	return crd
}

func generateInstance() *apiextv1beta1.CustomResourceDefinition {
	crd := generateCrd("Instance", "instances")
	dependProps := map[string]apiextv1beta1.JSONSchemaProps{
		"referenceName": apiextv1beta1.JSONSchemaProps{Type: "string", Description: "Name specifies the name of the dependency.  Referenced via this in defaults.config"},
		"version":       apiextv1beta1.JSONSchemaProps{Type: "string", Description: "Version captures the requirements for what versions of the above object are allowed Example: ^3.1.4"},
	}
	specProps := map[string]apiextv1beta1.JSONSchemaProps{
		"dependencies": apiextv1beta1.JSONSchemaProps{
			Type:        "array",
			Description: "Dependency references specific",
			Items: &apiextv1beta1.JSONSchemaPropsOrArray{Schema: &apiextv1beta1.JSONSchemaProps{
				Type:       "object",
				Required:   []string{"referenceName", "version"},
				Properties: dependProps,
			}, JSONSchemas: []apiextv1beta1.JSONSchemaProps{}},
		},
		"OperatorVersion": apiextv1beta1.JSONSchemaProps{Type: "object", Description: "Operator specifies a reference to a specific Operator object"},
		"parameters":      apiextv1beta1.JSONSchemaProps{Type: "object"},
	}
	statusProps := map[string]apiextv1beta1.JSONSchemaProps{
		"activePlan": apiextv1beta1.JSONSchemaProps{Type: "object"},
		"status":     apiextv1beta1.JSONSchemaProps{Type: "string"},
	}

	validationProps := map[string]apiextv1beta1.JSONSchemaProps{
		"apiVersion": apiextv1beta1.JSONSchemaProps{Type: "string"},
		"kind":       apiextv1beta1.JSONSchemaProps{Type: "string"},
		"meta":       apiextv1beta1.JSONSchemaProps{Type: "object"},
		"spec":       apiextv1beta1.JSONSchemaProps{Properties: specProps, Type: "object"},
		"status": apiextv1beta1.JSONSchemaProps{
			Type:       "object",
			Properties: statusProps,
		},
	}

	crd.Spec.Validation = &apiextv1beta1.CustomResourceValidation{
		OpenAPIV3Schema: &apiextv1beta1.JSONSchemaProps{
			Properties: validationProps,
		},
	}
	return crd
}

// PlanExecutionCrd provides the PlanExecution CRD manifest for printing
func PlanExecutionCrd() *apiextv1beta1.CustomResourceDefinition {
	crd := generatePlanExecution()
	crd.TypeMeta = metav1.TypeMeta{
		Kind:       "CustomResourceDefinition",
		APIVersion: "apiextensions.k8s.io/v1beta1",
	}
	return crd
}

func generatePlanExecution() *apiextv1beta1.CustomResourceDefinition {
	crd := generateCrd("PlanExecution", "planexecutions")
	specProps := map[string]apiextv1beta1.JSONSchemaProps{
		"instance": apiextv1beta1.JSONSchemaProps{Type: "object"},
		"planName": apiextv1beta1.JSONSchemaProps{Type: "string"},
		"suspend":  apiextv1beta1.JSONSchemaProps{Type: "boolean", Description: "This is copied from the CronJob Spec This flag tells the controller to suspend subsequent executions, it does not apply to already started executions.  Defaults to false."},
	}

	stepProps := map[string]apiextv1beta1.JSONSchemaProps{
		"delete": apiextv1beta1.JSONSchemaProps{Type: "boolean"},
		"name":   apiextv1beta1.JSONSchemaProps{Type: "string"},
		"state":  apiextv1beta1.JSONSchemaProps{Type: "string"},
	}

	phaseProps := map[string]apiextv1beta1.JSONSchemaProps{
		"name":  apiextv1beta1.JSONSchemaProps{Type: "string"},
		"state": apiextv1beta1.JSONSchemaProps{Type: "string"},
		"steps": apiextv1beta1.JSONSchemaProps{
			Type:        "array",
			Description: "Steps maps a step name to a list of templates objects stored as a string",
			Items: &apiextv1beta1.JSONSchemaPropsOrArray{Schema: &apiextv1beta1.JSONSchemaProps{
				Type:       "object",
				Properties: stepProps,
			}, JSONSchemas: []apiextv1beta1.JSONSchemaProps{}},
		},
		"strategy": apiextv1beta1.JSONSchemaProps{Type: "string"},
	}

	statusProps := map[string]apiextv1beta1.JSONSchemaProps{
		"name": apiextv1beta1.JSONSchemaProps{Type: "string", Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run 'make' to regenerate code after modifying this file"},
		"phases": apiextv1beta1.JSONSchemaProps{
			Type:        "array",
			Description: "Phases maps a phase name to a Phase object",
			Items: &apiextv1beta1.JSONSchemaPropsOrArray{Schema: &apiextv1beta1.JSONSchemaProps{
				Type:       "object",
				Required:   []string{"steps"},
				Properties: phaseProps,
			}, JSONSchemas: []apiextv1beta1.JSONSchemaProps{}},
		},
		"state":    apiextv1beta1.JSONSchemaProps{Type: "string"},
		"strategy": apiextv1beta1.JSONSchemaProps{Type: "string"},
	}

	validationProps := map[string]apiextv1beta1.JSONSchemaProps{
		"apiVersion": apiextv1beta1.JSONSchemaProps{Type: "string"},
		"kind":       apiextv1beta1.JSONSchemaProps{Type: "string"},
		"meta":       apiextv1beta1.JSONSchemaProps{Type: "object"},
		"spec": apiextv1beta1.JSONSchemaProps{
			Properties: specProps,
			Type:       "object",
			Required:   []string{"planName", "instance"},
		},
		"status": apiextv1beta1.JSONSchemaProps{
			Type:       "object",
			Properties: statusProps,
		},
	}

	crd.Spec.Validation = &apiextv1beta1.CustomResourceValidation{
		OpenAPIV3Schema: &apiextv1beta1.JSONSchemaProps{
			Properties: validationProps,
		},
	}
	return crd
}

// generateCrd provides a generic CRD object to be configured
func generateCrd(kind string, plural string) *apiextv1beta1.CustomResourceDefinition {
	plural = strings.ToLower(plural)
	name := plural + "." + group

	labels := generateLabels(map[string]string{"controller-tools.k8s.io": "1.0"})
	crd := &apiextv1beta1.CustomResourceDefinition{
		ObjectMeta: v1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
		Spec: apiextv1beta1.CustomResourceDefinitionSpec{
			Group:   group,
			Version: version,
			Names: apiextv1beta1.CustomResourceDefinitionNames{
				Plural:     plural,
				Singular:   strings.ToLower(kind),
				ShortNames: nil,
				Kind:       kind,
			},
			Scope: "Namespaced",
		},
		Status: apiextv1beta1.CustomResourceDefinitionStatus{
			Conditions:     []apiextv1beta1.CustomResourceDefinitionCondition{},
			StoredVersions: []string{},
		},
	}
	return crd
}

// CrdManifests provides a slice of strings for each CRD manifest
func CrdManifests() ([]string, error) {
	o := OperatorCrd()
	ov := OperatorVersionCrd()
	i := InstanceCrd()
	pe := PlanExecutionCrd()

	objs := []runtime.Object{o, ov, i, pe}

	manifests := make([]string, len(objs))
	for i, obj := range objs {
		o, err := yaml.Marshal(obj)
		if err != nil {
			return []string{}, err
		}
		manifests[i] = string(o)
	}

	return manifests, nil
}

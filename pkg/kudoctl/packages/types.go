package packages

import (
	kudoapi "github.com/kudobuilder/kudo/pkg/apis/kudo/v1beta1"
)

const (
	APIVersion = "kudo.dev/v1beta1"
)

// This is an abstraction which abstracts the underlying packages, which is likely file system or compressed file.
// There should be a complete separation between retrieving a packages if not local and working with a packages.

// Package is an abstraction of the collection of files that makes up a package.  It is anything we can retrieve the Resources from.
type Package struct {
	// transformed server view
	Resources *Resources
	// working with local package files
	Files *Files
}

// Resources is collection of CRDs that are used when installing operator
// during installation, package format is converted to this structure
type Resources struct {
	Operator        *kudoapi.Operator
	OperatorVersion *kudoapi.OperatorVersion
	Instance        *kudoapi.Instance
}

// Modified kudoapi.Parameter that allows for defaults provided as YAML.
type Parameter struct {
	DisplayName string                `json:"displayName,omitempty"`
	Name        string                `json:"name,omitempty"`
	Description string                `json:"description,omitempty"`
	Required    *bool                 `json:"required,omitempty"`
	Default     interface{}           `json:"default,omitempty"`
	Trigger     string                `json:"trigger,omitempty"`
	Type        kudoapi.ParameterType `json:"type,omitempty"`
	Immutable   *bool                 `json:"immutable,omitempty"`
	Enum        *[]string             `json:"enum,omitempty"`

	// The following fields are descriptive only and are not used in the OperatorVersion
	Group    string `json:"group,omitempty"`
	Advanced *bool  `json:"advanced,omitempty"`
	Hint     string `json:"hint,omitempty"`
}

func (p Parameter) IsImmutable() bool {
	return p.Immutable != nil && *p.Immutable
}

func (p Parameter) IsRequired() bool {
	return p.Required != nil && *p.Required
}

func (p Parameter) IsAdvanced() bool {
	return p.Advanced != nil && *p.Advanced
}

func (p Parameter) IsEnum() bool {
	return p.Enum != nil
}

func (p *Parameter) HasDefault() bool {
	return p.Default != nil
}

func (p *Parameter) EnumValues() []string {
	if p.IsEnum() {
		return *p.Enum
	}
	return []string{}
}

type Parameters []Parameter

// Len returns the number of params.
// This is needed to allow sorting of params.
func (p Parameters) Len() int { return len(p) }

// Swap swaps the position of two items in the params slice.
// This is needed to allow sorting of params.
func (p Parameters) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Less returns true if the name of a param a is less than the name of param b.
// This is needed to allow sorting of params.
func (p Parameters) Less(x, y int) bool {
	return p[x].Name < p[y].Name
}

// Templates is a map of file names and stringified files in the template folder of an operator
type Templates map[string]string

// Files represents the raw operator package format the way it is found in the tgz packages
type Files struct {
	Templates Templates
	Operator  *OperatorFile
	Params    *ParamsFile
}

// ParamsFile is a representation of the package params.yaml
type ParamsFile struct {
	APIVersion string     `json:"apiVersion,omitempty"`
	Parameters Parameters `json:"parameters"`
}

// OperatorFile is a representation of the package operator.yaml
type OperatorFile struct {
	APIVersion        string                  `json:"apiVersion,omitempty"`
	Name              string                  `json:"name"`
	Description       string                  `json:"description,omitempty"`
	OperatorVersion   string                  `json:"operatorVersion"`
	AppVersion        string                  `json:"appVersion,omitempty"`
	KUDOVersion       string                  `json:"kudoVersion,omitempty"`
	KubernetesVersion string                  `json:"kubernetesVersion,omitempty"`
	Maintainers       []*kudoapi.Maintainer   `json:"maintainers,omitempty"`
	URL               string                  `json:"url,omitempty"`
	Tasks             []kudoapi.Task          `json:"tasks"`
	Plans             map[string]kudoapi.Plan `json:"plans"`
	NamespaceManifest string                  `json:"namespaceManifest,omitempty"`
}

package verify

import (
	"fmt"
	"strings"

	"github.com/kudobuilder/kudo/pkg/kudoctl/packages"
	"github.com/kudobuilder/kudo/pkg/kudoctl/packages/verifier"
	"github.com/kudobuilder/kudo/pkg/kudoctl/packages/verifier/template"
)

var verifiers = []verifier.PackageVerifier{
	DuplicateVerifier{},
	InvalidCharVerifier{";,"},
	template.ParametersVerifier{},
	template.ReferenceVerifier{},
}

// PackageFiles verifies operator package files
func PackageFiles(pf *packages.Files) verifier.Result {
	res := verifier.NewResult()
	for _, vv := range verifiers {
		res.Merge(vv.Verify(pf))
	}
	return res
}

// DuplicateVerifier provides verification that there are no duplicates disallowing casing (Kudo and kudo are duplicates)
type DuplicateVerifier struct{}

func (DuplicateVerifier) Verify(pf *packages.Files) verifier.Result {
	res := verifier.NewResult()
	names := map[string]bool{}
	for _, param := range pf.Params.Parameters {
		name := strings.ToLower(param.Name)
		if names[name] {
			res.AddParamError(param, "has a duplicate")
		}
		names[name] = true
	}
	return res
}

type InvalidCharVerifier struct {
	InvalidChars string
}

func (v InvalidCharVerifier) Verify(pf *packages.Files) verifier.Result {
	res := verifier.NewResult()
	for _, param := range pf.Params.Parameters {
		name := strings.ToLower(param.Name)
		for _, char := range name {
			if strings.Contains(v.InvalidChars, strings.ToLower(string(char))) {
				res.AddParamError(param, fmt.Sprintf("contains invalid character %q", char))
			}
		}

	}

	return res
}

package nodejs

import (
	"bytes"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
)

func GenerateTypes(tool string, pkg *schema.Package) (string, error) {
	// Decode node-specific info
	if err := pkg.ImportLanguages(map[string]schema.Language{"nodejs": Importer}); err != nil {
		return "", err
	}
	info, _ := pkg.Language["nodejs"].(NodePackageInfo)

	modules, err := generateModuleContextMap(tool, pkg, info, nil)
	if err != nil {
		return "", err
	}

	for _, mod := range modules {
		if len(mod.types) > 0 {
			crdTypes := mod.genCRDNamespaces()
			return crdTypes, nil
		}
	}

	return "", errors.New("cannot find module with types")
}

func (mod *modContext) genCRDNamespaces() string {
	// Create inputs buffer
	inputs := &bytes.Buffer{}

	// Generate the header and imports
	inputs.WriteString("// *** This file was generated by crd2pulumi. ***\n\n")
	inputs.WriteString("import * as pulumi from \"@pulumi/pulumi\";\n")
	inputs.WriteString("import * as k8s from \"@pulumi/kubernetes\";\n\n")
	inputs.WriteString("type ObjectMeta = k8s.types.input.meta.v1.ObjectMeta;\n\n")

	// Get and generate the namespace
	namespaces := mod.getNamespaces()
	baseNamespace := namespaces[""]
	// Ignore the top-level namespace, we only want versions
	versionNamespaces := baseNamespace.children

	for _, versionNamespace := range versionNamespaces {
		mod.genNamespace(inputs, versionNamespace, true, 0)
	}

	return inputs.String()
}

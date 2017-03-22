/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generators

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

type unversionedGenerator struct {
	generator.DefaultGen
	imports               []string
	pkg                   *types.Package
	group                 string
	domain                string
	typesByVersionKind    map[string]map[string]*types.Type
	typesByKindVersion    map[string]map[string]*types.Type
	allTypesByKindVersion map[string]map[string]*types.Type
	versionedApiTypes     []string

	subresourceApiTypes []string
	subresources        map[string]SubResource
}

var _ generator.Generator = &unversionedGenerator{}

const unversioned = "unversioned"

func CreateUnversionedGenerator(
	context *generator.Context, pkg *types.Package, arguments *args.GeneratorArgs,
	group string, domain string) generator.Generator {
	subresources := GetSubresources(context, group)

	//typesByVersionKind, typesByKindVersion, unversionedApiTypes := GetIndexedTypes(context, group)

	versionedApiTypes, unversionedApiTypes, subresourceApiTypes := GetVersionedAndUnversioned(context, group)
	typesByVersionKind, typesByKindVersion := IndexByVersionAndKind(
		context, group, versionedApiTypes, unversionedApiTypes)

	allVersionedApiTypes, allUnversionedApiTypes := GetAllVersionedAndUnversioned(context, group)
	_, allTypesByKindVersion := IndexAllByVersionAndKind(
		context, group, allVersionedApiTypes, allUnversionedApiTypes)

	// calculate import statements
	toImport := sets.String{}
	toImport.Insert(
		"metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apimachinery/pkg/runtime/schema",
		"genericapirequest \"k8s.io/apiserver/pkg/endpoints/request\"",
		"genericregistry \"k8s.io/apiserver/pkg/registry/generic/registry\"",
		"metainternalversion \"k8s.io/apimachinery/pkg/apis/meta/internalversion\"",
		"k8s.io/apiserver-builder/pkg/defaults",
		"fmt",
		"k8s.io/apiserver/pkg/registry/rest",
		"reflect",
		"k8s.io/client-go/pkg/api",
	)

	return &unversionedGenerator{
		generator.DefaultGen{OptionalName: arguments.OutputFileBaseName},
		toImport.List(),
		pkg,
		group,
		domain,
		typesByVersionKind,
		typesByKindVersion,
		allTypesByKindVersion,
		versionedApiTypes.List(),
		subresourceApiTypes.List(),
		subresources,
	}
}

func (d *unversionedGenerator) Init(*generator.Context, io.Writer) error {
	return nil
}

func (d *unversionedGenerator) Filter(c *generator.Context, t *types.Type) bool {
	return true
}

func (d *unversionedGenerator) Namers(c *generator.Context) namer.NameSystems {
	return nil
}
func (d *unversionedGenerator) Imports(c *generator.Context) []string {
	return d.imports
}

func (d *unversionedGenerator) PackageVars(c *generator.Context) []string {
	vars := []string{}
	buffer := &bytes.Buffer{}

	types := []string{}
	for _, n := range d.versionedApiTypes {
		types = append(types, fmt.Sprintf("&%s{}", n), fmt.Sprintf("&%sList{}", n))
	}
	for _, n := range d.subresources {
		types = append(types, fmt.Sprintf("&%s{}", n.RequestKind))
	}

	t := strings.Join(types, ", ")
	sw := generator.NewSnippetWriter(buffer, c, "$", "$")
	sw.Do(fmt.Sprintf(`registerFn = func(scheme *runtime.Scheme) error {
		scheme.AddKnownTypes(SchemeGroupVersion, %s)
		metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
		return nil
	}%s`, t, "\n"), nil)
	vars = append(vars, buffer.String())
	buffer.Reset()

	sw = generator.NewSnippetWriter(buffer, c, "$", "$")
	sw.Do("SchemeGroupVersion = schema.GroupVersion{Group, runtime.APIVersionInternal}\n", nil)
	vars = append(vars, buffer.String())
	buffer.Reset()

	sw = generator.NewSnippetWriter(buffer, c, "$", "$")
	sw.Do("SchemeBuilder = runtime.NewSchemeBuilder(registerFn)\n", nil)
	vars = append(vars, buffer.String())
	buffer.Reset()

	sw = generator.NewSnippetWriter(buffer, c, "$", "$")
	sw.Do("AddToScheme = SchemeBuilder.AddToScheme\n", nil)
	vars = append(vars, buffer.String())
	buffer.Reset()

	return vars
}

func (d *unversionedGenerator) PackageConsts(c *generator.Context) []string {
	consts := []string{}
	buffer := &bytes.Buffer{}
	sw := generator.NewSnippetWriter(buffer, c, "$", "$")
	sw.Do(fmt.Sprintf("Group = \"%s.%s\"", d.group, d.domain), nil)
	consts = append(consts, buffer.String())
	return consts
}

const unversionedListType = "type %sList struct { \n" +
	"    metav1.TypeMeta \n" +
	"    metav1.ListMeta \n" +
	"    Items []%s" +
	"}\n\n"

type GenerateTypes struct {
	Kind  string
	IsApi bool
	Types []*types.Type
}

func (d *unversionedGenerator) CreateGenerateTypes(kind string, isAPI bool) GenerateTypes {
	types := GenerateTypes{
		Kind:  kind,
		IsApi: isAPI,
	}
	for _, t := range d.allTypesByKindVersion[kind] {
		for _, mt := range t.Members {
			if strings.Contains(mt.Tags, "patchStrategy=") {
				// Parse "patchStrategy" value
				// Emit patchStrategy and mergeKey extensions
			}
		}

		types.Types = append(types.Types, t)
	}
	return types
}

func (d *unversionedGenerator) CreateGenerateTypesList() []GenerateTypes {
	types := []GenerateTypes{}
	for _, k := range d.versionedApiTypes {
		types = append(types, d.CreateGenerateTypes(k, true))
	}
	for _, k := range d.subresourceApiTypes {
		types = append(types, d.CreateGenerateTypes(k, false))
	}
	return types
}

func (d *unversionedGenerator) Finalize(context *generator.Context, w io.Writer) error {
	for _, sr := range d.subresources {
		fmt.Printf("Doing Sub %s\n", sr.Path)
		Templates.subresourceTemplate.Execute(w, sr)
	}

	toGenerate := d.CreateGenerateTypesList()
	generated := sets.String{}

	// While there are types to process
	for len(toGenerate) > 0 {
		// Pop the next element from the list
		gen := toGenerate[0]
		toGenerate[0] = toGenerate[len(toGenerate)-1]
		toGenerate = toGenerate[:len(toGenerate)-1]

		// Already processed this type
		name := gen.Kind
		if generated.Has(name) {
			continue
		}

		nextGen, _ := d.DoType(context, gen, w)
		toGenerate = append(toGenerate, nextGen...)
	}

	// For each kind write the strategy, REST, storage, and Registry
	for _, k := range d.GetListOfKinds() {
		if err := Templates.kindTemplate.Execute(w, KindTemplateArgs{d.group, k, d.subresources}); err != nil {
			panic(errors.Errorf("Failed to execute template %v", err))
		}
	}

	return nil
}

// GetListOfKinds returns the list of unique kinds in the group
func (d *unversionedGenerator) GetListOfKinds() []string {
	kinds := sets.String{}
	for _, m := range d.typesByVersionKind {
		for _, t := range m {
			kinds.Insert(t.Name.Name)
		}
	}
	return kinds.List()
}

func (d *unversionedGenerator) DoType(c *generator.Context, versionedTypes GenerateTypes, w io.Writer) ([]GenerateTypes, error) {
	generatedMembers := map[string]types.Member{}
	additionalTypes := []GenerateTypes{}

	for _, t := range versionedTypes.Types {
		for _, m := range t.Members {

			if lm, f := generatedMembers[m.Name]; f {
				if lm.Type.Name.Name != m.Type.Name.Name {
					panic(errors.Errorf(
						"Cannot have different versions of the same kind with the same"+
							"field and different types.  Field name: %s, type: %s,"+
							"type %s", m.Name, lm.Type.Name.Name, m.Type.Name.Name))
				}
			} else {
				generatedMembers[m.Name] = m

				if IsGroup(m.Type, d.group) && !m.Type.IsPrimitive() {
					add := d.CreateGenerateTypes(m.Type.Name.Name, false)
					additionalTypes = append(additionalTypes, add)
				}
			}
		}
	}

	//version := GetVersion(t, d.group)
	//d.typesToGenerate = append(d.typesToGenerate, m.Type)

	sw := generator.NewSnippetWriter(w, c, "$", "$")
	o := ""
	if versionedTypes.IsApi {
		o = o + fmt.Sprintf("// +genclient=true\n\ntype %s struct {\n", versionedTypes.Kind)
	} else {
		o = o + fmt.Sprintf("type %s struct {\n", versionedTypes.Kind)
	}

	for _, m := range generatedMembers {
		group := GetGroup(m.Type)
		kind := m.Type.Name.Name

		// Special case meta data
		if kind == "TypeMeta" || kind == "ObjectMeta" {
			group = "metav1"
		}

		field := m.Name
		if m.Embedded {
			// Embedded fields don't have names
			field = ""
		}

		if m.Type.IsPrimitive() {
			o = o + fmt.Sprintf("\t%s %s\n\n", field, kind)
		} else if group == d.group {
			o = o + fmt.Sprintf("\t%s %s\n\n", field, kind)
		} else {
			o = o + fmt.Sprintf("\t%s %s.%s\n\n", field, group, kind)
		}
	}
	o = o + fmt.Sprintf("}\n\n")
	sw.Do(o, nil)

	if versionedTypes.IsApi {
		sw := generator.NewSnippetWriter(w, c, "$", "$")
		listType := fmt.Sprintf(unversionedListType, versionedTypes.Kind, versionedTypes.Kind)
		sw.Do(listType, nil)
	}
	return additionalTypes, nil
}

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

	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

type versionedGenerator struct {
	generator.DefaultGen
	pkg          *types.Package
	version      string
	group        string
	domain       string
	apiTypeNames []string
}

var _ generator.Generator = &versionedGenerator{}

func CreateVersionedGenerator(
	c *generator.Context, pkg *types.Package, arguments *args.GeneratorArgs,
	group string, version string, domain string) generator.Generator {
	return &versionedGenerator{
		generator.DefaultGen{OptionalName: arguments.OutputFileBaseName},
		pkg,
		version,
		group,
		domain,
		GetApiTypeNames(c, group),
	}
}

func (d *versionedGenerator) Filter(c *generator.Context, t *types.Type) bool {
	return true
}
func (d *versionedGenerator) Namers(c *generator.Context) namer.NameSystems {
	return nil
}
func (d *versionedGenerator) Imports(c *generator.Context) []string {
	return []string{
		"metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apimachinery/pkg/runtime/schema"}
}

func (d *versionedGenerator) PackageVars(c *generator.Context) []string {
	vars := []string{}
	buffer := &bytes.Buffer{}

	apiTypes := d.apiTypeNames
	types := []string{}
	for _, n := range apiTypes {
		types = append(types, fmt.Sprintf("&%s{}", n))
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
	sw.Do("SchemeGroupVersion = schema.GroupVersion{Group, Version}\n", nil)
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

func (d *versionedGenerator) PackageConsts(c *generator.Context) []string {
	consts := []string{}

	buffer := &bytes.Buffer{}
	sw := generator.NewSnippetWriter(buffer, c, "$", "$")
	sw.Do(fmt.Sprintf("Group = \"%s.%s\"", d.group, d.domain), nil)
	consts = append(consts, buffer.String())

	buffer.Reset()
	sw = generator.NewSnippetWriter(buffer, c, "$", "$")
	sw.Do(fmt.Sprintf("Version = \"%s\"", d.version), nil)
	consts = append(consts, buffer.String())

	return consts
}

const listType = "type %sList struct { \n" +
	"    metav1.TypeMeta `json:\",inline\"` \n" +
	"    metav1.ListMeta `json:\"metadata,omitempty\" protobuf:\"bytes,1,opt,name=metadata\"` \n" +
	"    Items []%s `json:\"items\" protobuf:\"bytes,2,rep,name=items\"`\n" +
	"}\n\n"

func (d *versionedGenerator) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	if !IsApiType(t) || !IsGroup(t, d.group) {
		return nil
	}
	name := t.Name.Name
	sw := generator.NewSnippetWriter(w, c, "$", "$")
	sw.Do(fmt.Sprintf(listType, name, name), nil)
	return nil
}

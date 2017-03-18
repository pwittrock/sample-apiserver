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
	"io"

	"k8s.io/gengo/generator"
	"text/template"
)

type versionedGenerator struct {
	generator.DefaultGen
	apiversion *APIVersion
	apigroup   *APIGroup
}

var _ generator.Generator = &versionedGenerator{}

func CreateVersionedGenerator(apiversion *APIVersion, apigroup *APIGroup, filename string) generator.Generator {
	return &versionedGenerator{
		generator.DefaultGen{OptionalName: filename},
		apiversion,
		apigroup,
	}
}

func (d *versionedGenerator) Imports(c *generator.Context) []string {
	return []string{
		"metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apiserver-builder/pkg/builders",
		d.apigroup.Pkg.Path,
	}
}

func (d *versionedGenerator) Finalize(context *generator.Context, w io.Writer) error {
	temp := template.Must(template.New("versioned-template").Parse(VersionedAPITemplate))
	return temp.Execute(w, d.apiversion)
}

var VersionedAPITemplate = `
var (
	ApiVersion = builders.NewVersionedApiBuilder("{{.Group}}.{{.Domain}}", "{{.Version}}").WithResources(
		{{ range $api := .Resources -}}
		builders.NewVersionedResourceWithStorage( //  Resource endpoint
			{{ $api.Group }}.{{ $api.Kind }}Singleton,
			func() runtime.Object { return &{{ $api.Kind }}{} },     // Register versioned resource
			func() runtime.Object { return &{{ $api.Kind }}List{} }, // Register versioned resource list
			&{{ $api.Group }}.{{ $api.Kind }}Strategy{builders.StorageStrategySingleton},
		),
		builders.NewVersionedResourceWithStorage( // Resource status endpoint
			{{ $api.Group }}.{{ $api.Kind }}StatusSingleton,
			func() runtime.Object { return &{{ $api.Kind }}{} },     // Register versioned resource
			func() runtime.Object { return &{{ $api.Kind }}List{} }, // Register versioned resource list
			&{{ $api.Group }}.{{ $api.Kind }}StatusStrategy{builders.StatusStorageStrategySingleton},
		),
		{{ range $subresource := $api.Subresources -}}
		builders.NewVersionedResourceWithoutStorage(
			{{ $api.Group }}.{{ $subresource.REST }}Singleton,
			func() runtime.Object { return &{{ $subresource.Request }}{} }, // Register versioned resource
			&{{ $api.Group }}.{{ $subresource.REST }}{},
		),
		{{ end -}}
		{{ end -}}
	)

	// Expected by generated deepcopy and conversion
	SchemeBuilder = ApiVersion.SchemaBuilder
)

{{ range $api := .Resources -}}
type {{$api.Kind}}List struct {
	metav1.TypeMeta ` + "`json:\",inline\"`" + `
	metav1.ListMeta ` + "`json:\"metadata,omitempty\"`" + `
	Items           []{{$api.Kind}} ` + "`json:\"items\"`" + `
}
{{ end -}}
`

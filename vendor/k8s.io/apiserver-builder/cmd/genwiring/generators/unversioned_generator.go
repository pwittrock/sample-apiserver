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
	"text/template"

	"k8s.io/gengo/generator"
)

type unversionedGenerator struct {
	generator.DefaultGen
	apigroup *APIGroup
}

var _ generator.Generator = &unversionedGenerator{}

func CreateUnversionedGenerator(apigroup *APIGroup, filename string) generator.Generator {
	return &unversionedGenerator{
		generator.DefaultGen{OptionalName: filename},
		apigroup,
	}
}

func (d *unversionedGenerator) Imports(c *generator.Context) []string {
	return []string{
		"fmt",
		"k8s.io/apimachinery/pkg/apis/meta/internalversion",
		"metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apiserver-builder/pkg/builders",
		"k8s.io/apiserver/pkg/endpoints/request",
		"k8s.io/apiserver/pkg/registry/rest",
		"k8s.io/client-go/pkg/api",
	}
}

func (d *unversionedGenerator) Finalize(context *generator.Context, w io.Writer) error {
	temp := template.Must(template.New("unversioned-wiring-template").Parse(UnversionedAPITemplate))
	err := temp.Execute(w, d.apigroup)
	if err != nil {
		return err
	}
	return err
}

var UnversionedAPITemplate = `
var (
	{{ range $api := .UnversionedResources -}}
	{{ $api.Kind }}Singleton = builders.NewUnversionedResource(
		"{{ $api.Resource }}",
		func() runtime.Object { return &{{ $api.Kind }}{} },
		func() runtime.Object { return &{{ $api.Kind }}List{} },
	)
	{{ $api.Kind }}StatusSingleton = builders.NewUnversionedStatus(
		"{{ $api.Resource }}",
		func() runtime.Object { return &{{ $api.Kind }}{} },
		func() runtime.Object { return &{{ $api.Kind }}List{} },
	)
	{{ range $subresource := .Subresources -}}
	{{$subresource.REST}}Singleton = builders.NewUnversionedSubresource(
		"{{$subresource.Resource}}", "{{$subresource.Path}}",
		func() runtime.Object { return &{{$subresource.Request}}{} },
	)
	{{ end -}}
	{{ end -}}

	// Registered resources and subresources
	ApiVersion = builders.NewUnVersionedApiBuilder("{{.Group}}.{{.Domain}}").WithKinds(
		{{ range $api := .UnversionedResources -}}
		{{$api.Kind}}Singleton,
		{{$api.Kind}}StatusSingleton,
		{{ range $subresource := $api.Subresources -}}
		{{$subresource.REST}}Singleton,
		{{ end -}}
		{{ end -}}
	)
	SchemeBuilder = ApiVersion.SchemaBuilder
)

{{ range $s := .Structs -}}
type {{ $s.Name }} struct {
{{ range $f := $s.Fields -}}
    {{ $f.Name }} {{ $f.Type }}
{{ end -}}
}

{{ end -}}

{{ range $api := .UnversionedResources -}}
//
// {{.Kind}} Functions and Structs
//
type {{.Kind}}Strategy struct {
	builders.DefaultStorageStrategy
}

type {{$api.Kind}}StatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

type {{$api.Kind}}List struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []{{$api.Kind}}
}

func ({{$api.Kind}}) NewStatus() interface{} {
	return {{$api.Kind}}Status{}
}

func (pc *{{$api.Kind}}) GetStatus() interface{} {
	return pc.Status
}

func (pc *{{$api.Kind}}) SetStatus(s interface{}) {
	pc.Status = s.({{$api.Kind}}Status)
}

func (pc *{{$api.Kind}}) GetSpec() interface{} {
	return pc.Status
}

func (pc *{{$api.Kind}}) SetSpec(s interface{}) {
	pc.Spec = s.({{$api.Kind}}Spec)
}

func (pc *{{$api.Kind}}) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *{{$api.Kind}}) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc {{$api.Kind}}) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store {{.Kind}}.
type {{.Kind}}Registry interface {
	List{{.Kind}}s(ctx request.Context, options *internalversion.ListOptions) (*{{.Kind}}List, error)
	Get{{.Kind}}(ctx request.Context, id string, options *metav1.GetOptions) (*{{.Kind}}, error)
	Create{{.Kind}}(ctx request.Context, id *{{.Kind}}) (*{{.Kind}}, error)
	Update{{.Kind}}(ctx request.Context, id *{{.Kind}}) (*{{.Kind}}, error)
	Delete{{.Kind}}(ctx request.Context, id string) error
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func New{{.Kind}}Registry(s rest.StandardStorage) {{.Kind}}Registry {
	return &storage{{.Kind}}{s}
}

// Implement Registry
// storage puts strong typing around storage calls
type storage{{.Kind}} struct {
	rest.StandardStorage
}

func (s *storage{{.Kind}}) List{{.Kind}}s(ctx request.Context, options *internalversion.ListOptions) (*{{.Kind}}List, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Kind}}List), err
}

func (s *storage{{.Kind}}) Get{{.Kind}}(ctx request.Context, id string, options *metav1.GetOptions) (*{{.Kind}}, error) {
	obj, err := s.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Kind}}), nil
}

func (s *storage{{.Kind}}) Create{{.Kind}}(ctx request.Context, object *{{.Kind}}) (*{{.Kind}}, error) {
	obj, err := s.Create(ctx, object)
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Kind}}), nil
}

func (s *storage{{.Kind}}) Update{{.Kind}}(ctx request.Context, object *{{.Kind}}) (*{{.Kind}}, error) {
	obj, _, err := s.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object, api.Scheme))
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Kind}}), nil
}

func (s *storage{{.Kind}}) Delete{{.Kind}}(ctx request.Context, id string) error {
	_, err := s.Delete(ctx, id, nil)
	return err
}

{{ end -}}
`

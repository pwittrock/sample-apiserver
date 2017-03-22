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

type KindTemplateArgs struct {
	Group string
	Kind  string
}

const kindTemplateName = "KindTemplate"
const kindTemplateString = `

///////////////////////////////////////////////////////////////////////////////
// {{.Kind}} user functions //
///////////////////////////////////////////////////////////////////////////////

// Add functions to this type in order to override the default behaviors
type {{.Kind}}Strategy struct {
	Default{{.Kind}}Strategy
}

// Add functions to this type in order to override the default behaviors
type {{.Kind}}Store struct {
	*genericregistry.Store
}

// Add functions to this type in order to override the default behaviors
type {{.Kind}}StatusStore struct {
	*genericregistry.Store
}

// Registry is an interface for things that know how to store {{.Kind}}.
type {{.Kind}}Registry interface {
	List{{.Kind}}s(ctx genericapirequest.Context, options *metainternalversion.ListOptions) (*{{.Kind}}List, error)
	Get{{.Kind}}(ctx genericapirequest.Context, id string, options *metav1.GetOptions) (*{{.Kind}}, error)
	Create{{.Kind}}(ctx genericapirequest.Context, id *{{.Kind}}) (*{{.Kind}}, error)
	Update{{.Kind}}(ctx genericapirequest.Context, id *{{.Kind}}) (*{{.Kind}}, error)
	Delete{{.Kind}}(ctx genericapirequest.Context, id string) error
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func New{{.Kind}}Registry(s rest.StandardStorage) {{.Kind}}Registry {
	return &storage{{.Kind}}{s}
}

///////////////////////////////////////////////////////////////////////////////
// {{.Kind}} System functions //
///////////////////////////////////////////////////////////////////////////////

// Use the override strategy and embedd the defaults for anything not override.
var {{.Kind}}StrategySingleton = &{{.Kind}}Strategy{
	Default{{.Kind}}Strategy{ // Overide some methods
		defaults.NewBasicStrategy(), // Use defaults
	},
}

// Default Strategy for {{.Kind}}
type Default{{.Kind}}Strategy struct {
	// Inherit the basic create, delete, update strategy.
	defaults.BasicCreateDeleteUpdateStrategy
}

// NewFunc returns a new empty {{.Kind}}
func (r Default{{.Kind}}Strategy) NewFunc() runtime.Object {
	return &{{.Kind}}{}
}

// NewListFunc returns a new empty List of {{.Kind}}
func (r Default{{.Kind}}Strategy) NewListFunc() runtime.Object {
	return &{{.Kind}}List{}
}

// ObjectNameFunc returns the name for a {{.Kind}}
func (r Default{{.Kind}}Strategy) ObjectNameFunc(obj runtime.Object) (string, error) {
	return obj.(*{{.Kind}}).Name, nil
}

func ({{.Kind}}Strategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
	o := obj.(*{{.Kind}})
	o.Status = {{.Kind}}Status{}
	o.Generation = 1
}

func ({{.Kind}}Strategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
	new{{.Kind}} := obj.(*{{.Kind}})
	old{{.Kind}} := old.(*{{.Kind}})
	new{{.Kind}}.Status = old{{.Kind}}.Status

	// Spec and annotation updates bump the generation.
	if !reflect.DeepEqual(new{{.Kind}}.Spec, old{{.Kind}}.Spec) ||
		!reflect.DeepEqual(new{{.Kind}}.Annotations, old{{.Kind}}.Annotations) {
		new{{.Kind}}.Generation = old{{.Kind}}.Generation + 1
	}
}

// Implement Status endpoint
// StatusREST implements the REST endpoint for changing the status of a deployment
type {{.Kind}}StatusStrategy struct {
	{{.Kind}}Strategy
}

// {{.Kind}}StatusStrategySingleton contains the cross-cutting storage
var {{.Kind}}StatusStrategySingleton = {{.Kind}}StatusStrategy{*{{.Kind}}StrategySingleton}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update of status
func ({{.Kind}}StatusStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
	new{{.Kind}} := obj.(*{{.Kind}})
	old{{.Kind}} := old.(*{{.Kind}})
	new{{.Kind}}.Spec = old{{.Kind}}.Spec
	new{{.Kind}}.Labels = old{{.Kind}}.Labels
}

// Implement Registry
// storage puts strong typing around storage calls
type storage{{.Kind}} struct {
	rest.StandardStorage
}

func (s *storage{{.Kind}}) List{{.Kind}}s(ctx genericapirequest.Context, options *metainternalversion.ListOptions) (*{{.Kind}}List, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Kind}}List), err
}

func (s *storage{{.Kind}}) Get{{.Kind}}(ctx genericapirequest.Context, id string, options *metav1.GetOptions) (*{{.Kind}}, error) {
	obj, err := s.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Kind}}), nil
}

func (s *storage{{.Kind}}) Create{{.Kind}}(ctx genericapirequest.Context, object *{{.Kind}}) (*{{.Kind}}, error) {
	obj, err := s.Create(ctx, object)
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Kind}}), nil
}

func (s *storage{{.Kind}}) Update{{.Kind}}(ctx genericapirequest.Context, object *{{.Kind}}) (*{{.Kind}}, error) {
	obj, _, err := s.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object, api.Scheme))
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Kind}}), nil
}

func (s *storage{{.Kind}}) Delete{{.Kind}}(ctx genericapirequest.Context, id string) error {
	_, err := s.Delete(ctx, id, nil)
	return err
}


`

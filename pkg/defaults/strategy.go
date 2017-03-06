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

package defaults

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/apiserver/pkg/registry/rest"
)

type BasicApiServerStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var _ rest.RESTCreateStrategy = &BasicApiServerStrategy{}
var _ rest.RESTDeleteStrategy = &BasicApiServerStrategy{}
var _ rest.RESTUpdateStrategy = &BasicApiServerStrategy{}
//var _ rest.RESTGracefulDeleteStrategy = &BasicApiServerStrategy{}
//var _ rest.RESTExportStrategy = &BasicApiServerStrategy{}

type HasObjectMeta interface {
	GetObjectMeta() *metav1.ObjectMeta
}

func NewBasicStrategy(typer runtime.ObjectTyper) BasicApiServerStrategy {
	return BasicApiServerStrategy{typer, names.SimpleNameGenerator}
}

func (BasicApiServerStrategy) NamespaceScoped() bool {
	return false
}

func (BasicApiServerStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
}

func (BasicApiServerStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
}

func (BasicApiServerStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
	// return validation.ValidateFlunder(obj.(*wardle.Flunder))
}

func (BasicApiServerStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (BasicApiServerStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (BasicApiServerStrategy) Canonicalize(obj runtime.Object) {
}

func (BasicApiServerStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
	// return validation.ValidateFlunderUpdate(obj.(*wardle.Flunder), old.(*wardle.Flunder))
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	switch t := obj.(type) {
	case HasObjectMeta:
		apiserver := obj.(HasObjectMeta)
		return labels.Set(apiserver.GetObjectMeta().Labels), GetSelectableFields(apiserver), nil
	default:
		return nil, nil, fmt.Errorf("given object type %v does not have ObjectMeta.", t)
	}
}

// MatchResource is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func (BasicApiServerStrategy) BasicMatch(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// GetSelectableFields returns a field set that represents the object.
func GetSelectableFields(obj HasObjectMeta) fields.Set {
	return generic.ObjectMetaFieldsSet(obj.GetObjectMeta(), true)
}
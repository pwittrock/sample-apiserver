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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// rest implements a RESTStorage for API services against etcd
type BasicREST struct {
	*registry.Store
}

type RESTFactory struct {
	Scheme     *runtime.Scheme
	OptsGetter generic.RESTOptionsGetter
}

type RestFunctions interface {
	NewFunc() runtime.Object
	NewListFunc() runtime.Object
	ObjectNameFunc(obj runtime.Object) (string, error)
}

// NewREST returns a RESTStorage object that will work against API services.
func (f *RESTFactory) NewBasicREST(groupResource schema.GroupResource, fns RestFunctions) *BasicREST {
	strategy := NewBasicStrategy(f.Scheme)

	store := &registry.Store{
		Copier:      f.Scheme,
		NewFunc:     fns.NewFunc,
		NewListFunc: fns.NewListFunc,
		ObjectNameFunc: fns.ObjectNameFunc,
		PredicateFunc:     strategy.BasicMatch,
		QualifiedResource: groupResource,

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,
	}
	options := &generic.StoreOptions{RESTOptions: f.OptsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}
	return &BasicREST{store}
}

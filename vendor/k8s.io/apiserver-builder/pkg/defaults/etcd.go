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
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

type StorageFactory struct {
	Scheme     *runtime.Scheme
	OptsGetter generic.RESTOptionsGetter
}

type StorageStrategy interface {
	NewFunc() runtime.Object
	NewListFunc() runtime.Object
	ObjectNameFunc(obj runtime.Object) (string, error)
}

// NewREST returns a RESTStorage object that will work against ResourceDefinition services.
func (f *StorageFactory) Create(groupResource schema.GroupResource, resourceDef *ResourceDefinition) rest.Storage {
	store := &registry.Store{
		Copier:            f.Scheme,
		NewFunc:           resourceDef.StorageStrategy.NewFunc,
		NewListFunc:       resourceDef.StorageStrategy.NewListFunc,
		ObjectNameFunc:    resourceDef.StorageStrategy.ObjectNameFunc,
		PredicateFunc:     resourceDef.PredicateFunc,
		CreateStrategy:    resourceDef.CreateStrategy,
		UpdateStrategy:    resourceDef.UpdateStrategy,
		DeleteStrategy:    resourceDef.DeleteStrategy,
		QualifiedResource: groupResource,
		WatchCacheSize:    1000,
	}

	options := &generic.StoreOptions{RESTOptions: f.OptsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}
	return resourceDef.CreateStorageFunc(store)
}

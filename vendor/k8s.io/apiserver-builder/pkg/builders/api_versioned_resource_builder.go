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

package builders

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/client-go/pkg/api"
)

//
// Versioned Kind Builder builds a versioned resource using unversioned strategy
//

// NewVersionedResourceWithStorage returns a new versionedResourceBuilder for registering endpoints for
// resources that are persisted to storage.
// strategy - unversionedBuilder from calling NewUnversionedXXX()
// new - function for creating new empty VERSIONED instances - e.g. func() runtime.Object { return &Deployment{} }
// newList - function for creating an empty list of VERSIONED instances - e.g. func() runtime.Object { return &DeploymentList{} }
// storeFunc - override storage defaults
func NewVersionedResourceWithStorage(
	unversionedBuilder UnversionedResourceBuilder,
	new, newList func() runtime.Object,
	storeBuilder StorageBuilder) *versionedResourceBuilder {

	if storeBuilder == nil {
		storeBuilder = StorageStrategySingleton
	}

	return &versionedResourceBuilder{
		unversionedBuilder, new, newList, storeBuilder, nil,
	}
}

// NewVersionedResourceWithoutStorage returns a new versionedResourceBuilder for registering endpoints for
// subresources that are not persisted to storage.
// strategy - unversionedBuilder from calling NewUnversionedXXX()
// new - function for creating new empty VERSIONED instances - e.g. func() runtime.Object { return &Deployment{} }
// newList - function for creating an empty list of VERSIONED instances - e.g. func() runtime.Object { return &DeploymentList{} }
// restFunc - returns the REST implementation for this resource
func NewVersionedResourceWithoutStorage(
	unversionedBuilder UnversionedResourceBuilder,
	new func() runtime.Object,
	storage rest.Storage) *versionedResourceBuilder {
	return &versionedResourceBuilder{
		unversionedBuilder, new, nil, nil, storage,
	}
}

type versionedResourceBuilder struct {
	Unversioned UnversionedResourceBuilder

	// NewFunc returns an empty unversioned instance of a resource
	NewFunc func() runtime.Object

	// NewListFunc returns and empty unversioned instance of a resource List
	NewListFunc func() runtime.Object

	// Store is used to modify the default storage, mutually exclusive with RESTFunc
	StorageBuilder StorageBuilder

	// REST a rest.Store implementation, mutually exclusive with StoreFunc
	REST rest.Storage
}

func (b *versionedResourceBuilder) New() runtime.Object {
	if b.NewFunc == nil {
		return nil
	}
	return b.NewFunc()
}

func (b *versionedResourceBuilder) NewList() runtime.Object {
	if b.NewListFunc == nil {
		return nil
	}
	return b.NewListFunc()
}

//
func (b *versionedResourceBuilder) Build(
	group string,
	optionsGetter generic.RESTOptionsGetter) rest.Storage {

	// Set a default strategy
	store := &registry.Store{
		Copier:            api.Scheme,
		NewFunc:           b.Unversioned.New,     // Use the unversioned type
		NewListFunc:       b.Unversioned.NewList, // Use the unversioned type
		QualifiedResource: b.getGroupResource(group),
		WatchCacheSize:    1000,
	}

	// Use default, requires
	options := &generic.StoreOptions{RESTOptions: optionsGetter}

	if b.StorageBuilder != nil {
		// Allow overriding the storage defaults
		b.StorageBuilder.Build(b.StorageBuilder, store, options)
	}

	if err := store.CompleteWithOptions(options); err != nil {
		panic(err) // TODO: Propagate error up
	}
	return store
}

// getGroupResource returns the GroupResource for this Resource and the provided Group
// group is the group the resource belongs to
func (b *versionedResourceBuilder) getGroupResource(group string) schema.GroupResource {
	return schema.GroupResource{group, b.Unversioned.GetName()}

}

// registerEndpoints registers the REST endpoints for this resource in the registry
// group is the group to register the resource under
// optionsGetter is the RESTOptionsGetter provided by a server.Config
// registry is the server.APIGroupInfo VersionedResourcesStorageMap used to register REST endpoints
func (b *versionedResourceBuilder) registerEndpoints(
	group string,
	optionsGetter generic.RESTOptionsGetter,
	registry map[string]rest.Storage) {

	// Register the endpoint
	path := b.Unversioned.GetPath()
	if len(path) > 0 {
		// Subresources appear after the resource
		path = b.Unversioned.GetName() + "/" + path
	} else {
		path = b.Unversioned.GetName()
	}

	if b.REST != nil {
		// Use the REST implementation directly.
		registry[path] = b.REST
	} else {
		// Create a new REST implementation wired to storage.
		registry[path] = b.Build(group, optionsGetter)
	}
}

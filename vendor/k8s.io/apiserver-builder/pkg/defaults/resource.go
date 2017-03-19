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
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
)

type ResourceDefinitionProvider interface {
	GetResourceDefinitions() []*ResourceDefinition
	GetLegacyCodec() []schema.GroupVersion
	GetGroupName() string
	GetVersionPreferenceOrder() []string
	GetImportPrefix() string
	SchemeFunc() announced.SchemeFunc
	VersionToSchemeFunc() announced.VersionToSchemeFunc
}

type ResourceDefinition struct {
	GroupVersionResource schema.GroupVersionResource
	StorageStrategy      StorageStrategy
	CreateStrategy       rest.RESTCreateStrategy
	DeleteStrategy       rest.RESTDeleteStrategy
	UpdateStrategy       rest.RESTUpdateStrategy
	SubResources         map[string]rest.Storage
	PredicateFunc        func(label labels.Selector, field fields.Selector) storage.SelectionPredicate
}

func (api *ResourceDefinition) GetExternalGroupVersionResource() schema.GroupVersionResource {
	return api.GroupVersionResource
}

func (api *ResourceDefinition) GetInternalGroupVersionResource() schema.GroupVersionResource {
	return api.GroupVersionResource.GroupResource().WithVersion(runtime.APIVersionInternal)
}

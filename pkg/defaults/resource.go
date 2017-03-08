package defaults

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/storage"
)

type ResourceDefinitionProvider interface {
	GetResourceDefinitions() []*ResourceDefinition
	GetLegacyCodec() []schema.GroupVersion
	GetGroupName()              string
	GetVersionPreferenceOrder() []string
	GetImportPrefix() string
	SchemeFunc() announced.SchemeFunc
	VersionToSchemeFunc() announced.VersionToSchemeFunc
}

type ResourceDefinition struct {
	GroupVersionResource   schema.GroupVersionResource
	StorageStrategy        StorageStrategy
	CreateStrategy         rest.RESTCreateStrategy
	DeleteStrategy         rest.RESTDeleteStrategy
	UpdateStrategy         rest.RESTUpdateStrategy
	PredicateFunc 		func(label labels.Selector, field fields.Selector) storage.SelectionPredicate
}

func (api *ResourceDefinition) GetExternalGroupVersionResource() schema.GroupVersionResource {
	return api.GroupVersionResource
}

func (api *ResourceDefinition) GetInternalGroupVersionResource() schema.GroupVersionResource {
	return api.GroupVersionResource.GroupResource().WithVersion(runtime.APIVersionInternal)
}

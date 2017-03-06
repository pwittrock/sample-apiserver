package defaults

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
)

type APIFactory struct {
	RestFactory          RESTFactory
	GroupFactoryRegistry announced.APIGroupFactoryRegistry
	Registry             *registered.APIRegistrationManager
	Scheme               *runtime.Scheme
	Codecs               serializer.CodecFactory
}

type API struct {
	GroupVersionResource schema.GroupVersionResource
	Functions RestFunctions
}

func (api *API) GetExternalGroupVersionResource() schema.GroupVersionResource {
	return api.GroupVersionResource
}


func (api *API) GetInternalGroupVersionResource() schema.GroupVersionResource {
	return api.GroupVersionResource.GroupResource().WithVersion(runtime.APIVersionInternal)
}

func (f *APIFactory) createApiGroup(api *API) *genericapiserver.APIGroupInfo {
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(
		api.GroupVersionResource.Group,
		f.Registry,
		f.Scheme,
		metav1.ParameterCodec,
		f.Codecs)
	apiGroupInfo.GroupMeta.GroupVersion = api.GetExternalGroupVersionResource().GroupVersion()
	return &apiGroupInfo
}

func (f *APIFactory) CreateAPI(api *API) (*genericapiserver.APIGroupInfo, error) {
	r := api.GroupVersionResource.Resource
	v := api.GetExternalGroupVersionResource().Version

	// Initialize group
	apiGroupInfo := f.createApiGroup(api)

	// Initialize version if it doesn't exist
	if _, found := apiGroupInfo.VersionedResourcesStorageMap[v]; !found {
		apiGroupInfo.VersionedResourcesStorageMap[v] = map[string]rest.Storage{}
	}
	// Initialize resource at version
	apiGroupInfo.VersionedResourcesStorageMap[v][r] = f.RestFactory.NewBasicREST(
		api.GetInternalGroupVersionResource().GroupResource(), api.Functions)

	return apiGroupInfo, nil
}

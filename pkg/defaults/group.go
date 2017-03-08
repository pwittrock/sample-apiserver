package defaults

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	"github.com/pkg/errors"
	"github.com/golang/glog"
)


// APIGroupFactory builds APIGroupInfos from ResourceDefinitions.  The APIGroupInfos
// can be registered with an apiserver to start serving the resource.
type APIGroupFactory struct {
	StorageFactory       StorageFactory
	GroupFactoryRegistry announced.APIGroupFactoryRegistry
	Registry             *registered.APIRegistrationManager
	Scheme               *runtime.Scheme
	Codecs               serializer.CodecFactory
}

// getGroup returns the group name for the set of resources.  If the resources are in
// different api groups, an error is returned.
func (f *APIGroupFactory) getGroup(resources []*ResourceDefinition) (string, error) {
	group := ""
	for _, resource := range resources {
		if len(group) > 0 && group != resource.GroupVersionResource.Group {
			return "", errors.Errorf(
				"APIGroupInfo cannot be created for resources in multiple groups [%s, %s]",
				group, resource.GroupVersionResource.Group)
		}
		group = resource.GroupVersionResource.Group
	}
	return group, nil
}

// newGroupInfo returns a new APIGroupInfo for the group name.
func (f *APIGroupFactory) newGroupInfo(group string) *genericapiserver.APIGroupInfo {
	g := genericapiserver.NewDefaultAPIGroupInfo(
		group,
		f.Registry,
		f.Scheme,
		metav1.ParameterCodec,
		f.Codecs)
	return &g
}

// Create takes a group of ResourceDefinitions and returns a APIGroupInfo to register them.
// All ResourceDefinitions must belong to the same api group.
// This function should only be called once per-group.
func (f *APIGroupFactory) Create(resources []*ResourceDefinition) (*genericapiserver.APIGroupInfo, error) {
	group, err := f.getGroup(resources)
	if err != nil {
		return nil, err
	}

	// Create the group info
	apiGroupInfo := f.newGroupInfo(group)

	glog.Infof("Creating group %v", group)

	// Add each of the resource definitions to the groupinfo
	for _, definition := range resources {
		resource := definition.GroupVersionResource.Resource
		version := definition.GetExternalGroupVersionResource().Version
		groupResource := definition.GetInternalGroupVersionResource().GroupResource()
		apiGroupInfo.GroupMeta.GroupVersion = definition.GetExternalGroupVersionResource().GroupVersion()

		// Initialize version if it doesn't exist
		if _, found := apiGroupInfo.VersionedResourcesStorageMap[version]; !found {
			apiGroupInfo.VersionedResourcesStorageMap[version] = map[string]rest.Storage{}
		}
		apiGroupInfo.VersionedResourcesStorageMap[version][resource] =
			f.StorageFactory.Create(groupResource, definition)
	}
	return apiGroupInfo, nil
}

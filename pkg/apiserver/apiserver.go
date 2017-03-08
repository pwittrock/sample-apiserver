/*
Copyright 2016 The Kubernetes Authors.

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

package apiserver

import (
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/apimachinery/registered"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/version"
	genericapiserver "k8s.io/apiserver/pkg/server"

	"k8s.io/sample-apiserver/pkg/defaults"
	"k8s.io/apimachinery/pkg/util/sets"
	"github.com/golang/glog"
)


type Installer struct {
	GroupFactoryRegistry announced.APIGroupFactoryRegistry
	Registry *registered.APIRegistrationManager
	Scheme *runtime.Scheme
}

func (c *Config) Init() *Config {

	i := Installer{defaults.GroupFactoryRegistry, defaults.Registry, defaults.Scheme}
	for _, provider := range defaults.APIProviders {
		i.Install(provider)
	}

	// we need to add the options to empty v1
	// TODO fix the server code to avoid this
	metav1.AddToGroupVersion(defaults.Scheme, schema.GroupVersion{Version: "v1"})

	// TODO: keep the generic ResourceDefinition server from wanting this
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	defaults.Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)

	return c
}

type Config struct {
	GenericConfig *genericapiserver.Config
}

// WardleServer contains state for a Kubernetes cluster master/api server.
type WardleServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

type completedConfig struct {
	*Config
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
func (c *Config) Complete() completedConfig {
	c.GenericConfig.Complete()

	c.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	return completedConfig{c}
}

func (c *Config) AddApi(provider defaults.ResourceDefinitionProvider) *Config {
	defaults.APIProviders = append(defaults.APIProviders, provider)
	return c
}

// SkipComplete provides a way to construct a server instance without config completion.
func (c *Config) SkipComplete() completedConfig {
	return completedConfig{c}
}

// New returns a new instance of WardleServer from the given config.
func (c completedConfig) New() (*WardleServer, error) {
	genericServer, err := c.Config.GenericConfig.SkipComplete().New() // completion is done in Complete, no need for a second time
	if err != nil {
		return nil, err
	}

	s := &WardleServer{
		GenericAPIServer: genericServer,
	}

	apiGroupFactory := &defaults.APIGroupFactory{
		defaults.StorageFactory{defaults.Scheme, c.GenericConfig.RESTOptionsGetter},
		defaults.GroupFactoryRegistry,
		defaults.Registry,
		defaults.Scheme,
		defaults.Codecs,
	}

	glog.Infof("Provider count %v", len(defaults.APIProviders))
	for _, provider := range defaults.APIProviders {
		apiGroupInfo, err := apiGroupFactory.Create(provider.GetResourceDefinitions())
		if err != nil {
			return nil, err
		}
		if err := s.GenericAPIServer.InstallAPIGroup(apiGroupInfo); err != nil {
			return nil, err
		}
	}
	return s, nil
}

// Install registers the ResourceDefinition group and adds types to a scheme
func (i *Installer) Install(group defaults.ResourceDefinitionProvider) {
	glog.Infof("Installing %s %s", group.GetGroupName(), group.GetImportPrefix())
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:                  group.GetGroupName(),
			RootScopedKinds:            sets.NewString("APIService"),
			VersionPreferenceOrder:     group.GetVersionPreferenceOrder(),
			ImportPrefix:               group.GetImportPrefix(),
			AddInternalObjectsToScheme: group.SchemeFunc(),
		},
		group.VersionToSchemeFunc(),
	).Announce(i.GroupFactoryRegistry).RegisterAndEnable(i.Registry, i.Scheme); err != nil {
		panic(err)
	}
}

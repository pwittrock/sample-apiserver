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

package server

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver-builder/pkg/apiserver"
	"k8s.io/apiserver-builder/pkg/defaults"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/openapi"
)

const defaultEtcdPathPrefix = "/registry/wardle.kubernetes.io"

var GetOpenApiDefinition openapi.GetOpenAPIDefinitions

type WardleServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions

	StdOut       io.Writer
	StdErr       io.Writer
	APIProviders []defaults.ResourceDefinitionProvider
}

func NewWardleServerOptions(out, errOut io.Writer, providers []defaults.ResourceDefinitionProvider) *WardleServerOptions {
	versions := []schema.GroupVersion{}
	for _, p := range providers {
		versions = append(versions, p.GetLegacyCodec()...)
	}

	o := &WardleServerOptions{
		RecommendedOptions: genericoptions.NewRecommendedOptions(defaultEtcdPathPrefix, defaults.Scheme, defaults.Codecs.LegacyCodec(versions...)),

		StdOut:       out,
		StdErr:       errOut,
		APIProviders: providers,
	}
	o.RecommendedOptions.SecureServing.ServingOptions.BindPort = 443

	return o
}

var printBearerToken = false

// NewCommandStartMaster provides a CLI handler for 'start master' command
func NewCommandStartWardleServer(out, errOut io.Writer, providers []defaults.ResourceDefinitionProvider, stopCh <-chan struct{}) *cobra.Command {
	o := NewWardleServerOptions(out, errOut, providers)

	cmd := &cobra.Command{
		Short: "Launch a wardle API server",
		Long:  "Launch a wardle API server",
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}
			if err := o.RunWardleServer(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&printBearerToken, "print-bearer-token", false,
		"If true, print a curl command with the bearer token to test the server")
	o.RecommendedOptions.AddFlags(flags)

	return cmd
}

func (o WardleServerOptions) Validate(args []string) error {
	return nil
}

func (o *WardleServerOptions) Complete() error {
	return nil
}

func (o WardleServerOptions) Config() (*apiserver.Config, error) {
	// TODO have a "real" external address
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost"); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	serverConfig := genericapiserver.NewConfig().WithSerializer(defaults.Codecs)
	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	config := &apiserver.Config{
		GenericConfig: serverConfig,
	}
	return config, nil
}

func (o WardleServerOptions) SetAuthOptions() error {
	return nil
	//config, err := o.Config()
	//if err != nil {
	//	return err
	//}
	//
	//config.GenericConfig.LoopbackClientConfig =
	//authorizationConfig := s.Authorization.ToAuthorizationConfig(sharedInformers)
	//
	//apiAuthenticator, securityDefinitions, err := authenticatorConfig.New()
	//if err != nil {
	//	return nil, nil, fmt.Errorf("invalid authentication config: %v", err)
	//}
	//
	//apiAuthorizer, err := authorizationConfig.New()
	//config.GenericConfig.Authenticator
}

func (o WardleServerOptions) RunWardleServer(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	config.GenericConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(GetOpenApiDefinition, defaults.Scheme)
	//config.GenericConfig.OpenAPIConfig.PostProcessSpec = postProcessOpenAPISpecForBackwardCompatibility
	//config.GenericConfig.OpenAPIConfig.SecurityDefinitions = securityDefinitions
	config.GenericConfig.OpenAPIConfig.Info.Title = "Wardle"

	if printBearerToken {
		glog.Infof("Serving on loopback...")
		glog.Infof("\n\n********************************\nTo test the server run:\n"+
			"curl -k -H \"Authorization: Bearer %s\" %s\n********************************\n\n",
			config.GenericConfig.LoopbackClientConfig.BearerToken,
			config.GenericConfig.LoopbackClientConfig.Host)
		glog.Infof("Local Authorization Token: %s", config.GenericConfig.LoopbackClientConfig.BearerToken)
	}

	for _, provider := range o.APIProviders {
		config.AddApi(provider)
	}

	server, err := config.Init().Complete().New()
	if err != nil {
		return err
	}
	server.GenericAPIServer.PrepareRun().Run(stopCh)

	return nil
}

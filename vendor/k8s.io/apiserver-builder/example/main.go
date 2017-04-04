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

package main

import (
	"flag"
	"os"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver-builder/pkg/builders"
	"k8s.io/apiserver-builder/pkg/cmd/server"
	"k8s.io/apiserver/pkg/util/logs"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	// ACTION REQUIRED: update these with your go import paths and uncomment
	//"github.com/org/repo/pkg/apis"
	//"github.com/org/repo/pkg/openapi"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()
	var apis []*builders.APIGroupBuilder

	// ACTION REQUIRED: uncomment this
	// RegisterTypes the openapi
	//server.GetOpenApiDefinition = openapi.GetOpenAPIDefinitions
	//apis = *apis.GetAllApiBuilders()

	// To disable providers, manually specify the list provided by getKnownProviders()
	cmd := server.NewCommandStartServer(os.Stdout, os.Stderr, apis, wait.NeverStop)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

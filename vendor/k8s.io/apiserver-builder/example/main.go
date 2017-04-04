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
	_ "k8s.io/client-go/plugin/pkg/client/auth" // Enable cloud provider auth
	// ACTION REQUIRED: update these with your go import paths and uncomment
	//"github.com/org/repo/pkg/apis"
	//"github.com/org/repo/pkg/openapi"
)

func main() {
	// ACTION REQUIRED: uncomment this
	//server.StartApiServer(apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions)
}

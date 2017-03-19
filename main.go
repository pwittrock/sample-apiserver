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
	"runtime"

	"github.com/pwittrock/apiserver-helloworld/apis"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver-builder/pkg/cmd/server"
	"k8s.io/apiserver-builder/pkg/defaults"
	"k8s.io/apiserver/pkg/util/logs"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

//go:generate go run vendor/k8s.io/apiserver-builder/cmd/genwiring/main.go --input-dirs ./apis/...
func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	providers := []defaults.ResourceDefinitionProvider{apis.GetMushroomkingdomProvider(), apis.GetHyruleProvider(), apis.GetWardleProvider()}
	cmd := server.NewCommandStartWardleServer(os.Stdout, os.Stderr, providers, wait.NeverStop)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

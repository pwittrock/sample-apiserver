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
	"fmt"
	"github.com/pwittrock/apiserver-helloworld/pkg/apis"
	v2mushroomkingdom "github.com/pwittrock/apiserver-helloworld/pkg/apis/mushroomkingdom/v2"
	"github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/clientset"
	"github.com/pwittrock/apiserver-helloworld/pkg/openapi"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver-builder/pkg/test"
)

func main() {
	t := test.NewTestEnvironment()
	config := t.Start(apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions)
	client := clientset.NewForConfigOrDie(config)

	pc := &v2mushroomkingdom.PeachesCastle{}
	pc.Name = "marios-place"
	pc.Spec.Mushrooms = 7
	_, err := client.MushroomkingdomV2Client.PeachesCastles("default").Create(pc)
	if err != nil {
		panic(err)
	}

	result, err := client.MushroomkingdomV2Client.PeachesCastles("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, c := range result.Items {
		fmt.Printf("Found peaches castle %s %d\n", c.Name, c.Spec.Mushrooms)
	}

	t.Stop()
}

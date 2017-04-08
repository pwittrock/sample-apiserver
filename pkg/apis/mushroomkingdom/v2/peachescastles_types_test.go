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

package v2

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver-builder/pkg/test"

	"github.com/pwittrock/apiserver-helloworld/pkg/apis"
	mushroomkingdomv2 "github.com/pwittrock/apiserver-helloworld/pkg/apis/mushroomkingdom/v2"
	"github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/clientset"
	"github.com/pwittrock/apiserver-helloworld/pkg/openapi"
)

func TestPeachesCastles(t *testing.T) {
	// Start test environment
	testenv := test.NewTestEnvironment()
	config := testenv.Start(apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions)
	client := clientset.NewForConfigOrDie(config)

	pc := &mushroomkingdomv2.PeachesCastle{}
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

	if len(result.Items) != 1 {
		t.Fatalf("Expected to find 1 PeachesCastle, found %d", len(result.Items))
	}
	actual := result.Items[0]
	if actual.Name != "marios-place" {
		t.Fatalf("Expected to find %s, found %s", "marios-place", actual.Name)
	}
	if actual.Spec.Mushrooms != 7 {
		t.Fatalf("Expected to find %d, found %d", 7, actual.Spec.Mushrooms)
	}
	t.Logf("Found PeachesCastle %+v\n", actual)

	// Stop test environment
	testenv.Stop()
}

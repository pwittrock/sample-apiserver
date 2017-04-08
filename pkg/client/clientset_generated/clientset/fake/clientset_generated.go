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

package fake

import (
	clientset "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/clientset"
	hyrulekingdomv3 "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/clientset/typed/hyrulekingdom/v3"
	fakehyrulekingdomv3 "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/clientset/typed/hyrulekingdom/v3/fake"
	mushroomkingdomv2 "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/clientset/typed/mushroomkingdom/v2"
	fakemushroomkingdomv2 "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/clientset/typed/mushroomkingdom/v2/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(registry, scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o, registry.RESTMapper()))

	fakePtr.AddWatchReactor("*", testing.DefaultWatchReactor(watch.NewFake(), nil))

	return &Clientset{fakePtr}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return &fakediscovery.FakeDiscovery{Fake: &c.Fake}
}

var _ clientset.Interface = &Clientset{}

// HyrulekingdomV3 retrieves the HyrulekingdomV3Client
func (c *Clientset) HyrulekingdomV3() hyrulekingdomv3.HyrulekingdomV3Interface {
	return &fakehyrulekingdomv3.FakeHyrulekingdomV3{Fake: &c.Fake}
}

// Hyrulekingdom retrieves the HyrulekingdomV3Client
func (c *Clientset) Hyrulekingdom() hyrulekingdomv3.HyrulekingdomV3Interface {
	return &fakehyrulekingdomv3.FakeHyrulekingdomV3{Fake: &c.Fake}
}

// MushroomkingdomV2 retrieves the MushroomkingdomV2Client
func (c *Clientset) MushroomkingdomV2() mushroomkingdomv2.MushroomkingdomV2Interface {
	return &fakemushroomkingdomv2.FakeMushroomkingdomV2{Fake: &c.Fake}
}

// Mushroomkingdom retrieves the MushroomkingdomV2Client
func (c *Clientset) Mushroomkingdom() mushroomkingdomv2.MushroomkingdomV2Interface {
	return &fakemushroomkingdomv2.FakeMushroomkingdomV2{Fake: &c.Fake}
}

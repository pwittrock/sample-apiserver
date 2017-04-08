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
	internalversion "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/internalclientset/typed/mushroomkingdom/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMushroomkingdom struct {
	*testing.Fake
}

func (c *FakeMushroomkingdom) PeachesCastles(namespace string) internalversion.PeachesCastleInterface {
	return &FakePeachesCastles{c, namespace}
}

func (c *FakeMushroomkingdom) PeachesCastleSpecs(namespace string) internalversion.PeachesCastleSpecInterface {
	return &FakePeachesCastleSpecs{c, namespace}
}

func (c *FakeMushroomkingdom) PeachesCastleStatuses(namespace string) internalversion.PeachesCastleStatusInterface {
	return &FakePeachesCastleStatuses{c, namespace}
}

func (c *FakeMushroomkingdom) ScaleCastles(namespace string) internalversion.ScaleCastleInterface {
	return &FakeScaleCastles{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMushroomkingdom) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

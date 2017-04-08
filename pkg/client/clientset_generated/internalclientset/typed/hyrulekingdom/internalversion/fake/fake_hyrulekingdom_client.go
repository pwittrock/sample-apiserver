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
	internalversion "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/internalclientset/typed/hyrulekingdom/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHyrulekingdom struct {
	*testing.Fake
}

func (c *FakeHyrulekingdom) ZeldasCastles(namespace string) internalversion.ZeldasCastleInterface {
	return &FakeZeldasCastles{c, namespace}
}

func (c *FakeHyrulekingdom) ZeldasCastleSpecs(namespace string) internalversion.ZeldasCastleSpecInterface {
	return &FakeZeldasCastleSpecs{c, namespace}
}

func (c *FakeHyrulekingdom) ZeldasCastleStatuses(namespace string) internalversion.ZeldasCastleStatusInterface {
	return &FakeZeldasCastleStatuses{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHyrulekingdom) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

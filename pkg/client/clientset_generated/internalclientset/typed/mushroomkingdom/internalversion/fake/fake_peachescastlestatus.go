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
	mushroomkingdom "github.com/pwittrock/apiserver-helloworld/pkg/apis/mushroomkingdom"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePeachesCastleStatuses implements PeachesCastleStatusInterface
type FakePeachesCastleStatuses struct {
	Fake *FakeMushroomkingdom
	ns   string
}

var peachescastlestatusesResource = schema.GroupVersionResource{Group: "mushroomkingdom", Version: "", Resource: "peachescastlestatuses"}

func (c *FakePeachesCastleStatuses) Create(peachesCastleStatus *mushroomkingdom.PeachesCastleStatus) (result *mushroomkingdom.PeachesCastleStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(peachescastlestatusesResource, c.ns, peachesCastleStatus), &mushroomkingdom.PeachesCastleStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastleStatus), err
}

func (c *FakePeachesCastleStatuses) Update(peachesCastleStatus *mushroomkingdom.PeachesCastleStatus) (result *mushroomkingdom.PeachesCastleStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(peachescastlestatusesResource, c.ns, peachesCastleStatus), &mushroomkingdom.PeachesCastleStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastleStatus), err
}

func (c *FakePeachesCastleStatuses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(peachescastlestatusesResource, c.ns, name), &mushroomkingdom.PeachesCastleStatus{})

	return err
}

func (c *FakePeachesCastleStatuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(peachescastlestatusesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &mushroomkingdom.PeachesCastleStatusList{})
	return err
}

func (c *FakePeachesCastleStatuses) Get(name string, options v1.GetOptions) (result *mushroomkingdom.PeachesCastleStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(peachescastlestatusesResource, c.ns, name), &mushroomkingdom.PeachesCastleStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastleStatus), err
}

func (c *FakePeachesCastleStatuses) List(opts v1.ListOptions) (result *mushroomkingdom.PeachesCastleStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(peachescastlestatusesResource, c.ns, opts), &mushroomkingdom.PeachesCastleStatusList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastleStatusList), err
}

// Watch returns a watch.Interface that watches the requested peachesCastleStatuses.
func (c *FakePeachesCastleStatuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(peachescastlestatusesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched peachesCastleStatus.
func (c *FakePeachesCastleStatuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mushroomkingdom.PeachesCastleStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(peachescastlestatusesResource, c.ns, name, data, subresources...), &mushroomkingdom.PeachesCastleStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastleStatus), err
}

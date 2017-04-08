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
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePeachesCastles implements PeachesCastleInterface
type FakePeachesCastles struct {
	Fake *FakeMushroomkingdom
	ns   string
}

var peachescastlesResource = schema.GroupVersionResource{Group: "mushroomkingdom", Version: "", Resource: "peachescastles"}

func (c *FakePeachesCastles) Create(peachesCastle *mushroomkingdom.PeachesCastle) (result *mushroomkingdom.PeachesCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(peachescastlesResource, c.ns, peachesCastle), &mushroomkingdom.PeachesCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastle), err
}

func (c *FakePeachesCastles) Update(peachesCastle *mushroomkingdom.PeachesCastle) (result *mushroomkingdom.PeachesCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(peachescastlesResource, c.ns, peachesCastle), &mushroomkingdom.PeachesCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastle), err
}

func (c *FakePeachesCastles) UpdateStatus(peachesCastle *mushroomkingdom.PeachesCastle) (*mushroomkingdom.PeachesCastle, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(peachescastlesResource, "status", c.ns, peachesCastle), &mushroomkingdom.PeachesCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastle), err
}

func (c *FakePeachesCastles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(peachescastlesResource, c.ns, name), &mushroomkingdom.PeachesCastle{})

	return err
}

func (c *FakePeachesCastles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(peachescastlesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &mushroomkingdom.PeachesCastleList{})
	return err
}

func (c *FakePeachesCastles) Get(name string, options v1.GetOptions) (result *mushroomkingdom.PeachesCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(peachescastlesResource, c.ns, name), &mushroomkingdom.PeachesCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastle), err
}

func (c *FakePeachesCastles) List(opts v1.ListOptions) (result *mushroomkingdom.PeachesCastleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(peachescastlesResource, c.ns, opts), &mushroomkingdom.PeachesCastleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mushroomkingdom.PeachesCastleList{}
	for _, item := range obj.(*mushroomkingdom.PeachesCastleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested peachesCastles.
func (c *FakePeachesCastles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(peachescastlesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched peachesCastle.
func (c *FakePeachesCastles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mushroomkingdom.PeachesCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(peachescastlesResource, c.ns, name, data, subresources...), &mushroomkingdom.PeachesCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastle), err
}

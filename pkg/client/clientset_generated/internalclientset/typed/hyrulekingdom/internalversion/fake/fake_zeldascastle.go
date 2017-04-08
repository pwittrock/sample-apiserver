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
	hyrulekingdom "github.com/pwittrock/apiserver-helloworld/pkg/apis/hyrulekingdom"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeZeldasCastles implements ZeldasCastleInterface
type FakeZeldasCastles struct {
	Fake *FakeHyrulekingdom
	ns   string
}

var zeldascastlesResource = schema.GroupVersionResource{Group: "hyrulekingdom", Version: "", Resource: "zeldascastles"}

func (c *FakeZeldasCastles) Create(zeldasCastle *hyrulekingdom.ZeldasCastle) (result *hyrulekingdom.ZeldasCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(zeldascastlesResource, c.ns, zeldasCastle), &hyrulekingdom.ZeldasCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastle), err
}

func (c *FakeZeldasCastles) Update(zeldasCastle *hyrulekingdom.ZeldasCastle) (result *hyrulekingdom.ZeldasCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(zeldascastlesResource, c.ns, zeldasCastle), &hyrulekingdom.ZeldasCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastle), err
}

func (c *FakeZeldasCastles) UpdateStatus(zeldasCastle *hyrulekingdom.ZeldasCastle) (*hyrulekingdom.ZeldasCastle, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(zeldascastlesResource, "status", c.ns, zeldasCastle), &hyrulekingdom.ZeldasCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastle), err
}

func (c *FakeZeldasCastles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(zeldascastlesResource, c.ns, name), &hyrulekingdom.ZeldasCastle{})

	return err
}

func (c *FakeZeldasCastles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(zeldascastlesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &hyrulekingdom.ZeldasCastleList{})
	return err
}

func (c *FakeZeldasCastles) Get(name string, options v1.GetOptions) (result *hyrulekingdom.ZeldasCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(zeldascastlesResource, c.ns, name), &hyrulekingdom.ZeldasCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastle), err
}

func (c *FakeZeldasCastles) List(opts v1.ListOptions) (result *hyrulekingdom.ZeldasCastleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(zeldascastlesResource, c.ns, opts), &hyrulekingdom.ZeldasCastleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hyrulekingdom.ZeldasCastleList{}
	for _, item := range obj.(*hyrulekingdom.ZeldasCastleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested zeldasCastles.
func (c *FakeZeldasCastles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(zeldascastlesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched zeldasCastle.
func (c *FakeZeldasCastles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hyrulekingdom.ZeldasCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(zeldascastlesResource, c.ns, name, data, subresources...), &hyrulekingdom.ZeldasCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastle), err
}

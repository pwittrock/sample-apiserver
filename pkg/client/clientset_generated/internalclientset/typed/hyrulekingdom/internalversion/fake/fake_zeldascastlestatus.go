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
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeZeldasCastleStatuses implements ZeldasCastleStatusInterface
type FakeZeldasCastleStatuses struct {
	Fake *FakeHyrulekingdom
	ns   string
}

var zeldascastlestatusesResource = schema.GroupVersionResource{Group: "hyrulekingdom", Version: "", Resource: "zeldascastlestatuses"}

func (c *FakeZeldasCastleStatuses) Create(zeldasCastleStatus *hyrulekingdom.ZeldasCastleStatus) (result *hyrulekingdom.ZeldasCastleStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(zeldascastlestatusesResource, c.ns, zeldasCastleStatus), &hyrulekingdom.ZeldasCastleStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastleStatus), err
}

func (c *FakeZeldasCastleStatuses) Update(zeldasCastleStatus *hyrulekingdom.ZeldasCastleStatus) (result *hyrulekingdom.ZeldasCastleStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(zeldascastlestatusesResource, c.ns, zeldasCastleStatus), &hyrulekingdom.ZeldasCastleStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastleStatus), err
}

func (c *FakeZeldasCastleStatuses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(zeldascastlestatusesResource, c.ns, name), &hyrulekingdom.ZeldasCastleStatus{})

	return err
}

func (c *FakeZeldasCastleStatuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(zeldascastlestatusesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &hyrulekingdom.ZeldasCastleStatusList{})
	return err
}

func (c *FakeZeldasCastleStatuses) Get(name string, options v1.GetOptions) (result *hyrulekingdom.ZeldasCastleStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(zeldascastlestatusesResource, c.ns, name), &hyrulekingdom.ZeldasCastleStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastleStatus), err
}

func (c *FakeZeldasCastleStatuses) List(opts v1.ListOptions) (result *hyrulekingdom.ZeldasCastleStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(zeldascastlestatusesResource, c.ns, opts), &hyrulekingdom.ZeldasCastleStatusList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastleStatusList), err
}

// Watch returns a watch.Interface that watches the requested zeldasCastleStatuses.
func (c *FakeZeldasCastleStatuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(zeldascastlestatusesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched zeldasCastleStatus.
func (c *FakeZeldasCastleStatuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hyrulekingdom.ZeldasCastleStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(zeldascastlestatusesResource, c.ns, name, data, subresources...), &hyrulekingdom.ZeldasCastleStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastleStatus), err
}

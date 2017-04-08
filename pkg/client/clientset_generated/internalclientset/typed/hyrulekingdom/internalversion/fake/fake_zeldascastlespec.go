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

// FakeZeldasCastleSpecs implements ZeldasCastleSpecInterface
type FakeZeldasCastleSpecs struct {
	Fake *FakeHyrulekingdom
	ns   string
}

var zeldascastlespecsResource = schema.GroupVersionResource{Group: "hyrulekingdom", Version: "", Resource: "zeldascastlespecs"}

func (c *FakeZeldasCastleSpecs) Create(zeldasCastleSpec *hyrulekingdom.ZeldasCastleSpec) (result *hyrulekingdom.ZeldasCastleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(zeldascastlespecsResource, c.ns, zeldasCastleSpec), &hyrulekingdom.ZeldasCastleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastleSpec), err
}

func (c *FakeZeldasCastleSpecs) Update(zeldasCastleSpec *hyrulekingdom.ZeldasCastleSpec) (result *hyrulekingdom.ZeldasCastleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(zeldascastlespecsResource, c.ns, zeldasCastleSpec), &hyrulekingdom.ZeldasCastleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastleSpec), err
}

func (c *FakeZeldasCastleSpecs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(zeldascastlespecsResource, c.ns, name), &hyrulekingdom.ZeldasCastleSpec{})

	return err
}

func (c *FakeZeldasCastleSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(zeldascastlespecsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &hyrulekingdom.ZeldasCastleSpecList{})
	return err
}

func (c *FakeZeldasCastleSpecs) Get(name string, options v1.GetOptions) (result *hyrulekingdom.ZeldasCastleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(zeldascastlespecsResource, c.ns, name), &hyrulekingdom.ZeldasCastleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastleSpec), err
}

func (c *FakeZeldasCastleSpecs) List(opts v1.ListOptions) (result *hyrulekingdom.ZeldasCastleSpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(zeldascastlespecsResource, c.ns, opts), &hyrulekingdom.ZeldasCastleSpecList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastleSpecList), err
}

// Watch returns a watch.Interface that watches the requested zeldasCastleSpecs.
func (c *FakeZeldasCastleSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(zeldascastlespecsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched zeldasCastleSpec.
func (c *FakeZeldasCastleSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hyrulekingdom.ZeldasCastleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(zeldascastlespecsResource, c.ns, name, data, subresources...), &hyrulekingdom.ZeldasCastleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hyrulekingdom.ZeldasCastleSpec), err
}

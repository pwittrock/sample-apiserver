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

// FakePeachesCastleSpecs implements PeachesCastleSpecInterface
type FakePeachesCastleSpecs struct {
	Fake *FakeMushroomkingdom
	ns   string
}

var peachescastlespecsResource = schema.GroupVersionResource{Group: "mushroomkingdom", Version: "", Resource: "peachescastlespecs"}

func (c *FakePeachesCastleSpecs) Create(peachesCastleSpec *mushroomkingdom.PeachesCastleSpec) (result *mushroomkingdom.PeachesCastleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(peachescastlespecsResource, c.ns, peachesCastleSpec), &mushroomkingdom.PeachesCastleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastleSpec), err
}

func (c *FakePeachesCastleSpecs) Update(peachesCastleSpec *mushroomkingdom.PeachesCastleSpec) (result *mushroomkingdom.PeachesCastleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(peachescastlespecsResource, c.ns, peachesCastleSpec), &mushroomkingdom.PeachesCastleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastleSpec), err
}

func (c *FakePeachesCastleSpecs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(peachescastlespecsResource, c.ns, name), &mushroomkingdom.PeachesCastleSpec{})

	return err
}

func (c *FakePeachesCastleSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(peachescastlespecsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &mushroomkingdom.PeachesCastleSpecList{})
	return err
}

func (c *FakePeachesCastleSpecs) Get(name string, options v1.GetOptions) (result *mushroomkingdom.PeachesCastleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(peachescastlespecsResource, c.ns, name), &mushroomkingdom.PeachesCastleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastleSpec), err
}

func (c *FakePeachesCastleSpecs) List(opts v1.ListOptions) (result *mushroomkingdom.PeachesCastleSpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(peachescastlespecsResource, c.ns, opts), &mushroomkingdom.PeachesCastleSpecList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastleSpecList), err
}

// Watch returns a watch.Interface that watches the requested peachesCastleSpecs.
func (c *FakePeachesCastleSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(peachescastlespecsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched peachesCastleSpec.
func (c *FakePeachesCastleSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mushroomkingdom.PeachesCastleSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(peachescastlespecsResource, c.ns, name, data, subresources...), &mushroomkingdom.PeachesCastleSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.PeachesCastleSpec), err
}

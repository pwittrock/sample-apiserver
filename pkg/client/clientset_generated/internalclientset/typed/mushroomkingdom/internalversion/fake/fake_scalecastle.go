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

// FakeScaleCastles implements ScaleCastleInterface
type FakeScaleCastles struct {
	Fake *FakeMushroomkingdom
	ns   string
}

var scalecastlesResource = schema.GroupVersionResource{Group: "mushroomkingdom", Version: "", Resource: "scalecastles"}

func (c *FakeScaleCastles) Create(scaleCastle *mushroomkingdom.ScaleCastle) (result *mushroomkingdom.ScaleCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scalecastlesResource, c.ns, scaleCastle), &mushroomkingdom.ScaleCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.ScaleCastle), err
}

func (c *FakeScaleCastles) Update(scaleCastle *mushroomkingdom.ScaleCastle) (result *mushroomkingdom.ScaleCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scalecastlesResource, c.ns, scaleCastle), &mushroomkingdom.ScaleCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.ScaleCastle), err
}

func (c *FakeScaleCastles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(scalecastlesResource, c.ns, name), &mushroomkingdom.ScaleCastle{})

	return err
}

func (c *FakeScaleCastles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scalecastlesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &mushroomkingdom.ScaleCastleList{})
	return err
}

func (c *FakeScaleCastles) Get(name string, options v1.GetOptions) (result *mushroomkingdom.ScaleCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scalecastlesResource, c.ns, name), &mushroomkingdom.ScaleCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.ScaleCastle), err
}

func (c *FakeScaleCastles) List(opts v1.ListOptions) (result *mushroomkingdom.ScaleCastleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scalecastlesResource, c.ns, opts), &mushroomkingdom.ScaleCastleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mushroomkingdom.ScaleCastleList{}
	for _, item := range obj.(*mushroomkingdom.ScaleCastleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scaleCastles.
func (c *FakeScaleCastles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scalecastlesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched scaleCastle.
func (c *FakeScaleCastles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mushroomkingdom.ScaleCastle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scalecastlesResource, c.ns, name, data, subresources...), &mushroomkingdom.ScaleCastle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mushroomkingdom.ScaleCastle), err
}

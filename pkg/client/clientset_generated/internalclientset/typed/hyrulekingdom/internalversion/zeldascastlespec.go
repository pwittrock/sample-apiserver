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

package internalversion

import (
	hyrulekingdom "github.com/pwittrock/apiserver-helloworld/pkg/apis/hyrulekingdom"
	scheme "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ZeldasCastleSpecsGetter has a method to return a ZeldasCastleSpecInterface.
// A group's client should implement this interface.
type ZeldasCastleSpecsGetter interface {
	ZeldasCastleSpecs(namespace string) ZeldasCastleSpecInterface
}

// ZeldasCastleSpecInterface has methods to work with ZeldasCastleSpec resources.
type ZeldasCastleSpecInterface interface {
	Create(*hyrulekingdom.ZeldasCastleSpec) (*hyrulekingdom.ZeldasCastleSpec, error)
	Update(*hyrulekingdom.ZeldasCastleSpec) (*hyrulekingdom.ZeldasCastleSpec, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*hyrulekingdom.ZeldasCastleSpec, error)
	List(opts v1.ListOptions) (*hyrulekingdom.ZeldasCastleSpecList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hyrulekingdom.ZeldasCastleSpec, err error)
	ZeldasCastleSpecExpansion
}

// zeldasCastleSpecs implements ZeldasCastleSpecInterface
type zeldasCastleSpecs struct {
	client rest.Interface
	ns     string
}

// newZeldasCastleSpecs returns a ZeldasCastleSpecs
func newZeldasCastleSpecs(c *HyrulekingdomClient, namespace string) *zeldasCastleSpecs {
	return &zeldasCastleSpecs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a zeldasCastleSpec and creates it.  Returns the server's representation of the zeldasCastleSpec, and an error, if there is any.
func (c *zeldasCastleSpecs) Create(zeldasCastleSpec *hyrulekingdom.ZeldasCastleSpec) (result *hyrulekingdom.ZeldasCastleSpec, err error) {
	result = &hyrulekingdom.ZeldasCastleSpec{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("zeldascastlespecs").
		Body(zeldasCastleSpec).
		Do().
		Into(result)
	return
}

// Update takes the representation of a zeldasCastleSpec and updates it. Returns the server's representation of the zeldasCastleSpec, and an error, if there is any.
func (c *zeldasCastleSpecs) Update(zeldasCastleSpec *hyrulekingdom.ZeldasCastleSpec) (result *hyrulekingdom.ZeldasCastleSpec, err error) {
	result = &hyrulekingdom.ZeldasCastleSpec{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("zeldascastlespecs").
		Name(zeldasCastleSpec.Name).
		Body(zeldasCastleSpec).
		Do().
		Into(result)
	return
}

// Delete takes name of the zeldasCastleSpec and deletes it. Returns an error if one occurs.
func (c *zeldasCastleSpecs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("zeldascastlespecs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *zeldasCastleSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("zeldascastlespecs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the zeldasCastleSpec, and returns the corresponding zeldasCastleSpec object, and an error if there is any.
func (c *zeldasCastleSpecs) Get(name string, options v1.GetOptions) (result *hyrulekingdom.ZeldasCastleSpec, err error) {
	result = &hyrulekingdom.ZeldasCastleSpec{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("zeldascastlespecs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ZeldasCastleSpecs that match those selectors.
func (c *zeldasCastleSpecs) List(opts v1.ListOptions) (result *hyrulekingdom.ZeldasCastleSpecList, err error) {
	result = &hyrulekingdom.ZeldasCastleSpecList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("zeldascastlespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested zeldasCastleSpecs.
func (c *zeldasCastleSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("zeldascastlespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched zeldasCastleSpec.
func (c *zeldasCastleSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hyrulekingdom.ZeldasCastleSpec, err error) {
	result = &hyrulekingdom.ZeldasCastleSpec{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("zeldascastlespecs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

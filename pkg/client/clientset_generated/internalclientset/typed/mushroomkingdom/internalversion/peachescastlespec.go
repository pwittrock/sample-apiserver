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
	mushroomkingdom "github.com/pwittrock/apiserver-helloworld/pkg/apis/mushroomkingdom"
	scheme "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PeachesCastleSpecsGetter has a method to return a PeachesCastleSpecInterface.
// A group's client should implement this interface.
type PeachesCastleSpecsGetter interface {
	PeachesCastleSpecs(namespace string) PeachesCastleSpecInterface
}

// PeachesCastleSpecInterface has methods to work with PeachesCastleSpec resources.
type PeachesCastleSpecInterface interface {
	Create(*mushroomkingdom.PeachesCastleSpec) (*mushroomkingdom.PeachesCastleSpec, error)
	Update(*mushroomkingdom.PeachesCastleSpec) (*mushroomkingdom.PeachesCastleSpec, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*mushroomkingdom.PeachesCastleSpec, error)
	List(opts v1.ListOptions) (*mushroomkingdom.PeachesCastleSpecList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mushroomkingdom.PeachesCastleSpec, err error)
	PeachesCastleSpecExpansion
}

// peachesCastleSpecs implements PeachesCastleSpecInterface
type peachesCastleSpecs struct {
	client rest.Interface
	ns     string
}

// newPeachesCastleSpecs returns a PeachesCastleSpecs
func newPeachesCastleSpecs(c *MushroomkingdomClient, namespace string) *peachesCastleSpecs {
	return &peachesCastleSpecs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a peachesCastleSpec and creates it.  Returns the server's representation of the peachesCastleSpec, and an error, if there is any.
func (c *peachesCastleSpecs) Create(peachesCastleSpec *mushroomkingdom.PeachesCastleSpec) (result *mushroomkingdom.PeachesCastleSpec, err error) {
	result = &mushroomkingdom.PeachesCastleSpec{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("peachescastlespecs").
		Body(peachesCastleSpec).
		Do().
		Into(result)
	return
}

// Update takes the representation of a peachesCastleSpec and updates it. Returns the server's representation of the peachesCastleSpec, and an error, if there is any.
func (c *peachesCastleSpecs) Update(peachesCastleSpec *mushroomkingdom.PeachesCastleSpec) (result *mushroomkingdom.PeachesCastleSpec, err error) {
	result = &mushroomkingdom.PeachesCastleSpec{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("peachescastlespecs").
		Name(peachesCastleSpec.Name).
		Body(peachesCastleSpec).
		Do().
		Into(result)
	return
}

// Delete takes name of the peachesCastleSpec and deletes it. Returns an error if one occurs.
func (c *peachesCastleSpecs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("peachescastlespecs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *peachesCastleSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("peachescastlespecs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the peachesCastleSpec, and returns the corresponding peachesCastleSpec object, and an error if there is any.
func (c *peachesCastleSpecs) Get(name string, options v1.GetOptions) (result *mushroomkingdom.PeachesCastleSpec, err error) {
	result = &mushroomkingdom.PeachesCastleSpec{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("peachescastlespecs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PeachesCastleSpecs that match those selectors.
func (c *peachesCastleSpecs) List(opts v1.ListOptions) (result *mushroomkingdom.PeachesCastleSpecList, err error) {
	result = &mushroomkingdom.PeachesCastleSpecList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("peachescastlespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested peachesCastleSpecs.
func (c *peachesCastleSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("peachescastlespecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched peachesCastleSpec.
func (c *peachesCastleSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mushroomkingdom.PeachesCastleSpec, err error) {
	result = &mushroomkingdom.PeachesCastleSpec{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("peachescastlespecs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

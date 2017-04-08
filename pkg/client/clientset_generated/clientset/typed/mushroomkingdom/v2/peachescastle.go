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

package v2

import (
	v2 "github.com/pwittrock/apiserver-helloworld/pkg/apis/mushroomkingdom/v2"
	scheme "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PeachesCastlesGetter has a method to return a PeachesCastleInterface.
// A group's client should implement this interface.
type PeachesCastlesGetter interface {
	PeachesCastles(namespace string) PeachesCastleInterface
}

// PeachesCastleInterface has methods to work with PeachesCastle resources.
type PeachesCastleInterface interface {
	Create(*v2.PeachesCastle) (*v2.PeachesCastle, error)
	Update(*v2.PeachesCastle) (*v2.PeachesCastle, error)
	UpdateStatus(*v2.PeachesCastle) (*v2.PeachesCastle, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v2.PeachesCastle, error)
	List(opts v1.ListOptions) (*v2.PeachesCastleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.PeachesCastle, err error)
	PeachesCastleExpansion
}

// peachesCastles implements PeachesCastleInterface
type peachesCastles struct {
	client rest.Interface
	ns     string
}

// newPeachesCastles returns a PeachesCastles
func newPeachesCastles(c *MushroomkingdomV2Client, namespace string) *peachesCastles {
	return &peachesCastles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a peachesCastle and creates it.  Returns the server's representation of the peachesCastle, and an error, if there is any.
func (c *peachesCastles) Create(peachesCastle *v2.PeachesCastle) (result *v2.PeachesCastle, err error) {
	result = &v2.PeachesCastle{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("peachescastles").
		Body(peachesCastle).
		Do().
		Into(result)
	return
}

// Update takes the representation of a peachesCastle and updates it. Returns the server's representation of the peachesCastle, and an error, if there is any.
func (c *peachesCastles) Update(peachesCastle *v2.PeachesCastle) (result *v2.PeachesCastle, err error) {
	result = &v2.PeachesCastle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("peachescastles").
		Name(peachesCastle.Name).
		Body(peachesCastle).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *peachesCastles) UpdateStatus(peachesCastle *v2.PeachesCastle) (result *v2.PeachesCastle, err error) {
	result = &v2.PeachesCastle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("peachescastles").
		Name(peachesCastle.Name).
		SubResource("status").
		Body(peachesCastle).
		Do().
		Into(result)
	return
}

// Delete takes name of the peachesCastle and deletes it. Returns an error if one occurs.
func (c *peachesCastles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("peachescastles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *peachesCastles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("peachescastles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the peachesCastle, and returns the corresponding peachesCastle object, and an error if there is any.
func (c *peachesCastles) Get(name string, options v1.GetOptions) (result *v2.PeachesCastle, err error) {
	result = &v2.PeachesCastle{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("peachescastles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PeachesCastles that match those selectors.
func (c *peachesCastles) List(opts v1.ListOptions) (result *v2.PeachesCastleList, err error) {
	result = &v2.PeachesCastleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("peachescastles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested peachesCastles.
func (c *peachesCastles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("peachescastles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched peachesCastle.
func (c *peachesCastles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.PeachesCastle, err error) {
	result = &v2.PeachesCastle{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("peachescastles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

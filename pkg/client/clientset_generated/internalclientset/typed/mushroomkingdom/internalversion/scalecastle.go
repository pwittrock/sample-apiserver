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

// ScaleCastlesGetter has a method to return a ScaleCastleInterface.
// A group's client should implement this interface.
type ScaleCastlesGetter interface {
	ScaleCastles(namespace string) ScaleCastleInterface
}

// ScaleCastleInterface has methods to work with ScaleCastle resources.
type ScaleCastleInterface interface {
	Create(*mushroomkingdom.ScaleCastle) (*mushroomkingdom.ScaleCastle, error)
	Update(*mushroomkingdom.ScaleCastle) (*mushroomkingdom.ScaleCastle, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*mushroomkingdom.ScaleCastle, error)
	List(opts v1.ListOptions) (*mushroomkingdom.ScaleCastleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mushroomkingdom.ScaleCastle, err error)
	ScaleCastleExpansion
}

// scaleCastles implements ScaleCastleInterface
type scaleCastles struct {
	client rest.Interface
	ns     string
}

// newScaleCastles returns a ScaleCastles
func newScaleCastles(c *MushroomkingdomClient, namespace string) *scaleCastles {
	return &scaleCastles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a scaleCastle and creates it.  Returns the server's representation of the scaleCastle, and an error, if there is any.
func (c *scaleCastles) Create(scaleCastle *mushroomkingdom.ScaleCastle) (result *mushroomkingdom.ScaleCastle, err error) {
	result = &mushroomkingdom.ScaleCastle{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("scalecastles").
		Body(scaleCastle).
		Do().
		Into(result)
	return
}

// Update takes the representation of a scaleCastle and updates it. Returns the server's representation of the scaleCastle, and an error, if there is any.
func (c *scaleCastles) Update(scaleCastle *mushroomkingdom.ScaleCastle) (result *mushroomkingdom.ScaleCastle, err error) {
	result = &mushroomkingdom.ScaleCastle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scalecastles").
		Name(scaleCastle.Name).
		Body(scaleCastle).
		Do().
		Into(result)
	return
}

// Delete takes name of the scaleCastle and deletes it. Returns an error if one occurs.
func (c *scaleCastles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scalecastles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scaleCastles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scalecastles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the scaleCastle, and returns the corresponding scaleCastle object, and an error if there is any.
func (c *scaleCastles) Get(name string, options v1.GetOptions) (result *mushroomkingdom.ScaleCastle, err error) {
	result = &mushroomkingdom.ScaleCastle{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scalecastles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ScaleCastles that match those selectors.
func (c *scaleCastles) List(opts v1.ListOptions) (result *mushroomkingdom.ScaleCastleList, err error) {
	result = &mushroomkingdom.ScaleCastleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scalecastles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested scaleCastles.
func (c *scaleCastles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("scalecastles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched scaleCastle.
func (c *scaleCastles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mushroomkingdom.ScaleCastle, err error) {
	result = &mushroomkingdom.ScaleCastle{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("scalecastles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

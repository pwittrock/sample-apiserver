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

package v3

import (
	v3 "github.com/pwittrock/apiserver-helloworld/pkg/apis/hyrulekingdom/v3"
	scheme "github.com/pwittrock/apiserver-helloworld/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ZeldasCastlesGetter has a method to return a ZeldasCastleInterface.
// A group's client should implement this interface.
type ZeldasCastlesGetter interface {
	ZeldasCastles(namespace string) ZeldasCastleInterface
}

// ZeldasCastleInterface has methods to work with ZeldasCastle resources.
type ZeldasCastleInterface interface {
	Create(*v3.ZeldasCastle) (*v3.ZeldasCastle, error)
	Update(*v3.ZeldasCastle) (*v3.ZeldasCastle, error)
	UpdateStatus(*v3.ZeldasCastle) (*v3.ZeldasCastle, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v3.ZeldasCastle, error)
	List(opts v1.ListOptions) (*v3.ZeldasCastleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ZeldasCastle, err error)
	ZeldasCastleExpansion
}

// zeldasCastles implements ZeldasCastleInterface
type zeldasCastles struct {
	client rest.Interface
	ns     string
}

// newZeldasCastles returns a ZeldasCastles
func newZeldasCastles(c *HyrulekingdomV3Client, namespace string) *zeldasCastles {
	return &zeldasCastles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a zeldasCastle and creates it.  Returns the server's representation of the zeldasCastle, and an error, if there is any.
func (c *zeldasCastles) Create(zeldasCastle *v3.ZeldasCastle) (result *v3.ZeldasCastle, err error) {
	result = &v3.ZeldasCastle{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("zeldascastles").
		Body(zeldasCastle).
		Do().
		Into(result)
	return
}

// Update takes the representation of a zeldasCastle and updates it. Returns the server's representation of the zeldasCastle, and an error, if there is any.
func (c *zeldasCastles) Update(zeldasCastle *v3.ZeldasCastle) (result *v3.ZeldasCastle, err error) {
	result = &v3.ZeldasCastle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("zeldascastles").
		Name(zeldasCastle.Name).
		Body(zeldasCastle).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *zeldasCastles) UpdateStatus(zeldasCastle *v3.ZeldasCastle) (result *v3.ZeldasCastle, err error) {
	result = &v3.ZeldasCastle{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("zeldascastles").
		Name(zeldasCastle.Name).
		SubResource("status").
		Body(zeldasCastle).
		Do().
		Into(result)
	return
}

// Delete takes name of the zeldasCastle and deletes it. Returns an error if one occurs.
func (c *zeldasCastles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("zeldascastles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *zeldasCastles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("zeldascastles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the zeldasCastle, and returns the corresponding zeldasCastle object, and an error if there is any.
func (c *zeldasCastles) Get(name string, options v1.GetOptions) (result *v3.ZeldasCastle, err error) {
	result = &v3.ZeldasCastle{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("zeldascastles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ZeldasCastles that match those selectors.
func (c *zeldasCastles) List(opts v1.ListOptions) (result *v3.ZeldasCastleList, err error) {
	result = &v3.ZeldasCastleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("zeldascastles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested zeldasCastles.
func (c *zeldasCastles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("zeldascastles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched zeldasCastle.
func (c *zeldasCastles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ZeldasCastle, err error) {
	result = &v3.ZeldasCastle{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("zeldascastles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

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

// ZeldasCastleStatusesGetter has a method to return a ZeldasCastleStatusInterface.
// A group's client should implement this interface.
type ZeldasCastleStatusesGetter interface {
	ZeldasCastleStatuses(namespace string) ZeldasCastleStatusInterface
}

// ZeldasCastleStatusInterface has methods to work with ZeldasCastleStatus resources.
type ZeldasCastleStatusInterface interface {
	Create(*hyrulekingdom.ZeldasCastleStatus) (*hyrulekingdom.ZeldasCastleStatus, error)
	Update(*hyrulekingdom.ZeldasCastleStatus) (*hyrulekingdom.ZeldasCastleStatus, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*hyrulekingdom.ZeldasCastleStatus, error)
	List(opts v1.ListOptions) (*hyrulekingdom.ZeldasCastleStatusList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hyrulekingdom.ZeldasCastleStatus, err error)
	ZeldasCastleStatusExpansion
}

// zeldasCastleStatuses implements ZeldasCastleStatusInterface
type zeldasCastleStatuses struct {
	client rest.Interface
	ns     string
}

// newZeldasCastleStatuses returns a ZeldasCastleStatuses
func newZeldasCastleStatuses(c *HyrulekingdomClient, namespace string) *zeldasCastleStatuses {
	return &zeldasCastleStatuses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a zeldasCastleStatus and creates it.  Returns the server's representation of the zeldasCastleStatus, and an error, if there is any.
func (c *zeldasCastleStatuses) Create(zeldasCastleStatus *hyrulekingdom.ZeldasCastleStatus) (result *hyrulekingdom.ZeldasCastleStatus, err error) {
	result = &hyrulekingdom.ZeldasCastleStatus{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("zeldascastlestatuses").
		Body(zeldasCastleStatus).
		Do().
		Into(result)
	return
}

// Update takes the representation of a zeldasCastleStatus and updates it. Returns the server's representation of the zeldasCastleStatus, and an error, if there is any.
func (c *zeldasCastleStatuses) Update(zeldasCastleStatus *hyrulekingdom.ZeldasCastleStatus) (result *hyrulekingdom.ZeldasCastleStatus, err error) {
	result = &hyrulekingdom.ZeldasCastleStatus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("zeldascastlestatuses").
		Name(zeldasCastleStatus.Name).
		Body(zeldasCastleStatus).
		Do().
		Into(result)
	return
}

// Delete takes name of the zeldasCastleStatus and deletes it. Returns an error if one occurs.
func (c *zeldasCastleStatuses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("zeldascastlestatuses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *zeldasCastleStatuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("zeldascastlestatuses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the zeldasCastleStatus, and returns the corresponding zeldasCastleStatus object, and an error if there is any.
func (c *zeldasCastleStatuses) Get(name string, options v1.GetOptions) (result *hyrulekingdom.ZeldasCastleStatus, err error) {
	result = &hyrulekingdom.ZeldasCastleStatus{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("zeldascastlestatuses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ZeldasCastleStatuses that match those selectors.
func (c *zeldasCastleStatuses) List(opts v1.ListOptions) (result *hyrulekingdom.ZeldasCastleStatusList, err error) {
	result = &hyrulekingdom.ZeldasCastleStatusList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("zeldascastlestatuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested zeldasCastleStatuses.
func (c *zeldasCastleStatuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("zeldascastlestatuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched zeldasCastleStatus.
func (c *zeldasCastleStatuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hyrulekingdom.ZeldasCastleStatus, err error) {
	result = &hyrulekingdom.ZeldasCastleStatus{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("zeldascastlestatuses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

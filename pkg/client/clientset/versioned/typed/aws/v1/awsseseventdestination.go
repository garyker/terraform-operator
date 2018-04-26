/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	scheme "github.com/trussle/terraform-operator/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AwsSesEventDestinationsGetter has a method to return a AwsSesEventDestinationInterface.
// A group's client should implement this interface.
type AwsSesEventDestinationsGetter interface {
	AwsSesEventDestinations(namespace string) AwsSesEventDestinationInterface
}

// AwsSesEventDestinationInterface has methods to work with AwsSesEventDestination resources.
type AwsSesEventDestinationInterface interface {
	Create(*v1.AwsSesEventDestination) (*v1.AwsSesEventDestination, error)
	Update(*v1.AwsSesEventDestination) (*v1.AwsSesEventDestination, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsSesEventDestination, error)
	List(opts meta_v1.ListOptions) (*v1.AwsSesEventDestinationList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsSesEventDestination, err error)
	AwsSesEventDestinationExpansion
}

// awsSesEventDestinations implements AwsSesEventDestinationInterface
type awsSesEventDestinations struct {
	client rest.Interface
	ns     string
}

// newAwsSesEventDestinations returns a AwsSesEventDestinations
func newAwsSesEventDestinations(c *TrussleV1Client, namespace string) *awsSesEventDestinations {
	return &awsSesEventDestinations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsSesEventDestination, and returns the corresponding awsSesEventDestination object, and an error if there is any.
func (c *awsSesEventDestinations) Get(name string, options meta_v1.GetOptions) (result *v1.AwsSesEventDestination, err error) {
	result = &v1.AwsSesEventDestination{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsseseventdestinations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsSesEventDestinations that match those selectors.
func (c *awsSesEventDestinations) List(opts meta_v1.ListOptions) (result *v1.AwsSesEventDestinationList, err error) {
	result = &v1.AwsSesEventDestinationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsseseventdestinations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsSesEventDestinations.
func (c *awsSesEventDestinations) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsseseventdestinations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsSesEventDestination and creates it.  Returns the server's representation of the awsSesEventDestination, and an error, if there is any.
func (c *awsSesEventDestinations) Create(awsSesEventDestination *v1.AwsSesEventDestination) (result *v1.AwsSesEventDestination, err error) {
	result = &v1.AwsSesEventDestination{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsseseventdestinations").
		Body(awsSesEventDestination).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsSesEventDestination and updates it. Returns the server's representation of the awsSesEventDestination, and an error, if there is any.
func (c *awsSesEventDestinations) Update(awsSesEventDestination *v1.AwsSesEventDestination) (result *v1.AwsSesEventDestination, err error) {
	result = &v1.AwsSesEventDestination{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsseseventdestinations").
		Name(awsSesEventDestination.Name).
		Body(awsSesEventDestination).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsSesEventDestination and deletes it. Returns an error if one occurs.
func (c *awsSesEventDestinations) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsseseventdestinations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsSesEventDestinations) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsseseventdestinations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsSesEventDestination.
func (c *awsSesEventDestinations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsSesEventDestination, err error) {
	result = &v1.AwsSesEventDestination{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsseseventdestinations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
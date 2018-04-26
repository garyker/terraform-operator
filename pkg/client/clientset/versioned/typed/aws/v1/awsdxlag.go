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

// AwsDxLagsGetter has a method to return a AwsDxLagInterface.
// A group's client should implement this interface.
type AwsDxLagsGetter interface {
	AwsDxLags(namespace string) AwsDxLagInterface
}

// AwsDxLagInterface has methods to work with AwsDxLag resources.
type AwsDxLagInterface interface {
	Create(*v1.AwsDxLag) (*v1.AwsDxLag, error)
	Update(*v1.AwsDxLag) (*v1.AwsDxLag, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsDxLag, error)
	List(opts meta_v1.ListOptions) (*v1.AwsDxLagList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDxLag, err error)
	AwsDxLagExpansion
}

// awsDxLags implements AwsDxLagInterface
type awsDxLags struct {
	client rest.Interface
	ns     string
}

// newAwsDxLags returns a AwsDxLags
func newAwsDxLags(c *TrussleV1Client, namespace string) *awsDxLags {
	return &awsDxLags{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsDxLag, and returns the corresponding awsDxLag object, and an error if there is any.
func (c *awsDxLags) Get(name string, options meta_v1.GetOptions) (result *v1.AwsDxLag, err error) {
	result = &v1.AwsDxLag{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdxlags").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsDxLags that match those selectors.
func (c *awsDxLags) List(opts meta_v1.ListOptions) (result *v1.AwsDxLagList, err error) {
	result = &v1.AwsDxLagList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsdxlags").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsDxLags.
func (c *awsDxLags) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsdxlags").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsDxLag and creates it.  Returns the server's representation of the awsDxLag, and an error, if there is any.
func (c *awsDxLags) Create(awsDxLag *v1.AwsDxLag) (result *v1.AwsDxLag, err error) {
	result = &v1.AwsDxLag{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsdxlags").
		Body(awsDxLag).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsDxLag and updates it. Returns the server's representation of the awsDxLag, and an error, if there is any.
func (c *awsDxLags) Update(awsDxLag *v1.AwsDxLag) (result *v1.AwsDxLag, err error) {
	result = &v1.AwsDxLag{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsdxlags").
		Name(awsDxLag.Name).
		Body(awsDxLag).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsDxLag and deletes it. Returns an error if one occurs.
func (c *awsDxLags) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdxlags").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsDxLags) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsdxlags").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsDxLag.
func (c *awsDxLags) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsDxLag, err error) {
	result = &v1.AwsDxLag{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsdxlags").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
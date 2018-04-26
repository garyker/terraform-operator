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

// AwsLbsGetter has a method to return a AwsLbInterface.
// A group's client should implement this interface.
type AwsLbsGetter interface {
	AwsLbs(namespace string) AwsLbInterface
}

// AwsLbInterface has methods to work with AwsLb resources.
type AwsLbInterface interface {
	Create(*v1.AwsLb) (*v1.AwsLb, error)
	Update(*v1.AwsLb) (*v1.AwsLb, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsLb, error)
	List(opts meta_v1.ListOptions) (*v1.AwsLbList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsLb, err error)
	AwsLbExpansion
}

// awsLbs implements AwsLbInterface
type awsLbs struct {
	client rest.Interface
	ns     string
}

// newAwsLbs returns a AwsLbs
func newAwsLbs(c *TrussleV1Client, namespace string) *awsLbs {
	return &awsLbs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsLb, and returns the corresponding awsLb object, and an error if there is any.
func (c *awsLbs) Get(name string, options meta_v1.GetOptions) (result *v1.AwsLb, err error) {
	result = &v1.AwsLb{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslbs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsLbs that match those selectors.
func (c *awsLbs) List(opts meta_v1.ListOptions) (result *v1.AwsLbList, err error) {
	result = &v1.AwsLbList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awslbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsLbs.
func (c *awsLbs) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awslbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsLb and creates it.  Returns the server's representation of the awsLb, and an error, if there is any.
func (c *awsLbs) Create(awsLb *v1.AwsLb) (result *v1.AwsLb, err error) {
	result = &v1.AwsLb{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awslbs").
		Body(awsLb).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsLb and updates it. Returns the server's representation of the awsLb, and an error, if there is any.
func (c *awsLbs) Update(awsLb *v1.AwsLb) (result *v1.AwsLb, err error) {
	result = &v1.AwsLb{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awslbs").
		Name(awsLb.Name).
		Body(awsLb).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsLb and deletes it. Returns an error if one occurs.
func (c *awsLbs) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslbs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsLbs) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awslbs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsLb.
func (c *awsLbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsLb, err error) {
	result = &v1.AwsLb{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awslbs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
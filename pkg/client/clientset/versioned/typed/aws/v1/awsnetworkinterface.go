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

// AwsNetworkInterfacesGetter has a method to return a AwsNetworkInterfaceInterface.
// A group's client should implement this interface.
type AwsNetworkInterfacesGetter interface {
	AwsNetworkInterfaces(namespace string) AwsNetworkInterfaceInterface
}

// AwsNetworkInterfaceInterface has methods to work with AwsNetworkInterface resources.
type AwsNetworkInterfaceInterface interface {
	Create(*v1.AwsNetworkInterface) (*v1.AwsNetworkInterface, error)
	Update(*v1.AwsNetworkInterface) (*v1.AwsNetworkInterface, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.AwsNetworkInterface, error)
	List(opts meta_v1.ListOptions) (*v1.AwsNetworkInterfaceList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsNetworkInterface, err error)
	AwsNetworkInterfaceExpansion
}

// awsNetworkInterfaces implements AwsNetworkInterfaceInterface
type awsNetworkInterfaces struct {
	client rest.Interface
	ns     string
}

// newAwsNetworkInterfaces returns a AwsNetworkInterfaces
func newAwsNetworkInterfaces(c *TrussleV1Client, namespace string) *awsNetworkInterfaces {
	return &awsNetworkInterfaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the awsNetworkInterface, and returns the corresponding awsNetworkInterface object, and an error if there is any.
func (c *awsNetworkInterfaces) Get(name string, options meta_v1.GetOptions) (result *v1.AwsNetworkInterface, err error) {
	result = &v1.AwsNetworkInterface{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsnetworkinterfaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AwsNetworkInterfaces that match those selectors.
func (c *awsNetworkInterfaces) List(opts meta_v1.ListOptions) (result *v1.AwsNetworkInterfaceList, err error) {
	result = &v1.AwsNetworkInterfaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsnetworkinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested awsNetworkInterfaces.
func (c *awsNetworkInterfaces) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsnetworkinterfaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a awsNetworkInterface and creates it.  Returns the server's representation of the awsNetworkInterface, and an error, if there is any.
func (c *awsNetworkInterfaces) Create(awsNetworkInterface *v1.AwsNetworkInterface) (result *v1.AwsNetworkInterface, err error) {
	result = &v1.AwsNetworkInterface{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsnetworkinterfaces").
		Body(awsNetworkInterface).
		Do().
		Into(result)
	return
}

// Update takes the representation of a awsNetworkInterface and updates it. Returns the server's representation of the awsNetworkInterface, and an error, if there is any.
func (c *awsNetworkInterfaces) Update(awsNetworkInterface *v1.AwsNetworkInterface) (result *v1.AwsNetworkInterface, err error) {
	result = &v1.AwsNetworkInterface{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsnetworkinterfaces").
		Name(awsNetworkInterface.Name).
		Body(awsNetworkInterface).
		Do().
		Into(result)
	return
}

// Delete takes name of the awsNetworkInterface and deletes it. Returns an error if one occurs.
func (c *awsNetworkInterfaces) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsnetworkinterfaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *awsNetworkInterfaces) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsnetworkinterfaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched awsNetworkInterface.
func (c *awsNetworkInterfaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AwsNetworkInterface, err error) {
	result = &v1.AwsNetworkInterface{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsnetworkinterfaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
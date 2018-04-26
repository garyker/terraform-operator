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

package fake

import (
	aws_v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAwsCustomerGateways implements AwsCustomerGatewayInterface
type FakeAwsCustomerGateways struct {
	Fake *FakeTrussleV1
	ns   string
}

var awscustomergatewaysResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awscustomergateways"}

var awscustomergatewaysKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsCustomerGateway"}

// Get takes name of the awsCustomerGateway, and returns the corresponding awsCustomerGateway object, and an error if there is any.
func (c *FakeAwsCustomerGateways) Get(name string, options v1.GetOptions) (result *aws_v1.AwsCustomerGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscustomergatewaysResource, c.ns, name), &aws_v1.AwsCustomerGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCustomerGateway), err
}

// List takes label and field selectors, and returns the list of AwsCustomerGateways that match those selectors.
func (c *FakeAwsCustomerGateways) List(opts v1.ListOptions) (result *aws_v1.AwsCustomerGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscustomergatewaysResource, awscustomergatewaysKind, c.ns, opts), &aws_v1.AwsCustomerGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsCustomerGatewayList{}
	for _, item := range obj.(*aws_v1.AwsCustomerGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsCustomerGateways.
func (c *FakeAwsCustomerGateways) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscustomergatewaysResource, c.ns, opts))

}

// Create takes the representation of a awsCustomerGateway and creates it.  Returns the server's representation of the awsCustomerGateway, and an error, if there is any.
func (c *FakeAwsCustomerGateways) Create(awsCustomerGateway *aws_v1.AwsCustomerGateway) (result *aws_v1.AwsCustomerGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscustomergatewaysResource, c.ns, awsCustomerGateway), &aws_v1.AwsCustomerGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCustomerGateway), err
}

// Update takes the representation of a awsCustomerGateway and updates it. Returns the server's representation of the awsCustomerGateway, and an error, if there is any.
func (c *FakeAwsCustomerGateways) Update(awsCustomerGateway *aws_v1.AwsCustomerGateway) (result *aws_v1.AwsCustomerGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscustomergatewaysResource, c.ns, awsCustomerGateway), &aws_v1.AwsCustomerGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCustomerGateway), err
}

// Delete takes name of the awsCustomerGateway and deletes it. Returns an error if one occurs.
func (c *FakeAwsCustomerGateways) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscustomergatewaysResource, c.ns, name), &aws_v1.AwsCustomerGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsCustomerGateways) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscustomergatewaysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsCustomerGatewayList{})
	return err
}

// Patch applies the patch and returns the patched awsCustomerGateway.
func (c *FakeAwsCustomerGateways) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsCustomerGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscustomergatewaysResource, c.ns, name, data, subresources...), &aws_v1.AwsCustomerGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsCustomerGateway), err
}
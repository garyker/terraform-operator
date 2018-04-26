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

// FakeAwsLightsailStaticIps implements AwsLightsailStaticIpInterface
type FakeAwsLightsailStaticIps struct {
	Fake *FakeTrussleV1
	ns   string
}

var awslightsailstaticipsResource = schema.GroupVersionResource{Group: "trussle.com", Version: "v1", Resource: "awslightsailstaticips"}

var awslightsailstaticipsKind = schema.GroupVersionKind{Group: "trussle.com", Version: "v1", Kind: "AwsLightsailStaticIp"}

// Get takes name of the awsLightsailStaticIp, and returns the corresponding awsLightsailStaticIp object, and an error if there is any.
func (c *FakeAwsLightsailStaticIps) Get(name string, options v1.GetOptions) (result *aws_v1.AwsLightsailStaticIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awslightsailstaticipsResource, c.ns, name), &aws_v1.AwsLightsailStaticIp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLightsailStaticIp), err
}

// List takes label and field selectors, and returns the list of AwsLightsailStaticIps that match those selectors.
func (c *FakeAwsLightsailStaticIps) List(opts v1.ListOptions) (result *aws_v1.AwsLightsailStaticIpList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awslightsailstaticipsResource, awslightsailstaticipsKind, c.ns, opts), &aws_v1.AwsLightsailStaticIpList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aws_v1.AwsLightsailStaticIpList{}
	for _, item := range obj.(*aws_v1.AwsLightsailStaticIpList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsLightsailStaticIps.
func (c *FakeAwsLightsailStaticIps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awslightsailstaticipsResource, c.ns, opts))

}

// Create takes the representation of a awsLightsailStaticIp and creates it.  Returns the server's representation of the awsLightsailStaticIp, and an error, if there is any.
func (c *FakeAwsLightsailStaticIps) Create(awsLightsailStaticIp *aws_v1.AwsLightsailStaticIp) (result *aws_v1.AwsLightsailStaticIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awslightsailstaticipsResource, c.ns, awsLightsailStaticIp), &aws_v1.AwsLightsailStaticIp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLightsailStaticIp), err
}

// Update takes the representation of a awsLightsailStaticIp and updates it. Returns the server's representation of the awsLightsailStaticIp, and an error, if there is any.
func (c *FakeAwsLightsailStaticIps) Update(awsLightsailStaticIp *aws_v1.AwsLightsailStaticIp) (result *aws_v1.AwsLightsailStaticIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awslightsailstaticipsResource, c.ns, awsLightsailStaticIp), &aws_v1.AwsLightsailStaticIp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLightsailStaticIp), err
}

// Delete takes name of the awsLightsailStaticIp and deletes it. Returns an error if one occurs.
func (c *FakeAwsLightsailStaticIps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awslightsailstaticipsResource, c.ns, name), &aws_v1.AwsLightsailStaticIp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsLightsailStaticIps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awslightsailstaticipsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &aws_v1.AwsLightsailStaticIpList{})
	return err
}

// Patch applies the patch and returns the patched awsLightsailStaticIp.
func (c *FakeAwsLightsailStaticIps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *aws_v1.AwsLightsailStaticIp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awslightsailstaticipsResource, c.ns, name, data, subresources...), &aws_v1.AwsLightsailStaticIp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aws_v1.AwsLightsailStaticIp), err
}
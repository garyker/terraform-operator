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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/trussle/terraform-operator/pkg/apis/aws/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AwsNetworkInterfaceAttachmentLister helps list AwsNetworkInterfaceAttachments.
type AwsNetworkInterfaceAttachmentLister interface {
	// List lists all AwsNetworkInterfaceAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsNetworkInterfaceAttachment, err error)
	// AwsNetworkInterfaceAttachments returns an object that can list and get AwsNetworkInterfaceAttachments.
	AwsNetworkInterfaceAttachments(namespace string) AwsNetworkInterfaceAttachmentNamespaceLister
	AwsNetworkInterfaceAttachmentListerExpansion
}

// awsNetworkInterfaceAttachmentLister implements the AwsNetworkInterfaceAttachmentLister interface.
type awsNetworkInterfaceAttachmentLister struct {
	indexer cache.Indexer
}

// NewAwsNetworkInterfaceAttachmentLister returns a new AwsNetworkInterfaceAttachmentLister.
func NewAwsNetworkInterfaceAttachmentLister(indexer cache.Indexer) AwsNetworkInterfaceAttachmentLister {
	return &awsNetworkInterfaceAttachmentLister{indexer: indexer}
}

// List lists all AwsNetworkInterfaceAttachments in the indexer.
func (s *awsNetworkInterfaceAttachmentLister) List(selector labels.Selector) (ret []*v1.AwsNetworkInterfaceAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsNetworkInterfaceAttachment))
	})
	return ret, err
}

// AwsNetworkInterfaceAttachments returns an object that can list and get AwsNetworkInterfaceAttachments.
func (s *awsNetworkInterfaceAttachmentLister) AwsNetworkInterfaceAttachments(namespace string) AwsNetworkInterfaceAttachmentNamespaceLister {
	return awsNetworkInterfaceAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsNetworkInterfaceAttachmentNamespaceLister helps list and get AwsNetworkInterfaceAttachments.
type AwsNetworkInterfaceAttachmentNamespaceLister interface {
	// List lists all AwsNetworkInterfaceAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsNetworkInterfaceAttachment, err error)
	// Get retrieves the AwsNetworkInterfaceAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsNetworkInterfaceAttachment, error)
	AwsNetworkInterfaceAttachmentNamespaceListerExpansion
}

// awsNetworkInterfaceAttachmentNamespaceLister implements the AwsNetworkInterfaceAttachmentNamespaceLister
// interface.
type awsNetworkInterfaceAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsNetworkInterfaceAttachments in the indexer for a given namespace.
func (s awsNetworkInterfaceAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsNetworkInterfaceAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsNetworkInterfaceAttachment))
	})
	return ret, err
}

// Get retrieves the AwsNetworkInterfaceAttachment from the indexer for a given namespace and name.
func (s awsNetworkInterfaceAttachmentNamespaceLister) Get(name string) (*v1.AwsNetworkInterfaceAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsnetworkinterfaceattachment"), name)
	}
	return obj.(*v1.AwsNetworkInterfaceAttachment), nil
}
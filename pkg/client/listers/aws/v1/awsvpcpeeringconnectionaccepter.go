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

// AwsVpcPeeringConnectionAccepterLister helps list AwsVpcPeeringConnectionAccepters.
type AwsVpcPeeringConnectionAccepterLister interface {
	// List lists all AwsVpcPeeringConnectionAccepters in the indexer.
	List(selector labels.Selector) (ret []*v1.AwsVpcPeeringConnectionAccepter, err error)
	// AwsVpcPeeringConnectionAccepters returns an object that can list and get AwsVpcPeeringConnectionAccepters.
	AwsVpcPeeringConnectionAccepters(namespace string) AwsVpcPeeringConnectionAccepterNamespaceLister
	AwsVpcPeeringConnectionAccepterListerExpansion
}

// awsVpcPeeringConnectionAccepterLister implements the AwsVpcPeeringConnectionAccepterLister interface.
type awsVpcPeeringConnectionAccepterLister struct {
	indexer cache.Indexer
}

// NewAwsVpcPeeringConnectionAccepterLister returns a new AwsVpcPeeringConnectionAccepterLister.
func NewAwsVpcPeeringConnectionAccepterLister(indexer cache.Indexer) AwsVpcPeeringConnectionAccepterLister {
	return &awsVpcPeeringConnectionAccepterLister{indexer: indexer}
}

// List lists all AwsVpcPeeringConnectionAccepters in the indexer.
func (s *awsVpcPeeringConnectionAccepterLister) List(selector labels.Selector) (ret []*v1.AwsVpcPeeringConnectionAccepter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsVpcPeeringConnectionAccepter))
	})
	return ret, err
}

// AwsVpcPeeringConnectionAccepters returns an object that can list and get AwsVpcPeeringConnectionAccepters.
func (s *awsVpcPeeringConnectionAccepterLister) AwsVpcPeeringConnectionAccepters(namespace string) AwsVpcPeeringConnectionAccepterNamespaceLister {
	return awsVpcPeeringConnectionAccepterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AwsVpcPeeringConnectionAccepterNamespaceLister helps list and get AwsVpcPeeringConnectionAccepters.
type AwsVpcPeeringConnectionAccepterNamespaceLister interface {
	// List lists all AwsVpcPeeringConnectionAccepters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AwsVpcPeeringConnectionAccepter, err error)
	// Get retrieves the AwsVpcPeeringConnectionAccepter from the indexer for a given namespace and name.
	Get(name string) (*v1.AwsVpcPeeringConnectionAccepter, error)
	AwsVpcPeeringConnectionAccepterNamespaceListerExpansion
}

// awsVpcPeeringConnectionAccepterNamespaceLister implements the AwsVpcPeeringConnectionAccepterNamespaceLister
// interface.
type awsVpcPeeringConnectionAccepterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AwsVpcPeeringConnectionAccepters in the indexer for a given namespace.
func (s awsVpcPeeringConnectionAccepterNamespaceLister) List(selector labels.Selector) (ret []*v1.AwsVpcPeeringConnectionAccepter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AwsVpcPeeringConnectionAccepter))
	})
	return ret, err
}

// Get retrieves the AwsVpcPeeringConnectionAccepter from the indexer for a given namespace and name.
func (s awsVpcPeeringConnectionAccepterNamespaceLister) Get(name string) (*v1.AwsVpcPeeringConnectionAccepter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("awsvpcpeeringconnectionaccepter"), name)
	}
	return obj.(*v1.AwsVpcPeeringConnectionAccepter), nil
}
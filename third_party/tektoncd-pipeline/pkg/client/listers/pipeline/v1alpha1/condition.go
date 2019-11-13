// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/google/kf/third_party/tektoncd-pipeline/pkg/apis/pipeline/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConditionLister helps list Conditions.
type ConditionLister interface {
	// List lists all Conditions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Condition, err error)
	// Conditions returns an object that can list and get Conditions.
	Conditions(namespace string) ConditionNamespaceLister
	ConditionListerExpansion
}

// conditionLister implements the ConditionLister interface.
type conditionLister struct {
	indexer cache.Indexer
}

// NewConditionLister returns a new ConditionLister.
func NewConditionLister(indexer cache.Indexer) ConditionLister {
	return &conditionLister{indexer: indexer}
}

// List lists all Conditions in the indexer.
func (s *conditionLister) List(selector labels.Selector) (ret []*v1alpha1.Condition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Condition))
	})
	return ret, err
}

// Conditions returns an object that can list and get Conditions.
func (s *conditionLister) Conditions(namespace string) ConditionNamespaceLister {
	return conditionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConditionNamespaceLister helps list and get Conditions.
type ConditionNamespaceLister interface {
	// List lists all Conditions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Condition, err error)
	// Get retrieves the Condition from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Condition, error)
	ConditionNamespaceListerExpansion
}

// conditionNamespaceLister implements the ConditionNamespaceLister
// interface.
type conditionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Conditions in the indexer for a given namespace.
func (s conditionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Condition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Condition))
	})
	return ret, err
}

// Get retrieves the Condition from the indexer for a given namespace and name.
func (s conditionNamespaceLister) Get(name string) (*v1alpha1.Condition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("condition"), name)
	}
	return obj.(*v1alpha1.Condition), nil
}

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

package defaults

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type BasicStorage struct {
	NewListFunc  func() runtime.Object
	NewFunc      func() runtime.Object
	GroupVersion schema.GroupVersion
}

func (s BasicStorage) ObjectNameFunc(obj runtime.Object) (string, error) {
	return obj.(BasicResource).GetObjectMeta().Name, nil
}

func (s BasicStorage) Register(scheme *runtime.Scheme) error {
	if t := s.NewFunc(); t != nil {
		scheme.AddKnownTypes(s.GroupVersion, t)
	}
	if t := s.NewListFunc(); t != nil {
		scheme.AddKnownTypes(s.GroupVersion, t)
	}
	metav1.AddToGroupVersion(scheme, s.GroupVersion)
	return nil
}

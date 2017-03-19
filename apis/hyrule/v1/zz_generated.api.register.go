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

// This file was autogenerated by genwiring. Do not edit it manually!

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// Package-wide variables from generator "zz_generated.api.register".
	registerFn = func(scheme *runtime.Scheme) error {
		scheme.AddKnownTypes(SchemeGroupVersion, &HyruleCastle{}, &HyruleCastleList{})
		metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
		return nil
	}

	SchemeGroupVersion = schema.GroupVersion{Group, Version}

	SchemeBuilder = runtime.NewSchemeBuilder(registerFn)

	AddToScheme = SchemeBuilder.AddToScheme
)

const (
	// Package-wide consts from generator "zz_generated.api.register".
	Group   = "hyrule.k8s.io"
	Version = "v1"
)

type HyruleCastleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []HyruleCastle `json:"items" protobuf:"bytes,2,rep,name=items"`
}

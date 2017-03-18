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

package v2

import (
	"github.com/pwittrock/apiserver-helloworld/pkg/apis/mushroomkingdom"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver-builder/pkg/builders"
)

var (
	ApiVersion = builders.NewVersionedApiBuilder("mushroomkingdom.k8s.io", "v2").WithResources(
		builders.NewVersionedResourceWithStorage( //  Resource endpoint
			mushroomkingdom.PeachesCastleSingleton,
			func() runtime.Object { return &PeachesCastle{} },     // Register versioned resource
			func() runtime.Object { return &PeachesCastleList{} }, // Register versioned resource list
			&mushroomkingdom.PeachesCastleStrategy{builders.StorageStrategySingleton},
		),
		builders.NewVersionedResourceWithStorage( // Resource status endpoint
			mushroomkingdom.PeachesCastleStatusSingleton,
			func() runtime.Object { return &PeachesCastle{} },     // Register versioned resource
			func() runtime.Object { return &PeachesCastleList{} }, // Register versioned resource list
			&mushroomkingdom.PeachesCastleStatusStrategy{builders.StatusStorageStrategySingleton},
		),
		builders.NewVersionedResourceWithoutStorage(
			mushroomkingdom.ScalePeachesCastleRESTSingleton,
			func() runtime.Object { return &ScaleCastle{} }, // Register versioned resource
			&mushroomkingdom.ScalePeachesCastleREST{},
		),
	)

	// Expected by generated deepcopy and conversion
	SchemeBuilder = ApiVersion.SchemaBuilder
)

type PeachesCastleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PeachesCastle `json:"items"`
}

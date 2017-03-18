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

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient=true
// +k8s:openapi-gen=true
// +resource=peachescastles
// +subresource=peachescastles/scalecastle,PeachesCastle,ScaleCastle,ScalePeachesCastleREST
type PeachesCastle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PeachesCastleSpec   `json:"spec,omitempty"`
	Status PeachesCastleStatus `json:"status,omitempty"`
}

type PeachesCastleSpec struct {
	Mushrooms int `json:"mushrooms,omitempty"`
}

type PeachesCastleStatus struct {
	Message string `json:"message,omitempty"`
}

// +subresource-request
type ScaleCastle struct {
	metav1.TypeMeta
	// Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.
	// +optional
	metav1.ObjectMeta
}

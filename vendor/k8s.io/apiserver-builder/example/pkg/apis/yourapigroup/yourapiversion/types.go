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

package yourapiversion

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Generating code from this types.go file will generate storage and status REST endpoints for
// YourResource.

// ACTION REQUIRED: replace the YourResource with the name of the resource you are creating
// ACTION REQUIRED: replace <yourresources> with the lower case pluralized name of your resource.
// ACTION REQUIRED: add / remove fields to the Spec and Status structs for your resource

// +genclient=true
// +k8s:openapi-gen=true
// +resource=<yourresources>
type YourResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   YourResourceSpec   `json:"spec,omitempty"`
	Status YourResourceStatus `json:"status,omitempty"`
}

// YourResourceSpec defines the desired state of YourResource
type YourResourceSpec struct {
	// Add your Spec fields here.  Spec fields define the desired state of a Resource.
	// The field comments will be used to generate the openapi-spec and reference
	// documentation for the field

	// specfield defines some desired state for the resource
	SpecField int `json:"specfield,omitempty"`
}

// YourResourceStatus defines the observed state of YourResource
type YourResourceStatus struct {
	// Add your Status fields here. Status fields are purely for recording status about a Resource
	// and should not be the source of truth for the data

	// statusfield provides status information about YourResource
	StatusField string `json:"statusfield,omitempty"`
}

// Do define a Subresource add the following comment to YourResource
//// +subresource=yourresources/subresourcename,YourResource,SubYourResource,SubYourResourceREST
// - "yourresources/subresourcename" is the path to your subresource
// - YourResource is the resource this is a subresource of
// - SubYourResource is the request type of your subresource
// - SubYourResourceREST is the REST implementation of your subresource (defined in the group)
// Then uncomment the subresource request type and keep the +subresource-request comment tag
//// +subresource-request
//type SubResource struct {
//	metav1.TypeMeta
//	// Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.
//	// +optional
//	metav1.ObjectMeta
//}

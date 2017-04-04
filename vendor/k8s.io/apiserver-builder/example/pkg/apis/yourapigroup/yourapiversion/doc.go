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

// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// ACTION REQUIRED: replace the yourapiversion package with the api version for your api.
// e.g. v1alpha1, v1alpha2, v1beta1, v1, etc
// ACTION REQUIRED: replace <apigroup-go-import> with the go import path of your apigroup
// e.g. github.com/org/repo/pkg/apis/yourapigroup

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=<apigroup-go-import>

package yourapiversion

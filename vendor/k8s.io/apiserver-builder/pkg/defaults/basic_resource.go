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
)

type BasicResource interface {
	HasObjectMeta
	HasSpec
	HasStatus
	HasGeneration
}

type HasStatus interface {
	NewStatus() runtime.Object
	GetStatus() runtime.Object
	SetStatus(status runtime.Object)
}

type HasSpec interface {
	GetSpec() runtime.Object
	SetSpec(spec runtime.Object)
}

type HasObjectMeta interface {
	GetObjectMeta() *metav1.ObjectMeta
}

type HasGeneration interface {
	SetGeneration(generation int)
	GetGeneration() int
}

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

package yourapigroup

// Uncomment this package and the corresponding type in yourapiversion/types.go to implement a subresource.
// Subresources are request endpoints that live under some resource - e.g. "yourresource/scale"
//import (
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/runtime"
//	"k8s.io/apiserver/pkg/endpoints/request"
//	"k8s.io/apiserver/pkg/registry/rest"
//)
//
//var _ rest.CreaterUpdater = &SubYourResourceREST{}
//var _ rest.Patcher = &SubYourResourceREST{}
//
//type SubYourResourceREST struct {
//}
//
//func (r *SubYourResourceREST) Create(ctx request.Context, obj runtime.Object) (runtime.Object, error) {
//	return nil, nil
//}
//
//// Get retrieves the object from the storage. It is required to support Patch.
//func (r *SubYourResourceREST) Get(ctx request.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
//	return nil, nil
//}
//
//// Update alters the status subset of an object.
//func (r *SubYourResourceREST) Update(ctx request.Context, name string, objInfo rest.UpdatedObjectInfo) (runtime.Object, bool, error) {
//	return nil, false, nil
//}
//
//func (r *SubYourResourceREST) New() runtime.Object {
//	return &SubYourResource{}
//}

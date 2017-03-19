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

package apis

//import (
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/runtime"
//	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
//	"k8s.io/apiserver/pkg/registry/rest"
//)

////+sub-resource=hyrulecastles/print
//type PrintHyruleCastleImpl struct {
//
//}
//
//type PrintHyruleCastle struct {
//	metav1.TypeMeta
//	// +optional
//	metav1.ObjectMeta
//
//	kind string
//}
//
//func (r *PrintHyruleCastleImpl) New() runtime.Object {
//	return &PrintHyruleCastle{}
//}
//
//// Get finds a resource in the storage by name and returns it.
//// Although it can return an arbitrary error value, IsNotFound(err) is true for the
//// returned error value err when the specified resource is not found.
//func (r *PrintHyruleCastleImpl) Get(ctx genericapirequest.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
//	return nil, nil
//}
//
//// Update finds a resource in the storage and updates it. Some implementations
//// may allow updates creates the object - they should set the created boolean
//// to true.
//func (r *PrintHyruleCastleImpl) Update(ctx genericapirequest.Context, name string, objInfo rest.UpdatedObjectInfo) (runtime.Object, bool, error) {
//	return nil, true, nil
//}

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

// IMPORTANT: zz_generated.* files must be regenerated after uncommenting this!
// Uncomment this to override the Flunder Strategy
// used to create, update, delete new Flunder instances
// Validation and defaulting may be added
// Sub-resource api endpoints maybe added (e.g. scale)
type FlunderStrategy struct {
	DefaultFlunderStrategy
}

// Uncomment this to make Flunder a global (unnamedspaced) resource (e.g. like Node)
//func (s *FlunderStrategy) NamespaceScoped() bool {
//	return false
//}

// Uncomment this to add validation logic on create
//func (s *FlunderStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
//	return nil
//}

// Uncomment this to add validation logic on update
//func (s *FlunderStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
//
//}

// Uncomment this to add defaulting or other mutation logic before the Flunder object is persisted
//func (s *FlunderStrategy) Canonicalize(obj runtime.Object) {
//
//}

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

package generators

type VersionKindTemplateArgs struct {
	Version      string
	Group        string
	Kind         string
	LowerKind    string
	Name         string
	Resource     string
	SubResources map[string]SubResource
}

const versionKindTemplateName = "VersionKindTemplate"
const versionKindTemplateString = `
// Definition used to register {{.Name}} with the apiserver
var {{.Name}}ApiDefinition = &defaults.ResourceDefinition{
	{{.Version}}.SchemeGroupVersion.WithResource("{{.Resource}}"),
	{{.Group}}.{{.Kind}}StrategySingleton,
	{{.Group}}.{{.Kind}}StrategySingleton,
	{{.Group}}.{{.Kind}}StrategySingleton,
	{{.Group}}.{{.Kind}}StrategySingleton,
	map[string]*defaults.ResourceDefinition{
		"{{.Resource}}/status": {{.Name}}StatusApiDefinition,
	},
	{{.Group}}.{{.Kind}}StrategySingleton.BasicMatch,
	func(store *genericregistry.Store) rest.Storage { return &{{.Group}}.{{.Kind}}Store{store} },
}

var {{.Name}}StatusApiDefinition = &defaults.ResourceDefinition{
	{{.Version}}.SchemeGroupVersion.WithResource("{{.Resource}}"),
	{{.Group}}.{{.Kind}}StatusStrategySingleton,
	{{.Group}}.{{.Kind}}StatusStrategySingleton,
	{{.Group}}.{{.Kind}}StatusStrategySingleton,
	{{.Group}}.{{.Kind}}StatusStrategySingleton,
	map[string]*defaults.ResourceDefinition{},
	{{.Group}}.{{.Kind}}StatusStrategySingleton.BasicMatch,
	func(store *genericregistry.Store) rest.Storage { return &{{.Group}}.{{.Kind}}StatusStore{store} },
}

`

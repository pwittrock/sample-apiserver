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

type GroupTemplateArgs struct {
	GroupTitle          string
	Group               string
	Domain              string
	VersionOrder        []string
	ImportPrefix        string
	ResourceDefinitions []string
}

const groupTemplateName = "ApiTemplate"
const groupTemplateString = `
// Order list of version preferences
var {{.Group}}VersionPreferenceOrder = []string{
	{{- range $v := .VersionOrder}}
	{{$v}}.SchemeGroupVersion.Version,
	{{- end -}}
}
var {{.Group}}VersionToSchemeFunc = announced.VersionToSchemeFunc{
	{{- range $v := .VersionOrder}}
	{{$v}}.SchemeGroupVersion.Version: {{$v}}.AddToScheme,
	{{- end -}}
}

type {{.Group}}Provider struct{}

func Get{{.GroupTitle}}Provider() defaults.ResourceDefinitionProvider {
	return &{{.Group}}Provider{}
}

func (w *{{.Group}}Provider) GetResourceDefinitions() []*defaults.ResourceDefinition {
	return []*defaults.ResourceDefinition{
		{{- range $d := .ResourceDefinitions}}
		{{$d}}ApiDefinition,
		{{- end -}}
	}
}

func (w *{{.Group}}Provider) GetLegacyCodec() []schema.GroupVersion {
	return []schema.GroupVersion{
		{{- range $v := .VersionOrder}}
		{{$v}}.SchemeGroupVersion,
		{{- end -}}
	}
}

func (w *{{.Group}}Provider) GetGroupName() string {
	return "{{.Group}}.{{.Domain}}"
}

func (w *{{.Group}}Provider) GetVersionPreferenceOrder() []string {
	return {{.Group}}VersionPreferenceOrder
}

func (w *{{.Group}}Provider) GetImportPrefix() string {
	return "{{.ImportPrefix}}"
}

func (w *{{.Group}}Provider) SchemeFunc() announced.SchemeFunc {
	return {{.Group}}.AddToScheme
}

func (w *{{.Group}}Provider) VersionToSchemeFunc() announced.VersionToSchemeFunc {
	return {{.Group}}VersionToSchemeFunc
}
`

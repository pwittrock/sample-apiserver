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

import (
	"text/template"
)

type templates struct {
	groupTemplate        *template.Template
	kindTemplate         *template.Template
	versionKindTemplate  *template.Template
	subresourceTemplate  *template.Template
	AllProvidersTemplate *template.Template
}

var Templates = GetTemplates()

func GetTemplates() *templates {
	groupTemplate := template.Must(template.New(groupTemplateName).Parse(groupTemplateString))
	kindTemplate := template.Must(template.New(kindTemplateName).Parse(kindTemplateString))
	versionKindTemplate := template.Must(template.New(versionKindTemplateName).Parse(versionKindTemplateString))
	subresourceTemplate := template.Must(template.New(subresourceTemplateName).Parse(subresourceTemplateString))
	allProvidersTemplate := template.Must(template.New(AllProvidersTemplateName).Parse(AllProvidersTemplateString))

	return &templates{
		groupTemplate,
		kindTemplate,
		versionKindTemplate,
		subresourceTemplate,
		allProvidersTemplate,
	}
}

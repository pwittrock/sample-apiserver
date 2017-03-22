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
	"github.com/pkg/errors"
	"html/template"
)

type templates struct {
	groupTemplate       *template.Template
	kindTemplate        *template.Template
	versionKindTemplate *template.Template
	subresourceTemplate *template.Template
}

var Templates = GetTemplates()

func GetTemplates() *templates {
	groupTemplate, err := template.New(groupTemplateName).Parse(groupTemplateString)
	if err != nil {
		panic(errors.Errorf("Could not parse %v %s", err, groupTemplateString))
	}
	kindTemplate, err := template.New(kindTemplateName).Parse(kindTemplateString)
	if err != nil {
		panic(errors.Errorf("Could not parse %v %s", err, kindTemplateString))
	}
	versionKindTemplate, err := template.New(versionKindTemplateName).Parse(versionKindTemplateString)
	if err != nil {
		panic(errors.Errorf("Could not parse %v %s", err, versionKindTemplateString))
	}
	subresourceTemplate, err := template.New(subresourceTemplateName).Parse(subresourceTemplateString)
	if err != nil {
		panic(errors.Errorf("Could not parse %v %s", err, subresourceTemplateString))
	}

	return &templates{
		groupTemplate,
		kindTemplate,
		versionKindTemplate,
		subresourceTemplate,
	}
}

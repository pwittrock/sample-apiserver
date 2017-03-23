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
	"io"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

type mainGenerator struct {
	generator.DefaultGen
	pkg *types.Package
}

var _ generator.Generator = &versionedGenerator{}

var ApiProviders = sets.String{}

func CreateAllProvidersGenerator(
	c *generator.Context, pkg *types.Package) generator.Generator {
	return &mainGenerator{generator.DefaultGen{OptionalName: "zz_generated.api.register"}, pkg}
}

func (d *mainGenerator) Finalize(context *generator.Context, w io.Writer) error {
	Templates.AllProvidersTemplate.Execute(w, AllProvidersTemplateArgs{ApiProviders.List()})
	return nil
}

func (d *mainGenerator) Namers(c *generator.Context) namer.NameSystems {
	return nil
}
func (d *mainGenerator) Imports(c *generator.Context) []string {
	return []string{"k8s.io/apiserver-builder/pkg/defaults"}
}

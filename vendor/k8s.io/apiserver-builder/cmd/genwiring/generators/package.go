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
	"path/filepath"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"

	"github.com/pkg/errors"
)

// CustomArgs is used tby the go2idl framework to pass args specific to this
// generator.
type CustomArgs struct{}

type Gen struct {
	p []generator.Package
}

func (g *Gen) Execute(arguments *args.GeneratorArgs) error {
	return arguments.Execute(
		g.NameSystems(),
		g.DefaultNameSystem(),
		g.Packages)
}

// DefaultNameSystem returns the default name system for ordering the types to be
// processed by the generators in this package.
func (g *Gen) DefaultNameSystem() string {
	return "public"
}

// NameSystems returns the name system used by the generators in this package.
func (g *Gen) NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public": namer.NewPublicNamer(1),
		"raw":    namer.NewRawNamer("", nil),
	}
}

func (g *Gen) ParsePackages(context *generator.Context, arguments *args.GeneratorArgs) (sets.String, sets.String, string, string) {
	versionedPkgs := sets.NewString()
	unversionedPkgs := sets.NewString()
	mainPkg := ""
	apisPkg := ""
	for _, o := range context.Order {
		if IsApiType(o) {
			versioned := o.Name.Package
			versionedPkgs.Insert(versioned)
			unversioned := filepath.Dir(versioned)
			unversionedPkgs.Insert(unversioned)

			if apis := filepath.Dir(unversioned); apis != apisPkg && len(apisPkg) > 0 {
				panic(errors.Errorf(
					"Found multiple apis directory paths: %v and %v", apisPkg, apis))
			} else {
				apisPkg = apis
				mainPkg = filepath.Dir(apisPkg)
			}
		}
	}
	return versionedPkgs, unversionedPkgs, apisPkg, mainPkg
}

func (g *Gen) Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	g.p = generator.Packages{}

	// Do the versioned packages
	versionedPkgs, unversionedPkgs, apisPkg, _ := g.ParsePackages(context, arguments)
	pkg := context.Universe[apisPkg]
	if pkg == nil {
		// If the input had no Go files, for example.
		panic(errors.Errorf("Missing apis package."))
	}
	comments := Comments(pkg.Comments)
	domain := comments.GetTag("domain")

	groups := []string{}

	for p := range versionedPkgs {
		//glog.Infof("Considering versioned pkg %q", p)
		pkg := context.Universe[p]
		if pkg == nil {
			// If the input had no Go files, for example.
			continue
		}
		factory := &packageFactory{pkg, arguments}
		version := filepath.Base(p)
		group := filepath.Base(filepath.Dir(p))
		gen := CreateVersionedGenerator(context, pkg, arguments, group, version, domain)
		g.p = append(g.p, factory.createPackage(gen))

		groups = append(groups, group)
	}

	// Do the unversioned packages
	for p := range unversionedPkgs {
		//glog.Infof("Considering unversioned pkg %q", p)
		pkg := context.Universe[p]
		if pkg == nil {
			// If the input had no Go files, for example.
			continue
		}
		factory := &packageFactory{pkg, arguments}
		group := filepath.Base(p)
		g.p = append(g.p, factory.createPackage(CreateUnversionedGenerator(
			context, pkg, arguments, group, domain)))
	}

	// Do the base Api package
	pkg = context.Universe[apisPkg]
	factory := &packageFactory{pkg, arguments}
	if pkg != nil {
		for _, group := range groups {
			g.p = append(g.p, factory.createPackage(CreateApisGenerator(
				context, pkg, group, domain)))
		}
		// Run this after the individual packages so the list of providers is populated
		g.p = append(g.p, factory.createPackage(CreateAllProvidersGenerator(context, pkg)))
	}

	return g.p
}

type packageFactory struct {
	pkg       *types.Package
	arguments *args.GeneratorArgs
}

// Creates a package with a generator
func (f *packageFactory) createPackage(gen generator.Generator) generator.Package {
	path := f.pkg.Path
	name := strings.Split(filepath.Base(f.pkg.Path), ".")[0]
	return &generator.DefaultPackage{
		PackageName: name,
		PackagePath: path,
		HeaderText:  f.getHeader(),
		GeneratorFunc: func(c *generator.Context) (generators []generator.Generator) {
			return []generator.Generator{gen}
		},
		FilterFunc: func(c *generator.Context, t *types.Type) bool {
			return t.Name.Package == f.pkg.Path
		},
	}
}

// Returns the header for generated files
func (f *packageFactory) getHeader() []byte {
	header := []byte(`/*
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

// This file was autogenerated by genwiring. Do not edit it manually!

`)
	return header
}

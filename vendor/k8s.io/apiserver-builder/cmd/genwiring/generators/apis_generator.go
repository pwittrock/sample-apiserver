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
	"fmt"
	"io"
	"path/filepath"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
	"sort"
	"strings"
)

type apisGenerator struct {
	generator.DefaultGen
	pkg    *types.Package
	group  string
	domain string

	imports sets.String
	vars    []string
	consts  []string

	typesByVersionKind  map[string]map[string]*types.Type
	typesByKindVersion  map[string]map[string]*types.Type
	unversionedApiTypes sets.String
	subresources        map[string]SubResource
}

var _ generator.Generator = &apisGenerator{}

func CreateApisGenerator(context *generator.Context, pkg *types.Package, group, domain string) generator.Generator {
	subresources := GetSubresources(context, group)

	typesByVersionKind, typesByKindVersion, unversionedApiTypes := GetIndexedTypes(context, group)

	imports := sets.NewString(
		//"reflect",
		//"fmt",
		//"k8s.io/client-go/pkg/api",
		"k8s.io/apiserver/pkg/registry/rest",
		"k8s.io/apimachinery/pkg/apimachinery/announced",
		//"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apimachinery/pkg/runtime/schema",
		"k8s.io/apiserver-builder/pkg/defaults",
		//"genericapirequest \"k8s.io/apiserver/pkg/endpoints/request\"",
		"genericregistry \"k8s.io/apiserver/pkg/registry/generic/registry\"",
		//"metainternalversion \"k8s.io/apimachinery/pkg/apis/meta/internalversion\"",
		//"metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"",
	)

	return &apisGenerator{
		generator.DefaultGen{OptionalName: fmt.Sprintf("zz_generated.api.%s", group)},
		pkg,
		group,
		domain,
		imports,
		[]string{},
		[]string{},
		typesByVersionKind,
		typesByKindVersion,
		unversionedApiTypes,
		subresources,
	}
}

func (d *apisGenerator) Imports(c *generator.Context) []string {
	return d.imports.List()
}

func (d *apisGenerator) PackageVars(c *generator.Context) []string {
	return d.vars
}

func (d *apisGenerator) PackageConsts(c *generator.Context) []string {
	return d.consts
}

type SubResource struct {
	Kind        string
	RequestKind string
	Path        string
	REST        string
}

func (d *apisGenerator) Finalize(context *generator.Context, w io.Writer) error {
	prefix := ""
	for _, k := range d.GetListOfKinds() {
		//// For each kind, write the kind scoped template
		//if err := Templates.kindTemplate.Execute(w, KindTemplateArgs{d.group, k}); err != nil {
		//	panic(errors.Errorf("Failed to execute template %v", err))
		//}

		// For each version for each kind, write the version-kind scoped templates
		for v, _ := range d.typesByKindVersion[k] {
			// Don't write versionKind template for unversioned
			var versioned bool
			var args VersionKindTemplateArgs
			if args, prefix, versioned = d.GetVersionKindTemplateArgs(context, v, k); versioned {
				if err := Templates.versionKindTemplate.Execute(w, args); err != nil {
					panic(errors.Errorf("Failed to execute template %v", err))

				}
			}
		}
	}

	// Write the group scoped template
	args := GroupTemplateArgs{
		strings.Title(d.group),
		d.group,
		d.domain,
		d.GetOrderedVersions(),
		prefix,
		d.GetApiDefinitions()}
	if err := Templates.groupTemplate.Execute(w, args); err != nil {
		panic(errors.Errorf("Failed to execute template %v", err))
	}
	return nil
}

// Versions orders versions by maturity
type Versions []string

func (s Versions) Len() int      { return len(s) }
func (s Versions) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Versions) Less(i, j int) bool {
	switch {
	case strings.Contains(s[i], "alpha") && !strings.Contains(s[j], "alpha"):
		return true
	case !strings.Contains(s[i], "alpha") && strings.Contains(s[j], "alpha"):
		return false
	case strings.Contains(s[i], "beta") && !strings.Contains(s[j], "beta"):
		return true
	case !strings.Contains(s[i], "beta") && strings.Contains(s[j], "beta"):
		return false
	}

	return strings.Compare(s[i], s[j]) > 0
}

// GetApiDefinitions returns a list of apis formatted as "$version$Kind"
func (d *apisGenerator) GetApiDefinitions() []string {
	defs := []string{}
	for _, v := range d.GetOrderedVersions() {
		for k, _ := range d.typesByVersionKind[v] {
			if v == "unversioned" {
				continue
			}
			defs = append(defs, fmt.Sprintf("%s%s", v, k))
		}
	}
	return defs
}

// GetOrderedVersions returns the list of versions ordered by maturity.  More mature api
// versions appear earlier in the list.
//
// ga < beta < alpha
// v1alpha1 > v1alpha2
// v1alpha1 > v2alpha1
// v1alpha2 > v2alpha1
func (d *apisGenerator) GetOrderedVersions() []string {
	versions := Versions{}
	for v, _ := range d.typesByVersionKind {
		if v == "unversioned" {
			continue
		}
		versions = append(versions, v)
	}
	sort.Sort(versions)
	return versions
}

// UseDefaultStrategy returns true if the user has NOT provided a strategy to override the default
func (d *apisGenerator) UseDefaultStrategy(k string) bool {
	if _, f := d.pkg.Types[fmt.Sprintf("%sStrategy", k)]; f {
		return false
	}
	return true
}

// GetListOfKinds returns the list of unique kinds in the group
func (d *apisGenerator) GetListOfKinds() []string {
	kinds := sets.String{}
	for _, m := range d.typesByVersionKind {
		for _, t := range m {
			kinds.Insert(t.Name.Name)
		}
	}
	return kinds.List()
}

// GetVersionKindTemplateArgs returns arguments to pass into the VersionKind template
func (d *apisGenerator) GetVersionKindTemplateArgs(context *generator.Context, v, k string) (
	VersionKindTemplateArgs, string, bool) {

	sr := map[string]SubResource{}
	for _, s := range d.subresources {
		if s.Kind == k {
			sr[s.Path] = s
		}
	}
	t, f := d.typesByKindVersion[k][v]
	if !f {
		panic(errors.Errorf("Could not find type for %s %s", v, k))
	}

	// Don't write VersionKind info for unversioned objects
	if IsUnversioned(t, d.group) {
		return VersionKindTemplateArgs{}, "", false
	}

	etcdPrefix := filepath.Dir(t.Name.Package)
	d.imports.Insert(etcdPrefix)
	d.imports.Insert(t.Name.Package)

	// Parse the resource name from the comment lines
	comments := Comments(t.CommentLines)
	resource := comments.GetTag("resource")
	if len(resource) == 0 {
		panic(errors.Errorf("Must specify +resource comment for type %v", t.Name))
	}

	return VersionKindTemplateArgs{
		v, d.group, k, strings.ToLower(k),
		fmt.Sprintf("%s%s", v, k), resource,
		sr,
	}, etcdPrefix, true
}

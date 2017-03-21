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
	"html/template"
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

	groupTemplate       *template.Template
	kindTemplate        *template.Template
	versionKindTemplate *template.Template
}

var _ generator.Generator = &apisGenerator{}

func CreateApisGenerator(context *generator.Context, pkg *types.Package, group, domain string) generator.Generator {

	typesByVersionKind, typesByKindVersion, unversionedApiTypes := GetIndexedTypes(context, group)
	imports := sets.NewString(
		"reflect",
		"fmt",
		"k8s.io/client-go/pkg/api",
		"k8s.io/apiserver/pkg/registry/rest",
		"k8s.io/apimachinery/pkg/apimachinery/announced",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apimachinery/pkg/runtime/schema",
		"k8s.io/apiserver-builder/pkg/defaults",
		"genericapirequest \"k8s.io/apiserver/pkg/endpoints/request\"",
		"genericregistry \"k8s.io/apiserver/pkg/registry/generic/registry\"",
		"metainternalversion \"k8s.io/apimachinery/pkg/apis/meta/internalversion\"",
		"metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"",
	)

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
		groupTemplate,
		kindTemplate,
		versionKindTemplate,
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

func (d *apisGenerator) Finalize(context *generator.Context, w io.Writer) error {
	prefix := ""
	for _, k := range d.GetListOfKinds() {
		defaultStrategy := d.UseDefaultStrategy(k)

		// For each kind, write the kind scoped template
		if err := d.kindTemplate.Execute(w, KindTemplateArgs{d.group, k, strings.ToLower(k), defaultStrategy}); err != nil {
			panic(errors.Errorf("Failed to execute template %v", err))
		}

		// For each version for each kind, write the version-kind scoped templates
		for v, _ := range d.typesByKindVersion[k] {
			// Don't write versionKind template for unversioned
			var versioned bool
			var args VersionKindTemplateArgs
			if args, prefix, versioned = d.GetVersionKindTemplateArgs(context, v, k); versioned {
				if err := d.versionKindTemplate.Execute(w, args); err != nil {
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
	if err := d.groupTemplate.Execute(w, args); err != nil {
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
func (d *apisGenerator) GetVersionKindTemplateArgs(context *generator.Context, v, k string) (VersionKindTemplateArgs, string, bool) {
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
		v, d.group, k, strings.ToLower(k), fmt.Sprintf("%s%s", v, k), resource,
	}, etcdPrefix, true
}

type VersionKindTemplateArgs struct {
	Version   string
	Group     string
	Kind      string
	LowerKind string
	Name      string
	Resource  string
}

const versionKindTemplateName = "VersionKindTemplate"
const versionKindTemplateString = `
// Definition used to register {{.Name}} with the apiserver
var {{.Name}}ApiDefinition = &defaults.ResourceDefinition{
	{{.Version}}.SchemeGroupVersion.WithResource("{{.Resource}}"),
	singleton{{.Kind}}Strategy,
	singleton{{.Kind}}Strategy,
	singleton{{.Kind}}Strategy,
	singleton{{.Kind}}Strategy,
	map[string]*defaults.ResourceDefinition{
		"{{.Resource}}/status": {{.Name}}StatusApiDefinition,
	},
	singleton{{.Kind}}Strategy.BasicMatch,
	func(store *genericregistry.Store) rest.Storage { return &{{.Kind}}Store{store} },
}

var {{.Name}}StatusApiDefinition = &defaults.ResourceDefinition{
	{{.Version}}.SchemeGroupVersion.WithResource("{{.Resource}}"),
	singleton{{.Kind}}StatusStrategy,
	singleton{{.Kind}}StatusStrategy,
	singleton{{.Kind}}StatusStrategy,
	singleton{{.Kind}}StatusStrategy,
	map[string]*defaults.ResourceDefinition{},
	singleton{{.Kind}}StatusStrategy.BasicMatch,
	func(store *genericregistry.Store) rest.Storage { return &{{.Kind}}StatusStore{store} },
}

`

type KindTemplateArgs struct {
	Group      string
	Kind       string
	LowerKind  string
	UseDefault bool
}

const kindTemplateName = "KindTemplate"
const kindTemplateString = `


///////////////////////////////////////////////////////////////////////////////
// {{.Kind}} End user functions //
///////////////////////////////////////////////////////////////////////////////

// Add functions to this type in order to override the default behaviors
type {{.Kind}}Strategy struct {
	Default{{.Kind}}Strategy
}

// Add functions to this type in order to override the default behaviors
type {{.Kind}}Store struct {
	*genericregistry.Store
}

// Add functions to this type in order to override the default behaviors
type {{.Kind}}StatusStore struct {
	*genericregistry.Store
}

// Registry is an interface for things that know how to store {{.Kind}}.
type {{.Kind}}Registry interface {
	List{{.Kind}}s(ctx genericapirequest.Context, options *metainternalversion.ListOptions) (*{{.Group}}.{{.Kind}}List, error)
	Get{{.Kind}}(ctx genericapirequest.Context, id string, options *metav1.GetOptions) (*{{.Group}}.{{.Kind}}, error)
	Create{{.Kind}}(ctx genericapirequest.Context, id *{{.Group}}.{{.Kind}}) (*{{.Group}}.{{.Kind}}, error)
	Update{{.Kind}}(ctx genericapirequest.Context, id *{{.Group}}.{{.Kind}}) (*{{.Group}}.{{.Kind}}, error)
	Delete{{.Kind}}(ctx genericapirequest.Context, id string) error
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func New{{.Kind}}Registry(s rest.StandardStorage) {{.Kind}}Registry {
	return &storage{{.Kind}}{s}
}

///////////////////////////////////////////////////////////////////////////////
// {{.Kind}} System functions //
///////////////////////////////////////////////////////////////////////////////

// Use the override strategy and embedd the defaults for anything not override.
var singleton{{.Kind}}Strategy = &{{.Kind}}Strategy{
	Default{{.Kind}}Strategy{ // Overide some methods
		defaults.NewBasicStrategy(), // Use defaults
	},
}

// Default Strategy for {{.Kind}}
type Default{{.Kind}}Strategy struct {
	// Inherit the basic create, delete, update strategy.
	defaults.BasicCreateDeleteUpdateStrategy
}

// NewFunc returns a new empty {{.Kind}}
func (r Default{{.Kind}}Strategy) NewFunc() runtime.Object {
	return &{{.Group}}.{{.Kind}}{}
}

// NewListFunc returns a new empty List of {{.Kind}}
func (r Default{{.Kind}}Strategy) NewListFunc() runtime.Object {
	return &{{.Group}}.{{.Kind}}List{}
}

// ObjectNameFunc returns the name for a {{.Kind}}
func (r Default{{.Kind}}Strategy) ObjectNameFunc(obj runtime.Object) (string, error) {
	return obj.(*{{.Group}}.{{.Kind}}).Name, nil
}

func ({{.Kind}}Strategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
	o := obj.(*{{.Group}}.{{.Kind}})
	o.Status = {{.Group}}.{{.Kind}}Status{}
	o.Generation = 1
}

func ({{.Kind}}Strategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
	new{{.Kind}} := obj.(*{{.Group}}.{{.Kind}})
	old{{.Kind}} := old.(*{{.Group}}.{{.Kind}})
	new{{.Kind}}.Status = old{{.Kind}}.Status

	// Spec and annotation updates bump the generation.
	if !reflect.DeepEqual(new{{.Kind}}.Spec, old{{.Kind}}.Spec) ||
		!reflect.DeepEqual(new{{.Kind}}.Annotations, old{{.Kind}}.Annotations) {
		new{{.Kind}}.Generation = old{{.Kind}}.Generation + 1
	}
}

// Implement Status endpoint
// StatusREST implements the REST endpoint for changing the status of a deployment
type {{.Kind}}StatusStrategy struct {
	{{.Kind}}Strategy
}

// singleton{{.Kind}}StatusStrategy contains the cross-cutting storage
var singleton{{.Kind}}StatusStrategy = {{.Kind}}StatusStrategy{*singleton{{.Kind}}Strategy}

//// {{.Kind}}StatusREST contains the REST method implementations
//type {{.Kind}}StatusREST struct {
//	store *genericregistry.Store
//}
//
//func (r {{.Kind}}StatusREST) New() runtime.Object {
//	return &{{.Group}}.{{.Kind}}{}
//}
//
//// Get retrieves the object from the storage. It is required to support Patch.
//func (r {{.Kind}}StatusREST) Get(ctx genericapirequest.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
//	return r.store.Get(ctx, name, options)
//}
//
//// Update alters the status subset of an object.
//func (r {{.Kind}}StatusREST) Update(ctx genericapirequest.Context, name string, objInfo rest.UpdatedObjectInfo) (runtime.Object, bool, error) {
//	return r.store.Update(ctx, name, objInfo)
//}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update of status
func ({{.Kind}}StatusStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
	new{{.Kind}} := obj.(*{{.Group}}.{{.Kind}})
	old{{.Kind}} := old.(*{{.Group}}.{{.Kind}})
	new{{.Kind}}.Spec = old{{.Kind}}.Spec
	new{{.Kind}}.Labels = old{{.Kind}}.Labels
}

// Implement Registry
// storage puts strong typing around storage calls
type storage{{.Kind}} struct {
	rest.StandardStorage
}

func (s *storage{{.Kind}}) List{{.Kind}}s(ctx genericapirequest.Context, options *metainternalversion.ListOptions) (*{{.Group}}.{{.Kind}}List, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Group}}.{{.Kind}}List), err
}

func (s *storage{{.Kind}}) Get{{.Kind}}(ctx genericapirequest.Context, id string, options *metav1.GetOptions) (*{{.Group}}.{{.Kind}}, error) {
	obj, err := s.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Group}}.{{.Kind}}), nil
}

func (s *storage{{.Kind}}) Create{{.Kind}}(ctx genericapirequest.Context, object *{{.Group}}.{{.Kind}}) (*{{.Group}}.{{.Kind}}, error) {
	obj, err := s.Create(ctx, object)
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Group}}.{{.Kind}}), nil
}

func (s *storage{{.Kind}}) Update{{.Kind}}(ctx genericapirequest.Context, object *{{.Group}}.{{.Kind}}) (*{{.Group}}.{{.Kind}}, error) {
	obj, _, err := s.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object, api.Scheme))
	if err != nil {
		return nil, err
	}
	return obj.(*{{.Group}}.{{.Kind}}), nil
}

func (s *storage{{.Kind}}) Delete{{.Kind}}(ctx genericapirequest.Context, id string) error {
	_, err := s.Delete(ctx, id, nil)
	return err
}


`

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

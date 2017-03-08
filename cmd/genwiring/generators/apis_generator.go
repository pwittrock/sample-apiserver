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
		"k8s.io/apimachinery/pkg/apimachinery/announced",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apimachinery/pkg/runtime/schema",
		"k8s.io/apiserver/pkg/storage/names",
		"k8s.io/sample-apiserver/pkg/defaults",
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
		if err := d.kindTemplate.Execute(w, KindTemplateArgs{d.group, k, defaultStrategy}); err != nil {
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
func (d *apisGenerator) GetVersionKindTemplateArgs(context *generator.Context, v string, k string) (VersionKindTemplateArgs, string, bool) {
	t, f := d.typesByKindVersion[k][v]
	if !f {
		panic(errors.Errorf("Could not find type for %s %s", v, k))
	}

	// Add the imports
	etcdPrefix := filepath.Dir(t.Name.Package)
	d.imports.Insert(etcdPrefix)
	d.imports.Insert(t.Name.Package)

	// Don't write VersionKind info for unversioned objects
	if IsUnversioned(t, d.group) {
		return VersionKindTemplateArgs{}, "", false
	}

	// Parse the resource name from the comment lines
	comments := Comments(t.CommentLines)
	resource := comments.GetTag("resource")
	if len(resource) == 0 {
		panic(errors.Errorf("Must specify +resource comment for type %v", t.Name))
	}

	return VersionKindTemplateArgs{
		v, d.group, k, fmt.Sprintf("%s%s", v, k), resource,
	}, etcdPrefix, true
}

type VersionKindTemplateArgs struct {
	Version  string
	Group    string
	Kind     string
	Name     string
	Resource string
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
	singleton{{.Kind}}Strategy.BasicMatch,
}
`

type KindTemplateArgs struct {
	Group      string
	Kind       string
	UseDefault bool
}

const kindTemplateName = "KindTemplate"
const kindTemplateString = `
{{if .UseDefault}}
// Use the default strategy.  To override - in another file - define the struct {{.Kind}}Strategy and regenerate code
var singleton{{.Kind}}Strategy = &Default{{.Kind}}Strategy{
	defaults.BasicCreateDeleteUpdateStrategy{defaults.Scheme, names.SimpleNameGenerator},
}{{else}}
// Use the override strategy and embedd the defaults for anything not override.
var singleton{{.Kind}}Strategy = &{{.Kind}}Strategy{Default{{.Kind}}Strategy{
	defaults.BasicCreateDeleteUpdateStrategy{defaults.Scheme, names.SimpleNameGenerator},
}}{{end}}

// Default Strategy for {{.Kind}}
type Default{{.Kind}}Strategy struct {
	// Inherit the basic create, delete, update strategy.
	defaults.BasicCreateDeleteUpdateStrategy
}

// NewFunc returns a new empty {{.Kind}}
func (r *Default{{.Kind}}Strategy) NewFunc() runtime.Object {
	return &{{.Group}}.{{.Kind}}{}
}

// NewListFunc returns a new empty List of {{.Kind}}
func (r *Default{{.Kind}}Strategy) NewListFunc() runtime.Object {
	return &{{.Group}}.{{.Kind}}List{}
}

// ObjectNameFunc returns the name for a {{.Kind}}
func (r *Default{{.Kind}}Strategy) ObjectNameFunc(obj runtime.Object) (string, error) {
	return obj.(*{{.Group}}.{{.Kind}}).Name, nil
}
`

type GroupTemplateArgs struct {
	Group               string
	Domain              string
	VersionOrder        []string
	ImportPrefix        string
	ResourceDefinitions []string
}

const groupTemplateName = "ApiTemplate"
const groupTemplateString = `
// Order list of version preferences
var versionPreferenceOrder = []string{
	{{- range $v := .VersionOrder}}
	{{$v}}.SchemeGroupVersion.Version,
	{{- end -}}
}
var versionToSchemeFunc = announced.VersionToSchemeFunc{
	{{- range $v := .VersionOrder}}
	{{$v}}.SchemeGroupVersion.Version: {{$v}}.AddToScheme,
	{{- end -}}
}

type {{.Group}}Provider struct{}

func GetWardleProvider() defaults.ResourceDefinitionProvider {
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
	return versionPreferenceOrder
}

func (w *{{.Group}}Provider) GetImportPrefix() string {
	return "{{.ImportPrefix}}"
}

func (w *{{.Group}}Provider) SchemeFunc() announced.SchemeFunc {
	return {{.Group}}.AddToScheme
}

func (w *{{.Group}}Provider) VersionToSchemeFunc() announced.VersionToSchemeFunc {
	return versionToSchemeFunc
}
`

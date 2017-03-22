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


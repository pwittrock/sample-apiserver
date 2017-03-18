/*
Copyright 2016 The Kubernetes Authors.

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

package api

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/go-openapi/loads"
)

var AllowErrors = flag.Bool("allow-errors", false, "If true, don't fail on errors.")
var GenOpenApiDir = flag.String("gen-open-api-dir", "gen_open_api/", "Directory containing open api files")
var ConfigDir = flag.String("config-dir", "", "Directory contain api files.")

func NewConfig() *Config {
	config := loadYamlConfig()
	specs := LoadOpenApiSpec()

	// Initialize all of the operations
	config.Definitions = GetDefinitions(specs)

	// Initialization for ToC resources only
	vistToc := func(resource *Resource, definition *Definition) {
		definition.InToc = true // Mark as in Toc
		resource.Definition = definition
		config.initDefExample(definition) // Init the example yaml
	}
	config.VisitResourcesInToc(config.Definitions, vistToc)

	// Get the map of operations appearing in the open-api spec keyed by id
	config.InitOperations(specs)
	config.CleanUp()

	return config
}

func verifyBlacklisted(operation Operation) {
	switch {
	case strings.Contains(operation.ID, "NamespacedScheduledJob"):
	case strings.Contains(operation.ID, "ScheduledJobForAllNamespaces"):
	case strings.Contains(operation.ID, "ScheduledJobListForAllNamespaces"):
	case strings.Contains(operation.ID, "V1beta1NamespacedReplicationcontrollersScale"):
	case strings.Contains(operation.ID, "NamespacedPodAttach"):
	case strings.Contains(operation.ID, "NamespacedPodWithPath"):
	case strings.Contains(operation.ID, "proxyCoreV1"):
	case strings.Contains(operation.ID, "NamespacedScaleScale"):
	case strings.Contains(operation.ID, "NamespacedBindingBinding"):
	case strings.Contains(operation.ID, "NamespacedPodExe"):
	case strings.Contains(operation.ID, "logFileHandler"):
	case strings.Contains(operation.ID, "logFileListHandler"):
	case strings.Contains(operation.ID, "replaceCoreV1NamespaceFinalize"):
	case strings.Contains(operation.ID, "NamespacedEvictionEviction"):
	case strings.Contains(operation.ID, "getCodeVersion"):
	case strings.Contains(operation.ID, "V1beta1CertificateSigningRequestApproval"):
	default:
		panic(fmt.Sprintf("No Definition found for %s [%s].  \n", operation.ID, operation.Path))
	}
}

// GetOperations returns all Operations found in the Documents
func (config *Config) InitOperations(specs []*loads.Document) {
	o := Operations{}
	VisitOperations(specs, func(operation Operation) {
		//fmt.Printf("Operation: %s\n", operation.ID)
		o[operation.ID] = &operation
	})
	config.Operations = o

	config.mapOperationsToDefinitions()
	VisitOperations(specs, func(operation Operation) {
		if o, found := config.Operations[operation.ID]; !found || o.Definition == nil {
			verifyBlacklisted(operation)
		}
	})
	config.Definitions.initializeOperationParameters(config.Operations)

	// Clear the operations.  We still have to calculate the operations because that is how we determine
	// the API Group for each definition.
	if !*BuildOps {
		config.Operations = Operations{}
		config.OperationCategories = []OperationCategory{}
		for _, d := range config.Definitions.GetAllDefinitions() {
			d.OperationCategories = []*OperationCategory{}
		}
	}
}

// CleanUp sorts and dedups fields
func (c *Config) CleanUp() {
	for _, d := range c.Definitions.GetAllDefinitions() {
		sort.Sort(d.AppearsIn)
		sort.Sort(d.Fields)
		dedup := SortDefinitionsByName{}
		last := ""
		for _, i := range d.AppearsIn {
			if i.Name == last {
				continue
			}
			last = i.Name
			dedup = append(dedup, i)
		}
		d.AppearsIn = dedup
	}
}

// loadYamlConfig reads the config yaml file into a struct
func loadYamlConfig() *Config {
	f := filepath.Join(*ConfigDir, "config.yaml")

	config := &Config{}
	contents, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Printf("Failed to read yaml file %s: %v", f, err)
		os.Exit(2)
	}

	err = yaml.Unmarshal(contents, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	writeCategory := OperationCategory{
		Name: "Write Operations",
		OperationTypes: []OperationType{
			{
				Name:  "Create",
				Match: "create${group}${version}(Namespaced)?${resource}",
			},
			{
				Name:  "Patch",
				Match: "patch${group}${version}(Namespaced)?${resource}",
			},
			{
				Name:  "Replace",
				Match: "replace${group}${version}(Namespaced)?${resource}",
			},
			{
				Name:  "Delete",
				Match: "delete${group}${version}(Namespaced)?${resource}",
			},
			{
				Name:  "Delete Collection",
				Match: "delete${group}${version}Collection(Namespaced)?${resource}",
			},
		},
	}

	readCategory := OperationCategory{
		Name: "Read Operations",
		OperationTypes: []OperationType{
			{
				Name:  "Read",
				Match: "read${group}${version}(Namespaced)?${resource}",
			},
			{
				Name:  "List",
				Match: "list${group}${version}(Namespaced)?${resource}",
			},
			{
				Name:  "List All Namespaces",
				Match: "list${group}${version}(Namespaced)?${resource}ForAllNamespaces",
			},
			{
				Name:  "Watch",
				Match: "watch${group}${version}(Namespaced)?${resource}",
			},
			{
				Name:  "Watch List",
				Match: "watch${group}${version}(Namespaced)?${resource}List",
			},
			{
				Name:  "Watch List All Namespaces",
				Match: "watch${group}${version}(Namespaced)?${resource}ListForAllNamespaces",
			},
		},
	}

	statusCategory := OperationCategory{
		Name: "Status Operations",
		OperationTypes: []OperationType{
			{
				Name:  "Patch Status",
				Match: "patch${group}${version}(Namespaced)?${resource}Status",
			},
			{
				Name:  "Read Status",
				Match: "read${group}${version}(Namespaced)?${resource}Status",
			},
			{
				Name:  "Replace Status",
				Match: "replace${group}${version}(Namespaced)?${resource}Status",
			},
		},
	}

	config.OperationCategories = append([]OperationCategory{writeCategory, readCategory, statusCategory}, config.OperationCategories...)

	return config
}

// initOpExample reads the example config for each operation and sets it
func (config *Config) initOpExample(o *Operation) {
	path := o.Type.Name + ".yaml"
	path = filepath.Join(*ConfigDir, config.ExampleLocation, o.Definition.Name, path)
	path = strings.Replace(path, " ", "_", -1)
	path = strings.ToLower(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, &o.ExampleConfig)
	if err != nil {
		panic(fmt.Sprintf("Could not Unmarshal ExampleConfig yaml: %s\n", content))
	}
}

func (config *Config) GetDefExampleFile(d *Definition) string {
	return strings.Replace(strings.ToLower(filepath.Join(*ConfigDir, config.ExampleLocation, d.Name, d.Name+".yaml")), " ", "_", -1)
}

func (config *Config) initDefExample(d *Definition) {
	content, err := ioutil.ReadFile(config.GetDefExampleFile(d))
	if err != nil || len(content) <= 0 {
		//fmt.Printf("Missing example: %s %v\n", d.Name, err)
		return
	}
	err = yaml.Unmarshal(content, &d.Sample)
	if err != nil {
		panic(fmt.Sprintf("Could not Unmarshal SampleConfig yaml: %s\n", content))
	}
}

func getOperationId(match string, group string, version ApiVersion, kind string) string {
	// Substitute the api definition group-version-kind into the operation template and look for a match
	v, k := doScaleIdHack(string(version), kind, match)
	match = strings.Replace(match, "${group}", string(group), -1)
	match = strings.Replace(match, "${version}", v, -1)
	match = strings.Replace(match, "${resource}", k, -1)
	return match
}

func (config *Config) setOperation(match, namespaceRep string,
	ot *OperationType, oc *OperationCategory, definition *Definition) {

	key := strings.Replace(match, "(Namespaced)?", namespaceRep, -1)
	if o, found := config.Operations[key]; found {
		// Each operation should have exactly 1 definition
		if o.Definition != nil {
			panic(fmt.Sprintf(
				"Found multiple matching defintions [%s/%s, %s/%s] for operation key: %s",
				definition.Version, definition.Name, o.Definition.Version, o.Definition.Name, key))
		}
		o.Type = *ot
		o.Definition = definition
		oc.Operations = append(oc.Operations, o)
		config.initOpExample(o)
	}
}

// mapOperationsToDefinitions adds operations to the definitions they operate
// This is done by - for each definition - look at all potentially matching operations from operation categories
func (config *Config) mapOperationsToDefinitions() {
	// Look for matching operations for each definition
	for _, definition := range config.Definitions.GetAllDefinitions() {
		// Inlined definitions don't have operations
		if definition.IsInlined {
			continue
		}

		// Iterate through categories
		for i := range config.OperationCategories {
			oc := config.OperationCategories[i]

			// Iterate through possible operation matches
			for j := range oc.OperationTypes {
				// Iterate through possible api groups since we don't know the api group of the definition
				ot := oc.OperationTypes[j]

				operationId := getOperationId(ot.Match, definition.GetOperationGroupName(), definition.Version, definition.Name)
				// Look for a matching operation and set on the definition if found
				config.setOperation(operationId, "Namespaced", &ot, &oc, definition)
				config.setOperation(operationId, "", &ot, &oc, definition)
			}

			// If we found operations for this category, add the category to the definition
			if len(oc.Operations) > 0 {
				definition.OperationCategories = append(definition.OperationCategories, &oc)
			}
		}
	}
}

func doScaleIdHack(version, name, match string) (string, string) {
	// Hack to get around ids
	if strings.HasSuffix(match, "${resource}Scale") && name != "Scale" {
		// Scale names don't generate properly
		name = strings.ToLower(name) + "s"
		out := []rune(name)
		out[0] = unicode.ToUpper(out[0])
		name = string(out)
	}
	out := []rune(version)
	out[0] = unicode.ToUpper(out[0])
	version = string(out)

	return version, name
}

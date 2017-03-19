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

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/sets"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

var repoPath string
var repoName string
var repoPackage string
var copyFrom string
var domain string
var types []string
var goPath string

func main() {
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	cmd.AddCommand(initCmd, addTypesCmd, genCmd)
	initCmd.Flags().StringVar(&repoPath, "repo-path", "/out", "path to repo")
	initCmd.Flags().StringVar(&repoName, "repo-name", "", "full name of repo")
	initCmd.Flags().StringVar(&copyFrom, "from-path", "/go/src/github.com/pwittrock/apiserver-helloworld/", "path to repo to copy from")
	initCmd.Flags().StringVar(&domain, "domain", "k8s.io", "domain group lives in")

	addTypesCmd.Flags().StringVar(&repoName, "repo-name", "", "full name of repo")
	addTypesCmd.Flags().StringVar(&repoPath, "repo-path", "/out", "path to repo")
	addTypesCmd.Flags().StringSliceVar(&types, "types", []string{}, "list of group/version/kind")
	addTypesCmd.Flags().StringVar(&repoPackage, "repo-package", "", "repo package")
	addTypesCmd.Flags().StringVar(&domain, "domain", "k8s.io", "domain group lives in")

	genCmd.Flags().StringVar(&repoName, "repo-name", "", "full name of repo")
	genCmd.Flags().StringVar(&goPath, "go-path", "/out", "gopath")

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func RunMain(cmd *cobra.Command, args []string) {
	cmd.Help()
}

var cmd = &cobra.Command{
	Use:   "kubec",
	Short: "kubec builds Kubernetes extensions",
	Long:  `kubec is a set of commands for building Kubernetes extensions`,
	Run:   RunMain,
}

var genCmd = &cobra.Command{
	Use:   "generate",
	Short: "Create generated files",
	Long:  `Create generated files`,
	Run:   RunGenCmd,
}

func RunGenCmd(cmd *cobra.Command, args []string) {
	c := exec.Command("./run.sh")
	c.Env = append(c.Env, fmt.Sprintf("REPO=%s", repoName))
	c.Env = append(c.Env, fmt.Sprintf("GOPATH=%s", goPath))
	out, err := c.CombinedOutput()
	if err != nil {
		panic(fmt.Errorf("Error generating files: %v %s", err, out))
	}
	fmt.Printf("%s", out)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a directory",
	Long:  `Initialize a directory`,
	Run:   RunInit,
}

func RunInit(cmd *cobra.Command, args []string) {
	repoPath = filepath.Join(repoPath, "src", repoName)
	out, _ := exec.Command("cp", "-r", filepath.Join(copyFrom, "vendor"), repoPath).CombinedOutput()
	fmt.Printf("%s", out)
	out, _ = exec.Command("cp", "-r", filepath.Join(copyFrom, "Godeps"), repoPath).CombinedOutput()
	fmt.Printf("%s", out)
	out, _ = exec.Command("cp", "-r", filepath.Join(copyFrom, "main.go"), repoPath).CombinedOutput()
	fmt.Printf("%s", out)
	out, _ = exec.Command("mkdir", "-p", filepath.Join(repoPath, "apis")).CombinedOutput()
	fmt.Printf("%s", out)
	out, _ = exec.Command("mkdir", "-p", filepath.Join(repoPath, "docs")).CombinedOutput()
	fmt.Printf("%s", out)
	out, _ = exec.Command("mkdir", "-p", filepath.Join(repoPath, "pkg/openapi")).CombinedOutput()
	fmt.Printf("%s", out)
	out, _ = exec.Command("cp", "-r", filepath.Join(copyFrom, "pkg/openapi/doc.go"), filepath.Join(repoPath, "pkg/openapi")).CombinedOutput()
	fmt.Printf("%s", out)
}

var addTypesCmd = &cobra.Command{
	Use:   "add-types",
	Short: "Create new entries for group/version/kind types",
	Long:  `Specify types using group/version/kind`,
	Run:   RunAddTypes,
}

func RunAddTypes(cmd *cobra.Command, args []string) {
	repoPath = filepath.Join(repoPath, "src", repoName)

	groups := sets.String{}
	groupVersions := sets.String{}
	kindsToGroupVersion := map[string]string{}
	for _, tuple := range types {
		groupVersionKind := strings.Split(tuple, "/")
		groups.Insert(groupVersionKind[0])
		gv := filepath.Join(groupVersionKind[0], groupVersionKind[1])
		kindsToGroupVersion[groupVersionKind[2]] = gv
		groupVersions.Insert(gv)
	}

	for _, gv := range groupVersions.List() {
		split := strings.Split(gv, "/")
		group := split[0]
		version := split[1]

		path := filepath.Join(repoPath, "apis", gv)
		_, err := os.Stat(path)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(fmt.Sprintf("Could not stat directory %s %v", path, err))
			}
			fmt.Printf("Creating directory %s\n", path)
			out, err := exec.Command("mkdir", "-p", path).CombinedOutput()
			if err != nil {
				fmt.Printf("Failed to create directory %s %v %s", path, err, out)
			}
		}

		apisDocPath := filepath.Join(repoPath, "apis/doc.go")
		_, err = os.Stat(apisDocPath)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(fmt.Sprintf("Could not stat file %s %v", apisDocPath, err))
			}
			fmt.Printf("Creating file %s\n", apisDocPath)
			t := template.Must(template.New("apis-doc-template").Parse(apisDocTemplate))
			f, err := os.Create(apisDocPath)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.Close()

			f, err = os.OpenFile(apisDocPath, os.O_WRONLY, 0)
			err = t.Execute(f, ApisDocTemplateArguments{
				Domain: domain,
			})
			if err != nil {
				fmt.Println(err)
			}
			f.Close()
		}

		typesgo := filepath.Join(path, "types.go")
		_, err = os.Stat(typesgo)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(fmt.Sprintf("Could not stat file %s %v", typesgo, err))
			}
			t := template.Must(template.New("new-types-template").Parse(newTypesTemplate))
			f, err := os.Create(typesgo)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.Close()

			f, err = os.OpenFile(typesgo, os.O_WRONLY, 0)
			err = t.Execute(f, NewTypesGoArguments{
				Package: version,
			})
			if err != nil {
				fmt.Println(err)
			}
			f.Close()
		}

		docgo := filepath.Join(path, "doc.go")
		_, err = os.Stat(docgo)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(fmt.Sprintf("Could not stat file %s %v", docgo, err))
			}

			t := template.Must(template.New("new-doc-template").Parse(newVersionDocTemplate))
			f, err := os.Create(docgo)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.Close()

			f, err = os.OpenFile(docgo, os.O_WRONLY, 0)
			err = t.Execute(f, NewDocTemplateArguments{version, filepath.Join(repoName, "apis", group), group})
			if err != nil {
				fmt.Println(err)
			}
			f.Close()
		}

		groupdocgo := filepath.Join(repoPath, "apis", group, "doc.go")
		_, err = os.Stat(groupdocgo)
		if err != nil {
			if !os.IsNotExist(err) {
				panic(fmt.Sprintf("Could not stat file %s %v", groupdocgo, err))
			}

			t := template.Must(template.New("new-group-doc-template").Parse(newGroupDocTemplate))
			f, err := os.Create(groupdocgo)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.Close()

			f, err = os.OpenFile(groupdocgo, os.O_WRONLY, 0)
			err = t.Execute(f, NewDocTemplateArguments{version, filepath.Join(repoPackage, "apis", group), group})
			if err != nil {
				fmt.Println(err)
			}
			f.Close()
		}
	}

	for k, gv := range kindsToGroupVersion {
		t := template.Must(template.New("add-types-template").Parse(addTypesTemplate))
		path := filepath.Join(repoPath, "apis", gv)

		typesgo := filepath.Join(path, "types.go")
		f, err := os.Open(typesgo)
		if err != nil {
			panic(err)
			return
		}

		contents, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
			return
		}
		if strings.Contains(string(contents), fmt.Sprintf("type %s struct {", k)) {
			fmt.Printf("Skipping kind %s\n", k)
			f.Close()
			continue
		}
		f.Close()

		f, err = os.OpenFile(typesgo, os.O_WRONLY|os.O_APPEND, 0)
		err = t.Execute(f, AddTypeArguments{
			Kind:     k,
			Resource: fmt.Sprintf("%ss", strings.ToLower(k)),
		})
		if err != nil {
			fmt.Println(err)
		}
		f.Close()
	}
}

type AddTypeArguments struct {
	Resource string
	Kind     string
}

var addTypesTemplate = (`
// +genclient=true
// +genapi=true
// +resource={{.Resource}}
// +k8s:openapi-gen=true
type {{.Kind}} struct {
	metav1.TypeMeta   ` + "`json:\",inline\"`" + `
	metav1.ObjectMeta ` + "`json:\"metadata,omitempty\"`" + `

	Spec   {{.Kind}}Spec   ` + "`json:\"spec,omitempty\"`" + `
	Status {{.Kind}}Status ` + "`json:\"status,omitempty\"`" + `
}

type {{.Kind}}Spec struct {
}

type {{.Kind}}Status struct {
}
`)

type NewTypesGoArguments struct {
	Package string
}

var newTypesTemplate = (`

package {{.Package}}

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

`)

type NewDocTemplateArguments struct {
	Version string
	Package string
	Group   string
}

var newVersionDocTemplate = `
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

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen={{.Package}}

package {{.Version}}
`

var newGroupDocTemplate = `
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

// +k8s:deepcopy-gen=package,register

// Package api is the internal version of the API.
package {{.Group}}

`

type ApisDocTemplateArguments struct {
	Domain string
}

var apisDocTemplate = `
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

//
// +domain={{.Domain}}

package apis

`

var openApiDoc = `/*
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

// Package openapi exists to hold generated openapi code
package openapi

`

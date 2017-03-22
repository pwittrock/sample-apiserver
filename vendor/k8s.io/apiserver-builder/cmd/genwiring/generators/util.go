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
	"path/filepath"
	"strings"

	"github.com/pkg/errors"

	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
	//"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/util/sets"
)

func IsApiType(t *types.Type) bool {
	for _, c := range t.CommentLines {
		if strings.Contains(c, "+resource") {
			return true
		}
	}
	return false
}

func HasSubResource(t *types.Type) bool {
	for _, c := range t.CommentLines {
		if strings.Contains(c, "+subresource") {
			return true
		}
	}
	return false
}

func IsUnversioned(t *types.Type, group string) bool {
	return IsApisDir(filepath.Base(filepath.Dir(t.Name.Package))) && GetGroup(t) == group
}

func IsVersioned(t *types.Type, group string) bool {
	dir := filepath.Base(filepath.Dir(filepath.Dir(t.Name.Package)))
	return IsApisDir(dir) && GetGroup(t) == group
}

func GetVersion(t *types.Type, group string) string {
	if !IsVersioned(t, group) {
		panic(errors.Errorf("Cannot get version for unversioned type %v", t.Name))
	}
	return filepath.Base(t.Name.Package)
}

func IsGroup(t *types.Type, group string) bool {
	return GetGroup(t) == group
}

func GetGroup(t *types.Type) string {
	return filepath.Base(GetGroupPackage(t))
}

func GetGroupPackage(t *types.Type) string {
	if IsApisDir(filepath.Base(filepath.Dir(t.Name.Package))) {
		return t.Name.Package
	}
	return filepath.Dir(t.Name.Package)
}

func GetKind(t *types.Type, group string) string {
	if !IsVersioned(t, group) && !IsUnversioned(t, group) {
		panic(errors.Errorf("Cannot get kind for type not in group %v", t.Name))
	}
	return t.Name.Name
}

func IsApisDir(dir string) bool {
	return dir == "apis" || dir == "api"
}

type Comments []string

func (c Comments) GetTag(name string) string {
	for _, c := range c {
		prefix := fmt.Sprintf("+%s=", name)
		if strings.HasPrefix(c, prefix) {
			//fmt.Printf("Checking %s has %s %v\n", c, prefix, strings.HasPrefix(c, prefix))
			return strings.TrimLeft(c, prefix)
		}
	}
	return ""
}

func (c Comments) GetTags(name string) []string {
	tags := []string{}
	for _, c := range c {
		prefix := fmt.Sprintf("+%s=", name)
		if strings.HasPrefix(c, prefix) {
			//fmt.Printf("Checking %s has %s %v\n", c, prefix, strings.HasPrefix(c, prefix))
			tags = append(tags, strings.TrimLeft(c, prefix))
		}
	}
	return tags
}

func GetApiTypes(c *generator.Context, group string) []*types.Type {
	types := []*types.Type{}
	for _, o := range c.Order {
		if IsApiType(o) {
			if IsGroup(o, group) {
				types = append(types, o)
			}
		}
	}
	return types
}

func GetApiTypeNames(c *generator.Context, group string) []string {
	types := []string{}
	for _, o := range GetApiTypes(c, group) {
		types = append(types, fmt.Sprintf("%s", o.Name.Name))
		types = append(types, fmt.Sprintf("%sList", o.Name.Name))
	}
	return types
}

func GetSubresources(context *generator.Context, group string) map[string]SubResource {
	subresources := map[string]SubResource{}
	// Find subresources
	for _, o := range context.Order {
		comments := Comments(o.CommentLines)
		subresourceList := comments.GetTags("subresource")
		if len(subresourceList) == 0 {
			// Not a subresource
			continue
		}
		if !IsGroup(o, group) {
			continue
		}
		for _, subresource := range subresourceList {
			args := strings.Split(subresource, ",")
			path, kind, requestKind, rest := args[0], args[1], args[2], args[3]
			imp := ""
			if strings.Contains(requestKind, ".") {
				last := strings.LastIndex(requestKind, ".")
				imp = requestKind[:last]

				// Set the request kind to the struct name
				requestKind = requestKind[last+1:]
				// Find the package
				pkg := filepath.Base(imp)
				// Prefix the struct name with the package it is in
				requestKind = strings.Join([]string{pkg, requestKind}, ".")
			}
			sr := SubResource{kind, requestKind, path, rest, imp}
			if v, f := subresources[path]; f {
				panic(errors.Errorf("Multiple subresources registered for path %s: %v %v",
					path, v, subresource))
			}
			subresources[path] = sr
		}

	}
	return subresources
}

func GetIndexedTypes(context *generator.Context, group string) (map[string]map[string]*types.Type, map[string]map[string]*types.Type, sets.String) {
	versionedApiTypes := sets.NewString()
	unversionedApiTypes := sets.NewString()
	for _, c := range context.Order {
		if IsUnversioned(c, group) {
			unversionedApiTypes.Insert(c.Name.Name)
		}
	}
	for _, c := range context.Order {
		if IsVersioned(c, group) && IsApiType(c) {
			versionedApiTypes.Insert(c.Name.Name)
		}
	}

	// Only keep api types
	unversionedApiTypes = unversionedApiTypes.Intersection(versionedApiTypes)

	// Find types that have versioned objects, but are missing unversioned objects
	typesByVersionKind := map[string]map[string]*types.Type{}
	typesByKindVersion := map[string]map[string]*types.Type{}
	for _, c := range context.Order {
		// Not in the group
		if GetGroup(c) != group {
			continue
		}
		// Not an api type
		if !versionedApiTypes.Has(c.Name.Name) && !unversionedApiTypes.Has(c.Name.Name) {
			continue
		}

		version := unversioned
		if IsVersioned(c, group) {
			version = GetVersion(c, group)
		}
		if _, f := typesByVersionKind[version]; !f {
			typesByVersionKind[version] = map[string]*types.Type{}
		}
		if _, f := typesByKindVersion[c.Name.Name]; !f {
			typesByKindVersion[c.Name.Name] = map[string]*types.Type{}
		}
		typesByVersionKind[version][c.Name.Name] = c
		typesByKindVersion[c.Name.Name][version] = c
	}

	return typesByVersionKind, typesByKindVersion, unversionedApiTypes
}

func GetVersionedAndUnversioned(context *generator.Context, group string) (
	sets.String, sets.String) {
	versionedApiTypes := sets.NewString()
	unversionedApiTypes := sets.NewString()
	for _, c := range context.Order {
		if IsVersioned(c, group) && IsApiType(c) {
			// Find versioned types that are API types
			versionedApiTypes.Insert(c.Name.Name)
		}
	}

	for _, c := range context.Order {
		if IsUnversioned(c, group) && versionedApiTypes.Has(c.Name.Name) {
			// The only way to tell if an unversioned type is an api type is by checking if there is a versioned
			// type with the same name
			unversionedApiTypes.Insert(c.Name.Name)
		}
	}
	return versionedApiTypes, unversionedApiTypes
}

func IndexByVersionAndKind(context *generator.Context, group string, versionedSet, unversionedSet sets.String) (
	map[string]map[string]*types.Type, map[string]map[string]*types.Type) {

	typesByVersionKind := map[string]map[string]*types.Type{}
	typesByKindVersion := map[string]map[string]*types.Type{}
	for _, c := range context.Order {
		// Not in the group
		if GetGroup(c) != group {
			continue
		}
		// Not an api type
		if !versionedSet.Has(c.Name.Name) && !unversionedSet.Has(c.Name.Name) {
			continue
		}

		version := unversioned
		if IsVersioned(c, group) {
			version = GetVersion(c, group)
		}
		if _, f := typesByVersionKind[version]; !f {
			typesByVersionKind[version] = map[string]*types.Type{}
		}
		if _, f := typesByKindVersion[c.Name.Name]; !f {
			typesByKindVersion[c.Name.Name] = map[string]*types.Type{}
		}
		typesByVersionKind[version][c.Name.Name] = c
		typesByKindVersion[c.Name.Name][version] = c
	}
	return typesByVersionKind, typesByKindVersion
}

func GetAllVersionedAndUnversioned(context *generator.Context, group string) (
	sets.String, sets.String) {
	versionedApiTypes := sets.NewString()
	unversionedApiTypes := sets.NewString()
	for _, c := range context.Order {
		if IsVersioned(c, group) {
			versionedApiTypes.Insert(c.Name.Name)
		}
	}

	for _, c := range context.Order {
		if IsUnversioned(c, group) && versionedApiTypes.Has(c.Name.Name) {
			unversionedApiTypes.Insert(c.Name.Name)
		}
	}
	return versionedApiTypes, unversionedApiTypes
}

func IndexAllByVersionAndKind(context *generator.Context, group string, versionedSet, unversionedSet sets.String) (
	map[string]map[string]*types.Type, map[string]map[string]*types.Type) {

	typesByVersionKind := map[string]map[string]*types.Type{}
	typesByKindVersion := map[string]map[string]*types.Type{}
	for _, c := range context.Order {
		// Not in the group
		if GetGroup(c) != group {
			continue
		}
		// Not an api type
		if !versionedSet.Has(c.Name.Name) && !unversionedSet.Has(c.Name.Name) {
			continue
		}

		version := unversioned
		if IsVersioned(c, group) {
			version = GetVersion(c, group)
		}
		if _, f := typesByVersionKind[version]; !f {
			typesByVersionKind[version] = map[string]*types.Type{}
		}
		if _, f := typesByKindVersion[c.Name.Name]; !f {
			typesByKindVersion[c.Name.Name] = map[string]*types.Type{}
		}
		typesByVersionKind[version][c.Name.Name] = c
		typesByKindVersion[c.Name.Name][version] = c
	}
	return typesByVersionKind, typesByKindVersion
}

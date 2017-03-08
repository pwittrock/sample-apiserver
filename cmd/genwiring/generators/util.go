package generators

import (
	"path/filepath"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"k8s.io/gengo/types"
	"k8s.io/gengo/generator"
	//"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/util/sets"
)

func IsApiType(t *types.Type) bool {
	for _, c := range t.CommentLines {
		if strings.Contains(c, "+genapi=true") {
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
			return strings.TrimLeft(c, prefix)
		}
	}
	return ""
}

func GetApiTypes(c *generator.Context) []*types.Type {
	types := []*types.Type{}
	for _, o := range c.Order {
		if IsApiType(o) {
			types = append(types, o)
		}
	}
	return types
}

func GetApiTypeNames(c *generator.Context) []string {
	types := []string{}
	for _, o := range GetApiTypes(c) {
		types = append(types, fmt.Sprintf("%s", o.Name.Name))
		types = append(types, fmt.Sprintf("%sList", o.Name.Name))
	}
	return types
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

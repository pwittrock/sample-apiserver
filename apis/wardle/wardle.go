package wardle

import (
	"k8s.io/apimachinery/pkg/runtime"
	w "k8s.io/sample-apiserver/pkg/apis/wardle"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/sample-apiserver/pkg/defaults"
)

type restImpl struct {}

func (r *restImpl) NewFunc() runtime.Object { return &w.Flunder{} }

func (r *restImpl) NewListFunc() runtime.Object { return &w.FlunderList{} }

func (r *restImpl) ObjectNameFunc(obj runtime.Object) (string, error) {
	return obj.(*w.Flunder).Name, nil
}

var SchemeGroupVersionResource = schema.GroupVersionResource{Resource: "flunders", Group: "wardle.k8s.io", Version: "v1alpha1"}

var API = &defaults.API{SchemeGroupVersionResource, &restImpl{}}
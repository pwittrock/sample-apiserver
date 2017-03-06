package wardle

import (
	"k8s.io/apimachinery/pkg/runtime"
	w "k8s.io/sample-apiserver/pkg/apis/wardle"
)

type RestImpl struct {}

func (r *RestImpl) NewFunc() runtime.Object { return &w.Flunder{} }

func (r *RestImpl) NewListFunc() runtime.Object { return &w.FlunderList{} }

func (r *RestImpl) ObjectNameFunc(obj runtime.Object) (string, error) {
	return obj.(*w.Flunder).Name, nil
}



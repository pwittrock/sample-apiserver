package defaults

import "k8s.io/apimachinery/pkg/runtime"

type BasicStorage struct {
	NewListFunc func() runtime.Object
	NewFunc     func() runtime.Object
}

func (BasicStorage) ObjectNameFunc(obj runtime.Object) (string, error) {
	return obj.(BasicResource).GetObjectMeta().Name, nil
}

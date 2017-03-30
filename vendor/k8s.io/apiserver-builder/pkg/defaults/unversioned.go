package defaults

import (
	"k8s.io/apimachinery/pkg/apimachinery/announced"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type UnversionedApi struct {
	NewListFunc func() runtime.Object
	NewFunc     func() runtime.Object
	Group       string
}

type UnversionedBuilder struct {
	SchemeGroupVersion schema.GroupVersion
	SchemeBuilder      runtime.SchemeBuilder
	AddToScheme        announced.SchemeFunc
}

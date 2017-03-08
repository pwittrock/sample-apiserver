package v1alpha1

import (
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/sample-apiserver/apis/wardle"
)

func Convert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec(in *wardle.FlunderSpec, out *FlunderSpec, s conversion.Scope) error {
	autoConvert_wardle_FlunderSpec_To_v1alpha1_FlunderSpec(in, out, s)
	return nil
}

// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package hyrulekingdom

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	builders "k8s.io/apiserver-builder/pkg/builders"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_hyrulekingdom_ZeldasCastle, InType: reflect.TypeOf(&ZeldasCastle{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_hyrulekingdom_ZeldasCastleList, InType: reflect.TypeOf(&ZeldasCastleList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_hyrulekingdom_ZeldasCastleSpec, InType: reflect.TypeOf(&ZeldasCastleSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_hyrulekingdom_ZeldasCastleStatus, InType: reflect.TypeOf(&ZeldasCastleStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_hyrulekingdom_ZeldasCastleStatusStrategy, InType: reflect.TypeOf(&ZeldasCastleStatusStrategy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_hyrulekingdom_ZeldasCastleStrategy, InType: reflect.TypeOf(&ZeldasCastleStrategy{})},
	)
}

func DeepCopy_hyrulekingdom_ZeldasCastle(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ZeldasCastle)
		out := out.(*ZeldasCastle)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		return nil
	}
}

func DeepCopy_hyrulekingdom_ZeldasCastleList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ZeldasCastleList)
		out := out.(*ZeldasCastleList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ZeldasCastle, len(*in))
			for i := range *in {
				if err := DeepCopy_hyrulekingdom_ZeldasCastle(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_hyrulekingdom_ZeldasCastleSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ZeldasCastleSpec)
		out := out.(*ZeldasCastleSpec)
		*out = *in
		return nil
	}
}

func DeepCopy_hyrulekingdom_ZeldasCastleStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ZeldasCastleStatus)
		out := out.(*ZeldasCastleStatus)
		*out = *in
		return nil
	}
}

func DeepCopy_hyrulekingdom_ZeldasCastleStatusStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ZeldasCastleStatusStrategy)
		out := out.(*ZeldasCastleStatusStrategy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.DefaultStatusStorageStrategy); err != nil {
			return err
		} else {
			out.DefaultStatusStorageStrategy = *newVal.(*builders.DefaultStatusStorageStrategy)
		}
		return nil
	}
}

func DeepCopy_hyrulekingdom_ZeldasCastleStrategy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ZeldasCastleStrategy)
		out := out.(*ZeldasCastleStrategy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.DefaultStorageStrategy); err != nil {
			return err
		} else {
			out.DefaultStorageStrategy = *newVal.(*builders.DefaultStorageStrategy)
		}
		return nil
	}
}

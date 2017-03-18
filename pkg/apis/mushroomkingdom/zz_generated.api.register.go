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

// This file was autogenerated by genwiring. Do not edit it manually!

package mushroomkingdom

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver-builder/pkg/builders"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/client-go/pkg/api"
)

var (
	PeachesCastleSingleton = builders.NewUnversionedResource(
		"peachescastles",
		func() runtime.Object { return &PeachesCastle{} },
		func() runtime.Object { return &PeachesCastleList{} },
	)
	PeachesCastleStatusSingleton = builders.NewUnversionedStatus(
		"peachescastles",
		func() runtime.Object { return &PeachesCastle{} },
		func() runtime.Object { return &PeachesCastleList{} },
	)
	ScalePeachesCastleRESTSingleton = builders.NewUnversionedSubresource(
		"peachescastles", "scalecastle",
		func() runtime.Object { return &ScaleCastle{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewUnVersionedApiBuilder("mushroomkingdom.k8s.io").WithKinds(
		PeachesCastleSingleton,
		PeachesCastleStatusSingleton,
		ScalePeachesCastleRESTSingleton,
	)
	SchemeBuilder = ApiVersion.SchemaBuilder
)

type PeachesCastle struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   PeachesCastleSpec
	Status PeachesCastleStatus
}

type ScaleCastle struct {
	metav1.TypeMeta
	metav1.ObjectMeta
}

type PeachesCastleStatus struct {
	Message string
}

type PeachesCastleSpec struct {
	Mushrooms int
}

//
// PeachesCastle Functions and Structs
//
type PeachesCastleStrategy struct {
	builders.DefaultStorageStrategy
}

type PeachesCastleStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

type PeachesCastleList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []PeachesCastle
}

func (PeachesCastle) NewStatus() interface{} {
	return PeachesCastleStatus{}
}

func (pc *PeachesCastle) GetStatus() interface{} {
	return pc.Status
}

func (pc *PeachesCastle) SetStatus(s interface{}) {
	pc.Status = s.(PeachesCastleStatus)
}

func (pc *PeachesCastle) GetSpec() interface{} {
	return pc.Status
}

func (pc *PeachesCastle) SetSpec(s interface{}) {
	pc.Spec = s.(PeachesCastleSpec)
}

func (pc *PeachesCastle) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *PeachesCastle) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc PeachesCastle) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store PeachesCastle.
type PeachesCastleRegistry interface {
	ListPeachesCastles(ctx request.Context, options *internalversion.ListOptions) (*PeachesCastleList, error)
	GetPeachesCastle(ctx request.Context, id string, options *metav1.GetOptions) (*PeachesCastle, error)
	CreatePeachesCastle(ctx request.Context, id *PeachesCastle) (*PeachesCastle, error)
	UpdatePeachesCastle(ctx request.Context, id *PeachesCastle) (*PeachesCastle, error)
	DeletePeachesCastle(ctx request.Context, id string) error
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewPeachesCastleRegistry(s rest.StandardStorage) PeachesCastleRegistry {
	return &storagePeachesCastle{s}
}

// Implement Registry
// storage puts strong typing around storage calls
type storagePeachesCastle struct {
	rest.StandardStorage
}

func (s *storagePeachesCastle) ListPeachesCastles(ctx request.Context, options *internalversion.ListOptions) (*PeachesCastleList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*PeachesCastleList), err
}

func (s *storagePeachesCastle) GetPeachesCastle(ctx request.Context, id string, options *metav1.GetOptions) (*PeachesCastle, error) {
	obj, err := s.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*PeachesCastle), nil
}

func (s *storagePeachesCastle) CreatePeachesCastle(ctx request.Context, object *PeachesCastle) (*PeachesCastle, error) {
	obj, err := s.Create(ctx, object)
	if err != nil {
		return nil, err
	}
	return obj.(*PeachesCastle), nil
}

func (s *storagePeachesCastle) UpdatePeachesCastle(ctx request.Context, object *PeachesCastle) (*PeachesCastle, error) {
	obj, _, err := s.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object, api.Scheme))
	if err != nil {
		return nil, err
	}
	return obj.(*PeachesCastle), nil
}

func (s *storagePeachesCastle) DeletePeachesCastle(ctx request.Context, id string) error {
	_, err := s.Delete(ctx, id, nil)
	return err
}

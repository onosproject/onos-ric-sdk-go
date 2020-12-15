// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package gnmi implements a gnmi server to mock a device with YANG models.
package model

import (
	"errors"
	"fmt"
	"reflect"
	"sort"

	"github.com/openconfig/goyang/pkg/yang"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"

	pb "github.com/openconfig/gnmi/proto/gnmi"
)

// JSONUnmarshaler is the signature of the Unmarshal() function in the GoStruct code generated by openconfig ygot library.
type JSONUnmarshaler func([]byte, ygot.GoStruct, ...ytypes.UnmarshalOpt) error

// GoStructEnumData is the data type to maintain GoStruct enum type.
type GoStructEnumData map[string]map[int64]ygot.EnumDefinition

// Model contains the model data and GoStruct information for the device to config.
type Model struct {
	modelData       []*pb.ModelData
	structRootType  reflect.Type
	schemaTreeRoot  *yang.Entry
	jsonUnmarshaler JSONUnmarshaler
	enumData        GoStructEnumData
}

func (m *Model) NewRootValue() interface{} {
	return reflect.New(m.structRootType.Elem()).Interface()
}

func (m *Model) GetSchemeTreeRoot() *yang.Entry {
	return m.schemaTreeRoot
}

func (m *Model) GetJsonUnmarshaler() JSONUnmarshaler {
	return m.jsonUnmarshaler
}

func (m *Model) GetEnumData() GoStructEnumData {
	return m.enumData
}

// NewModel returns an instance of Model struct.
func NewModel(m []*pb.ModelData, t reflect.Type, r *yang.Entry, f JSONUnmarshaler, e GoStructEnumData) *Model {
	return &Model{
		modelData:       m,
		structRootType:  t,
		schemaTreeRoot:  r,
		jsonUnmarshaler: f,
		enumData:        e,
	}
}

// NewConfigStruct creates a ValidatedGoStruct of this model from jsonConfig. If jsonConfig is nil, creates an empty GoStruct.
func (m *Model) NewConfigStruct(jsonConfig []byte) (ygot.ValidatedGoStruct, error) {
	rootStruct, ok := m.NewRootValue().(ygot.ValidatedGoStruct)
	if !ok {
		return nil, errors.New("root node is not a ygot.ValidatedGoStruct")
	}
	if jsonConfig != nil {
		if err := m.jsonUnmarshaler(jsonConfig, rootStruct); err != nil {
			return nil, err
		}
		if err := rootStruct.Validate(); err != nil {
			return nil, err
		}
	}
	return rootStruct, nil
}

// SupportedModels returns a list of supported models.
func (m *Model) SupportedModels() []string {
	mDesc := make([]string, len(m.modelData))
	for i, m := range m.modelData {
		mDesc[i] = fmt.Sprintf("%s %s", m.Name, m.Version)
	}
	sort.Strings(mDesc)
	return mDesc
}

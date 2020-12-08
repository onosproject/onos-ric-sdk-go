// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package target

import (
	"reflect"

	ricmodelplugin "github.com/onosproject/config-models/modelplugin/ric-1.0.0/modelplugin"
	ricmodelv1 "github.com/onosproject/config-models/modelplugin/ric-1.0.0/ric_1_0_0"
	"github.com/openconfig/ygot/ygot"
)

// ModelType
type ModelType int

const (
	RIC ModelType = iota
)

func (t ModelType) String() string {
	return [...]string{"ric"}[t]
}

// ModelInfo
type ModelInfo struct {
	ModelType ModelType
	Version   string
}

// getRicModel get ric model information
// TODO this should be configured via helm chart later for an xApp or service
func getRicModelV1() *Model {
	v1Models := NewModel(ricmodelplugin.ModelData,
		reflect.TypeOf((*ricmodelv1.Device)(nil)),
		ricmodelv1.SchemaTree["Device"],
		ricmodelv1.Unmarshal,
		map[string]map[int64]ygot.EnumDefinition{},
	)
	return v1Models
}

// GetModel initialize and returns a model
func GetModel(modelInfo ModelInfo) *Model {
	switch modelInfo.ModelType {
	case RIC:
		switch modelInfo.Version {
		case "1.0.0":
			return getRicModelV1()
		}
	}

	return getRicModelV1()
}

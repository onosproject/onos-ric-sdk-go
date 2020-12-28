// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package query

import (
	"github.com/openconfig/ygot/ygot"
	"github.com/tidwall/gjson"
)

// Query
type Query struct {
	Config *ygot.ValidatedGoStruct
}

// GetJsonConfig emits RFC7951 Json format of configuration
func (q *Query) GetJsonConfig() (string, error) {
	JsonConfig, err := ygot.EmitJSON(*q.Config, &ygot.EmitJSONConfig{
		Format: ygot.RFC7951,
		Indent: "  ",
		RFC7951Config: &ygot.RFC7951JSONConfig{
			AppendModuleName: true,
		},
	})

	if err != nil {
		return "", err
	}

	return JsonConfig, nil
}

// GetJsonValue returns the config value based on a given key
func (q *Query) GetJsonValue(key string) (interface{}, error) {
	jsonConfig, err := q.GetJsonConfig()
	if err != nil {
		return nil, err
	}
	value := gjson.Get(jsonConfig, key)
	return value, nil

}

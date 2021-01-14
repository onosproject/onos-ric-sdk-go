// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"fmt"
	"strconv"

	"github.com/onosproject/onos-lib-go/pkg/errors"

	gnmi "github.com/openconfig/gnmi/proto/gnmi"
)

// TODO Add more conversions as needed

// ToUint converts an interface value to uint64
func ToUint64(value interface{}) (uint64, error) {
	switch v := value.(type) {
	case *gnmi.TypedValue:
		return toGnmiTypedValue(value).GetUintVal(), nil

	case float64:
		val, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 64)
		if err != nil {
			return 0, err
		}
		return val, nil
	case uint64:
		val, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 64)
		if err != nil {
			return 0, err
		}
		return val, nil

	default:
		return 0, errors.New(errors.NotSupported, "Not supported type %v", v)
	}

}

// ToFloat converts an interface value to float
func ToFloat(value interface{}) (float32, error) {
	switch v := value.(type) {
	case *gnmi.TypedValue:
		return toGnmiTypedValue(value).GetFloatVal(), nil

	case float32:
		return float32(v), nil

	default:
		return 0, errors.New(errors.NotSupported, "Not supported type %v", v)
	}

}

// ToString converts value to string
func ToString(value interface{}) (string, error) {
	switch v := value.(type) {
	case *gnmi.TypedValue:
		return toGnmiTypedValue(value).GetStringVal(), nil

	case string:
		return value.(string), nil

	default:
		return "", errors.New(errors.NotSupported, "Not supported type %v", v)
	}

}

// ToGnmiTypedValue
func toGnmiTypedValue(value interface{}) *gnmi.TypedValue {
	return value.(*gnmi.TypedValue)
}

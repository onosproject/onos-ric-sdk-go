// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package value

import (
	"fmt"

	typedvalue "github.com/onosproject/onos-api/go/onos/config/change/device"
	"github.com/openconfig/gnmi/proto/gnmi"
)

// GnmiTypedValueToNativeType converts gnmi type based values in to native byte array devicechange
func GnmiTypedValueToNativeType(gnmiTv *gnmi.TypedValue) (*typedvalue.TypedValue, error) {

	switch v := gnmiTv.GetValue().(type) {
	case *gnmi.TypedValue_StringVal:
		return typedvalue.NewTypedValueString(v.StringVal), nil
	case *gnmi.TypedValue_AsciiVal:
		return typedvalue.NewTypedValueString(v.AsciiVal), nil
	case *gnmi.TypedValue_IntVal:
		return typedvalue.NewTypedValueInt64(int(v.IntVal)), nil
	case *gnmi.TypedValue_UintVal:
		return typedvalue.NewTypedValueUint64(uint(v.UintVal)), nil
	case *gnmi.TypedValue_BoolVal:
		return typedvalue.NewTypedValueBool(v.BoolVal), nil
	case *gnmi.TypedValue_BytesVal:
		return typedvalue.NewTypedValueBytes(v.BytesVal), nil
	case *gnmi.TypedValue_DecimalVal:
		return typedvalue.NewTypedValueDecimal64(v.DecimalVal.Digits, v.DecimalVal.Precision), nil
	case *gnmi.TypedValue_FloatVal:
		return typedvalue.NewTypedValueFloat(v.FloatVal), nil
	case *gnmi.TypedValue_LeaflistVal:
		// TODO
		break
	}
	return nil, fmt.Errorf("the type value not yet supported")
}

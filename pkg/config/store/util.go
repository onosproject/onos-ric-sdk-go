// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package store

import (
	"strconv"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/utils"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func put(node interface{}, path interface{}, entry Entry) error {
	keys, err := utils.ParseStringPath(path.(string))
	if err != nil {
		return err
	}
	for i := 0; i < len(keys)-1; i++ {
		node, err = search(node, keys[i])
		if err != nil {
			return err
		}
	}

	lastKey := keys[len(keys)-1]
	switch node.(type) {
	case map[string]interface{}:
		node.(map[string]interface{})[lastKey.(string)] = entry.Value
	}

	return nil
}

func get(node interface{}, path interface{}) (interface{}, error) {
	keys, err := utils.ParseStringPath(path.(string))
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		node, err = search(node, key)
		if err != nil {
			log.Info(err)
			return nil, err
		}
	}

	return node, nil
}

func search(node interface{}, key interface{}) (interface{}, error) {
	switch node.(type) {
	case []interface{}:
		switch key.(type) {
		case map[string]string:
			keys := key.(map[string]string)
			array := node.([]interface{})
			for k, v := range keys {
				for index, value := range array {
					switch vt := value.(type) {
					case map[string]interface{}:
						valueMap := value.(map[string]interface{})
						switch valueMap[k].(type) {
						case string:
							if valueMap[k] == v {
								node = array[index]
							}
						case float64:
							floatValue, _ := strconv.ParseFloat(v, 64)
							if valueMap[k].(float64) == floatValue {
								node = array[index]
							}
						default:
							return nil, errors.New(errors.NotSupported, "type %v is not supported", vt)

						}

					}
				}
			}
		}
	case map[string]interface{}:
		key, ok := key.(string)
		if !ok {
			return nil, errors.New(errors.Unknown, "key is not a string")
		}
		node = node.(map[string]interface{})[key]
	default:
		return nil, errors.New(errors.NotSupported, "node can only be of types map[string]interface{} or []interface{}")
	}

	if node == nil {
		return nil, errors.New(errors.NotFound, "cannot find the node")
	}

	return node, nil
}

// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package store

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func put(node interface{}, path interface{}, entry Entry) error {
	keys, err := breakdownPath(path)
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
	default:

	}

	return nil
}

func get(node interface{}, path interface{}) (interface{}, error) {
	keys, err := breakdownPath(path)
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
func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func breakdownPath(path interface{}) ([]interface{}, error) {
	var keys []interface{}
	switch path := path.(type) {
	case string:
		names := strings.Split(trimFirstRune(path), "/")
		keys = make([]interface{}, len(names))
		for i, v := range names {
			if n, err := strconv.Atoi(v); err == nil {
				keys[i] = n
			} else {
				keys[i] = v
			}
		}
	case []interface{}:
		keys = path
	default:
		return nil, errors.New(errors.NotSupported, "path can only be of type string of []interface{}")
	}
	return keys, nil
}

func search(node interface{}, key interface{}) (interface{}, error) {
	switch node.(type) {
	case []interface{}:
		// TODO in openconfig access to array is represented by name (e.g "/interfaces/interface[name=admin]/config/name")
		//  rather than index (e.g xPath = "/interfaces/interface/0/config/name")
		//  this part should be changed to support that.
		idx, ok := key.(int)
		if !ok {
			return nil, errors.New(errors.Forbidden, "index is not an integer")
		}
		array := node.([]interface{})
		if idx >= len(array) {
			return nil, errors.New(errors.Forbidden, "index out of range")
		} else {
			node = array[idx]
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

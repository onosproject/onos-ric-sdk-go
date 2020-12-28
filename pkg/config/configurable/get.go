// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package configurable

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/openconfig/gnmi/value"
	"github.com/openconfig/ygot/util"

	pb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Get gets config value
func (c *Config) Get(req GetRequest) (GetResponse, error) {
	log.Debugf("Get Callback is called for:%+v", req)
	notifications := make([]*pb.Notification, len(req.Paths))

	jsonType := "IETF"

	if req.EncodingType == pb.Encoding_JSON {
		jsonType = "Internal"
	}

	if req.Paths == nil && req.DataType != "" {
		notifications := make([]*pb.Notification, 1)
		path := pb.Path{}
		// Gets the whole config data tree
		node, err := ytypes.GetNode(c.model.GetSchemeTreeRoot(), *c.config, &path, nil)
		if isNil(node) || err != nil {
			return GetResponse{}, status.Errorf(codes.NotFound, "path %v not found", path.String())
		}

		nodeStruct, _ := node[0].Data.(ygot.GoStruct)
		jsonTree, _ := ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{AppendModuleName: true})

		jsonTree = pruneConfigData(jsonTree, strings.ToLower(req.DataType), &path).(map[string]interface{})
		jsonDump, err := json.Marshal(jsonTree)

		if err != nil {
			msg := fmt.Sprintf("error in marshaling %s JSON tree to bytes: %v", jsonType, err)
			log.Error(msg)
			return GetResponse{}, status.Error(codes.Internal, msg)
		}
		ts := time.Now().UnixNano()

		update := buildUpdate(jsonDump, &path, jsonType)
		notifications[0] = &pb.Notification{
			Timestamp: ts,
			Prefix:    req.Prefix,
			Update:    []*pb.Update{update},
		}
		return GetResponse{
			Notifications: notifications,
		}, nil
	}

	for i, path := range req.Paths {
		// Get schema node for path from config struct.
		fullPath := path
		if req.Prefix != nil {
			fullPath = gnmiFullPath(req.Prefix, path)
		}

		if fullPath.GetElem() == nil {
			return GetResponse{}, status.Error(codes.Unimplemented, "path element is nil")
		}

		nodes, err := ytypes.GetNode(c.model.GetSchemeTreeRoot(), *c.config, fullPath)
		if len(nodes) == 0 || err != nil || util.IsValueNil(nodes[0].Data) {
			return GetResponse{}, status.Errorf(codes.NotFound, "path %v not found: %v", fullPath, err)
		}
		node := nodes[0].Data

		ts := time.Now().UnixNano()

		nodeStruct, ok := node.(ygot.GoStruct)
		dataTypeFlag := false
		// Return leaf node.
		if !ok {
			elements := fullPath.GetElem()
			dataTypeString := strings.ToLower(req.DataType)
			if strings.Compare(dataTypeString, "all") == 0 {
				dataTypeFlag = true
			} else {
				for _, elem := range elements {
					if strings.Compare(dataTypeString, elem.GetName()) == 0 {
						dataTypeFlag = true
						break
					}

				}
			}
			if !dataTypeFlag {
				return GetResponse{}, status.Error(codes.Internal, "The requested dataType is not valid")
			}
			var val *pb.TypedValue
			switch kind := reflect.ValueOf(node).Kind(); kind {
			case reflect.Ptr, reflect.Interface:
				var err error
				val, err = value.FromScalar(reflect.ValueOf(node).Elem().Interface())
				if err != nil {
					msg := fmt.Sprintf("leaf node %v does not contain a scalar type value: %v", path, err)
					log.Error(msg)
					return GetResponse{}, status.Error(codes.Internal, msg)
				}

			case reflect.Slice:
				var err error
				switch kind := reflect.ValueOf(node).Kind(); kind {
				case reflect.Int64:
					enumMap, ok := c.model.GetEnumData()[reflect.TypeOf(node).Name()]
					if !ok {
						return GetResponse{}, status.Error(codes.Internal, "not a GoStruct enumeration type")
					}
					val = &pb.TypedValue{
						Value: &pb.TypedValue_StringVal{
							StringVal: enumMap[reflect.ValueOf(nodes[0].Data).Int()].Name,
						},
					}
				default:
					if !reflect.ValueOf(node).Elem().IsValid() {
						return GetResponse{}, status.Errorf(codes.NotFound, "path %v not found", path)
					}
					val, err = value.FromScalar(reflect.ValueOf(node).Elem().Interface())
					if err != nil {
						msg := fmt.Sprintf("leaf node %v does not contain a scalar type value: %v", path, err)
						log.Error(msg)
						return GetResponse{}, status.Error(codes.Internal, msg)
					}
				}

			default:
				return GetResponse{}, status.Errorf(codes.Internal, "unexpected kind of leaf node type: %v %v", node, kind)
			}

			update := &pb.Update{Path: path, Val: val}
			notifications[i] = &pb.Notification{
				Timestamp: ts,
				Prefix:    req.Prefix,
				Update:    []*pb.Update{update},
			}
			continue
		}

		var jsonTree map[string]interface{}

		if reflect.ValueOf(nodeStruct).Pointer() == 0 {
			return GetResponse{}, status.Error(codes.NotFound, "value is 0")

		}
		jsonTree, err = jsonEncoder(jsonType, nodeStruct)
		jsonTree = pruneConfigData(jsonTree, strings.ToLower(req.DataType), fullPath).(map[string]interface{})
		if err != nil {
			msg := fmt.Sprintf("error in constructing %s JSON tree from requested node: %v", jsonType, err)
			log.Error(msg)
			return GetResponse{}, status.Error(codes.Internal, msg)
		}

		jsonDump, err := json.Marshal(jsonTree)
		if err != nil {
			msg := fmt.Sprintf("error in marshaling %s JSON tree to bytes: %v", jsonType, err)
			log.Error(msg)
			return GetResponse{}, status.Error(codes.Internal, msg)
		}

		update := buildUpdate(jsonDump, path, jsonType)
		notifications[i] = &pb.Notification{
			Timestamp: ts,
			Prefix:    req.Prefix,
			Update:    []*pb.Update{update},
		}
	}

	return GetResponse{
		Notifications: notifications,
	}, nil
}

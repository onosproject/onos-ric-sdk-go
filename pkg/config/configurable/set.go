// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package configurable

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/openconfig/gnmi/value"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/ygot/ytypes"

	pb "github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// doDelete deletes the path from the json tree if the path exists. If success,
// it calls the callback function to apply the change to the device hardware.
func (c *Config) doDelete(jsonTree map[string]interface{}, prefix, path *pb.Path) (*pb.UpdateResult, error) {
	// Update json tree of the device config
	var curNode interface{} = jsonTree
	pathDeleted := false
	fullPath := gnmiFullPath(prefix, path)
	schema := c.model.GetSchemeTreeRoot()
	for i, elem := range fullPath.Elem { // Delete sub-tree or leaf node.
		node, ok := curNode.(map[string]interface{})
		if !ok {
			break
		}

		// Delete node
		if i == len(fullPath.Elem)-1 {
			if elem.GetKey() == nil {
				delete(node, elem.Name)
				pathDeleted = true
				break
			}
			pathDeleted = deleteKeyedListEntry(node, elem)
			break
		}

		if curNode, schema = getChildNode(node, schema, elem, false); curNode == nil {
			break
		}
	}
	if reflect.DeepEqual(fullPath, pbRootPath) { // Delete root
		for k := range jsonTree {
			delete(jsonTree, k)
		}
	}

	// Apply the validated operation to the config tree and device.
	if pathDeleted {
		newConfig, err := c.toGoStruct(jsonTree)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		*c.config = newConfig
	}
	return &pb.UpdateResult{
		Path: path,
		Op:   pb.UpdateResult_DELETE,
	}, nil
}

// doReplaceOrUpdate validates the replace or update operation to be applied to
// the device, modifies the json tree of the config struct, then calls the
// callback function to apply the operation to the device hardware.
func (c *Config) doReplaceOrUpdate(jsonTree map[string]interface{}, op pb.UpdateResult_Operation, prefix, path *pb.Path, val *pb.TypedValue) (*pb.UpdateResult, error) {
	// Validate the operation.
	fullPath := gnmiFullPath(prefix, path)
	emptyNode, _, err := ytypes.GetOrCreateNode(c.model.GetSchemeTreeRoot(), c.model.NewRootValue(), fullPath)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "path %v is not found in the config structure: %v", fullPath, err)
	}
	var nodeVal interface{}
	nodeStruct, ok := emptyNode.(ygot.ValidatedGoStruct)
	if ok {
		jsonUnmarshaller := c.model.GetJsonUnmarshaler()
		if err := jsonUnmarshaller(val.GetJsonIetfVal(), nodeStruct); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "unmarshaling json data to config struct fails: %v", err)
		}
		if err := nodeStruct.Validate(); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "config data validation fails: %v", err)
		}
		var err error
		if nodeVal, err = ygot.ConstructIETFJSON(nodeStruct, &ygot.RFC7951JSONConfig{}); err != nil {
			msg := fmt.Sprintf("error in constructing IETF JSON tree from config struct: %v", err)
			log.Error(msg)
			return nil, status.Error(codes.Internal, msg)
		}
	} else {
		var err error
		if nodeVal, err = value.ToScalar(val); err != nil {
			return nil, status.Errorf(codes.Internal, "cannot convert leaf node to scalar type: %v", err)
		}
	}

	// Update json tree of the device config.
	var curNode interface{} = jsonTree
	schema := c.model.GetSchemeTreeRoot()
	for i, elem := range fullPath.Elem {
		switch node := curNode.(type) {
		case map[string]interface{}:
			// Set node value.
			if i == len(fullPath.Elem)-1 {
				if elem.GetKey() == nil {
					if grpcStatusError := setPathWithoutAttribute(op, node, elem, nodeVal); grpcStatusError != nil {
						return nil, grpcStatusError
					}
					break
				}
				if grpcStatusError := setPathWithAttribute(op, node, elem, nodeVal); grpcStatusError != nil {
					return nil, grpcStatusError
				}
				break
			}

			if curNode, schema = getChildNode(node, schema, elem, true); curNode == nil {
				return nil, status.Errorf(codes.NotFound, "path elem not found: %v", elem)
			}
		case []interface{}:
			return nil, status.Errorf(codes.NotFound, "incompatible path elem: %v", elem)
		default:
			return nil, status.Errorf(codes.Internal, "wrong node type: %T", curNode)
		}
	}
	if reflect.DeepEqual(fullPath, pbRootPath) { // Replace/Update root.
		if op == pb.UpdateResult_UPDATE {
			return nil, status.Error(codes.Unimplemented, "update the root of config tree is unsupported")
		}
		nodeValAsTree, ok := nodeVal.(map[string]interface{})
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "expect a tree to replace the root, got a scalar value: %T", nodeVal)
		}
		for k := range jsonTree {
			delete(jsonTree, k)
		}
		for k, v := range nodeValAsTree {
			jsonTree[k] = v
		}
	}

	return &pb.UpdateResult{
		Path: path,
		Op:   op,
	}, nil
}

// Set
func (c *Config) Set(req SetRequest) (SetResponse, error) {
	log.Debugf("Set Callback is called for:%+v", req)
	jsonTree, err := ygot.ConstructIETFJSON(*c.config, &ygot.RFC7951JSONConfig{})
	if err != nil {
		return SetResponse{}, err
	}

	var results []*pb.UpdateResult

	for _, path := range req.DeletePaths {
		res, grpcStatusError := c.doDelete(jsonTree, req.Prefix, path)
		if grpcStatusError != nil {
			return SetResponse{}, grpcStatusError
		}
		results = append(results, res)
	}
	for _, upd := range req.ReplacePaths {
		res, grpcStatusError := c.doReplaceOrUpdate(jsonTree, pb.UpdateResult_REPLACE, req.Prefix, upd.GetPath(), upd.GetVal())
		if grpcStatusError != nil {
			return SetResponse{}, grpcStatusError
		}
		results = append(results, res)
	}
	for _, upd := range req.UpdatePaths {
		res, grpcStatusError := c.doReplaceOrUpdate(jsonTree, pb.UpdateResult_UPDATE, req.Prefix, upd.GetPath(), upd.GetVal())
		if grpcStatusError != nil {
			return SetResponse{}, grpcStatusError
		}
		results = append(results, res)
	}

	jsonDump, err := json.Marshal(jsonTree)
	if err != nil {
		return SetResponse{}, err
	}
	rootStruct, err := c.model.NewConfigStruct(jsonDump)
	if err != nil {
		return SetResponse{}, err
	}

	*c.config = rootStruct
	return SetResponse{
		Results: results,
	}, nil
}

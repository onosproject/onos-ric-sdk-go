// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package configurable

import (
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/utils"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

func (c *Config) Set(req SetRequest) (SetResponse, error) {
	log.Debugf("Set Callback is called for:%+v", req)
	var results []*pb.UpdateResult
	for _, upd := range req.ReplacePaths {
		fullPath := utils.GnmiFullPath(req.Prefix, upd.Path)
		xpath := utils.ToXPath(fullPath)

		entry := Entry{
			Key:       xpath,
			Value:     upd.GetVal(),
			EventType: pb.UpdateResult_REPLACE.String(),
		}
		err := c.config.put(xpath, entry)
		if err != nil {
			return SetResponse{}, err
		}

		update := &pb.UpdateResult{
			Op:   pb.UpdateResult_REPLACE,
			Path: upd.Path,
		}
		results = append(results, update)
	}

	for _, upd := range req.UpdatePaths {
		fullPath := utils.GnmiFullPath(req.Prefix, upd.Path)
		xpath := utils.ToXPath(fullPath)
		entry := Entry{
			Key:       xpath,
			Value:     upd.GetVal(),
			EventType: pb.UpdateResult_UPDATE.String(),
		}
		err := c.config.put(xpath, entry)
		if err != nil {
			return SetResponse{}, err
		}

		update := &pb.UpdateResult{
			Op:   pb.UpdateResult_UPDATE,
			Path: upd.Path,
		}
		results = append(results, update)

	}

	return SetResponse{
		Results: results,
	}, nil
}

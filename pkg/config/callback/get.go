// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package callback

import (
	"encoding/json"
	"strings"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/utils"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

func buildUpdate(b []byte, path *pb.Path, valType string) *pb.Update {
	var update *pb.Update

	if strings.Compare(valType, "Internal") == 0 {
		update = &pb.Update{Path: path, Val: &pb.TypedValue{Value: &pb.TypedValue_JsonVal{JsonVal: b}}}
		return update
	}
	update = &pb.Update{Path: path, Val: &pb.TypedValue{Value: &pb.TypedValue_JsonIetfVal{JsonIetfVal: b}}}

	return update
}

func (c *Config) Get(req configurable.GetRequest) (configurable.GetResponse, error) {
	log.Debugf("Get Callback is called for:%+v", req)
	notifications := make([]*pb.Notification, len(req.Paths))

	for i, path := range req.Paths {
		fullPath := utils.GnmiFullPath(req.Prefix, path)
		xPath := utils.ToXPath(fullPath)
		entry, err := c.config.Get(xPath)
		if err != nil {
			return configurable.GetResponse{}, err
		}

		jsonDump, err := json.Marshal(entry.Value)
		if err != nil {
			return configurable.GetResponse{}, err
		}

		update := buildUpdate(jsonDump, path, "IETF")
		notifications[i] = &pb.Notification{
			Prefix: req.Prefix,
			Update: []*pb.Update{update},
		}

	}

	return configurable.GetResponse{
		Notifications: notifications,
	}, nil

}

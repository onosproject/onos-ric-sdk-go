// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package gnmi

import (
	"context"

	"github.com/onosproject/onos-ric-sdk-go/pkg/gnmi/value"

	"github.com/onosproject/onos-ric-sdk-go/pkg/gnmi/path"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	pb "github.com/openconfig/gnmi/proto/gnmi"
)

var log = logging.GetLogger("sdk", "gnmi")

var _ Configurable = &Config{}

type WatchRequest struct {
	Paths []path.Path
}

type GetRequest struct {
	Paths []path.Path
}

type GetResponse struct {
	Response map[path.Path]interface{}
}

type Event struct {
}

// Get returns configuration values based on a given a set of paths
func (c *Config) Get(req GetRequest) (GetResponse, error) {
	log.Infof("Received Get operation request on paths: %v", req.Paths)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	encoding, ok := pb.Encoding_value["JSON_IETF"]
	if !ok {
		return GetResponse{}, nil
	}

	pbPathList, err := path.ParsePaths(req.Paths)
	if err != nil {
		return GetResponse{}, err
	}

	getRequest := &pb.GetRequest{
		Encoding: pb.Encoding(encoding),
		Path:     pbPathList,
	}

	getResponse, err := c.server.Get(ctx, getRequest)
	if err != nil {
		return GetResponse{}, err
	}

	var response GetResponse
	response.Response = make(map[path.Path]interface{})

	for _, notification := range getResponse.Notification {
		for index, update := range notification.Update {
			if pbPathList[index].String() == update.Path.String() {
				newPath := path.Path{
					Value: req.Paths[index].Value,
				}
				val, _ := value.GnmiTypedValueToNativeType(update.Val)
				response.Response[newPath] = val.ValueToString()
			}
		}
	}

	return response, nil

}

// Configurable
type Configurable interface {
	// Get gets a configuration value based on a given path
	Get(req GetRequest) (GetResponse, error)
}

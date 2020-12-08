// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package gnmi

import (
	"context"

	"github.com/onosproject/onos-ric-sdk-go/pkg/gnmi/value"

	"github.com/onosproject/onos-ric-sdk-go/pkg/gnmi/path"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/onosproject/onos-ric-sdk-go/pkg/gnmi/target"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

var log = logging.GetLogger("sdk", "gnmi")

var _ Configurable = &Config{}

type Config struct {
	server    *target.Server
	ModelInfo target.ModelInfo
}

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

func RegisterConfigurable(c *Config) target.GnmiService {
	service := target.NewService(c.ModelInfo)
	c.server = service.GetServer()

	return service
}

// Watch watches configuration changes
func (c *Config) Watch(ctx context.Context, req WatchRequest, ch chan<- Event) error {
	updateChannel := make(chan *pb.Update)
	pbPathList, err := path.ParsePaths(req.Paths)
	if err != nil {
		return err
	}
	go c.server.WatchConfigUpdates(updateChannel)
	go func() {
		for update := range updateChannel {
			for _, updatedPath := range pbPathList {
				if update.Path.String() == updatedPath.String() {
					log.Infof("Send update", updatedPath.String(), update.Path.String())
					ch <- Event{}
				}
			}
		}
	}()
	return nil

}

// Get returns configuration values based on a given a set of paths
func (c *Config) Get(ctx context.Context, req GetRequest) (GetResponse, error) {
	log.Infof("Received Get operation request on paths: %v", req.Paths)
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
	Get(ctx context.Context, req GetRequest) (GetResponse, error)
	// Watch watches configuration changes
	Watch(ctx context.Context, req WatchRequest, ch chan<- Event) error
}

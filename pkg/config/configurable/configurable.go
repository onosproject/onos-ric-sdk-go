// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package configurable

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/store"
	"github.com/openconfig/gnmi/proto/gnmi"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

var log = logging.GetLogger("configurable")

// GetRequest
type GetRequest struct {
	Paths        []*gnmi.Path
	Prefix       *gnmi.Path
	EncodingType gnmi.Encoding
	DataType     string
}

// GetResponse
type GetResponse struct {
	Notifications []*pb.Notification
}

// SetRequest
type SetRequest struct {
	DeletePaths  []*gnmi.Path
	ReplacePaths []*gnmi.Update
	UpdatePaths  []*gnmi.Update
	Prefix       *gnmi.Path
}

// SetResponse
type SetResponse struct {
	Results []*pb.UpdateResult
}

// Config
type Config struct {
	config *store.ConfigStore
}

//
func (c *Config) InitConfig(config *store.ConfigStore) {
	c.config = config

}

// Configurable interface between gnmi agent and app
type Configurable interface {
	Get(GetRequest) (GetResponse, error)
	Set(SetRequest) (SetResponse, error)
}

var _ Configurable = &Config{}

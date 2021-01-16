// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package configurable

import (
	"github.com/openconfig/gnmi/proto/gnmi"
)

// GetRequest
type GetRequest struct {
	Paths        []*gnmi.Path
	Prefix       *gnmi.Path
	EncodingType gnmi.Encoding
	DataType     string
}

// GetResponse
type GetResponse struct {
	Notifications []*gnmi.Notification
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
	Results []*gnmi.UpdateResult
}

// Configurable interface between gnmi agent and app
type Configurable interface {
	Get(GetRequest) (GetResponse, error)
	Set(SetRequest) (SetResponse, error)
}

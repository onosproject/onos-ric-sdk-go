// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package gnmi implements a gnmi server to mock a device with YANG models.
package target

import (
	"sync"

	pb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ygot/ygot"
)

// ConfigCallback is the signature of the function to apply a validated config to the physical device.
type ConfigCallback func(ygot.ValidatedGoStruct) error

var (
	pbRootPath         = &pb.Path{}
	supportedEncodings = []pb.Encoding{pb.Encoding_JSON, pb.Encoding_JSON_IETF}
	dataTypes          = []string{"config", "state", "operational", "all"}
)

// Server struct maintains the data structure for device config and implements the interface of gnmi server. It supports Capabilities, Get, and Set APIs.
type Server struct {
	model        *Model
	callback     ConfigCallback
	config       ygot.ValidatedGoStruct
	configUpdate chan *pb.Update
	mu           sync.RWMutex // mu is the RW lock to protect the access to config
	subscribers  map[string]*streamClient
}

var (
	lowestSampleInterval uint64 = 5000000000 // 5000000000 nanoseconds
)

type streamClient struct {
	//target         string
	sr     *pb.SubscribeRequest
	stream pb.GNMI_SubscribeServer
	//errChan        chan error
	UpdateChan     chan *pb.Update
	sampleInterval uint64
}

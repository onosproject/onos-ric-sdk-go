// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package gnmi implements a gnmi server to mock a device with YANG models.
package agent

import (
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
	pb "github.com/openconfig/gnmi/proto/gnmi"
	"golang.org/x/net/context"
)

// Get implements the Get RPC in gNMI spec.
func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	log.Debugf("Processing Get Request: %+v", req)
	s.mu.RLock()
	defer s.mu.RUnlock()
	getResponse, err := s.configurable.Get(configurable.GetRequest{
		Paths:        req.Path,
		EncodingType: req.GetEncoding(),
		Prefix:       req.GetPrefix(),
		DataType:     req.Type.String(),
	})

	if err != nil {
		return &pb.GetResponse{}, err
	}

	resp := &pb.GetResponse{Notification: getResponse.Notifications}
	return resp, nil
}

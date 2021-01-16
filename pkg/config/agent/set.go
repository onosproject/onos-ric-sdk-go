// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package agent implements a gnmi server
package agent

import (
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
	pb "github.com/openconfig/gnmi/proto/gnmi"
	"golang.org/x/net/context"
)

// Set implements the Set RPC in gNMI spec.
func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	log.Debugf("Processing Set Request:%+v", req)
	s.mu.Lock()
	defer s.mu.Unlock()

	setReq := configurable.SetRequest{
		DeletePaths:  req.GetDelete(),
		UpdatePaths:  req.GetUpdate(),
		ReplacePaths: req.GetReplace(),
		Prefix:       req.GetPrefix(),
	}
	resp, err := s.configurable.Set(setReq)
	if err != nil {
		return &pb.SetResponse{}, err
	}

	setResponse := &pb.SetResponse{
		Prefix:   req.GetPrefix(),
		Response: resp.Results,
	}

	return setResponse, nil
}

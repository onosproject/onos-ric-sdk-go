// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// Package agent implements a gnmi server to mock a device with YANG models.
package agent

import (
	pb "github.com/openconfig/gnmi/proto/gnmi"
	"golang.org/x/net/context"
)

// Capabilities returns supported encodings and supported models.
func (s *Server) Capabilities(ctx context.Context, req *pb.CapabilityRequest) (*pb.CapabilityResponse, error) {
	return &pb.CapabilityResponse{}, nil
}

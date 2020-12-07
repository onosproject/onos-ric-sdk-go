// SPDX-FileCopyrightText: ${year}-present Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

// Package gnmi implements a gnmi server to mock a device with YANG models.
package target

import (
	pb "github.com/openconfig/gnmi/proto/gnmi"
	"golang.org/x/net/context"
)

// Capabilities returns supported encodings and supported models.
func (s *Server) Capabilities(ctx context.Context, req *pb.CapabilityRequest) (*pb.CapabilityResponse, error) {
	// TODO fix this
	/*ver, err := getGNMIServiceVersion()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error in getting gnmi service version: %v", err)
	}*/
	return &pb.CapabilityResponse{
		SupportedModels:    s.model.modelData,
		SupportedEncodings: supportedEncodings,
		//GNMIVersion:        *ver,
	}, nil
}

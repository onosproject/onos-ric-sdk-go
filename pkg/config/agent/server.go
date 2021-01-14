// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package gnmi implements a gnmi server to mock a device with YANG models.
package agent

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
	api "github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log = logging.GetLogger("gnmi", "agent")

// NewService creates a new gnmi service
func NewService(configurable configurable.Configurable) GnmiService {
	return newService(configurable)
}

type GnmiService interface {
	northbound.Service
}

func newService(configurable configurable.Configurable) GnmiService {
	server := &Server{
		configurable: configurable,
	}

	return &Service{
		server: server,
	}
}

// Service is a Service implementation for gnmi service.
type Service struct {
	server *Server
}

// Register registers the Service with the gRPC server.
func (s *Service) Register(r *grpc.Server) {
	server := s.server
	api.RegisterGNMIServer(r, server)
	reflection.Register(r)
}

var _ GnmiService = &Service{}

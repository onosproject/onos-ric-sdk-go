// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package gnmi implements a gnmi server to mock a device with YANG models.
package target

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	api "github.com/openconfig/gnmi/proto/gnmi"
	pb "github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var log = logging.GetLogger("gnmi", "target")

// NewService creates a new gnmi service
func NewService(info ModelInfo) GnmiService {
	byteValue, err := load()
	if err != nil {
		log.Error("Failed to read initial config", err)
	}
	return newService(GetModel(info), byteValue, nil)
}

type GnmiService interface {
	northbound.Service
	GetServer() *Server
}

func newService(model *Model, config []byte, callback ConfigCallback) GnmiService {
	rootStruct, err := model.NewConfigStruct(config)
	if err != nil {
		log.Errorf("initial config cannot be initialized", err)
	}
	server := &Server{
		model:        model,
		config:       rootStruct,
		callback:     callback,
		configUpdate: make(chan *api.Update),
	}

	return &Service{
		model:    model,
		config:   config,
		callback: callback,
		server:   server,
	}
}

// Service is a Service implementation for gnmi service.
type Service struct {
	model    *Model
	config   []byte
	callback ConfigCallback
	server   *Server
}

func (s *Service) GetServer() *Server {
	return s.server
}

// Register registers the Service with the gRPC server.
func (s *Service) Register(r *grpc.Server) {
	server := s.server
	if server.config != nil && server.callback != nil {
		if err := server.callback(s.server.config); err != nil {
			return
		}
	}
	server.subscribers = make(map[string]*streamClient)
	server.configUpdate = make(chan *pb.Update)
	api.RegisterGNMIServer(r, server)
	reflection.Register(r)
}

var _ GnmiService = &Service{}

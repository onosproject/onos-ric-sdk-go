// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package gnmi implements a gnmi server to mock a device with YANG models.
package gnmi

import (
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	api "github.com/openconfig/gnmi/proto/gnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewService creates a new gnmi service
func NewService(info ModelInfo) northbound.Service {
	return newService(GetModel(info), nil, nil)
}

func newService(model *Model, config []byte, callback ConfigCallback) northbound.Service {
	return &Service{
		model:    model,
		config:   config,
		callback: callback,
	}
}

// Service is a Service implementation for gnmi service.
type Service struct {
	model    *Model
	config   []byte
	callback ConfigCallback
}

// Register registers the Service with the gRPC server.
func (s *Service) Register(r *grpc.Server) {
	rootStruct, err := s.model.NewConfigStruct(s.config)
	if err != nil {
		// TODO log the error
		return
	}

	server := &Server{
		model:        s.model,
		config:       rootStruct,
		callback:     s.callback,
		configUpdate: make(chan *api.Update),
	}
	if server.config != nil && server.callback != nil {
		if err := server.callback(rootStruct); err != nil {
			return
		}
	}
	server.subscribers = make(map[string]*streamClient)
	api.RegisterGNMIServer(r, server)
	reflection.Register(r)
}

var _ northbound.Service = &Service{}

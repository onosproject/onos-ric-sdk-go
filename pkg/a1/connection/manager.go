// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package a1connection

import (
	"context"
	gogotypes "github.com/gogo/protobuf/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	a1endpoint "github.com/onosproject/onos-ric-sdk-go/pkg/a1/endpoint"
	"github.com/onosproject/onos-ric-sdk-go/pkg/topo"
	"github.com/onosproject/onos-ric-sdk-go/pkg/utils"
)

var log = logging.GetLogger("a1", "manager")

// NewManager creates a new A1 manager
func NewManager(caPath string, keyPath string, certPath string, grpcPort int) (*Manager, error) {
	topoClient, err := topo.NewClient()
	if err != nil {
		return nil, err
	}
	return &Manager{
		id:         utils.GetXappTopoID(),
		server:     a1endpoint.NewServer(caPath, keyPath, certPath, grpcPort),
		topoClient: topoClient,
	}, nil
}

// Manager is a struct of A1 interface
type Manager struct {
	id             topoapi.ID
	server         a1endpoint.Server
	topoClient     topo.Client
}

// Start inits and starts A1 server
func (m *Manager) Start(ctx context.Context) {
	go func(ctx context.Context) {
		log.Infof("Start (or restart) A1 connection manager")
		err := m.AddXAppElementOnTopo(ctx)
		if err != nil {
			log.Warn(err)
			log.Warn()
		}
	}(ctx)
}

// GetID returns ID
func (m *Manager) GetID() topoapi.ID {
	return m.id
}

// AddXAppElementOnTopo adds XApp type on topo
func (m *Manager) AddXAppElementOnTopo(ctx context.Context) error {
	object := &topoapi.Object{
		ID:   m.id,
		Type: topoapi.Object_ENTITY,
		Obj: &topoapi.Object_Entity{
			Entity: &topoapi.Entity{
				KindID: topoapi.XAPP,
			},
		},
		Aspects: make(map[string]*gogotypes.Any),
		Labels:  map[string]string{},
	}
	interfaces := make([]*topoapi.Interface, 1)
	interfaces[0] = &topoapi.Interface{
		IP:   env.GetPodIP(),
		Port: uint32(m.server.GRPCPort),
		Type: topoapi.Interface_INTERFACE_A1_XAPP,
	}

	aspect := &topoapi.XAppInfo{
		Interfaces: interfaces,
	}

	err := object.SetAspect(aspect)
	if err != nil {
		return err
	}
	obj, err := m.topoClient.Get(ctx, m.id)
	if err != nil {
		err = m.topoClient.Create(ctx, object)
		return err
	}
	log.Warn("Update topo A1 XApp entity to have the latest aspects since there is already exist")
	err = obj.SetAspect(aspect)
	if err != nil {
		return err
	}
	err = m.topoClient.Update(ctx, obj)
	return err
}

// DeleteXAppElementOnTopo removes all aspects on topo
func (m *Manager) DeleteXAppElementOnTopo(ctx context.Context) error {
	obj, err := m.topoClient.Get(ctx, m.id)
	if err != nil {
		return err
	}
	aspect := &topoapi.XAppInfo{}
	err = obj.SetAspect(aspect)
	if err != nil {
		return err
	}
	err = m.topoClient.Update(ctx, obj)
	if err != nil && !errors.IsAlreadyExists(err) {
		return err
	}
	return nil
}

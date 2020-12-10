// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package topo contains facilities for RIC applications to monitor topology for changes
package topo

import (
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/southbound"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

// ServiceConfig is a topo service configuration
type ServiceConfig struct {
	// Host is the service host
	Host string
	// Port is the service port
	Port int

	// CAPath is path to certificate authority key
	CAPath string
	// KeyPath is path to certificate cert key
	KeyPath string
	// CertPath is path to certificate server certificate
	CertPath string
	// Insecure indicates insecure connection
	Insecure bool
}

// GetHost gets the service host
func (c ServiceConfig) GetHost() string {
	return c.Host
}

// GetPort gets the service port
func (c ServiceConfig) GetPort() int {
	if c.Port == 0 {
		return defaultServicePort
	}
	return c.Port
}

// Client is a topo client
type Client interface {
	// Watch provides a simple facility for the application to watch for changes in the topology
	Watch(ch chan<- topoapi.Event, options ...FilterOption) error
}

const defaultServiceHost = "onos-topo"
const defaultServicePort = 5150

// Current topology service configuration; starts with reasonable defaults
var DefaultServiceConfig = ServiceConfig{
	Host:     defaultServiceHost,
	Port:     defaultServicePort,
	Insecure: true,
}

// NewClient creates a new E2 client
func NewClient(config ServiceConfig) (Client, error) {
	return &topoClient{
		config: config,
	}, nil
}

// e2Client is the default E2 client implementation
type topoClient struct {
	config ServiceConfig
}

// newClient creates a new topo client
func newTopoClient(cfg ServiceConfig) (topoapi.TopoClient, error) {
	opts := []grpc.DialOption{
		grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(100 * time.Millisecond)),
	}
	if cfg.Insecure {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := getTopoConn("onos-topo", opts...)
	if err != nil {
		stat, ok := status.FromError(err)
		if ok {
			log.Error("Unable to connect to topology service", err)
			return nil, errors.FromStatus(stat)
		}
		return nil, err
	}
	return topoapi.NewTopoClient(conn), nil
}

// getTopoConn gets a gRPC connection to the topology service
func getTopoConn(topoEndpoint string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return grpc.Dial(topoEndpoint, opts...)
}

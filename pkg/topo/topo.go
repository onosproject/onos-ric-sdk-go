// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	"context"
	"fmt"

	"github.com/onosproject/onos-ric-sdk-go/pkg/topo/options"

	"github.com/onosproject/onos-ric-sdk-go/pkg/utils/creds"
	"google.golang.org/grpc/credentials"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"time"

	client "github.com/onosproject/onos-ric-sdk-go/pkg/topo/client"

	"github.com/onosproject/onos-ric-sdk-go/pkg/topo/connection"

	"github.com/onosproject/onos-lib-go/pkg/southbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("topo")

const (
	defaultServiceHost = "onos-topo"
	defaultServicePort = 5150
)

// Config topo client config
type Config struct {
	TopoService ServiceConfig
}

// ServiceConfig is a topo service configuration
type ServiceConfig struct {
	// Host is the service host
	Host string
	// Port is the service port
	Port int
	// Insecure indicates insecure connection
	Insecure bool
}

// GetHost gets the service host
func (c ServiceConfig) GetHost() string {
	if c.Host == "" {
		c.Host = defaultServiceHost
	}
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
	// Get gets a topo object with a given ID
	Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error)

	// Watch provides a simple facility for the application to watch for changes in the topology
	Watch(ctx context.Context, ch chan<- topoapi.Event, opts ...options.Option) error

	// List of topo objects
	List(ctx context.Context, opts ...options.Option) ([]topoapi.Object, error)
}

// NewClient creates a new E2 client
func NewClient(config Config) (Client, error) {
	opts := []grpc.DialOption{
		grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(100 * time.Millisecond)),
	}
	if config.TopoService.Insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		tlsConfig, err := creds.GetClientCredentials()
		if err != nil {
			log.Warn(err)
			return nil, err
		}

		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	}
	conns := connection.NewManager()
	topoEndpointAddr := fmt.Sprintf("%s:%d", config.TopoService.GetHost(), config.TopoService.GetPort())
	conn, err := conns.Connect(topoEndpointAddr, opts...)
	if err != nil {
		return nil, err
	}

	cl, err := client.NewClient(conn)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	return &topoClient{
		config: config,
		client: cl,
	}, nil
}

// topoClient is the topo client implementation
type topoClient struct {
	config Config
	client client.Client
}

func (t *topoClient) List(ctx context.Context, opts ...options.Option) ([]topoapi.Object, error) {
	options := options.Options{}

	for _, opt := range opts {
		opt.Apply(&options)
	}

	return t.client.List(ctx, options.List)

}

func (t *topoClient) Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error) {
	return t.client.Get(ctx, id)
}

func (t *topoClient) Watch(ctx context.Context, ch chan<- topoapi.Event, opts ...options.Option) error {
	options := options.Options{}

	for _, opt := range opts {
		opt.Apply(&options)
	}

	return t.client.Watch(ctx, ch, options.Watch)
}

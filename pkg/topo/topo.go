// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	"context"

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

// Client is a topo client
type Client interface {

	// Update updates a topo object
	Update(ctx context.Context, object *topoapi.Object) error

	// Get gets a topo object with a given ID
	Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error)

	// Watch provides a simple facility for the application to watch for changes in the topology
	Watch(ctx context.Context, ch chan<- topoapi.Event, opts ...options.Option) error

	// List of topo objects
	List(ctx context.Context, opts ...options.Option) ([]topoapi.Object, error)
}

// NewClient creates a new E2 client
func NewClient(opts ...options.Option) (Client, error) {

	options := options.Options{
		Service: options.ServiceOptions{
			Host: defaultServiceHost,
			Port: defaultServicePort,
		},
	}

	for _, opt := range opts {
		opt.Apply(&options)
	}

	dialOpts := []grpc.DialOption{
		grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(100 * time.Millisecond)),
	}
	if options.Service.Insecure {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	} else {
		tlsConfig, err := creds.GetClientCredentials()
		if err != nil {
			log.Warn(err)
			return nil, err
		}

		dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	}
	conns := connection.NewManager()
	conn, err := conns.Connect(options.Service.GetAddress(), dialOpts...)
	if err != nil {
		return nil, err
	}

	cl, err := client.NewClient(conn)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	return &topoClient{
		client: cl,
	}, nil
}

// topoClient is the topo client implementation
type topoClient struct {
	client client.Client
}

func (t *topoClient) Update(ctx context.Context, object *topoapi.Object) error {
	return t.client.Update(ctx, object)
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

// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	"context"

	"github.com/onosproject/onos-ric-sdk-go/pkg/utils/creds"
	"google.golang.org/grpc/credentials"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"time"

	"github.com/onosproject/onos-ric-sdk-go/pkg/topo/connection"

	"github.com/onosproject/onos-lib-go/pkg/southbound"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("topo")

// Client is a topo SDK client
type Client interface {

	// Update updates a topo object
	Update(ctx context.Context, object *topoapi.Object) error

	// Get gets a topo object with a given ID
	Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error)

	// Watch provides a simple facility for the application to watch for changes in the topology
	Watch(ctx context.Context, ch chan<- topoapi.Event, opts ...WatchOption) error

	// List of topo objects
	List(ctx context.Context, opts ...ListOption) ([]topoapi.Object, error)
}

// NewClient creates a new topo client
func NewClient(opts ...Option) (Client, error) {
	clientOptions := Options{
		Service: ServiceOptions{
			Host: DefaultServiceHost,
			Port: DefaultServicePort,
		},
	}

	for _, opt := range opts {
		opt.apply(&clientOptions)
	}

	dialOpts := []grpc.DialOption{
		grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(100 * time.Millisecond)),
		grpc.WithUnaryInterceptor(southbound.RetryingUnaryClientInterceptor()),
	}
	if clientOptions.Service.Insecure {
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
	conn, err := conns.Connect(clientOptions.Service.GetAddress(), dialOpts...)
	if err != nil {
		return nil, err
	}

	topoClient, err := NewTopoClient(conn)
	if err != nil {
		log.Warn(err)
		return nil, err
	}

	return &topo{
		topoClient: topoClient,
	}, nil
}

// topoClient is the topo client implementation
type topo struct {
	topoClient TopoClient
}

func (t *topo) Update(ctx context.Context, object *topoapi.Object) error {
	return t.topoClient.Update(ctx, object)
}

func (t *topo) List(ctx context.Context, opts ...ListOption) ([]topoapi.Object, error) {
	return t.topoClient.List(ctx, opts...)

}

func (t *topo) Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error) {
	return t.topoClient.Get(ctx, id)
}

func (t *topo) Watch(ctx context.Context, ch chan<- topoapi.Event, opts ...WatchOption) error {
	return t.topoClient.Watch(ctx, ch, opts...)
}

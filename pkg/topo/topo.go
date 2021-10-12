// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package topo

import (
	"context"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"io"

	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/grpc/status"

	"github.com/onosproject/onos-ric-sdk-go/pkg/utils/creds"
	"google.golang.org/grpc/credentials"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/onosproject/onos-ric-sdk-go/pkg/topo/connection"

	"google.golang.org/grpc"
)

var log = logging.GetLogger("topo")

// Client is a topo SDK client
type Client interface {
	// Create creates a topo object
	Create(ctx context.Context, object *topoapi.Object) error

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
	clientOptions := Options{}

	for _, opt := range opts {
		opt.apply(&clientOptions)
	}

	if clientOptions.Service.Host == "" || clientOptions.Service.Port == 0 {
		clientOptions.Service.Host = DefaultServiceHost
		clientOptions.Service.Port = DefaultServicePort
	}

	dialOpts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor()),
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor()),
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

	cl := topoapi.NewTopoClient(conn)

	return &topo{
		client: cl,
	}, nil
}

// topo is the topo client
type topo struct {
	client topoapi.TopoClient
}

// Create creates a topo object
func (t *topo) Create(ctx context.Context, object *topoapi.Object) error {
	response, err := t.client.Create(ctx, &topoapi.CreateRequest{
		Object: object,
	})
	if err != nil {
		return errors.FromGRPC(err)
	}
	*object = *response.Object
	return nil
}

// Update updates a given topo object
func (t *topo) Update(ctx context.Context, object *topoapi.Object) error {
	response, err := t.client.Update(ctx, &topoapi.UpdateRequest{
		Object: object,
	})
	if err != nil {
		return errors.FromGRPC(err)
	}

	*object = *response.Object
	return nil
}

// List lists all of topo objects
func (t *topo) List(ctx context.Context, opts ...ListOption) ([]topoapi.Object, error) {
	options := ListOptions{}

	for _, opt := range opts {
		opt.apply(&options)
	}

	response, err := t.client.List(ctx, &topoapi.ListRequest{
		Filters: options.GetFilters(),
	})
	if err != nil {
		return nil, errors.FromGRPC(err)
	}

	return response.GetObjects(), nil
}

// Get get a topo object based on a given ID
func (t *topo) Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error) {
	response, err := t.client.Get(ctx, &topoapi.GetRequest{
		ID: id,
	})
	if err != nil {
		return nil, errors.FromGRPC(err)
	}
	return response.GetObject(), nil
}

// Watch watches topology events
func (t *topo) Watch(ctx context.Context, ch chan<- topoapi.Event, opts ...WatchOption) error {

	options := WatchOptions{}
	for _, opt := range opts {
		opt.apply(&options)
	}

	req := topoapi.WatchRequest{
		Filters:  options.GetFilters(),
		Noreplay: options.GetNoReplay(),
	}
	stream, err := t.client.Watch(ctx, &req)
	if err != nil {
		defer close(ch)
		return errors.FromGRPC(err)
	}

	go func() {
		defer close(ch)
		for {
			resp, err := stream.Recv()
			if err == io.EOF || err == context.Canceled {
				break
			}

			if err != nil {
				stat, ok := status.FromError(err)
				if ok {
					err = errors.FromStatus(stat)
					if errors.IsCanceled(err) || errors.IsTimeout(err) {
						break
					}
				}
				log.Error("An error occurred in receiving topology changes", err)
			} else {
				ch <- resp.Event
			}
		}
	}()
	return nil

}

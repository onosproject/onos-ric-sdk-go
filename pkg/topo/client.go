// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	"context"

	"io"

	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/grpc/status"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"google.golang.org/grpc"
)

// TopoClient is a topo client interface
type TopoClient interface {

	// Create creates an R-NIB object
	Create(ctx context.Context, object *topoapi.Object) error

	// Update updates an existing R-NIB object
	Update(ctx context.Context, object *topoapi.Object) error

	// Get gets an R-NIB object
	Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error)

	// List lists R-NIB objects
	List(ctx context.Context, opts ...ListOption) ([]topoapi.Object, error)

	// Delete deletes an R-NIB object using the given ID
	Delete(ctx context.Context, id topoapi.ID) error

	// Watch watches topology events
	Watch(ctx context.Context, ch chan<- topoapi.Event, opts ...WatchOption) error
}

// NewTopoClient creates a new topo client
func NewTopoClient(conn *grpc.ClientConn) (TopoClient, error) {
	cl := topoapi.NewTopoClient(conn)
	return &topoClient{
		client: cl,
	}, nil

}

// topoClient is the default topo client implementation
type topoClient struct {
	client topoapi.TopoClient
}

func (t *topoClient) Create(ctx context.Context, object *topoapi.Object) error {
	_, err := t.client.Create(ctx, &topoapi.CreateRequest{
		Object: object,
	})
	if err != nil {
		return errors.FromGRPC(err)
	}

	return nil
}

func (t *topoClient) Update(ctx context.Context, object *topoapi.Object) error {
	_, err := t.client.Update(ctx, &topoapi.UpdateRequest{
		Object: object,
	})
	if err != nil {
		return errors.FromGRPC(err)
	}

	return nil

}

func (t *topoClient) Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error) {
	response, err := t.client.Get(ctx, &topoapi.GetRequest{
		ID: id,
	})
	if err != nil {
		return nil, errors.FromGRPC(err)
	}
	return response.GetObject(), nil
}

func (t *topoClient) List(ctx context.Context, opts ...ListOption) ([]topoapi.Object, error) {
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

func (t *topoClient) Delete(ctx context.Context, id topoapi.ID) error {
	_, err := t.client.Delete(ctx, &topoapi.DeleteRequest{
		ID: id,
	})
	if err != nil {
		return errors.FromGRPC(err)
	}
	return nil

}

func (t *topoClient) Watch(ctx context.Context, ch chan<- topoapi.Event, opts ...WatchOption) error {

	options := WatchOptions{}
	for _, opt := range opts {
		opt.apply(&options)
	}

	req := topoapi.WatchRequest{
		Filters: options.GetFilters(),
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

var _ TopoClient = &topoClient{}

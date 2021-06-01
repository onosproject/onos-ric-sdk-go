// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package client

import (
	"context"

	"github.com/onosproject/onos-ric-sdk-go/pkg/topo/options"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"io"

	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/grpc/status"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"

	"google.golang.org/grpc"
)

var log = logging.GetLogger("topo", "client")

// Topo is a topo client interface
type Topo interface {
	// Create creates an R-NIB object
	Create(ctx context.Context, object *topoapi.Object) error

	// Update updates an existing R-NIB object
	Update(ctx context.Context, object *topoapi.Object) error

	// Get gets an R-NIB object
	Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error)

	// List lists R-NIB objects
	List(ctx context.Context, opts options.ListOptions) ([]topoapi.Object, error)

	// Delete deletes an R-NIB object using the given ID
	Delete(ctx context.Context, id topoapi.ID) error

	// Watch watches topology events
	Watch(ctx context.Context, ch chan<- topoapi.Event, opts options.WatchOptions) error
}

// NewClient creates a new E2 client
func NewClient(conn *grpc.ClientConn) (Topo, error) {
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
		stat, ok := status.FromError(err)
		if ok {
			return errors.FromStatus(stat)
		}
		return err
	}

	return nil
}

func (t *topoClient) Update(ctx context.Context, object *topoapi.Object) error {
	_, err := t.client.Update(ctx, &topoapi.UpdateRequest{
		Object: object,
	})
	if err != nil {
		stat, ok := status.FromError(err)
		if ok {
			return errors.FromStatus(stat)
		}
		return err
	}

	return nil

}

func (t *topoClient) Get(ctx context.Context, id topoapi.ID) (*topoapi.Object, error) {
	response, err := t.client.Get(ctx, &topoapi.GetRequest{
		ID: id,
	})
	if err != nil {
		stat, ok := status.FromError(err)
		if ok {
			return nil, errors.FromStatus(stat)
		}
		return nil, err
	}
	return response.GetObject(), nil
}

func (t *topoClient) List(ctx context.Context, opts options.ListOptions) ([]topoapi.Object, error) {
	response, err := t.client.List(ctx, &topoapi.ListRequest{
		Filters: opts.GetFilters(),
	})
	if err != nil {
		stat, ok := status.FromError(err)
		if ok {
			return nil, errors.FromStatus(stat)
		}
		return nil, err
	}

	return response.GetObjects(), nil

}

func (t *topoClient) Delete(ctx context.Context, id topoapi.ID) error {
	_, err := t.client.Delete(ctx, &topoapi.DeleteRequest{
		ID: id,
	})
	if err != nil {
		stat, ok := status.FromError(err)
		if ok {
			return errors.FromStatus(stat)
		}
		return err
	}
	return nil

}

func (t *topoClient) Watch(ctx context.Context, ch chan<- topoapi.Event, opts options.WatchOptions) error {
	req := topoapi.WatchRequest{
		Filters: opts.GetFilters(),
	}
	stream, err := t.client.Watch(ctx, &req)
	if err != nil {
		defer close(ch)
		stat, ok := status.FromError(err)
		if ok {
			return errors.FromStatus(stat)
		}
		return err
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

var _ Topo = &topoClient{}
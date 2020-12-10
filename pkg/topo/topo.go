// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package topo contains facilities for RIC applications to monitor topology for changes
package topo

import (
	"context"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/southbound"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"time"
)

var log = logging.GetLogger("topo", "client")

// FilterOptions describes the criteria for the type of topology entities and events that the application should be
// notified about
type FilterOption interface {
	matches(e topoapi.Event) bool
}

type typeFilterOption struct {
	objectType topoapi.Object_Type
}

func (f *typeFilterOption) matches(e topoapi.Event) bool {
	return f.objectType == e.Object.Type
}

func TypeFilter(t topoapi.Object_Type) FilterOption {
	return &typeFilterOption{
		objectType: t,
	}
}

type kindFilterOption struct {
	kindID topoapi.ID
}

func KindFilter(id topoapi.ID) FilterOption {
	return &kindFilterOption{
		kindID: id,
	}
}

func (f *kindFilterOption) matches(e topoapi.Event) bool {
	return f.kindID == e.Object.GetEntity().KindID
}


// ServiceConfig is a topo service configuration
type ServiceConfig struct {
	// Host is the service host
	Host string
	// Port is the service port
	Port int
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
	Watch(ch chan<- topoapi.Object, options ...FilterOption) error
}

const defaultServiceHost = "onos-topo"
const defaultServicePort = 5150

// Current topology service configuration; starts with reasonable defaults
var topoServiceConfig = ServiceConfig{
	Host: defaultServiceHost,
	Port: defaultServicePort,
}

func SetTopoClientConfig(cfg ServiceConfig) {
	topoServiceConfig = cfg
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

// Watch provides a simple facility for the application to watch for changes in the topology.
func (c *topoClient) Watch(ch chan<- topoapi.Object, filter ...FilterOption) error {
	// Establish connection to the topology service
	client, err := newTopoClient(c.config)
	if err != nil {
		defer close(ch)
		return err
	}

	// Issue a watch request with replay and passing through any service-side filter criteria, i.e.
	// those supported by the topo.Watch method (TBD)
	stream, err := client.Watch(context.Background(), &topoapi.WatchRequest{})
	if err != nil {
		log.Error("Unable to issue topology watch request", err)
		stat, ok := status.FromError(err)
		if ok {
			return errors.FromStatus(stat)
		}
		return err
	}

	// Kick off a go routine that reads events from the topology watch stream and subjects them to additional
	// client-side filter criteria; objects that pass the filter get put on the application provided channel.
	go handleWatchStream(ch, stream, filter...)
	return nil
}

// handleWatchStream reads events from the topology watch stream and subjects them to additional
// client-side filter criteria; objects that pass the filter get put on the application provided channel.
func handleWatchStream(ch chan<- topoapi.Object, stream topoapi.Topo_WatchClient, filter ...FilterOption) {
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
			log.Error("An error occurred in receiving Subscription changes", err)
		} else if applyFilters(resp.Event, filter...) {
			ch <- resp.Event.Object
		}
	}
}

func applyFilters(event topoapi.Event, filter ...FilterOption) bool {
	for _, f := range filter {
		if !f.matches(event) {
			return false
		}
	}
	return true
}

// newClient creates a new topo client
func newTopoClient(cfg ServiceConfig) (topoapi.TopoClient, error) {
	opts := grpc.WithStreamInterceptor(southbound.RetryingStreamClientInterceptor(100 * time.Millisecond))
	conn, err := getTopoConn("onos-topo", opts)
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

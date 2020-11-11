// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package endpoint

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"io"

	regapi "github.com/onosproject/onos-e2sub/api/e2/registry/v1beta1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var log = logging.GetLogger("e2", "endpoint", "client")

// Client provides an E2 end-point client interface
type Client interface {
	// Add adds a TerminationEndPoint
	Add(ctx context.Context, endPoint *regapi.TerminationEndPoint) error

	// Remove removes a TerminationEndPoint
	Remove(ctx context.Context, endPoint *regapi.TerminationEndPoint) error

	// Get returns a TerminationEndPoint based on a given TerminationEndPoint ID
	Get(ctx context.Context, id regapi.ID) (*regapi.TerminationEndPoint, error)

	// List returns the list of existing TerminationEndPoints
	List(ctx context.Context) ([]regapi.TerminationEndPoint, error)

	// Watch watches the TerminationEndPoint changes
	Watch(ctx context.Context, ch chan<- regapi.Event) error
}

// localClient TerminationEndPoint client
type localClient struct {
	conn   *grpc.ClientConn
	client regapi.E2RegistryServiceClient
}

// Destination determines E2 registry service endpoint
type Destination struct {
	// Addrs a slice of addresses by which a TerminationEndPoint service may be reached.
	Addrs []string
}

// NewClient creates a new E2 termination registry service client
func NewClient(ctx context.Context, dst Destination) (Client, error) {
	tlsConfig, err := e2.GetClientCredentials()
	if err != nil {
		return &localClient{}, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	conn, err := grpc.DialContext(ctx, dst.Addrs[0], opts...)
	if err != nil {
		return &localClient{}, err
	}

	cl := regapi.NewE2RegistryServiceClient(conn)

	client := localClient{
		client: cl,
		conn:   conn,
	}

	return &client, nil
}

// Add adds a new E2 termination end-point
func (c *localClient) Add(ctx context.Context, endPoint *regapi.TerminationEndPoint) error {
	req := &regapi.AddTerminationRequest{
		EndPoint: endPoint,
	}

	_, err := c.client.AddTermination(ctx, req)
	if err != nil {
		return err
	}

	return nil

}

// Remove removes an E2 termination end-point
func (c *localClient) Remove(ctx context.Context, endPoint *regapi.TerminationEndPoint) error {
	req := &regapi.RemoveTerminationRequest{
		ID: endPoint.ID,
	}

	_, err := c.client.RemoveTermination(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// Get returns information about an E2 termination end-point
func (c *localClient) Get(ctx context.Context, id regapi.ID) (*regapi.TerminationEndPoint, error) {
	req := &regapi.GetTerminationRequest{
		ID: id,
	}

	resp, err := c.client.GetTermination(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.EndPoint, nil
}

// List returns the list of currently registered E2 termination end-points
func (c *localClient) List(ctx context.Context) ([]regapi.TerminationEndPoint, error) {
	req := &regapi.ListTerminationsRequest{}

	resp, err := c.client.ListTerminations(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.EndPoints, nil
}

// Watch watches for changes in the inventory of available E2T termination end-points
func (c *localClient) Watch(ctx context.Context, ch chan<- regapi.Event) error {
	req := regapi.WatchTerminationsRequest{}
	stream, err := c.client.WatchTerminations(ctx, &req)
	if err != nil {
		close(ch)
		return err
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF || err == context.Canceled {
				close(ch)
				break
			}

			if err != nil {
				log.Error("an error occurred in receiving TerminationEndPoint changes", err)
			}

			ch <- resp.Event

		}

	}()
	return nil
}

// Close closes the client connection
func (c *localClient) Close() error {
	return c.conn.Close()
}

var _ Client = &localClient{}

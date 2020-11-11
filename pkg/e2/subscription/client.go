// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2"
	"io"

	subapi "github.com/onosproject/onos-e2sub/api/e2/subscription/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var log = logging.GetLogger("e2", "subscription", "client")

// Client is an E2 subscription service client interface
type Client interface {
	// Add adds a subscription
	Add(ctx context.Context, subscription *subapi.Subscription) error

	// Remove removes a subscription
	Remove(ctx context.Context, subscription *subapi.Subscription) error

	// Get returns a subscription based on a given subscription ID
	Get(ctx context.Context, id subapi.ID) (*subapi.Subscription, error)

	// List returns the list of existing subscriptions
	List(ctx context.Context) ([]subapi.Subscription, error)

	// Watch watches the subscription changes
	Watch(ctx context.Context, ch chan<- subapi.Event) error
}

// localClient subscription client
type localClient struct {
	conn   *grpc.ClientConn
	client subapi.E2SubscriptionServiceClient
}

// Destination determines subscription service endpoint
type Destination struct {
	// Addrs a slice of addresses by which a subscription service may be reached.
	Addrs []string
}

// NewClient creates a new subscribe service client
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

	cl := subapi.NewE2SubscriptionServiceClient(conn)

	client := localClient{
		client: cl,
		conn:   conn,
	}

	return &client, nil
}

// Add adds a subscription
func (c *localClient) Add(ctx context.Context, subscription *subapi.Subscription) error {
	req := &subapi.AddSubscriptionRequest{
		Subscription: subscription,
	}

	_, err := c.client.AddSubscription(ctx, req)
	if err != nil {
		return err
	}

	return nil

}

// Remove removes a subscription
func (c *localClient) Remove(ctx context.Context, subscription *subapi.Subscription) error {
	req := &subapi.RemoveSubscriptionRequest{
		ID: subscription.ID,
	}

	_, err := c.client.RemoveSubscription(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// Get returns information about a subscription
func (c *localClient) Get(ctx context.Context, id subapi.ID) (*subapi.Subscription, error) {
	req := &subapi.GetSubscriptionRequest{
		ID: id,
	}

	resp, err := c.client.GetSubscription(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Subscription, nil
}

// List returns the list of all subscriptions
func (c *localClient) List(ctx context.Context) ([]subapi.Subscription, error) {
	req := &subapi.ListSubscriptionsRequest{}

	resp, err := c.client.ListSubscriptions(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.Subscriptions, nil
}

// Watch watches for changes in the set of subscriptions
func (c *localClient) Watch(ctx context.Context, ch chan<- subapi.Event) error {
	req := subapi.WatchSubscriptionsRequest{}
	stream, err := c.client.WatchSubscriptions(ctx, &req)
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
				log.Error("an error occurred in receiving subscription changes", err)
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

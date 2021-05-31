// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"google.golang.org/grpc"
	"io"
	"os"
	"sync"
)

// NodeID is an E2 node identifier
type NodeID string

// Node is an E2 node
type Node interface {
	// ID is the node identifier
	ID() NodeID

	// Subscribe creates a subscription from the given SubscriptionDetails
	// The Subscribe method will block until the subscription is successfully registered.
	// The context.Context represents the lifecycle of this initial subscription process.
	// Once the subscription has been created and the method returns, indications will be written
	// to the given channel.
	// If the subscription is successful, a subscription.Context will be returned. The subscription
	// context can be used to cancel the subscription by calling Close() on the subscription.Context.
	Subscribe(ctx context.Context, sub *e2api.Subscription, indCh chan<- e2api.Indication, errCh chan<- error) error

	// Control creates and sends a E2 control message and awaits the outcome
	Control(ctx context.Context, message *e2api.ControlMessage) (*e2api.ControlOutcome, error)
}

// NewNode creates a new E2 Node with the given ID
func NewNode(nodeID NodeID, opts ...Option) Node {
	options := Options{
		App: AppOptions{
			AppID:      AppID(os.Getenv("SERVICE_NAME")),
			InstanceID: InstanceID(os.Getenv("POD_NAME")),
		},
		Service: ServiceOptions{
			Host: "onos-e2t",
			Port: defaultServicePort,
		},
	}
	for _, opt := range opts {
		opt.apply(&options)
	}

	uuid.SetNodeID([]byte(fmt.Sprintf("%s:%s", options.App.AppID, options.App.InstanceID)))

	return &e2Node{
		nodeID:  nodeID,
		options: options,
	}
}

// e2Node is the default E2 node implementation
type e2Node struct {
	nodeID  NodeID
	options Options
	conn    *grpc.ClientConn
	mu      sync.RWMutex
}

func (n *e2Node) ID() NodeID {
	return n.nodeID
}

func (n *e2Node) connect(ctx context.Context) (*grpc.ClientConn, error) {
	n.mu.RLock()
	conn := n.conn
	n.mu.RUnlock()

	if conn != nil {
		return conn, nil
	}

	n.mu.Lock()
	defer n.mu.Unlock()

	if n.conn != nil {
		return n.conn, nil
	}

	conn, err := grpc.DialContext(ctx, n.options.Service.GetAddress(), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		return nil, err
	}
	n.conn = conn
	return conn, nil
}

func (n *e2Node) getRequestHeaders() e2api.RequestHeaders {
	return e2api.RequestHeaders{
		AppID:      e2api.AppID(n.options.App.AppID),
		InstanceID: e2api.InstanceID(n.options.App.InstanceID),
		NodeID:     e2api.NodeID(n.nodeID),
		ServiceModel: e2api.ServiceModel{
			Name:    e2api.ServiceModelName(n.options.ServiceModel.Name),
			Version: e2api.ServiceModelVersion(n.options.ServiceModel.Version),
		},
	}
}

func (n *e2Node) Control(ctx context.Context, message *e2api.ControlMessage) (*e2api.ControlOutcome, error) {
	conn, err := n.connect(ctx)
	if err != nil {
		return nil, err
	}
	client := e2api.NewControlServiceClient(conn)

	request := &e2api.ControlRequest{
		Headers: n.getRequestHeaders(),
	}
	response, err := client.Control(ctx, request)
	if err != nil {
		return nil, errors.FromGRPC(err)
	}
	return &response.Outcome, nil
}

func (n *e2Node) Subscribe(ctx context.Context, sub *e2api.Subscription, indCh chan<- e2api.Indication, errCh chan<- error) error {
	conn, err := n.connect(ctx)
	if err != nil {
		return err
	}
	client := e2api.NewSubscriptionServiceClient(conn)

	request := &e2api.SubscribeRequest{
		Headers:      n.getRequestHeaders(),
		Subscription: *sub,
	}
	stream, err := client.Subscribe(ctx, request)
	if err != nil {
		defer close(indCh)
		defer close(errCh)
		return errors.FromGRPC(err)
	}

	ackCh := make(chan error)
	go func() {
		defer close(indCh)
		defer close(errCh)

		acked := false
		for {
			response, err := stream.Recv()
			err = errors.FromGRPC(err)
			if err == io.EOF || err == context.Canceled || errors.IsCanceled(err) {
				break
			}

			if err != nil {
				if errors.IsCanceled(err) || errors.IsTimeout(err) {
					break
				}
				log.Error("An error occurred in receiving Subscription changes", err)
				if !acked {
					ackCh <- err
					close(ackCh)
					break
				} else {
					errCh <- err
				}
			} else {
				switch m := response.Message.(type) {
				case *e2api.SubscribeResponse_Ack:
					close(ackCh)
					acked = true
				case *e2api.SubscribeResponse_Indication:
					indCh <- *m.Indication
				}
			}
		}
		if !acked {
			close(ackCh)
		}
	}()

	select {
	case <-ackCh:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (n *e2Node) Close() error {
	return n.conn.Close()
}

// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	e2api "github.com/onosproject/onos-api/go/onos/e2t/e2/v1beta1"
	"github.com/onosproject/onos-lib-go/pkg/env"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/creds"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/v1beta1/e2errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"io"
	"sync"
	"time"
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
	Subscribe(ctx context.Context, name string, sub e2api.SubscriptionSpec, indCh chan<- e2api.Indication, opts ...SubscribeOption) (e2api.ChannelID, error)

	// Unsubscribe unsubscribes from the given subscription
	Unsubscribe(ctx context.Context, name string) error

	// Control creates and sends a E2 control message and awaits the outcome
	Control(ctx context.Context, message *e2api.ControlMessage) (*e2api.ControlOutcome, error)
}

// NewNode creates a new E2 Node with the given ID
func NewNode(nodeID NodeID, opts ...Option) Node {
	options := Options{
		App: AppOptions{
			AppID:      AppID(env.GetServiceName()),
			InstanceID: InstanceID(env.GetPodName()),
		},
		Service: ServiceOptions{
			Host: "onos-e2t",
			Port: defaultServicePort,
		},
		Topo: ServiceOptions{
			Host: "onos-topo",
			Port: defaultServicePort,
		},
		Encoding: ProtoEncoding,
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

type ackResult struct {
	err       error
	channelID e2api.ChannelID
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

	clientCreds, _ := creds.GetClientCredentials()
	conn, err := grpc.DialContext(ctx, "localhost:5151",
		grpc.WithTransportCredentials(credentials.NewTLS(clientCreds)),
		grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor(retry.WithRetryOn(codes.Unavailable))),
		grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithRetryOn(codes.Unavailable))))
	if err != nil {
		return nil, err
	}
	n.conn = conn
	return conn, nil
}

func getErrorFromGRPC(err error) error {
	if e2errors.IsE2APError(err) {
		return e2errors.FromGRPC(err)
	}
	return errors.FromGRPC(err)
}

func (n *e2Node) getRequestHeaders() e2api.RequestHeaders {
	var encoding e2api.Encoding
	switch n.options.Encoding {
	case ProtoEncoding:
		encoding = e2api.Encoding_PROTO
	case ASN1Encoding:
		encoding = e2api.Encoding_ASN1_PER
	}
	return e2api.RequestHeaders{
		AppID:         e2api.AppID(n.options.App.AppID),
		AppInstanceID: e2api.AppInstanceID(n.options.App.InstanceID),
		E2NodeID:      e2api.E2NodeID(n.nodeID),
		ServiceModel: e2api.ServiceModel{
			Name:    e2api.ServiceModelName(n.options.ServiceModel.Name),
			Version: e2api.ServiceModelVersion(n.options.ServiceModel.Version),
		},
		Encoding: encoding,
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
		Message: *message,
	}
	log.Debugf("Sending ControlRequest %+v", request)
	response, err := client.Control(ctx, request)
	if err != nil {
		log.Warnf("ControlRequest %+v failed: %v", request, err)
		return nil, getErrorFromGRPC(err)
	}
	log.Debugf("Received ControlResponse %+v", response)
	return &response.Outcome, nil
}

func (n *e2Node) Subscribe(ctx context.Context, name string, sub e2api.SubscriptionSpec, indCh chan<- e2api.Indication, opts ...SubscribeOption) (e2api.ChannelID, error) {
	conn, err := n.connect(ctx)
	if err != nil {
		return "", err
	}
	client := e2api.NewSubscriptionServiceClient(conn)

	options := SubscribeOptions{TransactionTimeout: 2 * time.Minute}
	for _, opt := range opts {
		opt.apply(&options)
	}

	request := &e2api.SubscribeRequest{
		Headers:            n.getRequestHeaders(),
		TransactionID:      e2api.TransactionID(name),
		Subscription:       sub,
		TransactionTimeout: &options.TransactionTimeout,
	}
	log.Debugf("Sending SubscribeRequest %+v", request)
	stream, err := client.Subscribe(ctx, request)
	if err != nil {
		defer close(indCh)
		return "", getErrorFromGRPC(err)
	}

	ackCh := make(chan ackResult)
	go func() {
		defer close(indCh)
		acked := false
		var channelID e2api.ChannelID
		for {
			response, err := stream.Recv()
			if err == io.EOF || err == context.Canceled {
				break
			}

			if err != nil {
				err = getErrorFromGRPC(err)
				if errors.IsCanceled(err) || errors.IsTimeout(err) {
					break
				}

				log.Warnf("SubscribeRequest %+v failed: %v", request, err)
				if !acked {
					ackCh <- ackResult{
						err: err,
					}
					close(ackCh)
					acked = true
					break
				}
			} else {
				log.Debugf("Received SubscribeResponse %+v", response)
				switch m := response.Message.(type) {
				case *e2api.SubscribeResponse_Ack:
					channelID = m.Ack.ChannelID
					if !acked {
						ackCh <- ackResult{
							channelID: channelID,
						}
						close(ackCh)
						acked = true
					}
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
	case result := <-ackCh:
		return result.channelID, result.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (n *e2Node) Unsubscribe(ctx context.Context, name string) error {
	conn, err := n.connect(ctx)
	if err != nil {
		return err
	}
	client := e2api.NewSubscriptionServiceClient(conn)

	request := &e2api.UnsubscribeRequest{
		Headers:       n.getRequestHeaders(),
		TransactionID: e2api.TransactionID(name),
	}
	log.Debugf("Sending UnsubscribeRequest %+v", request)
	response, err := client.Unsubscribe(ctx, request)
	if err != nil {
		log.Warnf("UnsubscribeRequest %+v failed: %v", request, err)
		return getErrorFromGRPC(err)
	}
	log.Debugf("Received UnsubscribeResponse %+v", response)
	return nil
}

func (n *e2Node) Close() error {
	return n.conn.Close()
}

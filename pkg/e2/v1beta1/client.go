// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("e2", "v1beta1")

// Client is an E2 client
type Client interface {
	// Node returns a Node with the given NodeID
	Node(nodeID NodeID) Node
}

// NewClient creates a new E2 client
func NewClient(opts ...Option) Client {
	return &e2Client{
		opts: opts,
	}
}

// e2Client is the default E2 client implementation
type e2Client struct {
	opts []Option
}

func (c *e2Client) Node(nodeID NodeID) Node {
	return NewNode(nodeID, c.opts...)
}

var _ Client = &e2Client{}

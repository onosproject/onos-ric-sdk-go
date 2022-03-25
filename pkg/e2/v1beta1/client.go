// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"sync"
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
		opts:  opts,
		nodes: make(map[NodeID]Node),
	}
}

// e2Client is the default E2 client implementation
type e2Client struct {
	opts  []Option
	nodes map[NodeID]Node
	mu    sync.RWMutex
}

func (c *e2Client) Node(nodeID NodeID) Node {
	c.mu.RLock()
	node, ok := c.nodes[nodeID]
	c.mu.RUnlock()
	if ok {
		return node
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	node, ok = c.nodes[nodeID]
	if ok {
		return node
	}

	node = NewNode(nodeID, c.opts...)
	c.nodes[nodeID] = node
	go func() {
		<-node.Context().Done()
		c.mu.Lock()
		delete(c.nodes, nodeID)
		c.mu.Unlock()
	}()
	return node
}

var _ Client = &e2Client{}

// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package connection

import (
	"sync"

	"google.golang.org/grpc"
)

// NewManager creates a new connection manager
func NewManager() *Manager {
	return &Manager{
		conns: make(map[string]*grpc.ClientConn),
	}
}

// Manager is a connection manager
type Manager struct {
	conns map[string]*grpc.ClientConn
	mu    sync.RWMutex
}

// Connect connects to the given address
func (m *Manager) Connect(address string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	m.mu.RLock()
	conn, ok := m.conns[address]
	m.mu.RUnlock()
	if ok {
		return conn, nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	conn, ok = m.conns[address]
	if ok {
		return conn, nil
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}
	m.conns[address] = conn
	return conn, nil
}

// Close closes the connection manager
func (m *Manager) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	var err error
	for _, conn := range m.conns {
		if e := conn.Close(); e != nil {
			err = e
		}
	}
	return err
}

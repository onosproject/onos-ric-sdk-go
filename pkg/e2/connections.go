// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import "github.com/onosproject/onos-lib-go/pkg/logging"

var log = logging.GetLogger("connections")

// Connections is a structure for tracking E2T connections
type Connections struct {
}

// NewManager creates a new manager
func NewConnections() *Connections {
	log.Info("Creating connections")
	return &Connections{
	}
}


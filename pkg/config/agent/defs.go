// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package gnmi implements a gnmi server to mock a device with YANG models.
package agent

import (
	"sync"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
)

// Server implements the interface of gnmi server. It supports Capabilities, Get, and Set APIs.
type Server struct {
	mu           sync.RWMutex // mu is the RW lock to protect the access to config
	configurable configurable.Configurable
}

// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package topo contains facilities for RIC applications to monitor topology for changes
package topo

import (
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

// WatchRequest describes the criteria for the type of topology entities and events that the application should be
// notified about
type WatchRequest struct {
	// Various filter criteria to be added here
}

// Watch topology provides a simple facility for the application to watch for changes in the topology.
func WatchTopology(req WatchRequest, ch <- chan *topoapi.Object) error {
	return nil
}

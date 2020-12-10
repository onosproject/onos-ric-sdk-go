// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopoWatch(t *testing.T) {
	t.Skip("WIP")
	
	// Start mock topology service

	client, err := NewClient(DefaultServiceConfig)
	assert.NoError(t, err, "unable to get new topology client")

	ch := make(chan topoapi.Event)
	err = client.Watch(ch, TypeFilter(topoapi.Object_ENTITY), KindFilter("eNB"))
	assert.NoError(t, err, "unable to start topology watch")
}
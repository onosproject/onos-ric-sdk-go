// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package subscription

import (
	"github.com/gogo/protobuf/proto"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/node"
	"github.com/onosproject/onos-ric-sdk-go/pkg/e2/sm"
)

// Subscription is an E2 subscription
type Subscription struct {
	// NodeID is the E2 node identifier
	NodeID node.ID
	// ServiceModel is the service model
	ServiceModel sm.ServiceModel
	// Payload is the subscription payload
	Payload Payload
}

// Payload is an E2 subscription payload
type Payload proto.Message

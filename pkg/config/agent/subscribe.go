// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

// Package agent implements a gnmi server to mock a device with YANG models.
package agent

import (
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

// Subscribe handle subscribe requests including POLL, STREAM, ONCE subscribe requests
func (s *Server) Subscribe(stream pb.GNMI_SubscribeServer) error {
	return nil

}

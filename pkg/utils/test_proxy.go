// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package utils

import (
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-proxy/pkg/manager"
)

var mgr *manager.Manager

// StartTestProxy starts onos-proxy instance for testing purposes
func StartTestProxy() {
	cfg := manager.Config{
		CAPath:   certs.OnfCaCrt,
		KeyPath:  certs.DefaultOnosConfigKey,
		CertPath: certs.DefaultOnosConfigCrt,
		GRPCPort: 5151,
	}
	mgr := manager.NewManager(cfg)
	mgr.Run()
}

// StopTestProxy stops test instance of onos-proxy
func StopTestProxy() {
	mgr.Close()
}

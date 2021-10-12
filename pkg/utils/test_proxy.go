// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"fmt"
	"github.com/onosproject/onos-lib-go/pkg/certs"
	"github.com/onosproject/onos-proxy/pkg/manager"
	"io/ioutil"
	"os"
)

var mgr *manager.Manager

const (
	caCrtFile = "/tmp/onos-proxy.cacrt"
	crtFile   = "/tmp/onos-proxy.crt"
	keyFile   = "/tmp/onos-proxy.key"
)

// StartTestProxy starts onos-proxy instance for testing purposes
func StartTestProxy() {
	writeFile(caCrtFile, certs.OnfCaCrt)
	writeFile(crtFile, certs.DefaultOnosConfigCrt)
	writeFile(keyFile, certs.DefaultOnosConfigKey)

	cfg := manager.Config{
		CAPath:   caCrtFile,
		KeyPath:  keyFile,
		CertPath: crtFile,
		GRPCPort: 5151,
	}

	mgr = manager.NewManager(cfg)
	mgr.Run()
}

func writeFile(file string, s string) {
	err := ioutil.WriteFile(file, []byte(s), 0644)
	if err != nil {
		fmt.Printf("error writing generated code to file: %s\n", err)
		os.Exit(-1)
	}
}

// StopTestProxy stops test instance of onos-proxy
func StopTestProxy() {
	mgr.Close()
}

// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package registry

import (
	"io/ioutil"
	"os"
)

// load loads the initial configuration
func loadConfig(jsonPath string) ([]byte, error) {
	//jsonFile, err := os.Open("/etc/onos/config/config.json")
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return byteValue, nil

}

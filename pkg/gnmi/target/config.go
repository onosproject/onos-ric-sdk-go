// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package target

import (
	"io/ioutil"
	"os"
)

type Config struct {
	Config map[string]interface{}
}

// load loads the initial configuration
func load() ([]byte, error) {
	jsonFile, err := os.Open("/etc/onos/config/config.json")
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

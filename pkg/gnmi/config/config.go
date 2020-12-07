// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const configDir = ".onos"

type Config struct {
}

// load loads the configuration
func load(config *Config) error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	// Set the file name of the configurations file
	viper.SetConfigName("gnmi")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./" + configDir + "/config")
	viper.AddConfigPath(home + "/" + configDir + "/config")
	viper.AddConfigPath("/etc/onos/config")
	viper.AddConfigPath(".")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return err
	}
	return nil
}

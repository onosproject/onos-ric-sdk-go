// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package callback

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/store"
)

var log = logging.GetLogger("config", "callback")

// Config
type Config struct {
	config *store.ConfigStore
}

// InitConfig
func (c *Config) InitConfig(config *store.ConfigStore) {
	c.config = config

}

var _ configurable.Configurable = &Config{}

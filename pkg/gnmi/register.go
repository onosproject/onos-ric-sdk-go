// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package gnmi

import (
	"github.com/onosproject/onos-ric-sdk-go/pkg/gnmi/target"
)

func RegisterConfigurable(c *Config) target.GnmiService {
	service := target.NewService(c.ModelInfo)
	c.server = service.GetServer()

	return service
}

type Config struct {
	server    *target.Server
	ModelInfo target.ModelInfo
}

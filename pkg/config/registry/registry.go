// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package registry

import (
	"encoding/json"

	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/agent"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
)

var log = logging.GetLogger("registry")

type RegisterRequest struct {
	server *agent.Server
}

type RegisterResponse struct {
	AgentService agent.GnmiService
	Config       *configurable.ConfigStore
}

func RegisterConfigurable(c *RegisterRequest) (RegisterResponse, error) {
	initialConfig, err := load()
	if err != nil {
		log.Error("Failed to read initial config", err)
		return RegisterResponse{}, err
	}

	config := configurable.NewConfigStore()
	err = json.Unmarshal(initialConfig, &config.ConfigTree)
	if err != nil {
		log.Error("Failed to unmarshal initial config to json")
		return RegisterResponse{}, err
	}

	configurableEntity := &configurable.Config{}
	configurableEntity.InitConfig(config)

	service := agent.NewService(configurableEntity)
	c.server = service.GetServer()
	response := RegisterResponse{
		AgentService: service,
		Config:       config,
	}

	return response, nil
}

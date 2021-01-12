// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package registry

import (
	"encoding/json"

	"github.com/onosproject/onos-lib-go/pkg/northbound"

	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/agent"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
)

var log = logging.GetLogger("registry")

const (
	// IANA reserved port for gNMI
	gnmiAgentPort = 9339
)

type RegisterRequest struct {
	server *agent.Server
}

type RegisterResponse struct {
	Config *configurable.ConfigStore
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

	s := northbound.NewServer(northbound.NewServerCfg(
		"",
		"",
		"",
		int16(gnmiAgentPort),
		true,
		northbound.SecurityConfig{}))

	service := agent.NewService(configurableEntity)
	s.AddService(service)

	c.server = service.GetServer()
	response := RegisterResponse{
		Config: config,
	}

	doneCh := make(chan error)
	go func() {
		err := s.Serve(func(started string) {
			log.Info("Started gNMI Agent on port ", started)
			close(doneCh)
		})
		if err != nil {
			doneCh <- err
		}
	}()

	return response, <-doneCh
}

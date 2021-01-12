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
}

type RegisterResponse struct {
	Config *configurable.ConfigStore
}

func startAgent(c configurable.Configurable) error {
	s := northbound.NewServer(northbound.NewServerCfg(
		"",
		"",
		"",
		int16(gnmiAgentPort),
		true,
		northbound.SecurityConfig{}))

	service := agent.NewService(c)
	s.AddService(service)

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
	return <-doneCh
}

func RegisterConfigurable(r *RegisterRequest) (RegisterResponse, error) {
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
	err = startAgent(configurableEntity)
	if err != nil {
		return RegisterResponse{}, err
	}
	response := RegisterResponse{
		Config: config,
	}

	return response, nil
}

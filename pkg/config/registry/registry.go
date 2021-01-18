// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package registry

import (
	"encoding/json"

	_default "github.com/onosproject/onos-ric-sdk-go/pkg/config/app/default"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/callback"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/store"

	"github.com/onosproject/onos-lib-go/pkg/northbound"

	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/agent"
)

var log = logging.GetLogger("registry")

const (
	// IANA reserved port for gNMI
	gnmiAgentPort = 9339
)

type RegisterRequest struct {
}

type RegisterResponse struct {
	Config interface{}
}

// startAgent stats gnmi agent server
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

// RegisterConfigurable registers a configurable entity and starts a gNMI agent server
func RegisterConfigurable(req *RegisterRequest) (RegisterResponse, error) {
	initialConfig, err := loadConfig()
	if err != nil {
		log.Error("Failed to read initial config", err)
		return RegisterResponse{}, err
	}

	config := store.NewConfigStore()
	err = json.Unmarshal(initialConfig, &config.ConfigTree)
	if err != nil {
		log.Error("Failed to unmarshal initial config to json")
		return RegisterResponse{}, err
	}

	configurableEntity := &callback.Config{}
	configurableEntity.InitConfig(config)
	err = startAgent(configurableEntity)
	if err != nil {
		return RegisterResponse{}, err
	}

	cfg := _default.NewConfig(config)

	response := RegisterResponse{
		Config: cfg,
	}

	return response, nil
}

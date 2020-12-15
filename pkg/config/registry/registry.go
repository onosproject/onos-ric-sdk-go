// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package registry

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/agent"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/configurable"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/model"
	"github.com/openconfig/ygot/ygot"
)

var log = logging.GetLogger("registry")

type RegisterRequest struct {
	ModelInfo model.ModelInfo
	server    *agent.Server
}

type RegisterResponse struct {
	AgentService agent.GnmiService
	Config       *ygot.ValidatedGoStruct
}

func RegisterConfigurable(c *RegisterRequest) (RegisterResponse, error) {
	initialConfig, err := load()
	if err != nil {
		log.Error("Failed to read initial config", err)
	}

	configurableEntity := &configurable.Config{}

	newModel := model.GetModel(c.ModelInfo)
	rootStruct, err := newModel.NewConfigStruct(initialConfig)
	if err != nil {
		log.Errorf("initial config cannot be initialized", err)
	}

	configurableEntity.InitConfig(newModel, &rootStruct)

	service := agent.NewService(configurableEntity)
	c.server = service.GetServer()
	response := RegisterResponse{
		Config:       &rootStruct,
		AgentService: service,
	}

	return response, nil
}

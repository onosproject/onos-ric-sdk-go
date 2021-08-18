// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"math/rand"
)

func init() {
	balancer.Register(base.NewBalancerBuilder(resolverName, &PickerBuilder{}, base.Config{}))
}

type PickerBuilder struct{}

func (p *PickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	var master balancer.SubConn
	var backups []balancer.SubConn
	for sc, scInfo := range info.ReadySCs {
		isMaster := scInfo.Address.Attributes.Value("is_master").(bool)
		if isMaster {
			master = sc
			continue
		}
		backups = append(backups, sc)
	}
	if master == nil && len(backups) > 0 {
		master = backups[rand.Intn(len(backups))]
	}
	log.Debugf("Built new picker. Master: %s, Backups: %s", master, backups)
	return &Picker{
		master: master,
	}
}

var _ base.PickerBuilder = (*PickerBuilder)(nil)

type Picker struct {
	master balancer.SubConn
}

func (p *Picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	var result balancer.PickResult
	if p.master == nil {
		return result, balancer.ErrNoSubConnAvailable
	}
	result.SubConn = p.master
	return result, nil
}

var _ balancer.Picker = (*Picker)(nil)

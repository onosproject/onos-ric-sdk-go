// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

func init() {
	//balancer.Register(base.NewBalancerBuilder(resolverName, &PickerBuilder{}, base.Config{}))
}

// PickerBuilder :
type PickerBuilder struct{}

// Build :
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
	log.Debugf("Built new picker. Master: %s, Backups: %s", master, backups)
	return &Picker{
		master: master,
	}
}

var _ base.PickerBuilder = (*PickerBuilder)(nil)

// Picker :
type Picker struct {
	master balancer.SubConn
}

// Pick :
func (p *Picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	var result balancer.PickResult
	if p.master == nil {
		return result, balancer.ErrNoSubConnAvailable
	}
	result.SubConn = p.master
	return result, nil
}

var _ balancer.Picker = (*Picker)(nil)

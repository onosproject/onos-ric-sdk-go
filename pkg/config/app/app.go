// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"
)

// Entry :
type Entry struct {
	Value interface{}
}

// Configuration :
type Configuration interface {
	Config() interface{}

	Get(key string) (Entry, error)

	Watch(ctx context.Context, ch chan event.Event) error
}

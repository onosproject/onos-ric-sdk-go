// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package app

import (
	"context"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/store"
)

type Entry struct {
	Value interface{}
}

type Configuration interface {
	Config() interface{}

	Get(key string) (Entry, error)

	Watch(ctx context.Context, ch chan event.Event) error
}

type Config struct {
	config *store.ConfigStore
}

func (c *Config) Config() interface{} {
	return c.config.ConfigTree
}

func NewConfig(config *store.ConfigStore) *Config {
	return &Config{
		config: config,
	}
}

func (c *Config) Get(key string) (Entry, error) {
	entry, err := c.config.Get(key)
	if err != nil {
		return Entry{}, err
	}

	return Entry{Value: entry.Value}, nil
}

func (c *Config) Watch(ctx context.Context, ch chan event.Event) error {
	return c.config.Watch(ctx, ch)
}

var _ Configuration = &Config{}

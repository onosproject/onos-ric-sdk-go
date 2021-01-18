// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package _default

import (
	"context"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/app"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/store"
)

// Config config data structure
type Config struct {
	config *store.ConfigStore
}

// Config returns config tree
func (c *Config) Config() interface{} {
	return c.config.ConfigTree
}

// NewConfig creates a new configuration data structure
func NewConfig(config *store.ConfigStore) *Config {
	return &Config{
		config: config,
	}
}

// Get gets config value based on a given key
func (c *Config) Get(key string) (app.Entry, error) {
	entry, err := c.config.Get(key)
	if err != nil {
		return app.Entry{}, err
	}

	return app.Entry{Value: entry.Value}, nil
}

// Watch monitors config changes
func (c *Config) Watch(ctx context.Context, ch chan event.Event) error {
	return c.config.Watch(ctx, ch)
}

var _ app.Configuration = &Config{}

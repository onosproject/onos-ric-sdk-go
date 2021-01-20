// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package store

import (
	"context"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"

	"github.com/onosproject/onos-lib-go/pkg/logging"

	"github.com/google/uuid"
)

var log = logging.GetLogger("config", "store")

// Entry config entry
type Entry struct {
	Key       string
	Value     interface{}
	EventType string
}

type Store interface {
	Put(key string, entry Entry) error

	Get(key string) (Entry, error)

	Watch(ctx context.Context, ch chan event.Event) error
}

type ConfigStore struct {
	ConfigTree map[string]interface{}
	eventBus   *EventBus
}

func NewConfigStore() *ConfigStore {
	return &ConfigStore{
		ConfigTree: make(map[string]interface{}),
		eventBus: &EventBus{
			watchers: []ConfigTreeWatcher{},
		},
	}
}

func (c *ConfigStore) Watch(ctx context.Context, ch chan event.Event) error {
	c.eventBus.rm.Lock()
	cw := ConfigTreeWatcher{
		id: uuid.New(),
		ch: ch,
	}

	c.eventBus.watchers = append(c.eventBus.watchers, cw)
	c.eventBus.rm.Unlock()

	go func() {
		<-ctx.Done()
		c.eventBus.rm.Lock()
		watchers := make([]ConfigTreeWatcher, 0, len(c.eventBus.watchers)-1)
		for _, watcher := range c.eventBus.watchers {
			if watcher.id != cw.id {
				watchers = append(watchers, watcher)

			}
		}
		c.eventBus.watchers = watchers
		c.eventBus.rm.Unlock()
		close(ch)

	}()
	return nil
}

func (c *ConfigStore) Put(key string, entry Entry) error {
	err := put(c.ConfigTree, key, entry)
	if err != nil {
		return err
	}
	c.eventBus.send(event.Event{
		Key:       key,
		Value:     entry.Value,
		EventType: entry.EventType,
	})

	return nil
}

func (c *ConfigStore) Get(key string) (Entry, error) {
	node, err := get(c.ConfigTree, key)
	if err != nil {
		return Entry{}, err
	}
	return Entry{
		Key:   key,
		Value: node,
	}, nil
}

var _ Store = &ConfigStore{}

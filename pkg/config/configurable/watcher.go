// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package configurable

import (
	"sync"

	"github.com/google/uuid"
)

// Event config event
type Event struct {
	Key       string
	Value     interface{}
	EventType string
}

// EventChannel is a channel which can accept an Event
type EventChannel chan Event

// EventBus stores the information about watchers
type EventBus struct {
	watchers []ConfigTreeWatcher
	rm       sync.RWMutex
}

type ConfigTreeWatcher struct {
	id uuid.UUID
	ch chan Event
}

func (eb *EventBus) send(event Event) {
	eb.rm.RLock()
	for _, watcher := range eb.watchers {
		watcher.ch <- event
	}
	eb.rm.RUnlock()
}

// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package store

import (
	"sync"

	"github.com/google/uuid"
	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"
)

// EventChannel is a channel which can accept an Event
type EventChannel chan event.Event

// EventBus stores the information about watchers
type EventBus struct {
	watchers []ConfigTreeWatcher
	rm       sync.RWMutex
}

type ConfigTreeWatcher struct {
	id uuid.UUID
	ch chan event.Event
}

func (eb *EventBus) send(event event.Event) {
	eb.rm.RLock()
	for _, watcher := range eb.watchers {
		watcher.ch <- event
	}
	eb.rm.RUnlock()
}

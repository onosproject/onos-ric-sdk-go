// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

// FilterOptions describes the criteria for the type of topology entities and events that the application should be
// notified about
type FilterOption interface {
	matches(e topoapi.Event) bool
}

type eventFilterOption struct {
	eventTypes []topoapi.EventType
}

func (f *eventFilterOption) matches(e topoapi.Event) bool {
	for _, t := range f.eventTypes {
		if t == e.Type {
			return true
		}
	}
	return false
}

// WithEventFilter matches events of the specified event type
func WithEventFilter(t ...topoapi.EventType) FilterOption {
	return &eventFilterOption{
		eventTypes: t,
	}
}

type typeFilterOption struct {
	objectTypes []topoapi.Object_Type
}

func (f *typeFilterOption) matches(e topoapi.Event) bool {
	for _, t := range f.objectTypes {
		if t == e.Object.Type {
			return true
		}
	}
	return false
}

// WithTypeFilter matches events for objects of the specified type
func WithTypeFilter(t ...topoapi.Object_Type) FilterOption {
	return &typeFilterOption{
		objectTypes: t,
	}
}

type kindFilterOption struct {
	kindIDs []topoapi.ID
}

// WithKindFilter matches events for objects of the specified kind
func WithKindFilter(ids ...topoapi.ID) FilterOption {
	return &kindFilterOption{
		kindIDs: ids,
	}
}

func (f *kindFilterOption) matches(e topoapi.Event) bool {
	for _, k := range f.kindIDs {
		ok := e.Object.GetEntity() != nil && k == e.Object.GetEntity().KindID ||
			e.Object.GetRelation() != nil && k == e.Object.GetRelation().KindID
		if ok {
			return true
		}
	}
	return false
}

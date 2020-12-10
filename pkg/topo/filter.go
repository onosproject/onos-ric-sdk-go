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
	eventType topoapi.EventType
}

func (f *eventFilterOption) matches(e topoapi.Event) bool {
	return f.eventType == e.Type
}

// EventFilter matches events of the specified event type
func EventFilter(t topoapi.EventType) FilterOption {
	return &eventFilterOption{
		eventType: t,
	}
}

type typeFilterOption struct {
	objectType topoapi.Object_Type
}

func (f *typeFilterOption) matches(e topoapi.Event) bool {
	return f.objectType == e.Object.Type
}

// TypeFilter matches events for objects of the specified type
func TypeFilter(t topoapi.Object_Type) FilterOption {
	return &typeFilterOption{
		objectType: t,
	}
}

type kindFilterOption struct {
	kindID topoapi.ID
}

// KindFilter matches events for objects of the specified kind
func KindFilter(id topoapi.ID) FilterOption {
	return &kindFilterOption{
		kindID: id,
	}
}

func (f *kindFilterOption) matches(e topoapi.Event) bool {
	return e.Object.GetEntity() != nil && f.kindID == e.Object.GetEntity().KindID ||
		e.Object.GetRelation() != nil && f.kindID == e.Object.GetRelation().KindID
}

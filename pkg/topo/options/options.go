// SPDX-FileCopyrightText: ${year}-present Open Networking Foundation <info@opennetworking.org>
// SPDX-License-Identifier: Apache-2.0

package options

import topoapi "github.com/onosproject/onos-api/go/onos/topo"

// WatchOptions topo client watch method options
type WatchOptions struct {
	filters *topoapi.Filters
}

// WatchOption topo watch option function
type WatchOption func(options *WatchOptions)

// GetFilters get filters
func (w *WatchOptions) GetFilters() *topoapi.Filters {
	return w.filters
}

// WithWatchFilters sets filters for watch method
func WithWatchFilters(filters *topoapi.Filters) func(options *WatchOptions) {
	return func(options *WatchOptions) {
		options.filters = filters
	}
}

// ListOptions topo client get method options
type ListOptions struct {
	filters *topoapi.Filters
}

// GetFilters get filters
func (l *ListOptions) GetFilters() *topoapi.Filters {
	return l.filters
}

// ListOption topo get option function
type ListOption func(options *ListOptions)

// WithListFilters sets filters for list method
func WithListFilters(filters *topoapi.Filters) func(options *ListOptions) {
	return func(options *ListOptions) {
		options.filters = filters
	}
}

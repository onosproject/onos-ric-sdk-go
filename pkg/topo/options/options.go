// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package options

import (
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

// Options topo SDK options
type Options struct {
	Watch WatchOptions

	List ListOptions
}

// Option topo client
type Option interface {
	Apply(*Options)
}

type funcOption struct {
	f func(*Options)
}

func (f funcOption) Apply(options *Options) {
	f.f(options)
}

func newOption(f func(*Options)) Option {
	return funcOption{
		f: f,
	}
}

// WithOptions sets the client options
func WithOptions(opts Options) Option {
	return newOption(func(options *Options) {
		*options = opts
	})
}

// WatchOptions topo client watch method options
type WatchOptions struct {
	filters *topoapi.Filters
}

// GetFilters get filters
func (w WatchOptions) GetFilters() *topoapi.Filters {
	return w.filters
}

// WithWatchFilters sets filters for watch method
func WithWatchFilters(filters *topoapi.Filters) Option {
	return newOption(func(o *Options) {
		o.Watch.filters = filters
	})
}

// ListOptions topo client get method options
type ListOptions struct {
	filters *topoapi.Filters
}

// GetFilters get filters
func (l ListOptions) GetFilters() *topoapi.Filters {
	return l.filters
}

// WithListFilters sets filters for list method
func WithListFilters(filters *topoapi.Filters) Option {
	return newOption(func(o *Options) {
		o.List.filters = filters
	})

}

// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package options

import (
	"fmt"

	topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

const (
	defaultServicePort = 5150
	defaultServiceHost = "onos-topo"
)

// Options topo SDK options
type Options struct {
	// Watch watch options
	Watch WatchOptions

	// List list options
	List ListOptions

	// Service service options
	Service ServiceOptions
}

// ServiceOptions are the options for a service
type ServiceOptions struct {
	// Host is the service host
	Host string
	// Port is the service port
	Port int

	Insecure bool
}

// GetHost gets the service host
func (o ServiceOptions) GetHost() string {
	return o.Host
}

// GetPort gets the service port
func (o ServiceOptions) GetPort() int {
	if o.Port == 0 {
		return defaultServicePort
	}
	return o.Port
}

// IsInsecure is topo connection secure
func (o ServiceOptions) IsInsecure() bool {
	return o.Insecure
}

// GetAddress gets the service address
func (o ServiceOptions) GetAddress() string {
	return fmt.Sprintf("%s:%d", o.GetHost(), o.GetPort())
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

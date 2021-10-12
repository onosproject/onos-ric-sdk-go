// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package e2

import (
	"fmt"
	"time"
)

const defaultServicePort = 5150

// Encoding :
type Encoding int

const (
	// ProtoEncoding protobuf
	ProtoEncoding Encoding = iota

	// ASN1Encoding asn1
	ASN1Encoding
)

// Option is an E2 client option
type Option interface {
	apply(*Options)
}

// SubscribeOption is an option for subscribe request
type SubscribeOption interface {
	apply(*SubscribeOptions)
}

// EmptyOption is an empty client option
type EmptyOption struct{}

func (EmptyOption) apply(*Options) {}

// Options is a set of E2 client options
type Options struct {
	// AppOptions are the options for the application
	App AppOptions
	// ServiceMode is service model options
	ServiceModel ServiceModelOptions
	// Service is the E2 termination service configuration
	Service ServiceOptions
	// Topo is the topology service configuration
	Topo ServiceOptions
	// Encoding is the default encoding
	Encoding Encoding
}

// AppID is an application identifier
type AppID string

// InstanceID is an app instance identifier
type InstanceID string

// AppOptions are the options for the application
type AppOptions struct {
	// AppID is the application identifier
	AppID AppID
	// InstanceID is the application instance identifier
	InstanceID InstanceID
}

// ServiceOptions are the options for a service
type ServiceOptions struct {
	// Host is the service host
	Host string
	// Port is the service port
	Port int
}

// SubscribeOptions are the options for a subscription
type SubscribeOptions struct {
	// Port is the service port
	TransactionTimeout time.Duration
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

// GetAddress gets the service address
func (o ServiceOptions) GetAddress() string {
	return fmt.Sprintf("%s:%d", o.GetHost(), o.GetPort())
}

// ServiceModelName is a service model identifier
type ServiceModelName string

// ServiceModelVersion string
type ServiceModelVersion string

// ServiceModelOptions is options for defining a service model
type ServiceModelOptions struct {
	// Name is the service model identifier
	Name ServiceModelName

	// Version is the service model version
	Version ServiceModelVersion
}

type funcSubscribeOption struct {
	f func(*SubscribeOptions)
}

func (f funcSubscribeOption) apply(options *SubscribeOptions) {
	f.f(options)
}

func newSubscribeOption(f func(*SubscribeOptions)) SubscribeOption {
	return funcSubscribeOption{
		f: f,
	}
}

type funcOption struct {
	f func(*Options)
}

func (f funcOption) apply(options *Options) {
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

// WithAppID sets the client application identifier
func WithAppID(appID AppID) Option {
	return newOption(func(options *Options) {
		options.App.AppID = appID
	})
}

// WithInstanceID sets the client instance identifier
func WithInstanceID(instanceID InstanceID) Option {
	return newOption(func(options *Options) {
		options.App.InstanceID = instanceID
	})
}

// WithServiceModel sets the client service model
func WithServiceModel(name ServiceModelName, version ServiceModelVersion) Option {
	return newOption(func(options *Options) {
		options.ServiceModel = ServiceModelOptions{
			Name:    name,
			Version: version,
		}
	})
}

// WithEncoding sets the client encoding
func WithEncoding(encoding Encoding) Option {
	return newOption(func(options *Options) {
		options.Encoding = encoding
	})
}

// WithProtoEncoding sets the client encoding to ProtoEncoding
func WithProtoEncoding() Option {
	return WithEncoding(ProtoEncoding)
}

// WithASN1Encoding sets the client encoding to ASN1Encoding
func WithASN1Encoding() Option {
	return WithEncoding(ASN1Encoding)
}

// WithE2TAddress sets the address for the E2T service
func WithE2TAddress(host string, port int) Option {
	return newOption(func(options *Options) {
		options.Service.Host = host
		options.Service.Port = port
	})
}

// WithE2THost sets the host for the E2T service
func WithE2THost(host string) Option {
	return newOption(func(options *Options) {
		options.Service.Host = host
	})
}

// WithE2TPort sets the port for the E2T service
func WithE2TPort(port int) Option {
	return newOption(func(options *Options) {
		options.Service.Port = port
	})
}

// WithTransactionTimeout sets a timeout value for subscriptions
func WithTransactionTimeout(transactionTimeout time.Duration) SubscribeOption {
	return newSubscribeOption(func(options *SubscribeOptions) {
		options.TransactionTimeout = transactionTimeout
	})
}

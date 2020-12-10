// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

// Package topo contains facilities for RIC applications to monitor topology for changes
package topo

import (
	"context"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"google.golang.org/grpc/status"
	"io"
)

var log = logging.GetLogger("topo", "watch")

// Watch provides a simple facility for the application to watch for changes in the topology.
func (c *topoClient) Watch(ctx context.Context, ch chan<- topoapi.Event, filter ...FilterOption) error {
	// Issue a watch request with replay and passing through any service-side filter criteria, i.e.
	// those supported by the topo.Watch method (TBD)
	stream, err := c.client.Watch(ctx, &topoapi.WatchRequest{})
	if err != nil {
		log.Error("Unable to issue topology watch request", err)
		stat, ok := status.FromError(err)
		if ok {
			return errors.FromStatus(stat)
		}
		return err
	}

	// Kick off a go routine that reads events from the topology watch stream and subjects them to additional
	// client-side filter criteria; objects that pass the filter get put on the application provided channel.
	go handleWatchStream(ch, stream, filter...)
	return nil
}

// handleWatchStream reads events from the topology watch stream and subjects them to additional
// client-side filter criteria; objects that pass the filter get put on the application provided channel.
func handleWatchStream(ch chan<- topoapi.Event, stream topoapi.Topo_WatchClient, filter ...FilterOption) {
	defer close(ch)
	for {
		resp, err := stream.Recv()
		if err == io.EOF || err == context.Canceled {
			break
		}

		if err != nil {
			stat, ok := status.FromError(err)
			if ok {
				err = errors.FromStatus(stat)
				if errors.IsCanceled(err) || errors.IsTimeout(err) {
					break
				}
			}
			log.Error("An error occurred in receiving Subscription changes", err)
		} else if applyFilters(resp.Event, filter...) {
			ch <- resp.Event
		}
	}
}

func applyFilters(event topoapi.Event, filter ...FilterOption) bool {
	for _, f := range filter {
		if !f.matches(event) {
			return false
		}
	}
	return true
}

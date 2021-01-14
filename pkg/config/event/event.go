// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package event

// Event config event
type Event struct {
	Key       string
	Value     interface{}
	EventType string
}

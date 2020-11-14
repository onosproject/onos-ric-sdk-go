// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package v1beta1

import "google.golang.org/grpc"

// E2TClientFactory : Default E2TClient creation.
var E2TClientFactory = func(cc *grpc.ClientConn) E2TServiceClient {
	return NewE2TServiceClient(cc)
}

// CreateE2ServiceClient creates and returns a new config admin client
func CreateE2TServiceClient(cc *grpc.ClientConn) E2TServiceClient {
	return E2TClientFactory(cc)
}

// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package a1endpoint

func NewServer(caPath string, keyPath string, certPath string, grpcPort int) Server {
	return Server{
		CAPath: caPath,
		KeyPath: keyPath,
		CertPath: certPath,
		GRPCPort: grpcPort,
	}
}

type Server struct {
	CAPath string
	KeyPath string
	CertPath string
	GRPCPort int
}

// ToDo: Add server code here - should define xApp-A1 interface first
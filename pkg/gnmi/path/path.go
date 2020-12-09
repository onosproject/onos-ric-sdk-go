// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package path

import (
	"github.com/google/gnxi/utils/xpath"
	pb "github.com/openconfig/gnmi/proto/gnmi"
)

// Path represents configuration path in human readable form (e.g. x/y/z)
type Path struct {
	Value string
}

// ParsePaths converts xpaths to gnmi paths
func ParsePaths(paths []Path) ([]*pb.Path, error) {
	var pbPathList []*pb.Path
	for _, xPath := range paths {
		pbPath, err := xpath.ToGNMIPath(xPath.Value)
		if err != nil {
			return pbPathList, err
		}
		pbPathList = append(pbPathList, pbPath)
	}
	return pbPathList, nil
}

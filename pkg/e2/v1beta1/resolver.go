// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package e2

import (
	"context"
	"fmt"
	"github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/grpc/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
)

const resolverName = "e2"

func newResolver(nodeID NodeID, opts Options) resolver.Builder {
	return &ResolverBuilder{
		nodeID: nodeID,
		opts:   opts,
	}
}

type ResolverBuilder struct {
	nodeID NodeID
	opts   Options
}

func (b *ResolverBuilder) Scheme() string {
	return resolverName
}

func (b *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	var dialOpts []grpc.DialOption
	if opts.DialCreds != nil {
		dialOpts = append(
			dialOpts,
			grpc.WithTransportCredentials(opts.DialCreds),
		)
	} else {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	}
	dialOpts = append(dialOpts, grpc.WithUnaryInterceptor(retry.RetryingUnaryClientInterceptor(retry.WithRetryOn(codes.Unavailable, codes.Unknown))))
	dialOpts = append(dialOpts, grpc.WithStreamInterceptor(retry.RetryingStreamClientInterceptor(retry.WithRetryOn(codes.Unavailable, codes.Unknown))))
	dialOpts = append(dialOpts, grpc.WithContextDialer(opts.Dialer))

	resolverConn, err := grpc.Dial(b.opts.Topo.GetAddress(), dialOpts...)
	if err != nil {
		return nil, err
	}

	serviceConfig := cc.ParseServiceConfig(
		fmt.Sprintf(`{"loadBalancingConfig":[{"%s":{}}]}`, resolverName),
	)

	resolver := &Resolver{
		nodeID:        b.nodeID,
		clientConn:    cc,
		resolverConn:  resolverConn,
		serviceConfig: serviceConfig,
	}
	err = resolver.start()
	if err != nil {
		return nil, err
	}
	return resolver, nil
}

var _ resolver.Builder = (*ResolverBuilder)(nil)

type Resolver struct {
	nodeID        NodeID
	clientConn    resolver.ClientConn
	resolverConn  *grpc.ClientConn
	serviceConfig *serviceconfig.ParseResult
}

func (r *Resolver) start() error {
	client := topo.NewTopoClient(r.resolverConn)
	request := &topo.WatchRequest{}
	stream, err := client.Watch(context.Background(), request)
	if err != nil {
		return err
	}
	go func() {
		var mastership *topo.MastershipState
		e2tNodes := make(map[topo.ID]string)
		for {
			response, err := stream.Recv()
			if err != nil {
				return
			}

			object := response.Event.Object
			if entity, ok := object.Obj.(*topo.Object_Entity); ok &&
				entity.Entity.KindID == topo.E2NODE &&
				object.ID == topo.ID(r.nodeID) {
				var m topo.MastershipState
				_ = object.GetAspect(&m)
				if m.NodeId != "" && (mastership == nil || m.Term > mastership.Term) {
					mastership = &m
					address, ok := e2tNodes[topo.ID(mastership.NodeId)]
					if ok {
						var addrs []resolver.Address
						addrs = append(addrs, resolver.Address{
							Addr: address,
							Attributes: attributes.New(
								"is_master",
								true,
							),
						})

						for nodeID, address := range e2tNodes {
							if nodeID != topo.ID(m.NodeId) {
								addrs = append(addrs, resolver.Address{
									Addr: address,
									Attributes: attributes.New(
										"is_master",
										false,
									),
								})
							}
						}

						r.clientConn.UpdateState(resolver.State{
							Addresses:     addrs,
							ServiceConfig: r.serviceConfig,
						})
					}
				}
			} else if entity, ok := object.Obj.(*topo.Object_Entity); ok &&
				entity.Entity.KindID == topo.E2T {
				switch response.Event.Type {
				case topo.EventType_REMOVED:
					delete(e2tNodes, object.ID)
				default:
					var info topo.E2TInfo
					_ = object.GetAspect(&info)
					for _, iface := range info.Interfaces {
						if iface.Type == topo.Interface_INTERFACE_E2T {
							e2tNodes[object.ID] = fmt.Sprintf("%s:%d", iface.IP, iface.Port)
						}
					}
				}
			}
		}
	}()
	return nil
}

func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {}

func (r *Resolver) Close() {
	if err := r.resolverConn.Close(); err != nil {
		log.Error("failed to close conn", err)
	}
}

var _ resolver.Resolver = (*Resolver)(nil)

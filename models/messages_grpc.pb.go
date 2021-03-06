// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package models

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClusterInfoClient is the client API for ClusterInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterInfoClient interface {
	GetClusterInfo(ctx context.Context, in *WhatYouKnow, opts ...grpc.CallOption) (*IKnow, error)
}

type clusterInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterInfoClient(cc grpc.ClientConnInterface) ClusterInfoClient {
	return &clusterInfoClient{cc}
}

func (c *clusterInfoClient) GetClusterInfo(ctx context.Context, in *WhatYouKnow, opts ...grpc.CallOption) (*IKnow, error) {
	out := new(IKnow)
	err := c.cc.Invoke(ctx, "/models.ClusterInfo/GetClusterInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterInfoServer is the server API for ClusterInfo service.
// All implementations must embed UnimplementedClusterInfoServer
// for forward compatibility
type ClusterInfoServer interface {
	GetClusterInfo(context.Context, *WhatYouKnow) (*IKnow, error)
	mustEmbedUnimplementedClusterInfoServer()
}

// UnimplementedClusterInfoServer must be embedded to have forward compatible implementations.
type UnimplementedClusterInfoServer struct {
}

func (UnimplementedClusterInfoServer) GetClusterInfo(context.Context, *WhatYouKnow) (*IKnow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterInfo not implemented")
}
func (UnimplementedClusterInfoServer) mustEmbedUnimplementedClusterInfoServer() {}

// UnsafeClusterInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterInfoServer will
// result in compilation errors.
type UnsafeClusterInfoServer interface {
	mustEmbedUnimplementedClusterInfoServer()
}

func RegisterClusterInfoServer(s grpc.ServiceRegistrar, srv ClusterInfoServer) {
	s.RegisterService(&ClusterInfo_ServiceDesc, srv)
}

func _ClusterInfo_GetClusterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhatYouKnow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterInfoServer).GetClusterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.ClusterInfo/GetClusterInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterInfoServer).GetClusterInfo(ctx, req.(*WhatYouKnow))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterInfo_ServiceDesc is the grpc.ServiceDesc for ClusterInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "models.ClusterInfo",
	HandlerType: (*ClusterInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClusterInfo",
			Handler:    _ClusterInfo_GetClusterInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages.proto",
}

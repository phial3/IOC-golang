// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package boot

import (
	context "context"

	grpc "google.golang.org/grpc"

	codes "google.golang.org/grpc/codes"

	status "google.golang.org/grpc/status"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DebugServiceClient is the client API for DebugService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DebugServiceClient interface {
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (DebugService_WatchClient, error)
	ListServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListServiceResponse, error)
	WatchEdit(ctx context.Context, opts ...grpc.CallOption) (DebugService_WatchEditClient, error)
}

type debugServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugServiceClient(cc grpc.ClientConnInterface) DebugServiceClient {
	return &debugServiceClient{cc}
}

func (c *debugServiceClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (DebugService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &DebugService_ServiceDesc.Streams[0], "/ioc_golang.boot.DebugService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DebugService_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type debugServiceWatchClient struct {
	grpc.ClientStream
}

func (x *debugServiceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugServiceClient) ListServices(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListServiceResponse, error) {
	out := new(ListServiceResponse)
	err := c.cc.Invoke(ctx, "/ioc_golang.boot.DebugService/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugServiceClient) WatchEdit(ctx context.Context, opts ...grpc.CallOption) (DebugService_WatchEditClient, error) {
	stream, err := c.cc.NewStream(ctx, &DebugService_ServiceDesc.Streams[1], "/ioc_golang.boot.DebugService/WatchEdit", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugServiceWatchEditClient{stream}
	return x, nil
}

type DebugService_WatchEditClient interface {
	Send(*WatchEditRequest) error
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type debugServiceWatchEditClient struct {
	grpc.ClientStream
}

func (x *debugServiceWatchEditClient) Send(m *WatchEditRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *debugServiceWatchEditClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DebugServiceServer is the server API for DebugService service.
// All implementations must embed UnimplementedDebugServiceServer
// for forward compatibility
type DebugServiceServer interface {
	Watch(*WatchRequest, DebugService_WatchServer) error
	ListServices(context.Context, *emptypb.Empty) (*ListServiceResponse, error)
	WatchEdit(DebugService_WatchEditServer) error
	mustEmbedUnimplementedDebugServiceServer()
}

// UnimplementedDebugServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDebugServiceServer struct {
}

func (UnimplementedDebugServiceServer) Watch(*WatchRequest, DebugService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedDebugServiceServer) ListServices(context.Context, *emptypb.Empty) (*ListServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServices not implemented")
}
func (UnimplementedDebugServiceServer) WatchEdit(DebugService_WatchEditServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchEdit not implemented")
}
func (UnimplementedDebugServiceServer) mustEmbedUnimplementedDebugServiceServer() {}

// UnsafeDebugServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DebugServiceServer will
// result in compilation errors.
type UnsafeDebugServiceServer interface {
	mustEmbedUnimplementedDebugServiceServer()
}

func RegisterDebugServiceServer(s grpc.ServiceRegistrar, srv DebugServiceServer) {
	s.RegisterService(&DebugService_ServiceDesc, srv)
}

func _DebugService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServiceServer).Watch(m, &debugServiceWatchServer{stream})
}

type DebugService_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type debugServiceWatchServer struct {
	grpc.ServerStream
}

func (x *debugServiceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DebugService_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServiceServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ioc_golang.boot.DebugService/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServiceServer).ListServices(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebugService_WatchEdit_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DebugServiceServer).WatchEdit(&debugServiceWatchEditServer{stream})
}

type DebugService_WatchEditServer interface {
	Send(*WatchResponse) error
	Recv() (*WatchEditRequest, error)
	grpc.ServerStream
}

type debugServiceWatchEditServer struct {
	grpc.ServerStream
}

func (x *debugServiceWatchEditServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *debugServiceWatchEditServer) Recv() (*WatchEditRequest, error) {
	m := new(WatchEditRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DebugService_ServiceDesc is the grpc.ServiceDesc for DebugService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DebugService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ioc_golang.boot.DebugService",
	HandlerType: (*DebugServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _DebugService_ListServices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _DebugService_Watch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WatchEdit",
			Handler:       _DebugService_WatchEdit_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "debug/api/ioc_golang/boot/debug.proto",
}
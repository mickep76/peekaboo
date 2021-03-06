// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/ake-persson/peekaboo/pkg/pb/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SystemServiceClient is the client API for SystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemServiceClient interface {
	GetSystem(ctx context.Context, in *GetSystemRequest, opts ...grpc.CallOption) (*resources.System, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error)
	ListFilesystems(ctx context.Context, in *ListFilesystemsRequest, opts ...grpc.CallOption) (*ListFilesystemsResponse, error)
}

type systemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemServiceClient(cc grpc.ClientConnInterface) SystemServiceClient {
	return &systemServiceClient{cc}
}

func (c *systemServiceClient) GetSystem(ctx context.Context, in *GetSystemRequest, opts ...grpc.CallOption) (*resources.System, error) {
	out := new(resources.System)
	err := c.cc.Invoke(ctx, "/peekaboo.v1.services.SystemService/GetSystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/peekaboo.v1.services.SystemService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error) {
	out := new(ListGroupsResponse)
	err := c.cc.Invoke(ctx, "/peekaboo.v1.services.SystemService/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ListFilesystems(ctx context.Context, in *ListFilesystemsRequest, opts ...grpc.CallOption) (*ListFilesystemsResponse, error) {
	out := new(ListFilesystemsResponse)
	err := c.cc.Invoke(ctx, "/peekaboo.v1.services.SystemService/ListFilesystems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServiceServer is the server API for SystemService service.
// All implementations must embed UnimplementedSystemServiceServer
// for forward compatibility
type SystemServiceServer interface {
	GetSystem(context.Context, *GetSystemRequest) (*resources.System, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error)
	ListFilesystems(context.Context, *ListFilesystemsRequest) (*ListFilesystemsResponse, error)
	mustEmbedUnimplementedSystemServiceServer()
}

// UnimplementedSystemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServiceServer struct {
}

func (UnimplementedSystemServiceServer) GetSystem(context.Context, *GetSystemRequest) (*resources.System, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystem not implemented")
}
func (UnimplementedSystemServiceServer) ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedSystemServiceServer) ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (UnimplementedSystemServiceServer) ListFilesystems(context.Context, *ListFilesystemsRequest) (*ListFilesystemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFilesystems not implemented")
}
func (UnimplementedSystemServiceServer) mustEmbedUnimplementedSystemServiceServer() {}

// UnsafeSystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemServiceServer will
// result in compilation errors.
type UnsafeSystemServiceServer interface {
	mustEmbedUnimplementedSystemServiceServer()
}

func RegisterSystemServiceServer(s grpc.ServiceRegistrar, srv SystemServiceServer) {
	s.RegisterService(&SystemService_ServiceDesc, srv)
}

func _SystemService_GetSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).GetSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peekaboo.v1.services.SystemService/GetSystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).GetSystem(ctx, req.(*GetSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peekaboo.v1.services.SystemService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peekaboo.v1.services.SystemService/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ListGroups(ctx, req.(*ListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ListFilesystems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesystemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ListFilesystems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peekaboo.v1.services.SystemService/ListFilesystems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ListFilesystems(ctx, req.(*ListFilesystemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemService_ServiceDesc is the grpc.ServiceDesc for SystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "peekaboo.v1.services.SystemService",
	HandlerType: (*SystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSystem",
			Handler:    _SystemService_GetSystem_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _SystemService_ListUsers_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _SystemService_ListGroups_Handler,
		},
		{
			MethodName: "ListFilesystems",
			Handler:    _SystemService_ListFilesystems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/v1/services/system.proto",
}

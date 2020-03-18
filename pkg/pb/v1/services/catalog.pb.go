// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/v1/services/catalog.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	resources "github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ListSystemsRequest struct {
	Filter               *resources.Filter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter" csv:"filter" yaml:"filter"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" csv:"-" yaml:"-"`
	XXX_unrecognized     []byte            `json:"-" csv:"-" yaml:"-"`
	XXX_sizecache        int32             `json:"-" csv:"-" yaml:"-"`
}

func (m *ListSystemsRequest) Reset()         { *m = ListSystemsRequest{} }
func (m *ListSystemsRequest) String() string { return proto.CompactTextString(m) }
func (*ListSystemsRequest) ProtoMessage()    {}
func (*ListSystemsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a211ff456e3853e, []int{0}
}

func (m *ListSystemsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSystemsRequest.Unmarshal(m, b)
}
func (m *ListSystemsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSystemsRequest.Marshal(b, m, deterministic)
}
func (m *ListSystemsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSystemsRequest.Merge(m, src)
}
func (m *ListSystemsRequest) XXX_Size() int {
	return xxx_messageInfo_ListSystemsRequest.Size(m)
}
func (m *ListSystemsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSystemsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSystemsRequest proto.InternalMessageInfo

func (m *ListSystemsRequest) GetFilter() *resources.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type SystemKeepAliveRequest struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname" csv:"hostname" yaml:"hostname"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" csv:"-" yaml:"-"`
	XXX_unrecognized     []byte   `json:"-" csv:"-" yaml:"-"`
	XXX_sizecache        int32    `json:"-" csv:"-" yaml:"-"`
}

func (m *SystemKeepAliveRequest) Reset()         { *m = SystemKeepAliveRequest{} }
func (m *SystemKeepAliveRequest) String() string { return proto.CompactTextString(m) }
func (*SystemKeepAliveRequest) ProtoMessage()    {}
func (*SystemKeepAliveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a211ff456e3853e, []int{1}
}

func (m *SystemKeepAliveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemKeepAliveRequest.Unmarshal(m, b)
}
func (m *SystemKeepAliveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemKeepAliveRequest.Marshal(b, m, deterministic)
}
func (m *SystemKeepAliveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemKeepAliveRequest.Merge(m, src)
}
func (m *SystemKeepAliveRequest) XXX_Size() int {
	return xxx_messageInfo_SystemKeepAliveRequest.Size(m)
}
func (m *SystemKeepAliveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemKeepAliveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SystemKeepAliveRequest proto.InternalMessageInfo

func (m *SystemKeepAliveRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterType((*ListSystemsRequest)(nil), "peekaboo.v1.services.ListSystemsRequest")
	proto.RegisterType((*SystemKeepAliveRequest)(nil), "peekaboo.v1.services.SystemKeepAliveRequest")
}

func init() { proto.RegisterFile("pb/v1/services/catalog.proto", fileDescriptor_2a211ff456e3853e) }

var fileDescriptor_2a211ff456e3853e = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xa9, 0x87, 0xa1, 0x19, 0x4c, 0x08, 0x32, 0xa4, 0x2a, 0xc8, 0x4e, 0x3b, 0x68, 0xe2,
	0xa6, 0x9e, 0x44, 0x44, 0x45, 0x3d, 0xcc, 0x53, 0x77, 0x50, 0xbc, 0x35, 0xe5, 0x2d, 0x0b, 0x6b,
	0x4d, 0xcc, 0x4b, 0x0b, 0xfb, 0x0e, 0x7e, 0x68, 0x59, 0xb3, 0x96, 0xa9, 0xed, 0x6e, 0x7d, 0x7d,
	0xff, 0xfe, 0xf8, 0xbf, 0x1f, 0x25, 0xc7, 0x46, 0xf0, 0x62, 0xc4, 0x11, 0x6c, 0xa1, 0x12, 0x40,
	0x9e, 0xc4, 0x2e, 0x4e, 0xb5, 0x64, 0xc6, 0x6a, 0xa7, 0xe9, 0x81, 0x01, 0x58, 0xc4, 0x42, 0x6b,
	0x56, 0x8c, 0x58, 0x95, 0x09, 0x8f, 0xa4, 0xd6, 0x32, 0x05, 0x5e, 0x66, 0x44, 0x3e, 0xe3, 0x90,
	0x19, 0xb7, 0xf4, 0x9f, 0x84, 0x6b, 0xa0, 0x05, 0xd4, 0xb9, 0x5d, 0x11, 0x67, 0x2a, 0x75, 0x60,
	0xdb, 0xb6, 0xb8, 0x44, 0x07, 0x99, 0xdf, 0x0e, 0x26, 0x84, 0xbe, 0x2a, 0x74, 0xd3, 0xf2, 0x1d,
	0x46, 0xf0, 0x95, 0x03, 0x3a, 0x7a, 0x4d, 0x3a, 0x9e, 0x71, 0x18, 0x9c, 0x06, 0xc3, 0xee, 0xf8,
	0x84, 0x6d, 0xb6, 0xaa, 0x51, 0xec, 0xb9, 0x0c, 0x45, 0xeb, 0xf0, 0xe0, 0x8a, 0xf4, 0x3d, 0x68,
	0x02, 0x60, 0xee, 0x53, 0x55, 0x40, 0x05, 0x0c, 0xc9, 0xee, 0x5c, 0xa3, 0xfb, 0x8c, 0x33, 0x28,
	0x91, 0x7b, 0x51, 0x3d, 0x8f, 0xbf, 0x77, 0x48, 0xef, 0xd1, 0x3b, 0x98, 0xfa, 0x7b, 0xe9, 0x0b,
	0xe9, 0x45, 0x20, 0x15, 0x3a, 0xb0, 0x1e, 0x48, 0xdb, 0x1a, 0xf8, 0x75, 0xd8, 0x67, 0x5e, 0x10,
	0xab, 0x04, 0xb1, 0xa7, 0x95, 0x20, 0xfa, 0x46, 0xf6, 0xff, 0x34, 0xa2, 0x67, 0xac, 0xc9, 0x30,
	0x6b, 0x2e, 0xde, 0x0a, 0x7e, 0x27, 0xdd, 0x0d, 0x6f, 0x74, 0xd8, 0x0c, 0xfd, 0xaf, 0x36, 0xdc,
	0x7e, 0xc8, 0x45, 0xf0, 0x70, 0xf7, 0x71, 0x2b, 0x95, 0x9b, 0xe7, 0x82, 0x25, 0x3a, 0xe3, 0x55,
	0xf8, 0x3c, 0x8d, 0x05, 0xd6, 0x13, 0x37, 0x0b, 0xc9, 0x7f, 0xff, 0x46, 0x37, 0xd5, 0x83, 0xe8,
	0x94, 0x55, 0x2f, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x54, 0xa8, 0x43, 0x68, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatalogServiceClient is the client API for CatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatalogServiceClient interface {
	RegisterSystem(ctx context.Context, in *resources.System, opts ...grpc.CallOption) (*empty.Empty, error)
	SystemKeepAlive(ctx context.Context, in *SystemKeepAliveRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListSystems(ctx context.Context, in *ListSystemsRequest, opts ...grpc.CallOption) (CatalogService_ListSystemsClient, error)
}

type catalogServiceClient struct {
	cc *grpc.ClientConn
}

func NewCatalogServiceClient(cc *grpc.ClientConn) CatalogServiceClient {
	return &catalogServiceClient{cc}
}

func (c *catalogServiceClient) RegisterSystem(ctx context.Context, in *resources.System, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/peekaboo.v1.services.CatalogService/RegisterSystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) SystemKeepAlive(ctx context.Context, in *SystemKeepAliveRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/peekaboo.v1.services.CatalogService/SystemKeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catalogServiceClient) ListSystems(ctx context.Context, in *ListSystemsRequest, opts ...grpc.CallOption) (CatalogService_ListSystemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CatalogService_serviceDesc.Streams[0], "/peekaboo.v1.services.CatalogService/ListSystems", opts...)
	if err != nil {
		return nil, err
	}
	x := &catalogServiceListSystemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CatalogService_ListSystemsClient interface {
	Recv() (*resources.System, error)
	grpc.ClientStream
}

type catalogServiceListSystemsClient struct {
	grpc.ClientStream
}

func (x *catalogServiceListSystemsClient) Recv() (*resources.System, error) {
	m := new(resources.System)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CatalogServiceServer is the server API for CatalogService service.
type CatalogServiceServer interface {
	RegisterSystem(context.Context, *resources.System) (*empty.Empty, error)
	SystemKeepAlive(context.Context, *SystemKeepAliveRequest) (*empty.Empty, error)
	ListSystems(*ListSystemsRequest, CatalogService_ListSystemsServer) error
}

// UnimplementedCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatalogServiceServer struct {
}

func (*UnimplementedCatalogServiceServer) RegisterSystem(ctx context.Context, req *resources.System) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSystem not implemented")
}
func (*UnimplementedCatalogServiceServer) SystemKeepAlive(ctx context.Context, req *SystemKeepAliveRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemKeepAlive not implemented")
}
func (*UnimplementedCatalogServiceServer) ListSystems(req *ListSystemsRequest, srv CatalogService_ListSystemsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListSystems not implemented")
}

func RegisterCatalogServiceServer(s *grpc.Server, srv CatalogServiceServer) {
	s.RegisterService(&_CatalogService_serviceDesc, srv)
}

func _CatalogService_RegisterSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.System)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).RegisterSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peekaboo.v1.services.CatalogService/RegisterSystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).RegisterSystem(ctx, req.(*resources.System))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_SystemKeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SystemKeepAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatalogServiceServer).SystemKeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peekaboo.v1.services.CatalogService/SystemKeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatalogServiceServer).SystemKeepAlive(ctx, req.(*SystemKeepAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CatalogService_ListSystems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListSystemsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CatalogServiceServer).ListSystems(m, &catalogServiceListSystemsServer{stream})
}

type CatalogService_ListSystemsServer interface {
	Send(*resources.System) error
	grpc.ServerStream
}

type catalogServiceListSystemsServer struct {
	grpc.ServerStream
}

func (x *catalogServiceListSystemsServer) Send(m *resources.System) error {
	return x.ServerStream.SendMsg(m)
}

var _CatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "peekaboo.v1.services.CatalogService",
	HandlerType: (*CatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterSystem",
			Handler:    _CatalogService_RegisterSystem_Handler,
		},
		{
			MethodName: "SystemKeepAlive",
			Handler:    _CatalogService_SystemKeepAlive_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListSystems",
			Handler:       _CatalogService_ListSystems_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/v1/services/catalog.proto",
}

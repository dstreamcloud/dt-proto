// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package agent_endpoint_pb

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

// V1Client is the client API for V1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type V1Client interface {
	RegisterV1(ctx context.Context, in *RegisterV1_Request, opts ...grpc.CallOption) (*RegisterV1_Response, error)
	SubscribeV1(ctx context.Context, in *SubscribeV1_Request, opts ...grpc.CallOption) (V1_SubscribeV1Client, error)
	ReportJobV1(ctx context.Context, in *ReportJobV1_Request, opts ...grpc.CallOption) (*ReportJobV1_Response, error)
	ReportStatV1(ctx context.Context, in *ReportStatV1_Request, opts ...grpc.CallOption) (*ReportStatV1_Response, error)
}

type v1Client struct {
	cc grpc.ClientConnInterface
}

func NewV1Client(cc grpc.ClientConnInterface) V1Client {
	return &v1Client{cc}
}

func (c *v1Client) RegisterV1(ctx context.Context, in *RegisterV1_Request, opts ...grpc.CallOption) (*RegisterV1_Response, error) {
	out := new(RegisterV1_Response)
	err := c.cc.Invoke(ctx, "/agent_endpoint.V1/RegisterV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1Client) SubscribeV1(ctx context.Context, in *SubscribeV1_Request, opts ...grpc.CallOption) (V1_SubscribeV1Client, error) {
	stream, err := c.cc.NewStream(ctx, &V1_ServiceDesc.Streams[0], "/agent_endpoint.V1/SubscribeV1", opts...)
	if err != nil {
		return nil, err
	}
	x := &v1SubscribeV1Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type V1_SubscribeV1Client interface {
	Recv() (*SubscribeV1_Response, error)
	grpc.ClientStream
}

type v1SubscribeV1Client struct {
	grpc.ClientStream
}

func (x *v1SubscribeV1Client) Recv() (*SubscribeV1_Response, error) {
	m := new(SubscribeV1_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *v1Client) ReportJobV1(ctx context.Context, in *ReportJobV1_Request, opts ...grpc.CallOption) (*ReportJobV1_Response, error) {
	out := new(ReportJobV1_Response)
	err := c.cc.Invoke(ctx, "/agent_endpoint.V1/ReportJobV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *v1Client) ReportStatV1(ctx context.Context, in *ReportStatV1_Request, opts ...grpc.CallOption) (*ReportStatV1_Response, error) {
	out := new(ReportStatV1_Response)
	err := c.cc.Invoke(ctx, "/agent_endpoint.V1/ReportStatV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// V1Server is the server API for V1 service.
// All implementations must embed UnimplementedV1Server
// for forward compatibility
type V1Server interface {
	RegisterV1(context.Context, *RegisterV1_Request) (*RegisterV1_Response, error)
	SubscribeV1(*SubscribeV1_Request, V1_SubscribeV1Server) error
	ReportJobV1(context.Context, *ReportJobV1_Request) (*ReportJobV1_Response, error)
	ReportStatV1(context.Context, *ReportStatV1_Request) (*ReportStatV1_Response, error)
	mustEmbedUnimplementedV1Server()
}

// UnimplementedV1Server must be embedded to have forward compatible implementations.
type UnimplementedV1Server struct {
}

func (UnimplementedV1Server) RegisterV1(context.Context, *RegisterV1_Request) (*RegisterV1_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterV1 not implemented")
}
func (UnimplementedV1Server) SubscribeV1(*SubscribeV1_Request, V1_SubscribeV1Server) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeV1 not implemented")
}
func (UnimplementedV1Server) ReportJobV1(context.Context, *ReportJobV1_Request) (*ReportJobV1_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportJobV1 not implemented")
}
func (UnimplementedV1Server) ReportStatV1(context.Context, *ReportStatV1_Request) (*ReportStatV1_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportStatV1 not implemented")
}
func (UnimplementedV1Server) mustEmbedUnimplementedV1Server() {}

// UnsafeV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to V1Server will
// result in compilation errors.
type UnsafeV1Server interface {
	mustEmbedUnimplementedV1Server()
}

func RegisterV1Server(s grpc.ServiceRegistrar, srv V1Server) {
	s.RegisterService(&V1_ServiceDesc, srv)
}

func _V1_RegisterV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterV1_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1Server).RegisterV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent_endpoint.V1/RegisterV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1Server).RegisterV1(ctx, req.(*RegisterV1_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1_SubscribeV1_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeV1_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(V1Server).SubscribeV1(m, &v1SubscribeV1Server{stream})
}

type V1_SubscribeV1Server interface {
	Send(*SubscribeV1_Response) error
	grpc.ServerStream
}

type v1SubscribeV1Server struct {
	grpc.ServerStream
}

func (x *v1SubscribeV1Server) Send(m *SubscribeV1_Response) error {
	return x.ServerStream.SendMsg(m)
}

func _V1_ReportJobV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportJobV1_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1Server).ReportJobV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent_endpoint.V1/ReportJobV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1Server).ReportJobV1(ctx, req.(*ReportJobV1_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _V1_ReportStatV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportStatV1_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(V1Server).ReportStatV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent_endpoint.V1/ReportStatV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(V1Server).ReportStatV1(ctx, req.(*ReportStatV1_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// V1_ServiceDesc is the grpc.ServiceDesc for V1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var V1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agent_endpoint.V1",
	HandlerType: (*V1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterV1",
			Handler:    _V1_RegisterV1_Handler,
		},
		{
			MethodName: "ReportJobV1",
			Handler:    _V1_ReportJobV1_Handler,
		},
		{
			MethodName: "ReportStatV1",
			Handler:    _V1_ReportStatV1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeV1",
			Handler:       _V1_SubscribeV1_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/agent-endpoint/v1.proto",
}

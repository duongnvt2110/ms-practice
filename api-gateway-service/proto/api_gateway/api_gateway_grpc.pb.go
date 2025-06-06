// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: api_gateway/api_gateway.proto

package api_gateway

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WebA_Proxy_FullMethodName = "/api_gateway.WebA/Proxy"
)

// WebAClient is the client API for WebA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The greeting service definition
type WebAClient interface {
	// Sends a greeting
	Proxy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type webAClient struct {
	cc grpc.ClientConnInterface
}

func NewWebAClient(cc grpc.ClientConnInterface) WebAClient {
	return &webAClient{cc}
}

func (c *webAClient) Proxy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, WebA_Proxy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebAServer is the server API for WebA service.
// All implementations must embed UnimplementedWebAServer
// for forward compatibility.
//
// The greeting service definition
type WebAServer interface {
	// Sends a greeting
	Proxy(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedWebAServer()
}

// UnimplementedWebAServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWebAServer struct{}

func (UnimplementedWebAServer) Proxy(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proxy not implemented")
}
func (UnimplementedWebAServer) mustEmbedUnimplementedWebAServer() {}
func (UnimplementedWebAServer) testEmbeddedByValue()              {}

// UnsafeWebAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebAServer will
// result in compilation errors.
type UnsafeWebAServer interface {
	mustEmbedUnimplementedWebAServer()
}

func RegisterWebAServer(s grpc.ServiceRegistrar, srv WebAServer) {
	// If the following call pancis, it indicates UnimplementedWebAServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WebA_ServiceDesc, srv)
}

func _WebA_Proxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebAServer).Proxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebA_Proxy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebAServer).Proxy(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// WebA_ServiceDesc is the grpc.ServiceDesc for WebA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api_gateway.WebA",
	HandlerType: (*WebAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proxy",
			Handler:    _WebA_Proxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_gateway/api_gateway.proto",
}

const (
	WebB_Proxy_FullMethodName = "/api_gateway.WebB/Proxy"
)

// WebBClient is the client API for WebB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebBClient interface {
	// Sends a greeting
	Proxy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type webBClient struct {
	cc grpc.ClientConnInterface
}

func NewWebBClient(cc grpc.ClientConnInterface) WebBClient {
	return &webBClient{cc}
}

func (c *webBClient) Proxy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, WebB_Proxy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebBServer is the server API for WebB service.
// All implementations must embed UnimplementedWebBServer
// for forward compatibility.
type WebBServer interface {
	// Sends a greeting
	Proxy(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedWebBServer()
}

// UnimplementedWebBServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWebBServer struct{}

func (UnimplementedWebBServer) Proxy(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Proxy not implemented")
}
func (UnimplementedWebBServer) mustEmbedUnimplementedWebBServer() {}
func (UnimplementedWebBServer) testEmbeddedByValue()              {}

// UnsafeWebBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebBServer will
// result in compilation errors.
type UnsafeWebBServer interface {
	mustEmbedUnimplementedWebBServer()
}

func RegisterWebBServer(s grpc.ServiceRegistrar, srv WebBServer) {
	// If the following call pancis, it indicates UnimplementedWebBServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WebB_ServiceDesc, srv)
}

func _WebB_Proxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebBServer).Proxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebB_Proxy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebBServer).Proxy(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// WebB_ServiceDesc is the grpc.ServiceDesc for WebB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api_gateway.WebB",
	HandlerType: (*WebBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Proxy",
			Handler:    _WebB_Proxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_gateway/api_gateway.proto",
}

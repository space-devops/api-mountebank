// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/planetservices.proto

package space

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PlanetServiceClient is the client API for PlanetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlanetServiceClient interface {
	GetWelcome(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WelcomeMessage, error)
	GetPlanetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PlanetList, error)
	GetPlanetDetails(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*PlanetDetails, error)
	GetSecrets(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SecretMessage, error)
}

type planetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlanetServiceClient(cc grpc.ClientConnInterface) PlanetServiceClient {
	return &planetServiceClient{cc}
}

func (c *planetServiceClient) GetWelcome(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*WelcomeMessage, error) {
	out := new(WelcomeMessage)
	err := c.cc.Invoke(ctx, "/space.PlanetService/getWelcome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planetServiceClient) GetPlanetList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PlanetList, error) {
	out := new(PlanetList)
	err := c.cc.Invoke(ctx, "/space.PlanetService/getPlanetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planetServiceClient) GetPlanetDetails(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*PlanetDetails, error) {
	out := new(PlanetDetails)
	err := c.cc.Invoke(ctx, "/space.PlanetService/getPlanetDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *planetServiceClient) GetSecrets(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SecretMessage, error) {
	out := new(SecretMessage)
	err := c.cc.Invoke(ctx, "/space.PlanetService/getSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlanetServiceServer is the server API for PlanetService service.
// All implementations must embed UnimplementedPlanetServiceServer
// for forward compatibility
type PlanetServiceServer interface {
	GetWelcome(context.Context, *emptypb.Empty) (*WelcomeMessage, error)
	GetPlanetList(context.Context, *emptypb.Empty) (*PlanetList, error)
	GetPlanetDetails(context.Context, *wrapperspb.StringValue) (*PlanetDetails, error)
	GetSecrets(context.Context, *emptypb.Empty) (*SecretMessage, error)
	mustEmbedUnimplementedPlanetServiceServer()
}

// UnimplementedPlanetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPlanetServiceServer struct {
}

func (UnimplementedPlanetServiceServer) GetWelcome(context.Context, *emptypb.Empty) (*WelcomeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWelcome not implemented")
}
func (UnimplementedPlanetServiceServer) GetPlanetList(context.Context, *emptypb.Empty) (*PlanetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlanetList not implemented")
}
func (UnimplementedPlanetServiceServer) GetPlanetDetails(context.Context, *wrapperspb.StringValue) (*PlanetDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlanetDetails not implemented")
}
func (UnimplementedPlanetServiceServer) GetSecrets(context.Context, *emptypb.Empty) (*SecretMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecrets not implemented")
}
func (UnimplementedPlanetServiceServer) mustEmbedUnimplementedPlanetServiceServer() {}

// UnsafePlanetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlanetServiceServer will
// result in compilation errors.
type UnsafePlanetServiceServer interface {
	mustEmbedUnimplementedPlanetServiceServer()
}

func RegisterPlanetServiceServer(s grpc.ServiceRegistrar, srv PlanetServiceServer) {
	s.RegisterService(&PlanetService_ServiceDesc, srv)
}

func _PlanetService_GetWelcome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanetServiceServer).GetWelcome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space.PlanetService/getWelcome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanetServiceServer).GetWelcome(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanetService_GetPlanetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanetServiceServer).GetPlanetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space.PlanetService/getPlanetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanetServiceServer).GetPlanetList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanetService_GetPlanetDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanetServiceServer).GetPlanetDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space.PlanetService/getPlanetDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanetServiceServer).GetPlanetDetails(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlanetService_GetSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlanetServiceServer).GetSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space.PlanetService/getSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlanetServiceServer).GetSecrets(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PlanetService_ServiceDesc is the grpc.ServiceDesc for PlanetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlanetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "space.PlanetService",
	HandlerType: (*PlanetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getWelcome",
			Handler:    _PlanetService_GetWelcome_Handler,
		},
		{
			MethodName: "getPlanetList",
			Handler:    _PlanetService_GetPlanetList_Handler,
		},
		{
			MethodName: "getPlanetDetails",
			Handler:    _PlanetService_GetPlanetDetails_Handler,
		},
		{
			MethodName: "getSecrets",
			Handler:    _PlanetService_GetSecrets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/planetservices.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: auth.proto

package gen

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

// AuthRegClient is the client API for AuthReg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthRegClient interface {
	CheckAuthorization(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthResponse, error)
	GetUserCredentials(ctx context.Context, in *UserCredentialsRequest, opts ...grpc.CallOption) (*UserCredentialsResponse, error)
}

type authRegClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthRegClient(cc grpc.ClientConnInterface) AuthRegClient {
	return &authRegClient{cc}
}

func (c *authRegClient) CheckAuthorization(ctx context.Context, in *CheckAuthRequest, opts ...grpc.CallOption) (*CheckAuthResponse, error) {
	out := new(CheckAuthResponse)
	err := c.cc.Invoke(ctx, "/main.AuthReg/CheckAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authRegClient) GetUserCredentials(ctx context.Context, in *UserCredentialsRequest, opts ...grpc.CallOption) (*UserCredentialsResponse, error) {
	out := new(UserCredentialsResponse)
	err := c.cc.Invoke(ctx, "/main.AuthReg/GetUserCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthRegServer is the server API for AuthReg service.
// All implementations must embed UnimplementedAuthRegServer
// for forward compatibility
type AuthRegServer interface {
	CheckAuthorization(context.Context, *CheckAuthRequest) (*CheckAuthResponse, error)
	GetUserCredentials(context.Context, *UserCredentialsRequest) (*UserCredentialsResponse, error)
	mustEmbedUnimplementedAuthRegServer()
}

// UnimplementedAuthRegServer must be embedded to have forward compatible implementations.
type UnimplementedAuthRegServer struct {
}

func (UnimplementedAuthRegServer) CheckAuthorization(context.Context, *CheckAuthRequest) (*CheckAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuthorization not implemented")
}
func (UnimplementedAuthRegServer) GetUserCredentials(context.Context, *UserCredentialsRequest) (*UserCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCredentials not implemented")
}
func (UnimplementedAuthRegServer) mustEmbedUnimplementedAuthRegServer() {}

// UnsafeAuthRegServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthRegServer will
// result in compilation errors.
type UnsafeAuthRegServer interface {
	mustEmbedUnimplementedAuthRegServer()
}

func RegisterAuthRegServer(s grpc.ServiceRegistrar, srv AuthRegServer) {
	s.RegisterService(&AuthReg_ServiceDesc, srv)
}

func _AuthReg_CheckAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthRegServer).CheckAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.AuthReg/CheckAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthRegServer).CheckAuthorization(ctx, req.(*CheckAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthReg_GetUserCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthRegServer).GetUserCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.AuthReg/GetUserCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthRegServer).GetUserCredentials(ctx, req.(*UserCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthReg_ServiceDesc is the grpc.ServiceDesc for AuthReg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthReg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.AuthReg",
	HandlerType: (*AuthRegServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAuthorization",
			Handler:    _AuthReg_CheckAuthorization_Handler,
		},
		{
			MethodName: "GetUserCredentials",
			Handler:    _AuthReg_GetUserCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

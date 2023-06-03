// Protobuf Version

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: modules/users/usersProto/users.proto

package neversuitup_e_commerce_test

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

const (
	UsersServices_FindOneUserByUsername_FullMethodName = "/UsersServices/FindOneUserByUsername"
)

// UsersServicesClient is the client API for UsersServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersServicesClient interface {
	FindOneUserByUsername(ctx context.Context, in *FindOneUserByUsernameReq, opts ...grpc.CallOption) (*UserForAll, error)
}

type usersServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServicesClient(cc grpc.ClientConnInterface) UsersServicesClient {
	return &usersServicesClient{cc}
}

func (c *usersServicesClient) FindOneUserByUsername(ctx context.Context, in *FindOneUserByUsernameReq, opts ...grpc.CallOption) (*UserForAll, error) {
	out := new(UserForAll)
	err := c.cc.Invoke(ctx, UsersServices_FindOneUserByUsername_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServicesServer is the server API for UsersServices service.
// All implementations must embed UnimplementedUsersServicesServer
// for forward compatibility
type UsersServicesServer interface {
	FindOneUserByUsername(context.Context, *FindOneUserByUsernameReq) (*UserForAll, error)
	mustEmbedUnimplementedUsersServicesServer()
}

// UnimplementedUsersServicesServer must be embedded to have forward compatible implementations.
type UnimplementedUsersServicesServer struct {
}

func (UnimplementedUsersServicesServer) FindOneUserByUsername(context.Context, *FindOneUserByUsernameReq) (*UserForAll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneUserByUsername not implemented")
}
func (UnimplementedUsersServicesServer) mustEmbedUnimplementedUsersServicesServer() {}

// UnsafeUsersServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServicesServer will
// result in compilation errors.
type UnsafeUsersServicesServer interface {
	mustEmbedUnimplementedUsersServicesServer()
}

func RegisterUsersServicesServer(s grpc.ServiceRegistrar, srv UsersServicesServer) {
	s.RegisterService(&UsersServices_ServiceDesc, srv)
}

func _UsersServices_FindOneUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneUserByUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServicesServer).FindOneUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersServices_FindOneUserByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServicesServer).FindOneUserByUsername(ctx, req.(*FindOneUserByUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersServices_ServiceDesc is the grpc.ServiceDesc for UsersServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UsersServices",
	HandlerType: (*UsersServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOneUserByUsername",
			Handler:    _UsersServices_FindOneUserByUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/users/usersProto/users.proto",
}

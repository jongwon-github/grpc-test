// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: config/config.proto

package config

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

// ConfigStoreClient is the client API for ConfigStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigStoreClient interface {
	Get(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type configStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigStoreClient(cc grpc.ClientConnInterface) ConfigStoreClient {
	return &configStoreClient{cc}
}

func (c *configStoreClient) Get(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigStore/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigStoreServer is the server API for ConfigStore service.
// All implementations must embed UnimplementedConfigStoreServer
// for forward compatibility
type ConfigStoreServer interface {
	Get(context.Context, *ConfigRequest) (*ConfigResponse, error)
	mustEmbedUnimplementedConfigStoreServer()
}

// UnimplementedConfigStoreServer must be embedded to have forward compatible implementations.
type UnimplementedConfigStoreServer struct {
}

func (UnimplementedConfigStoreServer) Get(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedConfigStoreServer) mustEmbedUnimplementedConfigStoreServer() {}

// UnsafeConfigStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigStoreServer will
// result in compilation errors.
type UnsafeConfigStoreServer interface {
	mustEmbedUnimplementedConfigStoreServer()
}

func RegisterConfigStoreServer(s grpc.ServiceRegistrar, srv ConfigStoreServer) {
	s.RegisterService(&ConfigStore_ServiceDesc, srv)
}

func _ConfigStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigStore/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigStoreServer).Get(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigStore_ServiceDesc is the grpc.ServiceDesc for ConfigStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.ConfigStore",
	HandlerType: (*ConfigStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ConfigStore_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/config.proto",
}

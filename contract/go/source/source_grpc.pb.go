// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package source

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

// IsWriteKeyValidClient is the client API for IsWriteKeyValid service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IsWriteKeyValidClient interface {
	IsWriteKeyValid(ctx context.Context, in *ValidateWriteKeyReq, opts ...grpc.CallOption) (*ValidateWriteKeyResp, error)
}

type isWriteKeyValidClient struct {
	cc grpc.ClientConnInterface
}

func NewIsWriteKeyValidClient(cc grpc.ClientConnInterface) IsWriteKeyValidClient {
	return &isWriteKeyValidClient{cc}
}

func (c *isWriteKeyValidClient) IsWriteKeyValid(ctx context.Context, in *ValidateWriteKeyReq, opts ...grpc.CallOption) (*ValidateWriteKeyResp, error) {
	out := new(ValidateWriteKeyResp)
	err := c.cc.Invoke(ctx, "/source.IsWriteKeyValid/IsWriteKeyValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IsWriteKeyValidServer is the server API for IsWriteKeyValid service.
// All implementations must embed UnimplementedIsWriteKeyValidServer
// for forward compatibility
type IsWriteKeyValidServer interface {
	IsWriteKeyValid(context.Context, *ValidateWriteKeyReq) (*ValidateWriteKeyResp, error)
	mustEmbedUnimplementedIsWriteKeyValidServer()
}

// UnimplementedIsWriteKeyValidServer must be embedded to have forward compatible implementations.
type UnimplementedIsWriteKeyValidServer struct {
}

func (UnimplementedIsWriteKeyValidServer) IsWriteKeyValid(context.Context, *ValidateWriteKeyReq) (*ValidateWriteKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsWriteKeyValid not implemented")
}
func (UnimplementedIsWriteKeyValidServer) mustEmbedUnimplementedIsWriteKeyValidServer() {}

// UnsafeIsWriteKeyValidServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IsWriteKeyValidServer will
// result in compilation errors.
type UnsafeIsWriteKeyValidServer interface {
	mustEmbedUnimplementedIsWriteKeyValidServer()
}

func RegisterIsWriteKeyValidServer(s grpc.ServiceRegistrar, srv IsWriteKeyValidServer) {
	s.RegisterService(&IsWriteKeyValid_ServiceDesc, srv)
}

func _IsWriteKeyValid_IsWriteKeyValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateWriteKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IsWriteKeyValidServer).IsWriteKeyValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/source.IsWriteKeyValid/IsWriteKeyValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IsWriteKeyValidServer).IsWriteKeyValid(ctx, req.(*ValidateWriteKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// IsWriteKeyValid_ServiceDesc is the grpc.ServiceDesc for IsWriteKeyValid service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IsWriteKeyValid_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "source.IsWriteKeyValid",
	HandlerType: (*IsWriteKeyValidServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsWriteKeyValid",
			Handler:    _IsWriteKeyValid_IsWriteKeyValid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "source/source.proto",
}

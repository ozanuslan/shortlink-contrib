// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// LinkQueryServiceClient is the client API for LinkQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkQueryServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type linkQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkQueryServiceClient(cc grpc.ClientConnInterface) LinkQueryServiceClient {
	return &linkQueryServiceClient{cc}
}

func (c *linkQueryServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/infrastructure.rpc.cqrs.link.v1.LinkQueryService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkQueryServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/infrastructure.rpc.cqrs.link.v1.LinkQueryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkQueryServiceServer is the server API for LinkQueryService service.
// All implementations must embed UnimplementedLinkQueryServiceServer
// for forward compatibility
type LinkQueryServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedLinkQueryServiceServer()
}

// UnimplementedLinkQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLinkQueryServiceServer struct {
}

func (UnimplementedLinkQueryServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLinkQueryServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLinkQueryServiceServer) mustEmbedUnimplementedLinkQueryServiceServer() {}

// UnsafeLinkQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkQueryServiceServer will
// result in compilation errors.
type UnsafeLinkQueryServiceServer interface {
	mustEmbedUnimplementedLinkQueryServiceServer()
}

func RegisterLinkQueryServiceServer(s grpc.ServiceRegistrar, srv LinkQueryServiceServer) {
	s.RegisterService(&LinkQueryService_ServiceDesc, srv)
}

func _LinkQueryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkQueryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infrastructure.rpc.cqrs.link.v1.LinkQueryService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkQueryServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkQueryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkQueryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infrastructure.rpc.cqrs.link.v1.LinkQueryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkQueryServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkQueryService_ServiceDesc is the grpc.ServiceDesc for LinkQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infrastructure.rpc.cqrs.link.v1.LinkQueryService",
	HandlerType: (*LinkQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _LinkQueryService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LinkQueryService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infrastructure/rpc/cqrs/link/v1/link_query.proto",
}

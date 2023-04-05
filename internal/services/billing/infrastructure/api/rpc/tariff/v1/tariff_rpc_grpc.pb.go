// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: infrastructure/api/rpc/tariff/v1/tariff_rpc.proto

package tariff_rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TariffService_Tariff_FullMethodName       = "/infrastructure.api.rpc.tariff.v1.TariffService/Tariff"
	TariffService_Tariffs_FullMethodName      = "/infrastructure.api.rpc.tariff.v1.TariffService/Tariffs"
	TariffService_TariffCreate_FullMethodName = "/infrastructure.api.rpc.tariff.v1.TariffService/TariffCreate"
	TariffService_TariffUpdate_FullMethodName = "/infrastructure.api.rpc.tariff.v1.TariffService/TariffUpdate"
	TariffService_TariffClose_FullMethodName  = "/infrastructure.api.rpc.tariff.v1.TariffService/TariffClose"
)

// TariffServiceClient is the client API for TariffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TariffServiceClient interface {
	Tariff(ctx context.Context, in *TariffRequest, opts ...grpc.CallOption) (*TariffResponse, error)
	Tariffs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TariffsResponse, error)
	TariffCreate(ctx context.Context, in *TariffCreateRequest, opts ...grpc.CallOption) (*TariffCreateResponse, error)
	TariffUpdate(ctx context.Context, in *TariffUpdateRequest, opts ...grpc.CallOption) (*TariffUpdateResponse, error)
	TariffClose(ctx context.Context, in *TariffCloseRequest, opts ...grpc.CallOption) (*TariffCloseResponse, error)
}

type tariffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTariffServiceClient(cc grpc.ClientConnInterface) TariffServiceClient {
	return &tariffServiceClient{cc}
}

func (c *tariffServiceClient) Tariff(ctx context.Context, in *TariffRequest, opts ...grpc.CallOption) (*TariffResponse, error) {
	out := new(TariffResponse)
	err := c.cc.Invoke(ctx, TariffService_Tariff_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffServiceClient) Tariffs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*TariffsResponse, error) {
	out := new(TariffsResponse)
	err := c.cc.Invoke(ctx, TariffService_Tariffs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffServiceClient) TariffCreate(ctx context.Context, in *TariffCreateRequest, opts ...grpc.CallOption) (*TariffCreateResponse, error) {
	out := new(TariffCreateResponse)
	err := c.cc.Invoke(ctx, TariffService_TariffCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffServiceClient) TariffUpdate(ctx context.Context, in *TariffUpdateRequest, opts ...grpc.CallOption) (*TariffUpdateResponse, error) {
	out := new(TariffUpdateResponse)
	err := c.cc.Invoke(ctx, TariffService_TariffUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tariffServiceClient) TariffClose(ctx context.Context, in *TariffCloseRequest, opts ...grpc.CallOption) (*TariffCloseResponse, error) {
	out := new(TariffCloseResponse)
	err := c.cc.Invoke(ctx, TariffService_TariffClose_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TariffServiceServer is the server API for TariffService service.
// All implementations must embed UnimplementedTariffServiceServer
// for forward compatibility
type TariffServiceServer interface {
	Tariff(context.Context, *TariffRequest) (*TariffResponse, error)
	Tariffs(context.Context, *emptypb.Empty) (*TariffsResponse, error)
	TariffCreate(context.Context, *TariffCreateRequest) (*TariffCreateResponse, error)
	TariffUpdate(context.Context, *TariffUpdateRequest) (*TariffUpdateResponse, error)
	TariffClose(context.Context, *TariffCloseRequest) (*TariffCloseResponse, error)
	mustEmbedUnimplementedTariffServiceServer()
}

// UnimplementedTariffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTariffServiceServer struct {
}

func (UnimplementedTariffServiceServer) Tariff(context.Context, *TariffRequest) (*TariffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tariff not implemented")
}
func (UnimplementedTariffServiceServer) Tariffs(context.Context, *emptypb.Empty) (*TariffsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tariffs not implemented")
}
func (UnimplementedTariffServiceServer) TariffCreate(context.Context, *TariffCreateRequest) (*TariffCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TariffCreate not implemented")
}
func (UnimplementedTariffServiceServer) TariffUpdate(context.Context, *TariffUpdateRequest) (*TariffUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TariffUpdate not implemented")
}
func (UnimplementedTariffServiceServer) TariffClose(context.Context, *TariffCloseRequest) (*TariffCloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TariffClose not implemented")
}
func (UnimplementedTariffServiceServer) mustEmbedUnimplementedTariffServiceServer() {}

// UnsafeTariffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TariffServiceServer will
// result in compilation errors.
type UnsafeTariffServiceServer interface {
	mustEmbedUnimplementedTariffServiceServer()
}

func RegisterTariffServiceServer(s grpc.ServiceRegistrar, srv TariffServiceServer) {
	s.RegisterService(&TariffService_ServiceDesc, srv)
}

func _TariffService_Tariff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TariffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServiceServer).Tariff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TariffService_Tariff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServiceServer).Tariff(ctx, req.(*TariffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TariffService_Tariffs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServiceServer).Tariffs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TariffService_Tariffs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServiceServer).Tariffs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TariffService_TariffCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TariffCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServiceServer).TariffCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TariffService_TariffCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServiceServer).TariffCreate(ctx, req.(*TariffCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TariffService_TariffUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TariffUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServiceServer).TariffUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TariffService_TariffUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServiceServer).TariffUpdate(ctx, req.(*TariffUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TariffService_TariffClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TariffCloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TariffServiceServer).TariffClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TariffService_TariffClose_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TariffServiceServer).TariffClose(ctx, req.(*TariffCloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TariffService_ServiceDesc is the grpc.ServiceDesc for TariffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TariffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infrastructure.api.rpc.tariff.v1.TariffService",
	HandlerType: (*TariffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tariff",
			Handler:    _TariffService_Tariff_Handler,
		},
		{
			MethodName: "Tariffs",
			Handler:    _TariffService_Tariffs_Handler,
		},
		{
			MethodName: "TariffCreate",
			Handler:    _TariffService_TariffCreate_Handler,
		},
		{
			MethodName: "TariffUpdate",
			Handler:    _TariffService_TariffUpdate_Handler,
		},
		{
			MethodName: "TariffClose",
			Handler:    _TariffService_TariffClose_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infrastructure/api/rpc/tariff/v1/tariff_rpc.proto",
}

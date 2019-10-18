// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/api/grpc-web/api.proto

package grpc_web

import (
	context "context"
	fmt "fmt"
	link "github.com/batazor/shortlink/pkg/link"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetLinkRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLinkRequest) Reset()         { *m = GetLinkRequest{} }
func (m *GetLinkRequest) String() string { return proto.CompactTextString(m) }
func (*GetLinkRequest) ProtoMessage()    {}
func (*GetLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d0e997a1f3c768, []int{0}
}

func (m *GetLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLinkRequest.Unmarshal(m, b)
}
func (m *GetLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLinkRequest.Marshal(b, m, deterministic)
}
func (m *GetLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLinkRequest.Merge(m, src)
}
func (m *GetLinkRequest) XXX_Size() int {
	return xxx_messageInfo_GetLinkRequest.Size(m)
}
func (m *GetLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLinkRequest proto.InternalMessageInfo

func (m *GetLinkRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type DeleteLinkRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteLinkRequest) Reset()         { *m = DeleteLinkRequest{} }
func (m *DeleteLinkRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteLinkRequest) ProtoMessage()    {}
func (*DeleteLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d0e997a1f3c768, []int{1}
}

func (m *DeleteLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLinkRequest.Unmarshal(m, b)
}
func (m *DeleteLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLinkRequest.Marshal(b, m, deterministic)
}
func (m *DeleteLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLinkRequest.Merge(m, src)
}
func (m *DeleteLinkRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteLinkRequest.Size(m)
}
func (m *DeleteLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLinkRequest proto.InternalMessageInfo

func (m *DeleteLinkRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type DeleteLinkResponse struct {
	Link                 *empty.Empty `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DeleteLinkResponse) Reset()         { *m = DeleteLinkResponse{} }
func (m *DeleteLinkResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteLinkResponse) ProtoMessage()    {}
func (*DeleteLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d0e997a1f3c768, []int{2}
}

func (m *DeleteLinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteLinkResponse.Unmarshal(m, b)
}
func (m *DeleteLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteLinkResponse.Marshal(b, m, deterministic)
}
func (m *DeleteLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteLinkResponse.Merge(m, src)
}
func (m *DeleteLinkResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteLinkResponse.Size(m)
}
func (m *DeleteLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteLinkResponse proto.InternalMessageInfo

func (m *DeleteLinkResponse) GetLink() *empty.Empty {
	if m != nil {
		return m.Link
	}
	return nil
}

type RedirectLinkRequest struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedirectLinkRequest) Reset()         { *m = RedirectLinkRequest{} }
func (m *RedirectLinkRequest) String() string { return proto.CompactTextString(m) }
func (*RedirectLinkRequest) ProtoMessage()    {}
func (*RedirectLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d0e997a1f3c768, []int{3}
}

func (m *RedirectLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedirectLinkRequest.Unmarshal(m, b)
}
func (m *RedirectLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedirectLinkRequest.Marshal(b, m, deterministic)
}
func (m *RedirectLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedirectLinkRequest.Merge(m, src)
}
func (m *RedirectLinkRequest) XXX_Size() int {
	return xxx_messageInfo_RedirectLinkRequest.Size(m)
}
func (m *RedirectLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RedirectLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RedirectLinkRequest proto.InternalMessageInfo

func (m *RedirectLinkRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type RedirectLinkResponse struct {
	Link                 *empty.Empty `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RedirectLinkResponse) Reset()         { *m = RedirectLinkResponse{} }
func (m *RedirectLinkResponse) String() string { return proto.CompactTextString(m) }
func (*RedirectLinkResponse) ProtoMessage()    {}
func (*RedirectLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_54d0e997a1f3c768, []int{4}
}

func (m *RedirectLinkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedirectLinkResponse.Unmarshal(m, b)
}
func (m *RedirectLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedirectLinkResponse.Marshal(b, m, deterministic)
}
func (m *RedirectLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedirectLinkResponse.Merge(m, src)
}
func (m *RedirectLinkResponse) XXX_Size() int {
	return xxx_messageInfo_RedirectLinkResponse.Size(m)
}
func (m *RedirectLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RedirectLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RedirectLinkResponse proto.InternalMessageInfo

func (m *RedirectLinkResponse) GetLink() *empty.Empty {
	if m != nil {
		return m.Link
	}
	return nil
}

func init() {
	proto.RegisterType((*GetLinkRequest)(nil), "grpc_web.GetLinkRequest")
	proto.RegisterType((*DeleteLinkRequest)(nil), "grpc_web.DeleteLinkRequest")
	proto.RegisterType((*DeleteLinkResponse)(nil), "grpc_web.DeleteLinkResponse")
	proto.RegisterType((*RedirectLinkRequest)(nil), "grpc_web.RedirectLinkRequest")
	proto.RegisterType((*RedirectLinkResponse)(nil), "grpc_web.RedirectLinkResponse")
}

func init() { proto.RegisterFile("pkg/api/grpc-web/api.proto", fileDescriptor_54d0e997a1f3c768) }

var fileDescriptor_54d0e997a1f3c768 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xf2, 0x40,
	0x10, 0xc7, 0x03, 0x1f, 0x9f, 0xc0, 0x60, 0x30, 0x0c, 0xc6, 0x90, 0x82, 0xc6, 0x34, 0x26, 0x2a,
	0x09, 0xbb, 0x11, 0x3d, 0x79, 0x22, 0xa8, 0xf1, 0xe2, 0xa9, 0xf1, 0xe4, 0x45, 0x5b, 0x1c, 0xdb,
	0x06, 0xe8, 0xd6, 0x76, 0x91, 0xa8, 0xf1, 0xe2, 0x2b, 0xf8, 0x2e, 0xbe, 0x88, 0xaf, 0xe0, 0x83,
	0x98, 0xdd, 0x96, 0x00, 0x8a, 0x21, 0x5e, 0x9a, 0xee, 0xcc, 0x7f, 0x66, 0xfe, 0xf3, 0xcb, 0x80,
	0x11, 0xf6, 0x5d, 0x6e, 0x87, 0x3e, 0x77, 0xa3, 0xb0, 0xd7, 0x1a, 0x93, 0xa3, 0x1e, 0x2c, 0x8c,
	0x84, 0x14, 0x58, 0x50, 0xb1, 0xeb, 0x31, 0x39, 0x46, 0xc3, 0x15, 0xc2, 0x1d, 0x90, 0x16, 0xda,
	0x41, 0x20, 0xa4, 0x2d, 0x7d, 0x11, 0xc4, 0x89, 0xce, 0xa8, 0xa7, 0x59, 0xfd, 0x72, 0x46, 0x77,
	0x9c, 0x86, 0xa1, 0x7c, 0x4c, 0x93, 0x55, 0x35, 0x60, 0xe0, 0x07, 0x7d, 0xfd, 0x49, 0x82, 0xe6,
	0x0e, 0x94, 0xcf, 0x49, 0x5e, 0xf8, 0x41, 0xdf, 0xa2, 0xfb, 0x11, 0xc5, 0x12, 0x11, 0x72, 0x9e,
	0x1d, 0x7b, 0xb5, 0xcc, 0x76, 0x66, 0xaf, 0x68, 0xe9, 0x7f, 0x73, 0x17, 0x2a, 0xa7, 0x34, 0x20,
	0x49, 0xcb, 0x84, 0x1d, 0xc0, 0x59, 0x61, 0x1c, 0x8a, 0x20, 0x26, 0x6c, 0x42, 0x4e, 0x8d, 0xd4,
	0xca, 0x52, 0x7b, 0x83, 0x25, 0x2e, 0xd9, 0xc4, 0x25, 0x3b, 0x53, 0x2e, 0x2d, 0xad, 0x31, 0xf7,
	0xa1, 0x6a, 0xd1, 0xad, 0x1f, 0x51, 0x6f, 0xa9, 0xab, 0x2e, 0xac, 0xcf, 0x4b, 0xff, 0x3e, 0xae,
	0xfd, 0x9e, 0x85, 0x9c, 0x2a, 0xc6, 0x0e, 0xe4, 0x53, 0x10, 0x58, 0x63, 0x13, 0xdc, 0x6c, 0x9e,
	0x8d, 0x01, 0x4c, 0xa3, 0x53, 0x21, 0x73, 0xed, 0xf5, 0xe3, 0xf3, 0x2d, 0x5b, 0xc4, 0x3c, 0x7f,
	0x56, 0x6e, 0x5e, 0xf0, 0x08, 0xe0, 0x24, 0x22, 0x3b, 0xd9, 0x1d, 0x67, 0xa4, 0x73, 0x65, 0x65,
	0x5d, 0x56, 0x30, 0xff, 0xf1, 0x87, 0x83, 0xe3, 0x4c, 0x13, 0x2f, 0x01, 0xa6, 0xc4, 0xb0, 0x3e,
	0x1d, 0xfd, 0x03, 0xb8, 0xd1, 0x58, 0x9c, 0x4c, 0xb6, 0x36, 0x4b, 0xba, 0xf1, 0xff, 0xa6, 0x6a,
	0x8c, 0x37, 0xb0, 0x3a, 0x8b, 0x06, 0x37, 0xa7, 0xa5, 0x0b, 0xe8, 0x1a, 0x5b, 0xbf, 0xa5, 0xd3,
	0xde, 0x15, 0xdd, 0xbb, 0x84, 0x45, 0x1e, 0xa7, 0xdb, 0x76, 0xf9, 0x55, 0xcb, 0xf5, 0xa5, 0x37,
	0x72, 0x58, 0x4f, 0x0c, 0xb9, 0x63, 0x4b, 0xfb, 0x49, 0x44, 0x3c, 0xf6, 0x44, 0x24, 0xf5, 0x8d,
	0x7d, 0xbf, 0x66, 0x67, 0x45, 0xf3, 0x3f, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x23, 0x22, 0xb6,
	0x9f, 0xe8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LinkClient is the client API for Link service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LinkClient interface {
	GetLink(ctx context.Context, in *GetLinkRequest, opts ...grpc.CallOption) (*link.Link, error)
	CreateLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*link.Link, error)
	DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...grpc.CallOption) (*DeleteLinkResponse, error)
	RedirectLink(ctx context.Context, in *RedirectLinkRequest, opts ...grpc.CallOption) (*RedirectLinkResponse, error)
}

type linkClient struct {
	cc *grpc.ClientConn
}

func NewLinkClient(cc *grpc.ClientConn) LinkClient {
	return &linkClient{cc}
}

func (c *linkClient) GetLink(ctx context.Context, in *GetLinkRequest, opts ...grpc.CallOption) (*link.Link, error) {
	out := new(link.Link)
	err := c.cc.Invoke(ctx, "/grpc_web.Link/GetLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkClient) CreateLink(ctx context.Context, in *link.Link, opts ...grpc.CallOption) (*link.Link, error) {
	out := new(link.Link)
	err := c.cc.Invoke(ctx, "/grpc_web.Link/CreateLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkClient) DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...grpc.CallOption) (*DeleteLinkResponse, error) {
	out := new(DeleteLinkResponse)
	err := c.cc.Invoke(ctx, "/grpc_web.Link/DeleteLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkClient) RedirectLink(ctx context.Context, in *RedirectLinkRequest, opts ...grpc.CallOption) (*RedirectLinkResponse, error) {
	out := new(RedirectLinkResponse)
	err := c.cc.Invoke(ctx, "/grpc_web.Link/RedirectLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkServer is the server API for Link service.
type LinkServer interface {
	GetLink(context.Context, *GetLinkRequest) (*link.Link, error)
	CreateLink(context.Context, *link.Link) (*link.Link, error)
	DeleteLink(context.Context, *DeleteLinkRequest) (*DeleteLinkResponse, error)
	RedirectLink(context.Context, *RedirectLinkRequest) (*RedirectLinkResponse, error)
}

// UnimplementedLinkServer can be embedded to have forward compatible implementations.
type UnimplementedLinkServer struct {
}

func (*UnimplementedLinkServer) GetLink(ctx context.Context, req *GetLinkRequest) (*link.Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLink not implemented")
}
func (*UnimplementedLinkServer) CreateLink(ctx context.Context, req *link.Link) (*link.Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLink not implemented")
}
func (*UnimplementedLinkServer) DeleteLink(ctx context.Context, req *DeleteLinkRequest) (*DeleteLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLink not implemented")
}
func (*UnimplementedLinkServer) RedirectLink(ctx context.Context, req *RedirectLinkRequest) (*RedirectLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedirectLink not implemented")
}

func RegisterLinkServer(s *grpc.Server, srv LinkServer) {
	s.RegisterService(&_Link_serviceDesc, srv)
}

func _Link_GetLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).GetLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_web.Link/GetLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).GetLink(ctx, req.(*GetLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Link_CreateLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(link.Link)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).CreateLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_web.Link/CreateLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).CreateLink(ctx, req.(*link.Link))
	}
	return interceptor(ctx, in, info, handler)
}

func _Link_DeleteLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).DeleteLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_web.Link/DeleteLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).DeleteLink(ctx, req.(*DeleteLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Link_RedirectLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedirectLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServer).RedirectLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_web.Link/RedirectLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServer).RedirectLink(ctx, req.(*RedirectLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Link_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_web.Link",
	HandlerType: (*LinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLink",
			Handler:    _Link_GetLink_Handler,
		},
		{
			MethodName: "CreateLink",
			Handler:    _Link_CreateLink_Handler,
		},
		{
			MethodName: "DeleteLink",
			Handler:    _Link_DeleteLink_Handler,
		},
		{
			MethodName: "RedirectLink",
			Handler:    _Link_RedirectLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/grpc-web/api.proto",
}

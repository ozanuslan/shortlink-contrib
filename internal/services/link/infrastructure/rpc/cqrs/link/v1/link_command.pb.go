// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: infrastructure/rpc/cqrs/link/v1/link_command.proto

package v1

import (
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link *v1.Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetLink() *v1.Link {
	if x != nil {
		return x.Link
	}
	return nil
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link *v1.Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetLink() *v1.Link {
	if x != nil {
		return x.Link
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link *v1.Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetLink() *v1.Link {
	if x != nil {
		return x.Link
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link *v1.Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResponse) GetLink() *v1.Link {
	if x != nil {
		return x.Link
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_infrastructure_rpc_cqrs_link_v1_link_command_proto protoreflect.FileDescriptor

var file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDesc = []byte{
	0x0a, 0x32, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x37, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x39,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x3a, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6c,
	0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x23, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x32, 0xb9, 0x02, 0x0a, 0x12, 0x4c,
	0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x62, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71,
	0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x2e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x52, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0xae, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x10,
	0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x04, 0x49, 0x52, 0x43, 0x4c, 0xaa,
	0x02, 0x1f, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x52, 0x70, 0x63, 0x2e, 0x43, 0x71, 0x72, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x1f, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x43, 0x71, 0x72, 0x73, 0x5c, 0x4c, 0x69, 0x6e, 0x6b,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2b, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x43, 0x71, 0x72, 0x73, 0x5c, 0x4c, 0x69,
	0x6e, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x23, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x43, 0x71, 0x72, 0x73, 0x3a, 0x3a, 0x4c,
	0x69, 0x6e, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescOnce sync.Once
	file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescData = file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDesc
)

func file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescGZIP() []byte {
	file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescOnce.Do(func() {
		file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescData)
	})
	return file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDescData
}

var file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_infrastructure_rpc_cqrs_link_v1_link_command_proto_goTypes = []any{
	(*AddRequest)(nil),     // 0: infrastructure.rpc.cqrs.link.v1.AddRequest
	(*AddResponse)(nil),    // 1: infrastructure.rpc.cqrs.link.v1.AddResponse
	(*UpdateRequest)(nil),  // 2: infrastructure.rpc.cqrs.link.v1.UpdateRequest
	(*UpdateResponse)(nil), // 3: infrastructure.rpc.cqrs.link.v1.UpdateResponse
	(*DeleteRequest)(nil),  // 4: infrastructure.rpc.cqrs.link.v1.DeleteRequest
	(*v1.Link)(nil),        // 5: domain.link.v1.Link
	(*emptypb.Empty)(nil),  // 6: google.protobuf.Empty
}
var file_infrastructure_rpc_cqrs_link_v1_link_command_proto_depIdxs = []int32{
	5, // 0: infrastructure.rpc.cqrs.link.v1.AddRequest.link:type_name -> domain.link.v1.Link
	5, // 1: infrastructure.rpc.cqrs.link.v1.AddResponse.link:type_name -> domain.link.v1.Link
	5, // 2: infrastructure.rpc.cqrs.link.v1.UpdateRequest.link:type_name -> domain.link.v1.Link
	5, // 3: infrastructure.rpc.cqrs.link.v1.UpdateResponse.link:type_name -> domain.link.v1.Link
	0, // 4: infrastructure.rpc.cqrs.link.v1.LinkCommandService.Add:input_type -> infrastructure.rpc.cqrs.link.v1.AddRequest
	2, // 5: infrastructure.rpc.cqrs.link.v1.LinkCommandService.Update:input_type -> infrastructure.rpc.cqrs.link.v1.UpdateRequest
	4, // 6: infrastructure.rpc.cqrs.link.v1.LinkCommandService.Delete:input_type -> infrastructure.rpc.cqrs.link.v1.DeleteRequest
	1, // 7: infrastructure.rpc.cqrs.link.v1.LinkCommandService.Add:output_type -> infrastructure.rpc.cqrs.link.v1.AddResponse
	3, // 8: infrastructure.rpc.cqrs.link.v1.LinkCommandService.Update:output_type -> infrastructure.rpc.cqrs.link.v1.UpdateResponse
	6, // 9: infrastructure.rpc.cqrs.link.v1.LinkCommandService.Delete:output_type -> google.protobuf.Empty
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_infrastructure_rpc_cqrs_link_v1_link_command_proto_init() }
func file_infrastructure_rpc_cqrs_link_v1_link_command_proto_init() {
	if File_infrastructure_rpc_cqrs_link_v1_link_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_rpc_cqrs_link_v1_link_command_proto_goTypes,
		DependencyIndexes: file_infrastructure_rpc_cqrs_link_v1_link_command_proto_depIdxs,
		MessageInfos:      file_infrastructure_rpc_cqrs_link_v1_link_command_proto_msgTypes,
	}.Build()
	File_infrastructure_rpc_cqrs_link_v1_link_command_proto = out.File
	file_infrastructure_rpc_cqrs_link_v1_link_command_proto_rawDesc = nil
	file_infrastructure_rpc_cqrs_link_v1_link_command_proto_goTypes = nil
	file_infrastructure_rpc_cqrs_link_v1_link_command_proto_depIdxs = nil
}

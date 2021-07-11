// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: infrastructure/api/rpc/payment/v1/billing_rpc.proto

package payment_rpc

import (
	v1 "github.com/batazor/shortlink/internal/services/billing/domain/billing/v1"
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

type PaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *v1.Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *PaymentRequest) GetPayment() *v1.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type PaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *v1.Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentResponse) GetPayment() *v1.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type PaymentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*v1.Payment `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PaymentsResponse) Reset() {
	*x = PaymentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentsResponse) ProtoMessage() {}

func (x *PaymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentsResponse.ProtoReflect.Descriptor instead.
func (*PaymentsResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentsResponse) GetList() []*v1.Payment {
	if x != nil {
		return x.List
	}
	return nil
}

type PaymentCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *v1.Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentCreateRequest) Reset() {
	*x = PaymentCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCreateRequest) ProtoMessage() {}

func (x *PaymentCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCreateRequest.ProtoReflect.Descriptor instead.
func (*PaymentCreateRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentCreateRequest) GetPayment() *v1.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type PaymentCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *v1.Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentCreateResponse) Reset() {
	*x = PaymentCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCreateResponse) ProtoMessage() {}

func (x *PaymentCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCreateResponse.ProtoReflect.Descriptor instead.
func (*PaymentCreateResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *PaymentCreateResponse) GetPayment() *v1.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type PaymentCloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *v1.Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentCloseRequest) Reset() {
	*x = PaymentCloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCloseRequest) ProtoMessage() {}

func (x *PaymentCloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCloseRequest.ProtoReflect.Descriptor instead.
func (*PaymentCloseRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *PaymentCloseRequest) GetPayment() *v1.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type PaymentCloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *v1.Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentCloseResponse) Reset() {
	*x = PaymentCloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCloseResponse) ProtoMessage() {}

func (x *PaymentCloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCloseResponse.ProtoReflect.Descriptor instead.
func (*PaymentCloseResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *PaymentCloseResponse) GetPayment() *v1.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

var File_infrastructure_api_rpc_payment_v1_billing_rpc_proto protoreflect.FileDescriptor

var file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x47,
	0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x14, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x15, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4b, 0x0a, 0x13, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x32, 0xea, 0x03, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x08, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x33,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x37, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x81, 0x01, 0x0a,
	0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x36, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescOnce sync.Once
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescData = file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDesc
)

func file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP() []byte {
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescOnce.Do(func() {
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescData)
	})
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescData
}

var file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_goTypes = []interface{}{
	(*PaymentRequest)(nil),        // 0: infrastructure.api.rpc.payment.v1.PaymentRequest
	(*PaymentResponse)(nil),       // 1: infrastructure.api.rpc.payment.v1.PaymentResponse
	(*PaymentsResponse)(nil),      // 2: infrastructure.api.rpc.payment.v1.PaymentsResponse
	(*PaymentCreateRequest)(nil),  // 3: infrastructure.api.rpc.payment.v1.PaymentCreateRequest
	(*PaymentCreateResponse)(nil), // 4: infrastructure.api.rpc.payment.v1.PaymentCreateResponse
	(*PaymentCloseRequest)(nil),   // 5: infrastructure.api.rpc.payment.v1.PaymentCloseRequest
	(*PaymentCloseResponse)(nil),  // 6: infrastructure.api.rpc.payment.v1.PaymentCloseResponse
	(*v1.Payment)(nil),            // 7: domain.billing.v1.Payment
	(*emptypb.Empty)(nil),         // 8: google.protobuf.Empty
}
var file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_depIdxs = []int32{
	7,  // 0: infrastructure.api.rpc.payment.v1.PaymentRequest.payment:type_name -> domain.billing.v1.Payment
	7,  // 1: infrastructure.api.rpc.payment.v1.PaymentResponse.payment:type_name -> domain.billing.v1.Payment
	7,  // 2: infrastructure.api.rpc.payment.v1.PaymentsResponse.list:type_name -> domain.billing.v1.Payment
	7,  // 3: infrastructure.api.rpc.payment.v1.PaymentCreateRequest.payment:type_name -> domain.billing.v1.Payment
	7,  // 4: infrastructure.api.rpc.payment.v1.PaymentCreateResponse.payment:type_name -> domain.billing.v1.Payment
	7,  // 5: infrastructure.api.rpc.payment.v1.PaymentCloseRequest.payment:type_name -> domain.billing.v1.Payment
	7,  // 6: infrastructure.api.rpc.payment.v1.PaymentCloseResponse.payment:type_name -> domain.billing.v1.Payment
	0,  // 7: infrastructure.api.rpc.payment.v1.PaymentService.Payment:input_type -> infrastructure.api.rpc.payment.v1.PaymentRequest
	8,  // 8: infrastructure.api.rpc.payment.v1.PaymentService.Payments:input_type -> google.protobuf.Empty
	3,  // 9: infrastructure.api.rpc.payment.v1.PaymentService.PaymentCreate:input_type -> infrastructure.api.rpc.payment.v1.PaymentCreateRequest
	5,  // 10: infrastructure.api.rpc.payment.v1.PaymentService.PaymentClose:input_type -> infrastructure.api.rpc.payment.v1.PaymentCloseRequest
	1,  // 11: infrastructure.api.rpc.payment.v1.PaymentService.Payment:output_type -> infrastructure.api.rpc.payment.v1.PaymentResponse
	2,  // 12: infrastructure.api.rpc.payment.v1.PaymentService.Payments:output_type -> infrastructure.api.rpc.payment.v1.PaymentsResponse
	4,  // 13: infrastructure.api.rpc.payment.v1.PaymentService.PaymentCreate:output_type -> infrastructure.api.rpc.payment.v1.PaymentCreateResponse
	6,  // 14: infrastructure.api.rpc.payment.v1.PaymentService.PaymentClose:output_type -> infrastructure.api.rpc.payment.v1.PaymentCloseResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_init() }
func file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_init() {
	if File_infrastructure_api_rpc_payment_v1_billing_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentRequest); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentResponse); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentsResponse); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCreateRequest); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCreateResponse); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCloseRequest); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCloseResponse); i {
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
			RawDescriptor: file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_goTypes,
		DependencyIndexes: file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_depIdxs,
		MessageInfos:      file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes,
	}.Build()
	File_infrastructure_api_rpc_payment_v1_billing_rpc_proto = out.File
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDesc = nil
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_goTypes = nil
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_depIdxs = nil
}

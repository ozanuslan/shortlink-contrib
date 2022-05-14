// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: shortdb/domain/table/v1/table.proto

package v1

import (
	v11 "github.com/batazor/shortlink/pkg/shortdb/domain/index/v1"
	v1 "github.com/batazor/shortlink/pkg/shortdb/domain/page/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Type int32

const (
	Type_TYPE_UNSPECIFIED Type = 0
	Type_TYPE_INTEGER     Type = 1
	Type_TYPE_STRING      Type = 2
	Type_TYPE_BOOLEAN     Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_INTEGER",
		2: "TYPE_STRING",
		3: "TYPE_BOOLEAN",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_INTEGER":     1,
		"TYPE_STRING":      2,
		"TYPE_BOOLEAN":     3,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_shortdb_domain_table_v1_table_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_shortdb_domain_table_v1_table_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_shortdb_domain_table_v1_table_proto_rawDescGZIP(), []int{0}
}

type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_domain_table_v1_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_domain_table_v1_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_shortdb_domain_table_v1_table_proto_rawDescGZIP(), []int{0}
}

func (x *Option) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fields map[string]Type       `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=shortdb.table.v1.Type"`
	Pages  map[int32]*v1.Page    `protobuf:"bytes,3,rep,name=pages,proto3" json:"pages,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Index  map[string]*v11.Index `protobuf:"bytes,4,rep,name=index,proto3" json:"index,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Stats  *TableStats           `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
	Option *Option               `protobuf:"bytes,6,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_domain_table_v1_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_domain_table_v1_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_shortdb_domain_table_v1_table_proto_rawDescGZIP(), []int{1}
}

func (x *Table) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Table) GetFields() map[string]Type {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Table) GetPages() map[int32]*v1.Page {
	if x != nil {
		return x.Pages
	}
	return nil
}

func (x *Table) GetIndex() map[string]*v11.Index {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *Table) GetStats() *TableStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Table) GetOption() *Option {
	if x != nil {
		return x.Option
	}
	return nil
}

var File_shortdb_domain_table_v1_table_proto protoreflect.FileDescriptor

var file_shortdb_domain_table_v1_table_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xa9, 0x04, 0x0a, 0x05,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x38, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x38, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x30,
	0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x51, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x4f, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x51, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x51, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x03, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72,
	0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortdb_domain_table_v1_table_proto_rawDescOnce sync.Once
	file_shortdb_domain_table_v1_table_proto_rawDescData = file_shortdb_domain_table_v1_table_proto_rawDesc
)

func file_shortdb_domain_table_v1_table_proto_rawDescGZIP() []byte {
	file_shortdb_domain_table_v1_table_proto_rawDescOnce.Do(func() {
		file_shortdb_domain_table_v1_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortdb_domain_table_v1_table_proto_rawDescData)
	})
	return file_shortdb_domain_table_v1_table_proto_rawDescData
}

var file_shortdb_domain_table_v1_table_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shortdb_domain_table_v1_table_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shortdb_domain_table_v1_table_proto_goTypes = []interface{}{
	(Type)(0),          // 0: shortdb.table.v1.Type
	(*Option)(nil),     // 1: shortdb.table.v1.Option
	(*Table)(nil),      // 2: shortdb.table.v1.Table
	nil,                // 3: shortdb.table.v1.Table.FieldsEntry
	nil,                // 4: shortdb.table.v1.Table.PagesEntry
	nil,                // 5: shortdb.table.v1.Table.IndexEntry
	(*TableStats)(nil), // 6: shortdb.table.v1.TableStats
	(*v1.Page)(nil),    // 7: shortdb.page.v1.Page
	(*v11.Index)(nil),  // 8: shortdb.index.v1.Index
}
var file_shortdb_domain_table_v1_table_proto_depIdxs = []int32{
	3, // 0: shortdb.table.v1.Table.fields:type_name -> shortdb.table.v1.Table.FieldsEntry
	4, // 1: shortdb.table.v1.Table.pages:type_name -> shortdb.table.v1.Table.PagesEntry
	5, // 2: shortdb.table.v1.Table.index:type_name -> shortdb.table.v1.Table.IndexEntry
	6, // 3: shortdb.table.v1.Table.stats:type_name -> shortdb.table.v1.TableStats
	1, // 4: shortdb.table.v1.Table.option:type_name -> shortdb.table.v1.Option
	0, // 5: shortdb.table.v1.Table.FieldsEntry.value:type_name -> shortdb.table.v1.Type
	7, // 6: shortdb.table.v1.Table.PagesEntry.value:type_name -> shortdb.page.v1.Page
	8, // 7: shortdb.table.v1.Table.IndexEntry.value:type_name -> shortdb.index.v1.Index
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_shortdb_domain_table_v1_table_proto_init() }
func file_shortdb_domain_table_v1_table_proto_init() {
	if File_shortdb_domain_table_v1_table_proto != nil {
		return
	}
	file_shortdb_domain_table_v1_stats_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shortdb_domain_table_v1_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Option); i {
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
		file_shortdb_domain_table_v1_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
			RawDescriptor: file_shortdb_domain_table_v1_table_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shortdb_domain_table_v1_table_proto_goTypes,
		DependencyIndexes: file_shortdb_domain_table_v1_table_proto_depIdxs,
		EnumInfos:         file_shortdb_domain_table_v1_table_proto_enumTypes,
		MessageInfos:      file_shortdb_domain_table_v1_table_proto_msgTypes,
	}.Build()
	File_shortdb_domain_table_v1_table_proto = out.File
	file_shortdb_domain_table_v1_table_proto_rawDesc = nil
	file_shortdb_domain_table_v1_table_proto_goTypes = nil
	file_shortdb_domain_table_v1_table_proto_depIdxs = nil
}

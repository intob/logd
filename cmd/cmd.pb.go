// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: cmd.proto

package cmd

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Name int32

const (
	Name_WRITE Name = 0
	Name_TAIL  Name = 1
	Name_PING  Name = 2
	Name_QUERY Name = 3
)

// Enum value maps for Name.
var (
	Name_name = map[int32]string{
		0: "WRITE",
		1: "TAIL",
		2: "PING",
		3: "QUERY",
	}
	Name_value = map[string]int32{
		"WRITE": 0,
		"TAIL":  1,
		"PING":  2,
		"QUERY": 3,
	}
)

func (x Name) Enum() *Name {
	p := new(Name)
	*p = x
	return p
}

func (x Name) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Name) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[0].Descriptor()
}

func (Name) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[0]
}

func (x Name) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Name.Descriptor instead.
func (Name) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0}
}

type Lvl int32

const (
	Lvl_LVL_UNKNOWN Lvl = 0
	Lvl_TRACE       Lvl = 1
	Lvl_DEBUG       Lvl = 2
	Lvl_INFO        Lvl = 3
	Lvl_WARN        Lvl = 4
	Lvl_ERROR       Lvl = 5
	Lvl_FATAL       Lvl = 6
)

// Enum value maps for Lvl.
var (
	Lvl_name = map[int32]string{
		0: "LVL_UNKNOWN",
		1: "TRACE",
		2: "DEBUG",
		3: "INFO",
		4: "WARN",
		5: "ERROR",
		6: "FATAL",
	}
	Lvl_value = map[string]int32{
		"LVL_UNKNOWN": 0,
		"TRACE":       1,
		"DEBUG":       2,
		"INFO":        3,
		"WARN":        4,
		"ERROR":       5,
		"FATAL":       6,
	}
)

func (x Lvl) Enum() *Lvl {
	p := new(Lvl)
	*p = x
	return p
}

func (x Lvl) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Lvl) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_proto_enumTypes[1].Descriptor()
}

func (Lvl) Type() protoreflect.EnumType {
	return &file_cmd_proto_enumTypes[1]
}

func (x Lvl) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Lvl.Descriptor instead.
func (Lvl) EnumDescriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{1}
}

type Cmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        Name         `protobuf:"varint,1,opt,name=name,proto3,enum=Name" json:"name,omitempty"`
	Msg         *Msg         `protobuf:"bytes,2,opt,name=msg,proto3,oneof" json:"msg,omitempty"`
	QueryParams *QueryParams `protobuf:"bytes,3,opt,name=queryParams,proto3,oneof" json:"queryParams,omitempty"`
}

func (x *Cmd) Reset() {
	*x = Cmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cmd) ProtoMessage() {}

func (x *Cmd) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cmd.ProtoReflect.Descriptor instead.
func (*Cmd) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{0}
}

func (x *Cmd) GetName() Name {
	if x != nil {
		return x.Name
	}
	return Name_WRITE
}

func (x *Cmd) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *Cmd) GetQueryParams() *QueryParams {
	if x != nil {
		return x.QueryParams
	}
	return nil
}

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
	Lvl Lvl                    `protobuf:"varint,6,opt,name=lvl,proto3,enum=Lvl" json:"lvl,omitempty"`
	Txt string                 `protobuf:"bytes,7,opt,name=txt,proto3" json:"txt,omitempty"`
	Key string                 `protobuf:"bytes,12,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{1}
}

func (x *Msg) GetT() *timestamppb.Timestamp {
	if x != nil {
		return x.T
	}
	return nil
}

func (x *Msg) GetLvl() Lvl {
	if x != nil {
		return x.Lvl
	}
	return Lvl_LVL_UNKNOWN
}

func (x *Msg) GetTxt() string {
	if x != nil {
		return x.Txt
	}
	return ""
}

func (x *Msg) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type QueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset    *uint32                `protobuf:"varint,1,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
	Limit     *uint32                `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	TStart    *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=tStart,proto3,oneof" json:"tStart,omitempty"`
	TEnd      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=tEnd,proto3,oneof" json:"tEnd,omitempty"`
	Lvl       *Lvl                   `protobuf:"varint,8,opt,name=lvl,proto3,enum=Lvl,oneof" json:"lvl,omitempty"`
	KeyPrefix *string                `protobuf:"bytes,13,opt,name=keyPrefix,proto3,oneof" json:"keyPrefix,omitempty"`
}

func (x *QueryParams) Reset() {
	*x = QueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParams) ProtoMessage() {}

func (x *QueryParams) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParams.ProtoReflect.Descriptor instead.
func (*QueryParams) Descriptor() ([]byte, []int) {
	return file_cmd_proto_rawDescGZIP(), []int{2}
}

func (x *QueryParams) GetOffset() uint32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *QueryParams) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *QueryParams) GetTStart() *timestamppb.Timestamp {
	if x != nil {
		return x.TStart
	}
	return nil
}

func (x *QueryParams) GetTEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.TEnd
	}
	return nil
}

func (x *QueryParams) GetLvl() Lvl {
	if x != nil && x.Lvl != nil {
		return *x.Lvl
	}
	return Lvl_LVL_UNKNOWN
}

func (x *QueryParams) GetKeyPrefix() string {
	if x != nil && x.KeyPrefix != nil {
		return *x.KeyPrefix
	}
	return ""
}

var File_cmd_proto protoreflect.FileDescriptor

var file_cmd_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a,
	0x03, 0x43, 0x6d, 0x64, 0x12, 0x19, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x05, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4d,
	0x73, 0x67, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x0b,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48,
	0x01, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x88, 0x01,
	0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x73, 0x67, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x6b, 0x0a, 0x03, 0x4d, 0x73, 0x67,
	0x12, 0x28, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x01, 0x74, 0x12, 0x16, 0x0a, 0x03, 0x6c, 0x76,
	0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x04, 0x2e, 0x4c, 0x76, 0x6c, 0x52, 0x03, 0x6c,
	0x76, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x78, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x74, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xb2, 0x02, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x37,
	0x0a, 0x06, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x06, 0x74, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x45, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x03, 0x52, 0x04, 0x74, 0x45, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x03,
	0x6c, 0x76, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x04, 0x2e, 0x4c, 0x76, 0x6c, 0x48,
	0x04, 0x52, 0x03, 0x6c, 0x76, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x6b, 0x65, 0x79,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x09,
	0x6b, 0x65, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x74, 0x45, 0x6e, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6c, 0x76, 0x6c, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x6b, 0x65, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2a, 0x30, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x03, 0x2a, 0x56, 0x0a,
	0x03, 0x4c, 0x76, 0x6c, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x56, 0x4c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x4e, 0x46, 0x4f, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x04, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41,
	0x54, 0x41, 0x4c, 0x10, 0x06, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x63, 0x6d, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_proto_rawDescOnce sync.Once
	file_cmd_proto_rawDescData = file_cmd_proto_rawDesc
)

func file_cmd_proto_rawDescGZIP() []byte {
	file_cmd_proto_rawDescOnce.Do(func() {
		file_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_proto_rawDescData)
	})
	return file_cmd_proto_rawDescData
}

var file_cmd_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmd_proto_goTypes = []interface{}{
	(Name)(0),                     // 0: Name
	(Lvl)(0),                      // 1: Lvl
	(*Cmd)(nil),                   // 2: Cmd
	(*Msg)(nil),                   // 3: Msg
	(*QueryParams)(nil),           // 4: QueryParams
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_cmd_proto_depIdxs = []int32{
	0, // 0: Cmd.name:type_name -> Name
	3, // 1: Cmd.msg:type_name -> Msg
	4, // 2: Cmd.queryParams:type_name -> QueryParams
	5, // 3: Msg.t:type_name -> google.protobuf.Timestamp
	1, // 4: Msg.lvl:type_name -> Lvl
	5, // 5: QueryParams.tStart:type_name -> google.protobuf.Timestamp
	5, // 6: QueryParams.tEnd:type_name -> google.protobuf.Timestamp
	1, // 7: QueryParams.lvl:type_name -> Lvl
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cmd_proto_init() }
func file_cmd_proto_init() {
	if File_cmd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cmd); i {
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
		file_cmd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_cmd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParams); i {
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
	file_cmd_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_cmd_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_proto_goTypes,
		DependencyIndexes: file_cmd_proto_depIdxs,
		EnumInfos:         file_cmd_proto_enumTypes,
		MessageInfos:      file_cmd_proto_msgTypes,
	}.Build()
	File_cmd_proto = out.File
	file_cmd_proto_rawDesc = nil
	file_cmd_proto_goTypes = nil
	file_cmd_proto_depIdxs = nil
}

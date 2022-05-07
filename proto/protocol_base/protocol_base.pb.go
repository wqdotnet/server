// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/protocol_base.proto

package protocol_base

import (
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

//消息号
type MSG_BASE int32

const (
	MSG_BASE_PLACEHOLDER_BASE MSG_BASE = 0 //占位
	//模块id
	MSG_BASE_Module MSG_BASE = 100 //基础模块
	//消息method id
	MSG_BASE_HeartBeat MSG_BASE = 101 //心跳
	MSG_BASE_NoticeMsg MSG_BASE = 103 //错误提示
)

// Enum value maps for MSG_BASE.
var (
	MSG_BASE_name = map[int32]string{
		0:   "PLACEHOLDER_BASE",
		100: "Module",
		101: "HeartBeat",
		103: "NoticeMsg",
	}
	MSG_BASE_value = map[string]int32{
		"PLACEHOLDER_BASE": 0,
		"Module":           100,
		"HeartBeat":        101,
		"NoticeMsg":        103,
	}
)

func (x MSG_BASE) Enum() *MSG_BASE {
	p := new(MSG_BASE)
	*p = x
	return p
}

func (x MSG_BASE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSG_BASE) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_protocol_base_proto_enumTypes[0].Descriptor()
}

func (MSG_BASE) Type() protoreflect.EnumType {
	return &file_proto_protocol_base_proto_enumTypes[0]
}

func (x MSG_BASE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG_BASE.Descriptor instead.
func (MSG_BASE) EnumDescriptor() ([]byte, []int) {
	return file_proto_protocol_base_proto_rawDescGZIP(), []int{0}
}

//心跳  1
type C2S_HeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2S_HeartBeat) Reset() {
	*x = C2S_HeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocol_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_HeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_HeartBeat) ProtoMessage() {}

func (x *C2S_HeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocol_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_HeartBeat.ProtoReflect.Descriptor instead.
func (*C2S_HeartBeat) Descriptor() ([]byte, []int) {
	return file_proto_protocol_base_proto_rawDescGZIP(), []int{0}
}

type S2C_HeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode   int32 `protobuf:"zigzag32,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Timestamp int64 `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"` //服务器时间
}

func (x *S2C_HeartBeat) Reset() {
	*x = S2C_HeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocol_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_HeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_HeartBeat) ProtoMessage() {}

func (x *S2C_HeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocol_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_HeartBeat.ProtoReflect.Descriptor instead.
func (*S2C_HeartBeat) Descriptor() ([]byte, []int) {
	return file_proto_protocol_base_proto_rawDescGZIP(), []int{1}
}

func (x *S2C_HeartBeat) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *S2C_HeartBeat) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

//错误提示消息
type S2C_NoticeMsg_S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode   int32  `protobuf:"zigzag32,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	NoticeMsg string `protobuf:"bytes,2,opt,name=NoticeMsg,proto3" json:"NoticeMsg,omitempty"`
}

func (x *S2C_NoticeMsg_S) Reset() {
	*x = S2C_NoticeMsg_S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_protocol_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_NoticeMsg_S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_NoticeMsg_S) ProtoMessage() {}

func (x *S2C_NoticeMsg_S) ProtoReflect() protoreflect.Message {
	mi := &file_proto_protocol_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_NoticeMsg_S.ProtoReflect.Descriptor instead.
func (*S2C_NoticeMsg_S) Descriptor() ([]byte, []int) {
	return file_proto_protocol_base_proto_rawDescGZIP(), []int{2}
}

func (x *S2C_NoticeMsg_S) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *S2C_NoticeMsg_S) GetNoticeMsg() string {
	if x != nil {
		return x.NoticeMsg
	}
	return ""
}

var File_proto_protocol_base_proto protoreflect.FileDescriptor

var file_proto_protocol_base_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x63, 0x32,
	0x73, 0x5f, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x22, 0x47, 0x0a, 0x0d, 0x73,
	0x32, 0x63, 0x5f, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x49, 0x0a, 0x0f, 0x73, 0x32, 0x63, 0x5f, 0x4e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x4d, 0x73, 0x67, 0x5f, 0x53, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x2a,
	0x4a, 0x0a, 0x08, 0x4d, 0x53, 0x47, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x12, 0x14, 0x0a, 0x10, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x10, 0x64, 0x12, 0x0d, 0x0a,
	0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x10, 0x65, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x4d, 0x73, 0x67, 0x10, 0x67, 0x42, 0x15, 0x5a, 0x13, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x62, 0x61,
	0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_protocol_base_proto_rawDescOnce sync.Once
	file_proto_protocol_base_proto_rawDescData = file_proto_protocol_base_proto_rawDesc
)

func file_proto_protocol_base_proto_rawDescGZIP() []byte {
	file_proto_protocol_base_proto_rawDescOnce.Do(func() {
		file_proto_protocol_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_protocol_base_proto_rawDescData)
	})
	return file_proto_protocol_base_proto_rawDescData
}

var file_proto_protocol_base_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_protocol_base_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_protocol_base_proto_goTypes = []interface{}{
	(MSG_BASE)(0),           // 0: protocol_base.MSG_BASE
	(*C2S_HeartBeat)(nil),   // 1: protocol_base.c2s_HeartBeat
	(*S2C_HeartBeat)(nil),   // 2: protocol_base.s2c_HeartBeat
	(*S2C_NoticeMsg_S)(nil), // 3: protocol_base.s2c_NoticeMsg_S
}
var file_proto_protocol_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_protocol_base_proto_init() }
func file_proto_protocol_base_proto_init() {
	if File_proto_protocol_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_protocol_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_HeartBeat); i {
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
		file_proto_protocol_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_HeartBeat); i {
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
		file_proto_protocol_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_NoticeMsg_S); i {
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
			RawDescriptor: file_proto_protocol_base_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_protocol_base_proto_goTypes,
		DependencyIndexes: file_proto_protocol_base_proto_depIdxs,
		EnumInfos:         file_proto_protocol_base_proto_enumTypes,
		MessageInfos:      file_proto_protocol_base_proto_msgTypes,
	}.Build()
	File_proto_protocol_base_proto = out.File
	file_proto_protocol_base_proto_rawDesc = nil
	file_proto_protocol_base_proto_goTypes = nil
	file_proto_protocol_base_proto_depIdxs = nil
}

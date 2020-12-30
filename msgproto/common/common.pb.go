// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: common.proto

package common

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//部队状态
type TroopsState int32

const (
	TroopsState_StandBy   TroopsState = 0 //等待命令
	TroopsState_Move      TroopsState = 1 //移动中
	TroopsState_Pause     TroopsState = 2 //暂停
	TroopsState_Stationed TroopsState = 3 //原地驻扎
	TroopsState_fight     TroopsState = 4 //战斗
)

// Enum value maps for TroopsState.
var (
	TroopsState_name = map[int32]string{
		0: "StandBy",
		1: "Move",
		2: "Pause",
		3: "Stationed",
		4: "fight",
	}
	TroopsState_value = map[string]int32{
		"StandBy":   0,
		"Move":      1,
		"Pause":     2,
		"Stationed": 3,
		"fight":     4,
	}
)

func (x TroopsState) Enum() *TroopsState {
	p := new(TroopsState)
	*p = x
	return p
}

func (x TroopsState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TroopsState) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (TroopsState) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x TroopsState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TroopsState.Descriptor instead.
func (TroopsState) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type NetworkMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module   int32  `protobuf:"varint,1,opt,name=Module,proto3" json:"Module,omitempty"`
	Method   int32  `protobuf:"varint,2,opt,name=Method,proto3" json:"Method,omitempty"`
	MsgBytes []byte `protobuf:"bytes,3,opt,name=MsgBytes,proto3" json:"MsgBytes,omitempty"`
}

func (x *NetworkMsg) Reset() {
	*x = NetworkMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkMsg) ProtoMessage() {}

func (x *NetworkMsg) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkMsg.ProtoReflect.Descriptor instead.
func (*NetworkMsg) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkMsg) GetModule() int32 {
	if x != nil {
		return x.Module
	}
	return 0
}

func (x *NetworkMsg) GetMethod() int32 {
	if x != nil {
		return x.Method
	}
	return 0
}

func (x *NetworkMsg) GetMsgBytes() []byte {
	if x != nil {
		return x.MsgBytes
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x58, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x4d, 0x73, 0x67, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x2a, 0x49, 0x0a, 0x0b, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x42, 0x79, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x4d, 0x6f, 0x76, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x61, 0x75, 0x73, 0x65, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x10, 0x03,
	0x12, 0x09, 0x0a, 0x05, 0x66, 0x69, 0x67, 0x68, 0x74, 0x10, 0x04, 0x42, 0x1b, 0x5a, 0x19, 0x73,
	0x6c, 0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x73, 0x67, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_common_proto_goTypes = []interface{}{
	(TroopsState)(0),   // 0: common.TroopsState
	(*NetworkMsg)(nil), // 1: common.NetworkMsg
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkMsg); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}

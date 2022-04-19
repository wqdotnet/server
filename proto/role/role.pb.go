// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/role.proto

package role

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

//用户游戏信息
type Pb_RoleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleID        int32           `protobuf:"varint,1,opt,name=roleID,proto3" json:"roleID,omitempty"`
	RoleName      string          `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName,omitempty"`
	Exp           int64           `protobuf:"varint,3,opt,name=exp,proto3" json:"exp,omitempty"`
	Level         int32           `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	Country       int32           `protobuf:"varint,5,opt,name=country,proto3" json:"country,omitempty"`                                                                                                      //所属国家
	TesourcesList map[int32]int32 `protobuf:"bytes,6,rep,name=TesourcesList,proto3" json:"TesourcesList,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` //玩家资源列表
	Settings      *RoleSettings   `protobuf:"bytes,7,opt,name=Settings,proto3" json:"Settings,omitempty"`                                                                                                     //游戏内设置
}

func (x *Pb_RoleInfo) Reset() {
	*x = Pb_RoleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pb_RoleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_RoleInfo) ProtoMessage() {}

func (x *Pb_RoleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_RoleInfo.ProtoReflect.Descriptor instead.
func (*Pb_RoleInfo) Descriptor() ([]byte, []int) {
	return file_proto_role_proto_rawDescGZIP(), []int{0}
}

func (x *Pb_RoleInfo) GetRoleID() int32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *Pb_RoleInfo) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *Pb_RoleInfo) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *Pb_RoleInfo) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Pb_RoleInfo) GetCountry() int32 {
	if x != nil {
		return x.Country
	}
	return 0
}

func (x *Pb_RoleInfo) GetTesourcesList() map[int32]int32 {
	if x != nil {
		return x.TesourcesList
	}
	return nil
}

func (x *Pb_RoleInfo) GetSettings() *RoleSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

//游戏设置
type RoleSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutoSelectTactics bool `protobuf:"varint,1,opt,name=AutoSelectTactics,proto3" json:"AutoSelectTactics,omitempty"` //自动选择战术
}

func (x *RoleSettings) Reset() {
	*x = RoleSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleSettings) ProtoMessage() {}

func (x *RoleSettings) ProtoReflect() protoreflect.Message {
	mi := &file_proto_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleSettings.ProtoReflect.Descriptor instead.
func (*RoleSettings) Descriptor() ([]byte, []int) {
	return file_proto_role_proto_rawDescGZIP(), []int{1}
}

func (x *RoleSettings) GetAutoSelectTactics() bool {
	if x != nil {
		return x.AutoSelectTactics
	}
	return false
}

var File_proto_role_proto protoreflect.FileDescriptor

var file_proto_role_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0xc1, 0x02, 0x0a, 0x0b, 0x50, 0x62, 0x5f,
	0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x4a,
	0x0a, 0x0d, 0x54, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x62, 0x5f,
	0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x54, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x54, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3c, 0x0a, 0x0c,
	0x52, 0x6f, 0x6c, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x11,
	0x41, 0x75, 0x74, 0x6f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x61, 0x63, 0x74, 0x69, 0x63,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x54, 0x61, 0x63, 0x74, 0x69, 0x63, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_role_proto_rawDescOnce sync.Once
	file_proto_role_proto_rawDescData = file_proto_role_proto_rawDesc
)

func file_proto_role_proto_rawDescGZIP() []byte {
	file_proto_role_proto_rawDescOnce.Do(func() {
		file_proto_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_role_proto_rawDescData)
	})
	return file_proto_role_proto_rawDescData
}

var file_proto_role_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_role_proto_goTypes = []interface{}{
	(*Pb_RoleInfo)(nil),  // 0: role.Pb_RoleInfo
	(*RoleSettings)(nil), // 1: role.RoleSettings
	nil,                  // 2: role.Pb_RoleInfo.TesourcesListEntry
}
var file_proto_role_proto_depIdxs = []int32{
	2, // 0: role.Pb_RoleInfo.TesourcesList:type_name -> role.Pb_RoleInfo.TesourcesListEntry
	1, // 1: role.Pb_RoleInfo.Settings:type_name -> role.RoleSettings
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_role_proto_init() }
func file_proto_role_proto_init() {
	if File_proto_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pb_RoleInfo); i {
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
		file_proto_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleSettings); i {
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
			RawDescriptor: file_proto_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_role_proto_goTypes,
		DependencyIndexes: file_proto_role_proto_depIdxs,
		MessageInfos:      file_proto_role_proto_msgTypes,
	}.Build()
	File_proto_role_proto = out.File
	file_proto_role_proto_rawDesc = nil
	file_proto_role_proto_goTypes = nil
	file_proto_role_proto_depIdxs = nil
}

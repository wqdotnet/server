// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/account.proto

package account

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	role "server/proto/role"
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
type MSG_ACCOUNT int32

const (
	MSG_ACCOUNT_PLACEHOLDER MSG_ACCOUNT = 0 //占位
	//账号模块
	MSG_ACCOUNT_Module MSG_ACCOUNT = 1000
	//method
	MSG_ACCOUNT_C2S_Login          MSG_ACCOUNT = 1001 //用户登陆
	MSG_ACCOUNT_S2C_Login          MSG_ACCOUNT = 1002
	MSG_ACCOUNT_C2S_CreateRole     MSG_ACCOUNT = 1003 //创建角色
	MSG_ACCOUNT_S2C_CreateRole     MSG_ACCOUNT = 1004
	MSG_ACCOUNT_C2S_UpdateRoleName MSG_ACCOUNT = 1005 //修改角色名
	MSG_ACCOUNT_S2C_UpdateRoleName MSG_ACCOUNT = 1006
	MSG_ACCOUNT_S2C_RoleAddExp     MSG_ACCOUNT = 1010 //角色获取经验
)

// Enum value maps for MSG_ACCOUNT.
var (
	MSG_ACCOUNT_name = map[int32]string{
		0:    "PLACEHOLDER",
		1000: "Module",
		1001: "C2S_Login",
		1002: "S2C_Login",
		1003: "C2S_CreateRole",
		1004: "S2C_CreateRole",
		1005: "C2S_UpdateRoleName",
		1006: "S2C_UpdateRoleName",
		1010: "S2C_RoleAddExp",
	}
	MSG_ACCOUNT_value = map[string]int32{
		"PLACEHOLDER":        0,
		"Module":             1000,
		"C2S_Login":          1001,
		"S2C_Login":          1002,
		"C2S_CreateRole":     1003,
		"S2C_CreateRole":     1004,
		"C2S_UpdateRoleName": 1005,
		"S2C_UpdateRoleName": 1006,
		"S2C_RoleAddExp":     1010,
	}
)

func (x MSG_ACCOUNT) Enum() *MSG_ACCOUNT {
	p := new(MSG_ACCOUNT)
	*p = x
	return p
}

func (x MSG_ACCOUNT) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSG_ACCOUNT) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_account_proto_enumTypes[0].Descriptor()
}

func (MSG_ACCOUNT) Type() protoreflect.EnumType {
	return &file_proto_account_proto_enumTypes[0]
}

func (x MSG_ACCOUNT) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG_ACCOUNT.Descriptor instead.
func (MSG_ACCOUNT) EnumDescriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{0}
}

//用户登陆
type C2S_Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *C2S_Login) Reset() {
	*x = C2S_Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_Login) ProtoMessage() {}

func (x *C2S_Login) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_Login.ProtoReflect.Descriptor instead.
func (*C2S_Login) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_Login) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *C2S_Login) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type S2C_Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode  int32             `protobuf:"zigzag32,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	RoleInfo *role.Pb_RoleInfo `protobuf:"bytes,2,opt,name=RoleInfo,proto3" json:"RoleInfo,omitempty"`
}

func (x *S2C_Login) Reset() {
	*x = S2C_Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_Login) ProtoMessage() {}

func (x *S2C_Login) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_Login.ProtoReflect.Descriptor instead.
func (*S2C_Login) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{1}
}

func (x *S2C_Login) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *S2C_Login) GetRoleInfo() *role.Pb_RoleInfo {
	if x != nil {
		return x.RoleInfo
	}
	return nil
}

// 创建角色
type C2S_CreateRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleName string `protobuf:"bytes,1,opt,name=RoleName,proto3" json:"RoleName,omitempty"`
	Country  int32  `protobuf:"varint,2,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *C2S_CreateRole) Reset() {
	*x = C2S_CreateRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_CreateRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_CreateRole) ProtoMessage() {}

func (x *C2S_CreateRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_CreateRole.ProtoReflect.Descriptor instead.
func (*C2S_CreateRole) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{2}
}

func (x *C2S_CreateRole) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *C2S_CreateRole) GetCountry() int32 {
	if x != nil {
		return x.Country
	}
	return 0
}

type S2C_CreateRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"zigzag32,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Roleid  int32 `protobuf:"varint,2,opt,name=roleid,proto3" json:"roleid,omitempty"`
}

func (x *S2C_CreateRole) Reset() {
	*x = S2C_CreateRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_CreateRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_CreateRole) ProtoMessage() {}

func (x *S2C_CreateRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_CreateRole.ProtoReflect.Descriptor instead.
func (*S2C_CreateRole) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{3}
}

func (x *S2C_CreateRole) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *S2C_CreateRole) GetRoleid() int32 {
	if x != nil {
		return x.Roleid
	}
	return 0
}

var File_proto_account_proto protoreflect.FileDescriptor

var file_proto_account_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x41, 0x0a, 0x09, 0x63, 0x32, 0x73, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x54, 0x0a, 0x09, 0x73, 0x32, 0x63, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x62, 0x5f, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x08, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x46, 0x0a, 0x0e, 0x63, 0x32, 0x73,
	0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52,
	0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52,
	0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x22, 0x42, 0x0a, 0x0e, 0x73, 0x32, 0x63, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x69, 0x64, 0x2a, 0xbc, 0x01, 0x0a, 0x0b, 0x4d, 0x53, 0x47, 0x5f, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f,
	0x4c, 0x44, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x10, 0xe8, 0x07, 0x12, 0x0e, 0x0a, 0x09, 0x43, 0x32, 0x53, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x10, 0xe9, 0x07, 0x12, 0x0e, 0x0a, 0x09, 0x53, 0x32, 0x43, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x10, 0xea, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x32, 0x53, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x10, 0xeb, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x53, 0x32, 0x43, 0x5f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x10, 0xec, 0x07, 0x12, 0x17, 0x0a,
	0x12, 0x43, 0x32, 0x53, 0x5f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x10, 0xed, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x53, 0x32, 0x43, 0x5f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0xee, 0x07, 0x12,
	0x13, 0x0a, 0x0e, 0x53, 0x32, 0x43, 0x5f, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x64, 0x64, 0x45, 0x78,
	0x70, 0x10, 0xf2, 0x07, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_proto_rawDescOnce sync.Once
	file_proto_account_proto_rawDescData = file_proto_account_proto_rawDesc
)

func file_proto_account_proto_rawDescGZIP() []byte {
	file_proto_account_proto_rawDescOnce.Do(func() {
		file_proto_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_proto_rawDescData)
	})
	return file_proto_account_proto_rawDescData
}

var file_proto_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_account_proto_goTypes = []interface{}{
	(MSG_ACCOUNT)(0),         // 0: account.MSG_ACCOUNT
	(*C2S_Login)(nil),        // 1: account.c2s_Login
	(*S2C_Login)(nil),        // 2: account.s2c_Login
	(*C2S_CreateRole)(nil),   // 3: account.c2s_CreateRole
	(*S2C_CreateRole)(nil),   // 4: account.s2c_CreateRole
	(*role.Pb_RoleInfo)(nil), // 5: role.Pb_RoleInfo
}
var file_proto_account_proto_depIdxs = []int32{
	5, // 0: account.s2c_Login.RoleInfo:type_name -> role.Pb_RoleInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_account_proto_init() }
func file_proto_account_proto_init() {
	if File_proto_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_Login); i {
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
		file_proto_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_Login); i {
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
		file_proto_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_CreateRole); i {
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
		file_proto_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_CreateRole); i {
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
			RawDescriptor: file_proto_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_account_proto_goTypes,
		DependencyIndexes: file_proto_account_proto_depIdxs,
		EnumInfos:         file_proto_account_proto_enumTypes,
		MessageInfos:      file_proto_account_proto_msgTypes,
	}.Build()
	File_proto_account_proto = out.File
	file_proto_account_proto_rawDesc = nil
	file_proto_account_proto_goTypes = nil
	file_proto_account_proto_depIdxs = nil
}

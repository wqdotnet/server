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
	MSG_ACCOUNT_Login      MSG_ACCOUNT = 1001 //用户登陆
	MSG_ACCOUNT_Register   MSG_ACCOUNT = 1002 //注册账号
	MSG_ACCOUNT_CreateRole MSG_ACCOUNT = 1003 //创建角色
)

// Enum value maps for MSG_ACCOUNT.
var (
	MSG_ACCOUNT_name = map[int32]string{
		0:    "PLACEHOLDER",
		1000: "Module",
		1001: "Login",
		1002: "Register",
		1003: "CreateRole",
	}
	MSG_ACCOUNT_value = map[string]int32{
		"PLACEHOLDER": 0,
		"Module":      1000,
		"Login":       1001,
		"Register":    1002,
		"CreateRole":  1003,
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

//用户账号信息
type P_Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account            string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password           string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Equipment          string `protobuf:"bytes,3,opt,name=equipment,proto3" json:"equipment,omitempty"`                   //设备信息
	RegistrationSource string `protobuf:"bytes,4,opt,name=registrationSource,proto3" json:"registrationSource,omitempty"` //注册来源
	RegistrationTime   string `protobuf:"bytes,5,opt,name=registrationTime,proto3" json:"registrationTime,omitempty"`     //注册时间
}

func (x *P_Account) Reset() {
	*x = P_Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P_Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P_Account) ProtoMessage() {}

func (x *P_Account) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use P_Account.ProtoReflect.Descriptor instead.
func (*P_Account) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{0}
}

func (x *P_Account) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *P_Account) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *P_Account) GetEquipment() string {
	if x != nil {
		return x.Equipment
	}
	return ""
}

func (x *P_Account) GetRegistrationSource() string {
	if x != nil {
		return x.RegistrationSource
	}
	return ""
}

func (x *P_Account) GetRegistrationTime() string {
	if x != nil {
		return x.RegistrationTime
	}
	return ""
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
		mi := &file_proto_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_Login) ProtoMessage() {}

func (x *C2S_Login) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use C2S_Login.ProtoReflect.Descriptor instead.
func (*C2S_Login) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{1}
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
	Settings map[uint32]string `protobuf:"bytes,3,rep,name=Settings,proto3" json:"Settings,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //游戏内设置
	ItemList map[uint32]string `protobuf:"bytes,7,rep,name=ItemList,proto3" json:"ItemList,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //背包物品列表
}

func (x *S2C_Login) Reset() {
	*x = S2C_Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_Login) ProtoMessage() {}

func (x *S2C_Login) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use S2C_Login.ProtoReflect.Descriptor instead.
func (*S2C_Login) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{2}
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

func (x *S2C_Login) GetSettings() map[uint32]string {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *S2C_Login) GetItemList() map[uint32]string {
	if x != nil {
		return x.ItemList
	}
	return nil
}

type C2S_Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account   string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	CDK       string `protobuf:"bytes,3,opt,name=CDK,proto3" json:"CDK,omitempty"`             //注册码
	Source    string `protobuf:"bytes,4,opt,name=Source,proto3" json:"Source,omitempty"`       //注册来源
	Equipment string `protobuf:"bytes,5,opt,name=Equipment,proto3" json:"Equipment,omitempty"` //设备信息
}

func (x *C2S_Register) Reset() {
	*x = C2S_Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_Register) ProtoMessage() {}

func (x *C2S_Register) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use C2S_Register.ProtoReflect.Descriptor instead.
func (*C2S_Register) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{3}
}

func (x *C2S_Register) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *C2S_Register) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *C2S_Register) GetCDK() string {
	if x != nil {
		return x.CDK
	}
	return ""
}

func (x *C2S_Register) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *C2S_Register) GetEquipment() string {
	if x != nil {
		return x.Equipment
	}
	return ""
}

type S2C_Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"zigzag32,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *S2C_Register) Reset() {
	*x = S2C_Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_Register) ProtoMessage() {}

func (x *S2C_Register) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_Register.ProtoReflect.Descriptor instead.
func (*S2C_Register) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{4}
}

func (x *S2C_Register) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

// 创建角色
type C2S_CreateRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleName string `protobuf:"bytes,1,opt,name=RoleName,proto3" json:"RoleName,omitempty"`
	Sex      uint32 `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`
	HeadID   uint32 `protobuf:"varint,3,opt,name=headID,proto3" json:"headID,omitempty"`
}

func (x *C2S_CreateRole) Reset() {
	*x = C2S_CreateRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_CreateRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_CreateRole) ProtoMessage() {}

func (x *C2S_CreateRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[5]
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
	return file_proto_account_proto_rawDescGZIP(), []int{5}
}

func (x *C2S_CreateRole) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *C2S_CreateRole) GetSex() uint32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *C2S_CreateRole) GetHeadID() uint32 {
	if x != nil {
		return x.HeadID
	}
	return 0
}

type S2C_CreateRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode  int32             `protobuf:"zigzag32,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	RoleInfo *role.Pb_RoleInfo `protobuf:"bytes,2,opt,name=RoleInfo,proto3" json:"RoleInfo,omitempty"`
}

func (x *S2C_CreateRole) Reset() {
	*x = S2C_CreateRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_CreateRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_CreateRole) ProtoMessage() {}

func (x *S2C_CreateRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[6]
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
	return file_proto_account_proto_rawDescGZIP(), []int{6}
}

func (x *S2C_CreateRole) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *S2C_CreateRole) GetRoleInfo() *role.Pb_RoleInfo {
	if x != nil {
		return x.RoleInfo
	}
	return nil
}

var File_proto_account_proto protoreflect.FileDescriptor

var file_proto_account_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbb, 0x01, 0x0a, 0x09, 0x50, 0x5f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x41,
	0x0a, 0x09, 0x63, 0x32, 0x73, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0xca, 0x02, 0x0a, 0x09, 0x73, 0x32, 0x63, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x52, 0x6f, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x50, 0x62, 0x5f, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x73, 0x32, 0x63, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x73, 0x32, 0x63, 0x5f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8c,
	0x01, 0x0a, 0x0c, 0x63, 0x32, 0x73, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x44, 0x4b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x43, 0x44, 0x4b, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x28, 0x0a,
	0x0c, 0x73, 0x32, 0x63, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x56, 0x0a, 0x0e, 0x63, 0x32, 0x73, 0x5f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x6f, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x49, 0x44, 0x22,
	0x59, 0x0a, 0x0e, 0x73, 0x32, 0x63, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x52,
	0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x50, 0x62, 0x5f, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x57, 0x0a, 0x0b, 0x4d, 0x53,
	0x47, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4c, 0x41,
	0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x10, 0xe8, 0x07, 0x12, 0x0a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x10, 0xe9, 0x07, 0x12, 0x0d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x10,
	0xea, 0x07, 0x12, 0x0f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x10, 0xeb, 0x07, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_proto_account_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_account_proto_goTypes = []interface{}{
	(MSG_ACCOUNT)(0),         // 0: account.MSG_ACCOUNT
	(*P_Account)(nil),        // 1: account.P_Account
	(*C2S_Login)(nil),        // 2: account.c2s_Login
	(*S2C_Login)(nil),        // 3: account.s2c_Login
	(*C2S_Register)(nil),     // 4: account.c2s_Register
	(*S2C_Register)(nil),     // 5: account.s2c_Register
	(*C2S_CreateRole)(nil),   // 6: account.c2s_CreateRole
	(*S2C_CreateRole)(nil),   // 7: account.s2c_CreateRole
	nil,                      // 8: account.s2c_Login.SettingsEntry
	nil,                      // 9: account.s2c_Login.ItemListEntry
	(*role.Pb_RoleInfo)(nil), // 10: role.Pb_RoleInfo
}
var file_proto_account_proto_depIdxs = []int32{
	10, // 0: account.s2c_Login.RoleInfo:type_name -> role.Pb_RoleInfo
	8,  // 1: account.s2c_Login.Settings:type_name -> account.s2c_Login.SettingsEntry
	9,  // 2: account.s2c_Login.ItemList:type_name -> account.s2c_Login.ItemListEntry
	10, // 3: account.s2c_CreateRole.RoleInfo:type_name -> role.Pb_RoleInfo
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_account_proto_init() }
func file_proto_account_proto_init() {
	if File_proto_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P_Account); i {
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
		file_proto_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_Register); i {
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
		file_proto_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_Register); i {
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
		file_proto_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   9,
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

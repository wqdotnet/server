// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type NetworkMsg struct {
	Module               int32    `protobuf:"varint,1,opt,name=Module,proto3" json:"Module,omitempty"`
	Method               int32    `protobuf:"varint,2,opt,name=Method,proto3" json:"Method,omitempty"`
	MsgBytes             []byte   `protobuf:"bytes,3,opt,name=MsgBytes,proto3" json:"MsgBytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkMsg) Reset()         { *m = NetworkMsg{} }
func (m *NetworkMsg) String() string { return proto.CompactTextString(m) }
func (*NetworkMsg) ProtoMessage()    {}
func (*NetworkMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *NetworkMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkMsg.Unmarshal(m, b)
}
func (m *NetworkMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkMsg.Marshal(b, m, deterministic)
}
func (m *NetworkMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkMsg.Merge(m, src)
}
func (m *NetworkMsg) XXX_Size() int {
	return xxx_messageInfo_NetworkMsg.Size(m)
}
func (m *NetworkMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkMsg.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkMsg proto.InternalMessageInfo

func (m *NetworkMsg) GetModule() int32 {
	if m != nil {
		return m.Module
	}
	return 0
}

func (m *NetworkMsg) GetMethod() int32 {
	if m != nil {
		return m.Method
	}
	return 0
}

func (m *NetworkMsg) GetMsgBytes() []byte {
	if m != nil {
		return m.MsgBytes
	}
	return nil
}

type UserLoginTos struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Index                int32    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginTos) Reset()         { *m = UserLoginTos{} }
func (m *UserLoginTos) String() string { return proto.CompactTextString(m) }
func (*UserLoginTos) ProtoMessage()    {}
func (*UserLoginTos) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *UserLoginTos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginTos.Unmarshal(m, b)
}
func (m *UserLoginTos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginTos.Marshal(b, m, deterministic)
}
func (m *UserLoginTos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginTos.Merge(m, src)
}
func (m *UserLoginTos) XXX_Size() int {
	return xxx_messageInfo_UserLoginTos.Size(m)
}
func (m *UserLoginTos) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginTos.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginTos proto.InternalMessageInfo

func (m *UserLoginTos) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserLoginTos) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type UserLoginToc struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginToc) Reset()         { *m = UserLoginToc{} }
func (m *UserLoginToc) String() string { return proto.CompactTextString(m) }
func (*UserLoginToc) ProtoMessage()    {}
func (*UserLoginToc) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *UserLoginToc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginToc.Unmarshal(m, b)
}
func (m *UserLoginToc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginToc.Marshal(b, m, deterministic)
}
func (m *UserLoginToc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginToc.Merge(m, src)
}
func (m *UserLoginToc) XXX_Size() int {
	return xxx_messageInfo_UserLoginToc.Size(m)
}
func (m *UserLoginToc) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginToc.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginToc proto.InternalMessageInfo

func (m *UserLoginToc) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UserLoginToc) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type MoveTos struct {
	Status               *PStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveTos) Reset()         { *m = MoveTos{} }
func (m *MoveTos) String() string { return proto.CompactTextString(m) }
func (*MoveTos) ProtoMessage()    {}
func (*MoveTos) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *MoveTos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveTos.Unmarshal(m, b)
}
func (m *MoveTos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveTos.Marshal(b, m, deterministic)
}
func (m *MoveTos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveTos.Merge(m, src)
}
func (m *MoveTos) XXX_Size() int {
	return xxx_messageInfo_MoveTos.Size(m)
}
func (m *MoveTos) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveTos.DiscardUnknown(m)
}

var xxx_messageInfo_MoveTos proto.InternalMessageInfo

func (m *MoveTos) GetStatus() *PStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type PStatus struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Direction            int32    `protobuf:"varint,3,opt,name=direction,proto3" json:"direction,omitempty"`
	Lenght               int32    `protobuf:"varint,4,opt,name=lenght,proto3" json:"lenght,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PStatus) Reset()         { *m = PStatus{} }
func (m *PStatus) String() string { return proto.CompactTextString(m) }
func (*PStatus) ProtoMessage()    {}
func (*PStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *PStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PStatus.Unmarshal(m, b)
}
func (m *PStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PStatus.Marshal(b, m, deterministic)
}
func (m *PStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PStatus.Merge(m, src)
}
func (m *PStatus) XXX_Size() int {
	return xxx_messageInfo_PStatus.Size(m)
}
func (m *PStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PStatus proto.InternalMessageInfo

func (m *PStatus) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *PStatus) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *PStatus) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *PStatus) GetLenght() int32 {
	if m != nil {
		return m.Lenght
	}
	return 0
}

func init() {
	proto.RegisterType((*NetworkMsg)(nil), "NetworkMsg")
	proto.RegisterType((*UserLoginTos)(nil), "UserLogin_tos")
	proto.RegisterType((*UserLoginToc)(nil), "UserLogin_toc")
	proto.RegisterType((*MoveTos)(nil), "move_tos")
	proto.RegisterType((*PStatus)(nil), "p_status")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x89, 0x73, 0xb5, 0x79, 0x4e, 0x90, 0x20, 0x52, 0xd4, 0xc3, 0xec, 0x69, 0x17, 0x77,
	0xd0, 0xa3, 0xb7, 0x9d, 0xed, 0x0e, 0x01, 0xc1, 0x83, 0x50, 0x6a, 0xfb, 0xc8, 0x8a, 0x6b, 0x32,
	0xf2, 0x52, 0x6d, 0xff, 0x7b, 0x69, 0x92, 0x29, 0xde, 0xf2, 0xfb, 0x85, 0xf7, 0xbd, 0xe4, 0x03,
	0xde, 0x91, 0x5a, 0x1f, 0xac, 0x71, 0x26, 0x7f, 0x03, 0xd8, 0xa2, 0xfb, 0x36, 0xf6, 0xb3, 0x20,
	0x25, 0xae, 0x21, 0x29, 0x4c, 0xd3, 0xef, 0x31, 0x63, 0x4b, 0xb6, 0x9a, 0xcb, 0x48, 0xde, 0xa3,
	0xdb, 0x99, 0x26, 0x3b, 0x89, 0xde, 0x93, 0xb8, 0x81, 0xb4, 0x20, 0xb5, 0x19, 0x1d, 0x52, 0x36,
	0x5b, 0xb2, 0xd5, 0x42, 0xfe, 0x72, 0xbe, 0x81, 0x8b, 0x57, 0x42, 0xfb, 0x62, 0x54, 0xab, 0x4b,
	0x67, 0x48, 0xdc, 0x02, 0xef, 0x09, 0x6d, 0xa9, 0xab, 0x2e, 0xe4, 0x73, 0x99, 0x4e, 0x62, 0x5b,
	0x75, 0x28, 0xae, 0x60, 0xde, 0xea, 0x06, 0x87, 0xb8, 0x20, 0x40, 0xfe, 0xfc, 0x3f, 0xa3, 0x16,
	0x19, 0x9c, 0x51, 0x5f, 0xd7, 0x48, 0xe4, 0x13, 0x52, 0x79, 0x44, 0x71, 0x09, 0xb3, 0x8e, 0x94,
	0x1f, 0xe7, 0x72, 0x3a, 0xe6, 0x0f, 0x90, 0x76, 0xe6, 0x0b, 0xfd, 0xee, 0x7b, 0x48, 0xc8, 0x55,
	0xae, 0x0f, 0x63, 0xe7, 0x8f, 0x7c, 0x7d, 0x28, 0x83, 0x90, 0xf1, 0x22, 0x7f, 0x87, 0xf4, 0xe8,
	0xc4, 0x02, 0xd8, 0x10, 0x2b, 0x60, 0xc3, 0x44, 0x63, 0x7c, 0x17, 0x1b, 0xc5, 0x1d, 0xf0, 0xa6,
	0xb5, 0x58, 0xbb, 0xd6, 0x68, 0xff, 0xe9, 0xb9, 0xfc, 0x13, 0x53, 0x53, 0x7b, 0xd4, 0x6a, 0xe7,
	0xb2, 0xd3, 0xd0, 0x54, 0xa0, 0x8f, 0xc4, 0xd7, 0xfd, 0xf4, 0x13, 0x00, 0x00, 0xff, 0xff, 0x52,
	0xb2, 0x29, 0x0c, 0x7b, 0x01, 0x00, 0x00,
}

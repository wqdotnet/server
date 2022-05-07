// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: proto/item.proto

package item

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

//策划配置表属性
type Item_Type int32

const (
	Item_Type_UnKnown       Item_Type = 0  //未知
	Item_Type_Money         Item_Type = 1  //货币
	Item_Type_Equip         Item_Type = 2  //装备
	Item_Type_Prop          Item_Type = 3  //道具
	Item_Type_Drug          Item_Type = 4  //丹药
	Item_Type_Trump         Item_Type = 5  //法宝
	Item_Type_Rune          Item_Type = 6  //符文
	Item_Type_Material      Item_Type = 7  //材料
	Item_Type_Exp           Item_Type = 8  //经验
	Item_Type_Pokemon       Item_Type = 9  //宠物
	Item_Type_TrumpMaterial Item_Type = 10 //法宝材料
)

// Enum value maps for Item_Type.
var (
	Item_Type_name = map[int32]string{
		0:  "UnKnown",
		1:  "Money",
		2:  "Equip",
		3:  "Prop",
		4:  "Drug",
		5:  "Trump",
		6:  "Rune",
		7:  "Material",
		8:  "Exp",
		9:  "Pokemon",
		10: "TrumpMaterial",
	}
	Item_Type_value = map[string]int32{
		"UnKnown":       0,
		"Money":         1,
		"Equip":         2,
		"Prop":          3,
		"Drug":          4,
		"Trump":         5,
		"Rune":          6,
		"Material":      7,
		"Exp":           8,
		"Pokemon":       9,
		"TrumpMaterial": 10,
	}
)

func (x Item_Type) Enum() *Item_Type {
	p := new(Item_Type)
	*p = x
	return p
}

func (x Item_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Item_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_item_proto_enumTypes[0].Descriptor()
}

func (Item_Type) Type() protoreflect.EnumType {
	return &file_proto_item_proto_enumTypes[0]
}

func (x Item_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Item_Type.Descriptor instead.
func (Item_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{0}
}

//消息号
type MSG_ITEM int32

const (
	MSG_ITEM_PLACEHOLDER MSG_ITEM = 0 //占位
	//账号模块
	MSG_ITEM_Module          MSG_ITEM = 3000
	MSG_ITEM_GetBackpackInfo MSG_ITEM = 3001 //获取背包信息
)

// Enum value maps for MSG_ITEM.
var (
	MSG_ITEM_name = map[int32]string{
		0:    "PLACEHOLDER",
		3000: "Module",
		3001: "GetBackpackInfo",
	}
	MSG_ITEM_value = map[string]int32{
		"PLACEHOLDER":     0,
		"Module":          3000,
		"GetBackpackInfo": 3001,
	}
)

func (x MSG_ITEM) Enum() *MSG_ITEM {
	p := new(MSG_ITEM)
	*p = x
	return p
}

func (x MSG_ITEM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSG_ITEM) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_item_proto_enumTypes[1].Descriptor()
}

func (MSG_ITEM) Type() protoreflect.EnumType {
	return &file_proto_item_proto_enumTypes[1]
}

func (x MSG_ITEM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG_ITEM.Descriptor instead.
func (MSG_ITEM) EnumDescriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{1}
}

type PbItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid            string          `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` //道具id
	ItemId          uint32          `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	ItemNumber      int64           `protobuf:"varint,3,opt,name=item_number,json=itemNumber,proto3" json:"item_number,omitempty"`                                                                                                           //拥有数量
	ItemLock        uint32          `protobuf:"varint,4,opt,name=item_lock,json=itemLock,proto3" json:"item_lock,omitempty"`                                                                                                                 //是否锁定0没锁
	UseLocation     uint32          `protobuf:"varint,5,opt,name=use_location,json=useLocation,proto3" json:"use_location,omitempty"`                                                                                                        //是否装备，0否，1~6装备位置
	EquipBaseAttr   map[int64]int64 `protobuf:"bytes,6,rep,name=equip_base_attr,json=equipBaseAttr,proto3" json:"equip_base_attr,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`       //基本属性
	EquipExtendAttr map[int64]int64 `protobuf:"bytes,7,rep,name=equip_extend_attr,json=equipExtendAttr,proto3" json:"equip_extend_attr,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` //扩展属性
	TrumpCurStar    uint32          `protobuf:"varint,8,opt,name=trump_cur_star,json=trumpCurStar,proto3" json:"trump_cur_star,omitempty"`                                                                                                   //法宝当前星级
	TrumpCurExp     uint32          `protobuf:"varint,9,opt,name=trump_cur_exp,json=trumpCurExp,proto3" json:"trump_cur_exp,omitempty"`                                                                                                      //法宝当前经验
	AttributeJson   string          `protobuf:"bytes,10,opt,name=AttributeJson,proto3" json:"AttributeJson,omitempty"`                                                                                                                       //扩展属性 丹方 丹药
	Name            string          `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	Desc            string          `protobuf:"bytes,12,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *PbItem) Reset() {
	*x = PbItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PbItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PbItem) ProtoMessage() {}

func (x *PbItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PbItem.ProtoReflect.Descriptor instead.
func (*PbItem) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{0}
}

func (x *PbItem) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PbItem) GetItemId() uint32 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *PbItem) GetItemNumber() int64 {
	if x != nil {
		return x.ItemNumber
	}
	return 0
}

func (x *PbItem) GetItemLock() uint32 {
	if x != nil {
		return x.ItemLock
	}
	return 0
}

func (x *PbItem) GetUseLocation() uint32 {
	if x != nil {
		return x.UseLocation
	}
	return 0
}

func (x *PbItem) GetEquipBaseAttr() map[int64]int64 {
	if x != nil {
		return x.EquipBaseAttr
	}
	return nil
}

func (x *PbItem) GetEquipExtendAttr() map[int64]int64 {
	if x != nil {
		return x.EquipExtendAttr
	}
	return nil
}

func (x *PbItem) GetTrumpCurStar() uint32 {
	if x != nil {
		return x.TrumpCurStar
	}
	return 0
}

func (x *PbItem) GetTrumpCurExp() uint32 {
	if x != nil {
		return x.TrumpCurExp
	}
	return 0
}

func (x *PbItem) GetAttributeJson() string {
	if x != nil {
		return x.AttributeJson
	}
	return ""
}

func (x *PbItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PbItem) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type C2S_GetBackpackInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2S_GetBackpackInfo) Reset() {
	*x = C2S_GetBackpackInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_GetBackpackInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_GetBackpackInfo) ProtoMessage() {}

func (x *C2S_GetBackpackInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_GetBackpackInfo.ProtoReflect.Descriptor instead.
func (*C2S_GetBackpackInfo) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{1}
}

type S2C_GetBackpackInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *S2C_GetBackpackInfo) Reset() {
	*x = S2C_GetBackpackInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_GetBackpackInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_GetBackpackInfo) ProtoMessage() {}

func (x *S2C_GetBackpackInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_GetBackpackInfo.ProtoReflect.Descriptor instead.
func (*S2C_GetBackpackInfo) Descriptor() ([]byte, []int) {
	return file_proto_item_proto_rawDescGZIP(), []int{2}
}

var File_proto_item_proto protoreflect.FileDescriptor

var file_proto_item_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0xcc, 0x04, 0x0a, 0x06, 0x70, 0x62, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x21,
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x47, 0x0a, 0x0f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x74, 0x65,
	0x6d, 0x2e, 0x70, 0x62, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x42, 0x61,
	0x73, 0x65, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x42, 0x61, 0x73, 0x65, 0x41, 0x74, 0x74, 0x72, 0x12, 0x4d, 0x0a, 0x11, 0x65, 0x71,
	0x75, 0x69, 0x70, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x62, 0x49,
	0x74, 0x65, 0x6d, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x41,
	0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x74, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x72, 0x75,
	0x6d, 0x70, 0x5f, 0x63, 0x75, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x74, 0x72, 0x75, 0x6d, 0x70, 0x43, 0x75, 0x72, 0x53, 0x74, 0x61, 0x72, 0x12,
	0x22, 0x0a, 0x0d, 0x74, 0x72, 0x75, 0x6d, 0x70, 0x5f, 0x63, 0x75, 0x72, 0x5f, 0x65, 0x78, 0x70,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x6d, 0x70, 0x43, 0x75, 0x72,
	0x45, 0x78, 0x70, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x1a, 0x40, 0x0a, 0x12, 0x45, 0x71, 0x75, 0x69, 0x70, 0x42, 0x61, 0x73, 0x65, 0x41, 0x74,
	0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x42, 0x0a, 0x14, 0x45, 0x71, 0x75, 0x69, 0x70, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x41, 0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x15, 0x0a, 0x13, 0x63, 0x32, 0x73, 0x5f, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x70, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x15,
	0x0a, 0x13, 0x73, 0x32, 0x63, 0x5f, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x70, 0x61, 0x63,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x8e, 0x01, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x5f, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x72, 0x6f, 0x70, 0x10, 0x03,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x72, 0x75, 0x67, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x72,
	0x75, 0x6d, 0x70, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x75, 0x6e, 0x65, 0x10, 0x06, 0x12,
	0x0c, 0x0a, 0x08, 0x4d, 0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x10, 0x07, 0x12, 0x07, 0x0a,
	0x03, 0x45, 0x78, 0x70, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x6f, 0x6b, 0x65, 0x6d, 0x6f,
	0x6e, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x72, 0x75, 0x6d, 0x70, 0x4d, 0x61, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x10, 0x0a, 0x2a, 0x3e, 0x0a, 0x08, 0x4d, 0x53, 0x47, 0x5f, 0x49, 0x54,
	0x45, 0x4d, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f, 0x4c, 0x44, 0x45,
	0x52, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x10, 0xb8, 0x17,
	0x12, 0x14, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x61, 0x63, 0x6b, 0x70, 0x61, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x10, 0xb9, 0x17, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x69, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_item_proto_rawDescOnce sync.Once
	file_proto_item_proto_rawDescData = file_proto_item_proto_rawDesc
)

func file_proto_item_proto_rawDescGZIP() []byte {
	file_proto_item_proto_rawDescOnce.Do(func() {
		file_proto_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_item_proto_rawDescData)
	})
	return file_proto_item_proto_rawDescData
}

var file_proto_item_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_item_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_item_proto_goTypes = []interface{}{
	(Item_Type)(0),              // 0: item.Item_Type
	(MSG_ITEM)(0),               // 1: item.MSG_ITEM
	(*PbItem)(nil),              // 2: item.pbItem
	(*C2S_GetBackpackInfo)(nil), // 3: item.c2s_GetBackpackInfo
	(*S2C_GetBackpackInfo)(nil), // 4: item.s2c_GetBackpackInfo
	nil,                         // 5: item.pbItem.EquipBaseAttrEntry
	nil,                         // 6: item.pbItem.EquipExtendAttrEntry
}
var file_proto_item_proto_depIdxs = []int32{
	5, // 0: item.pbItem.equip_base_attr:type_name -> item.pbItem.EquipBaseAttrEntry
	6, // 1: item.pbItem.equip_extend_attr:type_name -> item.pbItem.EquipExtendAttrEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_item_proto_init() }
func file_proto_item_proto_init() {
	if File_proto_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PbItem); i {
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
		file_proto_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_GetBackpackInfo); i {
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
		file_proto_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_GetBackpackInfo); i {
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
			RawDescriptor: file_proto_item_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_item_proto_goTypes,
		DependencyIndexes: file_proto_item_proto_depIdxs,
		EnumInfos:         file_proto_item_proto_enumTypes,
		MessageInfos:      file_proto_item_proto_msgTypes,
	}.Build()
	File_proto_item_proto = out.File
	file_proto_item_proto_rawDesc = nil
	file_proto_item_proto_goTypes = nil
	file_proto_item_proto_depIdxs = nil
}

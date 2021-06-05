// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: troops.proto

package troops

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	common "server/msgproto/common"
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

//消息号
type MSG_TROOPS int32

const (
	MSG_TROOPS_PLACEHOLDER_TROOPS MSG_TROOPS = 0 //占位
	//大地图模块
	MSG_TROOPS_Module_TROOPS MSG_TROOPS = 3000
	//部队信息
	MSG_TROOPS_S2C_TroopsList       MSG_TROOPS = 3001 //
	MSG_TROOPS_S2C_UpdateTroopsInfo MSG_TROOPS = 3002 //
	MSG_TROOPS_C2S_Behavior         MSG_TROOPS = 3011 //部队行为
	MSG_TROOPS_S2C_Behavior         MSG_TROOPS = 3012
	MSG_TROOPS_S2C_TroopsAddExp     MSG_TROOPS = 3020
)

// Enum value maps for MSG_TROOPS.
var (
	MSG_TROOPS_name = map[int32]string{
		0:    "PLACEHOLDER_TROOPS",
		3000: "Module_TROOPS",
		3001: "S2C_TroopsList",
		3002: "S2C_UpdateTroopsInfo",
		3011: "C2S_Behavior",
		3012: "S2C_Behavior",
		3020: "S2C_TroopsAddExp",
	}
	MSG_TROOPS_value = map[string]int32{
		"PLACEHOLDER_TROOPS":   0,
		"Module_TROOPS":        3000,
		"S2C_TroopsList":       3001,
		"S2C_UpdateTroopsInfo": 3002,
		"C2S_Behavior":         3011,
		"S2C_Behavior":         3012,
		"S2C_TroopsAddExp":     3020,
	}
)

func (x MSG_TROOPS) Enum() *MSG_TROOPS {
	p := new(MSG_TROOPS)
	*p = x
	return p
}

func (x MSG_TROOPS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MSG_TROOPS) Descriptor() protoreflect.EnumDescriptor {
	return file_troops_proto_enumTypes[0].Descriptor()
}

func (MSG_TROOPS) Type() protoreflect.EnumType {
	return &file_troops_proto_enumTypes[0]
}

func (x MSG_TROOPS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MSG_TROOPS.Descriptor instead.
func (MSG_TROOPS) EnumDescriptor() ([]byte, []int) {
	return file_troops_proto_rawDescGZIP(), []int{0}
}

type TroopsBehavior int32

const (
	TroopsBehavior_AddExp  TroopsBehavior = 0
	TroopsBehavior_Recruit TroopsBehavior = 1
	TroopsBehavior_OnStage TroopsBehavior = 2
	TroopsBehavior_Exit    TroopsBehavior = 3
)

// Enum value maps for TroopsBehavior.
var (
	TroopsBehavior_name = map[int32]string{
		0: "AddExp",
		1: "Recruit",
		2: "OnStage",
		3: "Exit",
	}
	TroopsBehavior_value = map[string]int32{
		"AddExp":  0,
		"Recruit": 1,
		"OnStage": 2,
		"Exit":    3,
	}
)

func (x TroopsBehavior) Enum() *TroopsBehavior {
	p := new(TroopsBehavior)
	*p = x
	return p
}

func (x TroopsBehavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TroopsBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_troops_proto_enumTypes[1].Descriptor()
}

func (TroopsBehavior) Type() protoreflect.EnumType {
	return &file_troops_proto_enumTypes[1]
}

func (x TroopsBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TroopsBehavior.Descriptor instead.
func (TroopsBehavior) EnumDescriptor() ([]byte, []int) {
	return file_troops_proto_rawDescGZIP(), []int{1}
}

type P_RoleTroops struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roleid     int32       `protobuf:"varint,1,opt,name=roleid,proto3" json:"roleid,omitempty"`
	TroopsList []*P_Troops `protobuf:"bytes,2,rep,name=TroopsList,proto3" json:"TroopsList,omitempty"`
}

func (x *P_RoleTroops) Reset() {
	*x = P_RoleTroops{}
	if protoimpl.UnsafeEnabled {
		mi := &file_troops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P_RoleTroops) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P_RoleTroops) ProtoMessage() {}

func (x *P_RoleTroops) ProtoReflect() protoreflect.Message {
	mi := &file_troops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P_RoleTroops.ProtoReflect.Descriptor instead.
func (*P_RoleTroops) Descriptor() ([]byte, []int) {
	return file_troops_proto_rawDescGZIP(), []int{0}
}

func (x *P_RoleTroops) GetRoleid() int32 {
	if x != nil {
		return x.Roleid
	}
	return 0
}

func (x *P_RoleTroops) GetTroopsList() []*P_Troops {
	if x != nil {
		return x.TroopsList
	}
	return nil
}

//部队信息
type P_Troops struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TroopsID   int32              `protobuf:"varint,1,opt,name=TroopsID,proto3" json:"TroopsID,omitempty"` //部队名/id
	Name       string             `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Country    int32              `protobuf:"varint,3,opt,name=country,proto3" json:"country,omitempty"`                     //国家归属
	AreasList  []int32            `protobuf:"varint,4,rep,packed,name=AreasList,proto3" json:"AreasList,omitempty"`          //移动路径 区域id list
	AreasIndex int32              `protobuf:"varint,5,opt,name=AreasIndex,proto3" json:"AreasIndex,omitempty"`               //当前区域ID
	State      common.TroopsState `protobuf:"varint,6,opt,name=State,proto3,enum=common.TroopsState" json:"State,omitempty"` //状态 0:未出动    1:移动  2:驻扎(暂停)  3:战斗
	RowHP      []int32            `protobuf:"varint,7,rep,packed,name=RowHP,proto3" json:"RowHP,omitempty"`                  //每排血量
	//兵种组成
	Type           int32   `protobuf:"varint,8,opt,name=Type,proto3" json:"Type,omitempty"`           //部队类型
	MaxNumber      int32   `protobuf:"varint,9,opt,name=MaxNumber,proto3" json:"MaxNumber,omitempty"` //部队上限
	Number         int32   `protobuf:"varint,10,opt,name=Number,proto3" json:"Number,omitempty"`      //当前数量
	Level          int32   `protobuf:"varint,11,opt,name=Level,proto3" json:"Level,omitempty"`        //等级
	Roleid         int32   `protobuf:"varint,12,opt,name=roleid,proto3" json:"roleid,omitempty"`
	Attack         int32   `protobuf:"varint,13,opt,name=Attack,proto3" json:"Attack,omitempty"`                 // 攻击	att_A
	Defensive      int32   `protobuf:"varint,14,opt,name=Defensive,proto3" json:"Defensive,omitempty"`           // 防御	def_A
	AttackSuper    int32   `protobuf:"varint,15,opt,name=AttackSuper,proto3" json:"AttackSuper,omitempty"`       // 强攻	att_damage
	DefensiveSuper int32   `protobuf:"varint,17,opt,name=DefensiveSuper,proto3" json:"DefensiveSuper,omitempty"` // 强防	def_damage
	Strong         int32   `protobuf:"varint,18,opt,name=Strong,proto3" json:"Strong,omitempty"`                 // 强壮	att_B
	Control        int32   `protobuf:"varint,19,opt,name=Control,proto3" json:"Control,omitempty"`               // 掌控	def_B
	Leader         int32   `protobuf:"varint,20,opt,name=Leader,proto3" json:"Leader,omitempty"`                 // 统帅	leader
	Strength       int32   `protobuf:"varint,21,opt,name=Strength,proto3" json:"Strength,omitempty"`             // 勇气	strength
	Politics       int32   `protobuf:"varint,22,opt,name=politics,proto3" json:"politics,omitempty"`             //政治
	Intelligence   int32   `protobuf:"varint,23,opt,name=Intelligence,proto3" json:"Intelligence,omitempty"`     //智力
	SkillID        int32   `protobuf:"varint,24,opt,name=SkillID,proto3" json:"SkillID,omitempty"`               //战法id
	TalentID       int32   `protobuf:"varint,25,opt,name=TalentID,proto3" json:"TalentID,omitempty"`             //天赋id
	TacticsID      int32   `protobuf:"varint,26,opt,name=TacticsID,proto3" json:"TacticsID,omitempty"`           //战术id
	QueueNum       int32   `protobuf:"varint,27,opt,name=QueueNum,proto3" json:"QueueNum,omitempty"`             //队列编号(第几个进入队列的)
	FightType      int32   `protobuf:"varint,28,opt,name=FightType,proto3" json:"FightType,omitempty"`           //战斗状态  0:待战   1:上阵
	FightState     int32   `protobuf:"varint,29,opt,name=FightState,proto3" json:"FightState,omitempty"`         //战斗类型  0:国战   1:副本  2:剧本
	StageNumber    int32   `protobuf:"varint,30,opt,name=StageNumber,proto3" json:"StageNumber,omitempty"`       //上阵编号
	SelectTactics  int32   `protobuf:"varint,31,opt,name=SelectTactics,proto3" json:"SelectTactics,omitempty"`   //已选择的战术(战斗中)
	Bufflist       []int32 `protobuf:"varint,32,rep,packed,name=Bufflist,proto3" json:"Bufflist,omitempty"`      //buf 列表
}

func (x *P_Troops) Reset() {
	*x = P_Troops{}
	if protoimpl.UnsafeEnabled {
		mi := &file_troops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P_Troops) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P_Troops) ProtoMessage() {}

func (x *P_Troops) ProtoReflect() protoreflect.Message {
	mi := &file_troops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P_Troops.ProtoReflect.Descriptor instead.
func (*P_Troops) Descriptor() ([]byte, []int) {
	return file_troops_proto_rawDescGZIP(), []int{1}
}

func (x *P_Troops) GetTroopsID() int32 {
	if x != nil {
		return x.TroopsID
	}
	return 0
}

func (x *P_Troops) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *P_Troops) GetCountry() int32 {
	if x != nil {
		return x.Country
	}
	return 0
}

func (x *P_Troops) GetAreasList() []int32 {
	if x != nil {
		return x.AreasList
	}
	return nil
}

func (x *P_Troops) GetAreasIndex() int32 {
	if x != nil {
		return x.AreasIndex
	}
	return 0
}

func (x *P_Troops) GetState() common.TroopsState {
	if x != nil {
		return x.State
	}
	return common.TroopsState_StandBy
}

func (x *P_Troops) GetRowHP() []int32 {
	if x != nil {
		return x.RowHP
	}
	return nil
}

func (x *P_Troops) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *P_Troops) GetMaxNumber() int32 {
	if x != nil {
		return x.MaxNumber
	}
	return 0
}

func (x *P_Troops) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *P_Troops) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *P_Troops) GetRoleid() int32 {
	if x != nil {
		return x.Roleid
	}
	return 0
}

func (x *P_Troops) GetAttack() int32 {
	if x != nil {
		return x.Attack
	}
	return 0
}

func (x *P_Troops) GetDefensive() int32 {
	if x != nil {
		return x.Defensive
	}
	return 0
}

func (x *P_Troops) GetAttackSuper() int32 {
	if x != nil {
		return x.AttackSuper
	}
	return 0
}

func (x *P_Troops) GetDefensiveSuper() int32 {
	if x != nil {
		return x.DefensiveSuper
	}
	return 0
}

func (x *P_Troops) GetStrong() int32 {
	if x != nil {
		return x.Strong
	}
	return 0
}

func (x *P_Troops) GetControl() int32 {
	if x != nil {
		return x.Control
	}
	return 0
}

func (x *P_Troops) GetLeader() int32 {
	if x != nil {
		return x.Leader
	}
	return 0
}

func (x *P_Troops) GetStrength() int32 {
	if x != nil {
		return x.Strength
	}
	return 0
}

func (x *P_Troops) GetPolitics() int32 {
	if x != nil {
		return x.Politics
	}
	return 0
}

func (x *P_Troops) GetIntelligence() int32 {
	if x != nil {
		return x.Intelligence
	}
	return 0
}

func (x *P_Troops) GetSkillID() int32 {
	if x != nil {
		return x.SkillID
	}
	return 0
}

func (x *P_Troops) GetTalentID() int32 {
	if x != nil {
		return x.TalentID
	}
	return 0
}

func (x *P_Troops) GetTacticsID() int32 {
	if x != nil {
		return x.TacticsID
	}
	return 0
}

func (x *P_Troops) GetQueueNum() int32 {
	if x != nil {
		return x.QueueNum
	}
	return 0
}

func (x *P_Troops) GetFightType() int32 {
	if x != nil {
		return x.FightType
	}
	return 0
}

func (x *P_Troops) GetFightState() int32 {
	if x != nil {
		return x.FightState
	}
	return 0
}

func (x *P_Troops) GetStageNumber() int32 {
	if x != nil {
		return x.StageNumber
	}
	return 0
}

func (x *P_Troops) GetSelectTactics() int32 {
	if x != nil {
		return x.SelectTactics
	}
	return 0
}

func (x *P_Troops) GetBufflist() []int32 {
	if x != nil {
		return x.Bufflist
	}
	return nil
}

type S2C_TroopsAddExp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddExp   int64 `protobuf:"varint,1,opt,name=addExp,proto3" json:"addExp,omitempty"`
	NewExp   int64 `protobuf:"varint,2,opt,name=NewExp,proto3" json:"NewExp,omitempty"`
	NewLevel int64 `protobuf:"varint,3,opt,name=NewLevel,proto3" json:"NewLevel,omitempty"`
}

func (x *S2C_TroopsAddExp) Reset() {
	*x = S2C_TroopsAddExp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_troops_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_TroopsAddExp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_TroopsAddExp) ProtoMessage() {}

func (x *S2C_TroopsAddExp) ProtoReflect() protoreflect.Message {
	mi := &file_troops_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_TroopsAddExp.ProtoReflect.Descriptor instead.
func (*S2C_TroopsAddExp) Descriptor() ([]byte, []int) {
	return file_troops_proto_rawDescGZIP(), []int{2}
}

func (x *S2C_TroopsAddExp) GetAddExp() int64 {
	if x != nil {
		return x.AddExp
	}
	return 0
}

func (x *S2C_TroopsAddExp) GetNewExp() int64 {
	if x != nil {
		return x.NewExp
	}
	return 0
}

func (x *S2C_TroopsAddExp) GetNewLevel() int64 {
	if x != nil {
		return x.NewLevel
	}
	return 0
}

//部队行为
//AddExp:加经验  parValue-> 加经验值
//Recruit:招募  parValue-> 兵种类型
//OnStage:上阵  parValue-> 上阵编号(1-5)
//Exit:下阵
type C2S_Behavior struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BehaviorID TroopsBehavior `protobuf:"varint,1,opt,name=behaviorID,proto3,enum=troops.TroopsBehavior" json:"behaviorID,omitempty"`
	TroopsID   int32          `protobuf:"varint,2,opt,name=troopsID,proto3" json:"troopsID,omitempty"`
	ParValue   []int64        `protobuf:"varint,3,rep,packed,name=parValue,proto3" json:"parValue,omitempty"`
}

func (x *C2S_Behavior) Reset() {
	*x = C2S_Behavior{}
	if protoimpl.UnsafeEnabled {
		mi := &file_troops_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_Behavior) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_Behavior) ProtoMessage() {}

func (x *C2S_Behavior) ProtoReflect() protoreflect.Message {
	mi := &file_troops_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_Behavior.ProtoReflect.Descriptor instead.
func (*C2S_Behavior) Descriptor() ([]byte, []int) {
	return file_troops_proto_rawDescGZIP(), []int{3}
}

func (x *C2S_Behavior) GetBehaviorID() TroopsBehavior {
	if x != nil {
		return x.BehaviorID
	}
	return TroopsBehavior_AddExp
}

func (x *C2S_Behavior) GetTroopsID() int32 {
	if x != nil {
		return x.TroopsID
	}
	return 0
}

func (x *C2S_Behavior) GetParValue() []int64 {
	if x != nil {
		return x.ParValue
	}
	return nil
}

//部队行为
//AddExp:加经验  item -> addExp
//Recruit:招募  item -> TroopsID
//OnStage:上阵  item -> StageNumber(上阵编号 1-5)
//Exit:下阵
type S2C_Behavior struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TroopsID   int32            `protobuf:"varint,1,opt,name=troopsID,proto3" json:"troopsID,omitempty"`
	BehaviorID TroopsBehavior   `protobuf:"varint,2,opt,name=behaviorID,proto3,enum=troops.TroopsBehavior" json:"behaviorID,omitempty"`
	Item       map[string]int64 `protobuf:"bytes,3,rep,name=item,proto3" json:"item,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Msg        string           `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"` //错误信息
}

func (x *S2C_Behavior) Reset() {
	*x = S2C_Behavior{}
	if protoimpl.UnsafeEnabled {
		mi := &file_troops_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_Behavior) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_Behavior) ProtoMessage() {}

func (x *S2C_Behavior) ProtoReflect() protoreflect.Message {
	mi := &file_troops_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_Behavior.ProtoReflect.Descriptor instead.
func (*S2C_Behavior) Descriptor() ([]byte, []int) {
	return file_troops_proto_rawDescGZIP(), []int{4}
}

func (x *S2C_Behavior) GetTroopsID() int32 {
	if x != nil {
		return x.TroopsID
	}
	return 0
}

func (x *S2C_Behavior) GetBehaviorID() TroopsBehavior {
	if x != nil {
		return x.BehaviorID
	}
	return TroopsBehavior_AddExp
}

func (x *S2C_Behavior) GetItem() map[string]int64 {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *S2C_Behavior) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

//部队信息
type S2C_TroopsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TroopsList []*P_Troops `protobuf:"bytes,1,rep,name=TroopsList,proto3" json:"TroopsList,omitempty"`
}

func (x *S2C_TroopsList) Reset() {
	*x = S2C_TroopsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_troops_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_TroopsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_TroopsList) ProtoMessage() {}

func (x *S2C_TroopsList) ProtoReflect() protoreflect.Message {
	mi := &file_troops_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_TroopsList.ProtoReflect.Descriptor instead.
func (*S2C_TroopsList) Descriptor() ([]byte, []int) {
	return file_troops_proto_rawDescGZIP(), []int{5}
}

func (x *S2C_TroopsList) GetTroopsList() []*P_Troops {
	if x != nil {
		return x.TroopsList
	}
	return nil
}

//单个部队信息更新
type S2C_UpdateTroopsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TroopsInfo *P_Troops `protobuf:"bytes,1,opt,name=TroopsInfo,proto3" json:"TroopsInfo,omitempty"`
}

func (x *S2C_UpdateTroopsInfo) Reset() {
	*x = S2C_UpdateTroopsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_troops_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_UpdateTroopsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_UpdateTroopsInfo) ProtoMessage() {}

func (x *S2C_UpdateTroopsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_troops_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_UpdateTroopsInfo.ProtoReflect.Descriptor instead.
func (*S2C_UpdateTroopsInfo) Descriptor() ([]byte, []int) {
	return file_troops_proto_rawDescGZIP(), []int{6}
}

func (x *S2C_UpdateTroopsInfo) GetTroopsInfo() *P_Troops {
	if x != nil {
		return x.TroopsInfo
	}
	return nil
}

var File_troops_proto protoreflect.FileDescriptor

var file_troops_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0c, 0x50, 0x5f, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x72,
	0x6f, 0x6f, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x0a,
	0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x2e, 0x50, 0x5f, 0x54, 0x72, 0x6f, 0x6f,
	0x70, 0x73, 0x52, 0x0a, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x83,
	0x07, 0x0a, 0x08, 0x50, 0x5f, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x72, 0x6f, 0x6f, 0x70, 0x73, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x54,
	0x72, 0x6f, 0x6f, 0x70, 0x73, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x72, 0x65, 0x61, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x41, 0x72, 0x65, 0x61, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x72, 0x65, 0x61, 0x73, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x41, 0x72, 0x65, 0x61, 0x73, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x29, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x6f, 0x6f,
	0x70, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x52, 0x6f, 0x77, 0x48, 0x50, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x52,
	0x6f, 0x77, 0x48, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x61, 0x78, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x4d, 0x61, 0x78,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x69, 0x76,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x69,
	0x76, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x75, 0x70, 0x65,
	0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x73, 0x69, 0x76,
	0x65, 0x53, 0x75, 0x70, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x44, 0x65,
	0x66, 0x65, 0x6e, 0x73, 0x69, 0x76, 0x65, 0x53, 0x75, 0x70, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74,
	0x72, 0x6f, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x74, 0x69, 0x63, 0x73, 0x18, 0x16,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x74, 0x69, 0x63, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x6c, 0x6c, 0x69, 0x67, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x44, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x19, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x61, 0x63, 0x74,
	0x69, 0x63, 0x73, 0x49, 0x44, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x54, 0x61, 0x63,
	0x74, 0x69, 0x63, 0x73, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4e,
	0x75, 0x6d, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x1c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x46, 0x69, 0x67, 0x68, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x1d,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x46, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x61, 0x63, 0x74,
	0x69, 0x63, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x54, 0x61, 0x63, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x75, 0x66, 0x66,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x20, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x42, 0x75, 0x66, 0x66,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x5e, 0x0a, 0x10, 0x73, 0x32, 0x63, 0x5f, 0x54, 0x72, 0x6f, 0x6f,
	0x70, 0x73, 0x41, 0x64, 0x64, 0x45, 0x78, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x64, 0x45,
	0x78, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x64, 0x64, 0x45, 0x78, 0x70,
	0x12, 0x16, 0x0a, 0x06, 0x4e, 0x65, 0x77, 0x45, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x4e, 0x65, 0x77, 0x45, 0x78, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x4e, 0x65, 0x77, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x22, 0x7e, 0x0a, 0x0c, 0x63, 0x32, 0x73, 0x5f, 0x42, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x0a, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x6f, 0x6f, 0x70,
	0x73, 0x2e, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x52, 0x0a, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x0c, 0x73, 0x32, 0x63, 0x5f, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x49,
	0x44, 0x12, 0x36, 0x0a, 0x0a, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x2e, 0x54,
	0x72, 0x6f, 0x6f, 0x70, 0x73, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x0a, 0x62,
	0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73,
	0x2e, 0x73, 0x32, 0x63, 0x5f, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x1a,
	0x37, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x42, 0x0a, 0x0e, 0x73, 0x32, 0x63, 0x5f,
	0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x54, 0x72,
	0x6f, 0x6f, 0x70, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x2e, 0x50, 0x5f, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73,
	0x52, 0x0a, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x14,
	0x73, 0x32, 0x63, 0x5f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x30, 0x0a, 0x0a, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x72, 0x6f, 0x6f, 0x70,
	0x73, 0x2e, 0x50, 0x5f, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x52, 0x0a, 0x54, 0x72, 0x6f, 0x6f,
	0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0xa5, 0x01, 0x0a, 0x0a, 0x4d, 0x53, 0x47, 0x5f, 0x54,
	0x52, 0x4f, 0x4f, 0x50, 0x53, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x48, 0x4f,
	0x4c, 0x44, 0x45, 0x52, 0x5f, 0x54, 0x52, 0x4f, 0x4f, 0x50, 0x53, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x54, 0x52, 0x4f, 0x4f, 0x50, 0x53, 0x10, 0xb8,
	0x17, 0x12, 0x13, 0x0a, 0x0e, 0x53, 0x32, 0x43, 0x5f, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x10, 0xb9, 0x17, 0x12, 0x19, 0x0a, 0x14, 0x53, 0x32, 0x43, 0x5f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0xba,
	0x17, 0x12, 0x11, 0x0a, 0x0c, 0x43, 0x32, 0x53, 0x5f, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x10, 0xc3, 0x17, 0x12, 0x11, 0x0a, 0x0c, 0x53, 0x32, 0x43, 0x5f, 0x42, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x10, 0xc4, 0x17, 0x12, 0x15, 0x0a, 0x10, 0x53, 0x32, 0x43, 0x5f, 0x54,
	0x72, 0x6f, 0x6f, 0x70, 0x73, 0x41, 0x64, 0x64, 0x45, 0x78, 0x70, 0x10, 0xcc, 0x17, 0x2a, 0x40,
	0x0a, 0x0e, 0x54, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x45, 0x78, 0x70, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x52, 0x65, 0x63, 0x72, 0x75, 0x69, 0x74, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x6e, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x78, 0x69, 0x74, 0x10, 0x03,
	0x42, 0x1b, 0x5a, 0x19, 0x73, 0x6c, 0x67, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6d, 0x73,
	0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x6f, 0x6f, 0x70, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_troops_proto_rawDescOnce sync.Once
	file_troops_proto_rawDescData = file_troops_proto_rawDesc
)

func file_troops_proto_rawDescGZIP() []byte {
	file_troops_proto_rawDescOnce.Do(func() {
		file_troops_proto_rawDescData = protoimpl.X.CompressGZIP(file_troops_proto_rawDescData)
	})
	return file_troops_proto_rawDescData
}

var file_troops_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_troops_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_troops_proto_goTypes = []interface{}{
	(MSG_TROOPS)(0),              // 0: troops.MSG_TROOPS
	(TroopsBehavior)(0),          // 1: troops.TroopsBehavior
	(*P_RoleTroops)(nil),         // 2: troops.P_RoleTroops
	(*P_Troops)(nil),             // 3: troops.P_Troops
	(*S2C_TroopsAddExp)(nil),     // 4: troops.s2c_TroopsAddExp
	(*C2S_Behavior)(nil),         // 5: troops.c2s_Behavior
	(*S2C_Behavior)(nil),         // 6: troops.s2c_Behavior
	(*S2C_TroopsList)(nil),       // 7: troops.s2c_TroopsList
	(*S2C_UpdateTroopsInfo)(nil), // 8: troops.s2c_UpdateTroopsInfo
	nil,                          // 9: troops.s2c_Behavior.ItemEntry
	(common.TroopsState)(0),      // 10: common.TroopsState
}
var file_troops_proto_depIdxs = []int32{
	3,  // 0: troops.P_RoleTroops.TroopsList:type_name -> troops.P_Troops
	10, // 1: troops.P_Troops.State:type_name -> common.TroopsState
	1,  // 2: troops.c2s_Behavior.behaviorID:type_name -> troops.TroopsBehavior
	1,  // 3: troops.s2c_Behavior.behaviorID:type_name -> troops.TroopsBehavior
	9,  // 4: troops.s2c_Behavior.item:type_name -> troops.s2c_Behavior.ItemEntry
	3,  // 5: troops.s2c_TroopsList.TroopsList:type_name -> troops.P_Troops
	3,  // 6: troops.s2c_UpdateTroopsInfo.TroopsInfo:type_name -> troops.P_Troops
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_troops_proto_init() }
func file_troops_proto_init() {
	if File_troops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_troops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P_RoleTroops); i {
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
		file_troops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P_Troops); i {
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
		file_troops_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_TroopsAddExp); i {
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
		file_troops_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_Behavior); i {
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
		file_troops_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_Behavior); i {
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
		file_troops_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_TroopsList); i {
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
		file_troops_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_UpdateTroopsInfo); i {
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
			RawDescriptor: file_troops_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_troops_proto_goTypes,
		DependencyIndexes: file_troops_proto_depIdxs,
		EnumInfos:         file_troops_proto_enumTypes,
		MessageInfos:      file_troops_proto_msgTypes,
	}.Build()
	File_troops_proto = out.File
	file_troops_proto_rawDesc = nil
	file_troops_proto_goTypes = nil
	file_troops_proto_depIdxs = nil
}
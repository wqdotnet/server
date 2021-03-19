package commonstruct

import (
	"fmt"
	"slgserver/gserver/cfg"
	"slgserver/msgproto/common"
	"slgserver/msgproto/troops"
	"slgserver/tool"

	log "github.com/sirupsen/logrus"
)

//TroopsStruct 部队信息
type TroopsStruct struct {
	TroopsID int32 //部队名/id
	Roleid   int32
	Name     string //归属角色名
	Pic      int32  //头像
	Country  int32  //国家归属

	Exp       int64
	Level     int32 //等级
	Attribute int32 //队伍属性 0 正常 1 幻影  3 npc

	//出战编号
	StageNumber int32 // 出战编号   0:未出战  1-5：出战

	//===============移动状态信息===============

	AreasList  []int32 //移动路径 区域id list
	AreasIndex int32   //当前区域ID

	MoveStamp   int64 //上次移动时间戳
	MoveNum     int32 //上次移动路径列表(AreasList)位置索引
	ArrivalTime int64 //预计到达时间

	//===============部队状态===========================
	State       common.TroopsState //部队状态  0:未出动    1:移动  2:驻扎(暂停)  3:战斗
	FitghtState int32              //战斗状态  0:待战   1:上阵
	FightType   int32              //战斗类型  0:国战   1:副本  2:剧本
	QueueNum    int32              //队列编号(第几个进入)

	Scene PositionScene //部队所在场景

	//========================战斗记录====================
	RoundWins      bool  //回合状态 胜、负
	SkillUseNumber int32 //技能施放次数
	KillNumber     int32 //连杀数
	SelectTactics  int32 //已选择的战法

	BuffID   []int32       //buf 效果
	FightSet *FightSetting //部队设置

	//===============基本属性================================
	//兵种组成
	Type      int32 //部队类型
	Quality   int32 //品质
	MaxNumber int32 //最大数量
	Number    int32 //当前数量

	RowHP []int32 //每排HP

	Attack         int32 // 攻击	att_A
	Defensive      int32 // 防御	def_A
	AttackSuper    int32 // 强攻	att_damage
	DefensiveSuper int32 // 强防	def_damage
	Strong         int32 // 强壮	att_B
	Control        int32 // 掌控	def_B

	//2极属性
	Leader   int32 // 统帅	leader
	Strength int32 // 勇气	strength

	Politics     int32 //政治
	Intelligence int32 //智力

	SkillID   int32   //战法id
	TalentID  int32   //天赋id
	TacticsID []int32 //战术id
}

//PositionScene 当前位置所在的场景
type PositionScene int32

const (
	//SceneNULL 空场景
	SceneNULL PositionScene = 0
	//SceneBigMap 大地图
	SceneBigMap PositionScene = 1
)

//NewTroops create troops
func NewTroops(name string, troopsid, attribute, country, typeid int32) *TroopsStruct {
	troopscfg := cfg.GetTroopsCfg(typeid)
	if troopscfg == nil {
		log.Error("配置中无此部队ID:", typeid)
		//troopscfg = cfg.GetTroopsCfg(1)
		return nil
	}

	//每排血量
	blood := int32(troopscfg.OriginBlood)

	//初始化技能(战术)
	Strategylist := []int32{
		cfg.GetGlobalInt("StrategieScheme"),
		cfg.GetGlobalInt("StrategiesDef"),
		cfg.GetGlobalInt("StrategiesAtt"),
	}
	for _, v := range troopscfg.StrategyID {
		Strategylist = append(Strategylist, int32(v))
	}

	return &TroopsStruct{
		TroopsID:    troopsid,
		Attribute:   attribute,
		Name:        fmt.Sprintf("%v - %v", name, troopscfg.Name),
		Country:     country,
		State:       common.TroopsState_StandBy,
		FitghtState: 0,
		Scene:       SceneNULL,

		//兵种组成
		Level:          1,                        //等级
		Type:           typeid,                   //部队类型
		Quality:        int32(troopscfg.Quality), //品质
		MaxNumber:      blood,                    //每排最大血量
		Number:         blood * 4,                //总数量
		RowHP:          []int32{blood, blood, blood, blood},
		Attack:         int32(troopscfg.OriginAtt), //攻击力
		Defensive:      int32(troopscfg.OriginDef), //防御力
		AttackSuper:    0,                          // 强攻	att_damage
		DefensiveSuper: 0,                          // 强防	def_damage
		Strong:         0,                          // 强壮	att_B 战法攻击
		Control:        0,                          // 掌控	def_B 战法防御

		Leader:       int32(troopscfg.Leader),       // 统帅	leader
		Strength:     int32(troopscfg.Strength),     // 勇气	strength
		Politics:     int32(troopscfg.Politics),     //政治
		Intelligence: int32(troopscfg.Intelligence), //智力

		SkillID:   int32(troopscfg.SkillID), //战法id
		TalentID:  0,                        //天赋id
		TacticsID: Strategylist,             //战术id
		BuffID:    make([]int32, 0),
		//战斗设置
		FightSet: &FightSetting{},
	}
}

//CalculationAttribute 计算属性
//装备道具加成  buff效果
func (t *TroopsStruct) CalculationAttribute() {
	troopscfg := cfg.GetTroopsCfg(t.Type)

	growAtt := cfg.GetGlobalInt("GrowAtt")
	growDef := cfg.GetGlobalInt("GrowDef")
	growBlood := cfg.GetGlobalInt("GrowBlood")

	//血量
	blood := int32(troopscfg.OriginBlood) + t.Level*growAtt
	t.MaxNumber = blood  //每排最大血量
	t.Number = blood * 4 //总数量
	t.RowHP = []int32{blood, blood, blood, blood}

	t.Attack = int32(troopscfg.OriginAtt) + t.Level*growDef
	t.Defensive = int32(troopscfg.OriginDef) + t.Level*growBlood
	// t.Leader = int32(troopscfg.Leader) + t.Level*0             // 统帅	leader
	// t.Strength = int32(troopscfg.Strength) + t.Level*0         // 勇气	strength
	// t.Politics = int32(troopscfg.Politics) + t.Level*0         //政治
	// t.Intelligence = int32(troopscfg.Intelligence) + t.Level*0 //智力

	t.CalculationBuff()
}

//CalculationBuff 计算BUFF 加成属性
func (t *TroopsStruct) CalculationBuff() {
	for _, bufid := range t.BuffID {
		bufcfg := cfg.GetBuffCfg(bufid)

		if bufcfg.EffectType == 1 && bufcfg.EffectSubtypes1 == 1 && bufcfg.Key1 == 5 {
			t.Control = int32(float64(t.Control) * bufcfg.Value1)
		}

	}
}

//CleanBuf 清除BUF
func (t *TroopsStruct) CleanBuf(bufid int32) {
	if bufid == 0 {
		t.BuffID = make([]int32, 0)
		return
	}
	t.BuffID = tool.DelList(t.BuffID, bufid)
}

//SkillCD 部队 是否可施放战术战法   0 都可选 1 战术可选 2 都不可选
func (t *TroopsStruct) SkillCD() int32 {
	if !t.RoundWins {

		if t.SkillID == 0 {
			return 1
		}

		skillcfg := cfg.GetSkillCfg(t.SkillID)
		if t.SkillUseNumber < skillcfg.UseTime && skillcfg.Skilltype == 1 {
			return 0
		}
		return 1
	}
	return 2
}

//AddExp add exp
func (t *TroopsStruct) AddExp(addexp int64) (isUp bool) {
	oldLelve := t.Level
	newLevel, newExp := cfg.AddTroopsExp(int64(t.Level), int64(t.Exp), addexp)
	t.Level = int32(newLevel)
	t.Exp = newExp

	//升级后属性计算
	if t.Level > oldLelve {
		t.CalculationAttribute()
		return true
	}

	return false
}

//AddBuff 添加BUF
func (t *TroopsStruct) AddBuff(bufid int32) {
	if t.BuffID == nil {
		t.BuffID = []int32{bufid}
	} else {
		t.BuffID = append(t.BuffID, bufid)
	}
}

//ConvertTroopsProto 格式转换
func (t *TroopsStruct) ConvertTroopsProto() *troops.P_Troops {
	return &troops.P_Troops{
		TroopsID:    t.TroopsID,
		Country:     t.Country,
		AreasList:   t.AreasList,
		AreasIndex:  t.AreasIndex,
		Name:        t.Name,
		RowHP:       t.RowHP,
		QueueNum:    t.QueueNum,
		FightType:   t.FightType,
		FightState:  t.FitghtState,
		StageNumber: t.StageNumber,

		SkillID:   t.SkillID,
		TacticsID: t.TalentID,

		State:     common.TroopsState(t.State),
		Type:      t.Type,
		MaxNumber: t.MaxNumber,
		Number:    t.Number,
		Level:     t.Level,
		Roleid:    t.Roleid,

		SelectTactics: t.SelectTactics,
		Bufflist:      t.BuffID,
	}
}

package bigmapmanage

import (
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"slgserver/gserver/fightmod"
	"slgserver/gserver/process"
	"slgserver/msgproto/bigmap"
	"slgserver/msgproto/fight"

	log "github.com/sirupsen/logrus"
)

//区域内战斗
func (areas *AreasInfo) areasFighting(unix int64) {
	//战斗计算
	var attackTroops, defenseTroops *commonstruct.TroopsStruct   //第一位置队伍
	var attackTroops2, defenseTroops2 *commonstruct.TroopsStruct //第二上阵队伍（可能不存在）

	if aTroops, aTroops2, dTroops, dTroops2 := areas.getQueueTroops(); aTroops != nil && dTroops != nil {
		attackTroops = aTroops
		defenseTroops = dTroops

		attackTroops2 = aTroops2
		defenseTroops2 = dTroops2

	} else {
		log.Warn("未找到队列部队数据")
	}

	attackTroops.FightSet.AutoSelect = fightSetAuto[attackTroops.Roleid]
	defenseTroops.FightSet.AutoSelect = fightSetAuto[defenseTroops.Roleid]

	//启动准备阶段
	if areas.TimeStamp == 0 {
		areas.refreshbaseWaitTime(nil, attackTroops, defenseTroops)
		areas.TimeStamp = unix
		attackTroops.RoundWins = false
		defenseTroops.RoundWins = false
		return
	}

	skillselecttime := areas.refreshFightTimeStamp(attackTroops, defenseTroops)

	//===========================================
	//未到碰撞时间
	if unix < areas.TimeStamp+areas.BaseStamp+skillselecttime {
		return
	}
	areas.TimeStamp = unix

	log.Infof("now=> 战斗设置 [%v %v %v %v]  [%v %v %v %v]", attackTroops.TroopsID, attackTroops.FightSet.AutoSelect, defenseTroops.FightSet.SkillID, attackTroops.FightSet.TacticsID,
		defenseTroops.TroopsID, defenseTroops.FightSet.AutoSelect, defenseTroops.FightSet.SkillID, defenseTroops.FightSet.TacticsID)

	var recordlist []*fight.P_FightRecord

	//本回合战斗计算
	if attackTroops.RoundWins {
		recordlist = fightmod.FightCalculation(defenseTroops, attackTroops, attackTroops2, defenseTroops.FightSet, attackTroops.FightSet)

		var tmp []*fight.P_FightItem
		for _, v := range recordlist {
			tmp = v.Attack
			v.Attack = v.Defense
			v.Defense = tmp
		}
	} else {
		recordlist = fightmod.FightCalculation(attackTroops, defenseTroops, defenseTroops2, attackTroops.FightSet, defenseTroops.FightSet)
	}

	//重置队伍技能选择
	attackTroops.FightSet.SkillID = 0
	if !attackTroops.RoundWins {
		attackTroops.FightSet.TacticsID = 0
	}

	defenseTroops.FightSet.SkillID = 0
	if !defenseTroops.RoundWins {
		defenseTroops.FightSet.TacticsID = 0
	}

	//发送战报
	pushFightRecord := &fight.S2C_FightRecordPush{
		Record:           recordlist,
		FightStamp:       areas.TimeStamp + areas.BaseStamp + areas.refreshFightTimeStamp(attackTroops, defenseTroops),
		AttackTacticsID:  attackTroops.SelectTactics,
		DefenseTacticsID: defenseTroops.SelectTactics,
	}
	areas.pushFightRecord(pushFightRecord)
	areas.refreshbaseWaitTime(pushFightRecord, attackTroops, defenseTroops)

	//战斗效果处理
	for _, record := range recordlist {
		//50%几率吓退敌方后50个部队
		if record.Type == 1 && record.Value == 202 {
			log.Info("技能吓退敌方后50个部队")
			beside := int32(cfg.GetBeside(areas.AreasIndex)[0])
			tmpare := getAreasInfo(beside)
			log.Info("退出地区1：", tmpare.DefenseQueue, tmpare.AttackQueue)

			if record.TroopsID == attackTroops.TroopsID {
				areas.DefenseQueue = skillExit(areas.DefenseQueue, beside)
				areas.DefenseQueueNum = int32(len(areas.DefenseQueue))
			} else if record.TroopsID == defenseTroops.TroopsID {
				areas.AttackQueue = skillExit(areas.AttackQueue, beside)
				areas.AttackQueueNum = int32(len(areas.AttackQueue))
			}
			tmpare = getAreasInfo(beside)
			log.Info("退出地区2：", tmpare.DefenseQueue, tmpare.AttackQueue)

		}
	}

	//=================== 部队 角色加经验 ===========
	troopslist := map[int32]*commonstruct.TroopsStruct{
		attackTroops.TroopsID:  attackTroops,
		defenseTroops.TroopsID: defenseTroops,
	}

	if attackTroops2 != nil {
		troopslist[attackTroops2.TroopsID] = attackTroops2
	}

	if defenseTroops2 != nil {
		troopslist[defenseTroops2.TroopsID] = defenseTroops2
	}

	troopsExplist := make(map[int32]int32)
	for _, fightrecord := range recordlist {
		for _, fitem := range fightrecord.Attack {
			troopsExplist[fitem.TroopsID] = troopsExplist[fitem.TroopsID] + fitem.Loss
		}
		for _, fitem := range fightrecord.Defense {
			troopsExplist[fitem.TroopsID] = troopsExplist[fitem.TroopsID] + fitem.Loss*cfg.GetGlobalInt("expTroops")
		}
	}
	//经验列表
	for k, v := range troopsExplist {
		if troops, ok := troopslist[k]; ok && troops.Attribute != 3 {
			addexp := int64(v) * cfg.GetGlobalInt64("expTroops")
			attackTroops.AddExp(addexp)
			//部队获取经验通知
			process.SendMsg(attackTroops.Roleid, commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgAddExp,
				Data: commonstruct.AddExpItem{Type: 1, Key: k, AddExp: addexp, NewLevel: int64(attackTroops.Level),
					NewExp: attackTroops.Exp, LostNum: v}})

			//角色获取经验通知
			process.SendMsg(attackTroops.Roleid, commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgAddExp,
				Data: commonstruct.AddExpItem{Type: 0, Key: k, AddExp: int64(v) * cfg.GetGlobalInt64("expRole")}})

		}
	}

	//==============================================

	//胜负结果--输掉队伍退出区域
	if _, v := getHPfun(attackTroops.RowHP); v == 0 {
		areas.AttackQueue = areas.AttackQueue[1:]
		areas.troopsOutBigmap(attackTroops)
	} else {
		saveMapTroops(attackTroops)
	}

	if _, v := getHPfun(defenseTroops.RowHP); v == 0 {
		areas.DefenseQueue = areas.DefenseQueue[1:]
		areas.troopsOutBigmap(defenseTroops)
	} else {
		saveMapTroops(defenseTroops)
	}

	//第二位置
	if attackTroops2 != nil {
		saveMapTroops(attackTroops2)
	}

	if defenseTroops2 != nil {
		saveMapTroops(defenseTroops2)
	}

	//战斗结束 区域状态变化
	//队列 双方剩余队伍数量
	attackNum := len(areas.AttackQueue)
	defenseNum := len(areas.DefenseQueue)
	if attackNum == 0 && defenseNum == 0 {
		log.Info("打平") //打平 区域归属归属不变
		areas.TimeStamp = 0
		areas.State = 0
		areas.fightOver(1)
		sendAllRolesAreasStateChange(areas.AreasIndex, areas.Occupy, areas.State)
	} else if attackNum == 0 {
		log.Info("守方胜")
		//守方胜
		areas.TimeStamp = 0
		areas.Occupy = defenseTroops.Country
		areas.State = 0
		areas.fightOver(1)
		sendAllRolesAreasStateChange(areas.AreasIndex, areas.Occupy, areas.State)
	} else if defenseNum == 0 {
		log.Info("攻方胜")
		//攻方胜
		areas.TimeStamp = 0
		areas.Occupy = attackTroops.Country
		areas.State = 0
		areas.fightOver(2)
		sendAllRolesAreasStateChange(areas.AreasIndex, areas.Occupy, areas.State)
		//队伍两个国家拆分继续战斗
		areas.queueSplit()
	}
	// else {
	// 	// if attackNum > (areas.pushNum - 1) {
	// 	// 	pushFightRecord.AttackTroops = commonstruct.ConvertTroopsProto(areas.attackQueue[areas.pushNum-1])
	// 	// }
	// 	// if defenseNum > (areas.pushNum - 1) {
	// 	// 	pushFightRecord.DefenseTroops = commonstruct.ConvertTroopsProto(areas.defenseQueue[areas.pushNum-1])
	// 	// }
	// }

	// log.Info("=====>", attackTroops.TroopsID, pushFightRecord.AttackTacticsID, attackTroops.RoundWins, attackTroops.RowHP)
	// log.Info("=====>", defenseTroops.TroopsID, pushFightRecord.DefenseTacticsID, defenseTroops.RoundWins, defenseTroops.RowHP)
	// for _, v := range recordlist {
	// 	log.Info("=====>", v)
	// }

	areas.pushSkillSelect()
	//刷新队列 下一位上阵
	//log.Info("每轮战斗后 区域状态：", areas.State)
	areas.refreshTop2TroopsFightState()
}

//两支战斗部队 下回合技能施放情况
func (areas *AreasInfo) pushSkillSelect() {
	if areas.State == 0 {
		return
	}
	attackTroops, _, defenseTroops, _ := areas.getQueueTroops()

	if attackTroops != nil && attackTroops.Attribute == 0 && areas.isSubscribe(attackTroops.Roleid) {
		//推送给下轮战斗双方 技能施放选择情况
		process.SendSocketMsg(attackTroops.Roleid,
			int32(fight.MSG_FIGHT_Module_FIGHT),
			int32(fight.MSG_FIGHT_S2C_Select_Tactics), &fight.S2C_SelectTactics{TroopsID: attackTroops.TroopsID, SelectSkill: attackTroops.SkillCD()})
	}

	if defenseTroops != nil && defenseTroops.Attribute == 0 && areas.isSubscribe(defenseTroops.Roleid) {
		process.SendSocketMsg(defenseTroops.Roleid,
			int32(fight.MSG_FIGHT_Module_FIGHT),
			int32(fight.MSG_FIGHT_S2C_Select_Tactics), &fight.S2C_SelectTactics{TroopsID: defenseTroops.TroopsID, SelectSkill: defenseTroops.SkillCD()})

	}

}

func (areas *AreasInfo) getQueueTroops() (attackTroops1, attackTroops2, defenseTroops1, defenseTroops2 *commonstruct.TroopsStruct) {
	var aTroops, aTroops2, dTroops, dTroops2 *commonstruct.TroopsStruct
	if troops, ok := GetMapTroopsInfo(areas.AttackQueue[0]); ok {
		aTroops = &troops
	} else {
		log.Info("未找到：", areas.AttackQueue[0])
		//数据修正
		//1.清除出队列
		//2.队列为空时结束掉区域战斗状态
	}

	if len(areas.AttackQueue) > 1 {
		if troops, ok := GetMapTroopsInfo(areas.AttackQueue[1]); ok {
			aTroops2 = &troops
		}
	}

	if troops, ok := GetMapTroopsInfo(areas.DefenseQueue[0]); ok {
		dTroops = &troops
	} else {
		log.Info("未找到：", areas.DefenseQueue[0])
		//数据修正
		//1.清除出队列
		//2.队列为空时结束掉区域战斗状态
	}

	if len(areas.DefenseQueue) > 1 {
		if troops, ok := GetMapTroopsInfo(areas.DefenseQueue[1]); ok {
			dTroops2 = &troops
		}
	}

	return aTroops, aTroops2, dTroops, dTroops2
}

func (areas *AreasInfo) fightOver(victory int32) {
	for _, v := range areas.MsgSubscribe {
		log.Infof("发送战斗结果：消息订阅:[%v]  区域ID，胜方：[%v %v] ", v, areas.AreasIndex, victory)
		process.SendSocketMsg(v,
			int32(bigmap.MSG_BIGMAP_Module_BIGMAP),
			int32(bigmap.MSG_BIGMAP_S2C_OverFight), &bigmap.S2C_OverFight{AreasIndex: areas.AreasIndex, Victory: victory})
	}
	areas.MsgSubscribe = make(map[int32]int32)
}

func getHPfun(list []int32) (int32, int32) {
	for k, v := range list {
		if v > 0 {
			return int32(k), v
		}
	}
	return 0, 0
}

//发送战斗信息
func (areas *AreasInfo) pushFightRecord(fightrecord *fight.S2C_FightRecordPush) {
	for _, v := range areas.MsgSubscribe {
		process.SendSocketMsg(v,
			int32(fight.MSG_FIGHT_Module_FIGHT),
			int32(fight.MSG_FIGHT_S2C_FightRecordPush), fightrecord)
	}
}

//队伍战败退出大地图
func (areas *AreasInfo) troopsOutBigmap(troops *commonstruct.TroopsStruct) {
	//清除出大地图
	cleanBigmapTroops(troops)

	//npc 不通知
	if troops.Attribute == 3 {
		return
	}

	//部队离开队列 发送消息给订阅用户
	for _, v := range areas.MsgSubscribe {
		process.SendSocketMsg(v,
			int32(bigmap.MSG_BIGMAP_Module_BIGMAP),
			int32(bigmap.MSG_BIGMAP_S2C_LeaveQueue), &bigmap.S2C_LeaveQueue{AreasIndex: areas.AreasIndex, TroopsID: troops.TroopsID})
	}
	log.Debugf("战败队伍退出大地图：[%v] [%v]", troops.Roleid, troops.TroopsID)

	//1.通知客户端 更新部队信息
	ok := process.SendMsg(troops.Roleid, commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgUpdateTroopsInfo, Data: *troops})
	if !ok {
		log.Info("玩家没在线:", troops)
		//2.玩家没在线则直接修改数据库里数据
	}
}

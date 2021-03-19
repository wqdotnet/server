package fightmod

import (
	"math/rand"
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"slgserver/msgproto/fight"
	"slgserver/tool"
	"time"

	log "github.com/sirupsen/logrus"
)

//FightCalculation 战斗计算
//attackTroops, defenseTroops 攻守双方队列上阵队伍
//attackset, defenseset 攻守双方战斗选择配置
func FightCalculation(attackTroops, defenseTroops, defenseTroops2 *commonstruct.TroopsStruct, attackset, defenseset *commonstruct.FightSetting) (fightRecord []*fight.P_FightRecord) {
	//attackset, defenseset 施计战术战法处理
	skillSelect(attackTroops, defenseTroops, attackset, defenseset)

	recordlist := []*fight.P_FightRecord{}
	//攻方主动技能 或 触发被动技能
	if attackset.SkillID != 0 {
		record, fightover := skillCalculation(attackTroops, defenseTroops, defenseTroops2, attackset.SkillID, defenseset.SkillID)
		recordlist = append(recordlist, record...)
		if fightover {
			return recordlist
		}
	}

	if defenseset.SkillID != 0 {
		record, fightover := skillCalculation(defenseTroops, attackTroops, defenseTroops2, defenseset.SkillID, attackset.SkillID)
		recordlist = append(recordlist, record...)
		if fightover {
			return recordlist
		}
	}

	if attackset.SkillID == 0 && defenseset.SkillID == 0 {
		//无技能  则战术触发
		record, fightover := tacticsCalculation(attackTroops, defenseTroops, attackset.TacticsID, defenseset.TacticsID)
		recordlist = append(recordlist, record)
		if fightover {
			return recordlist
		}
	}

	//平砍伤害
	for _, v := range attacklogic(attackTroops, defenseTroops) {
		recordlist = append(recordlist, v)
	}

	return recordlist
}

//施放技能 和 战术 选择判断
func skillSelect(attackTroops, defenseTroops *commonstruct.TroopsStruct, attackset, defenseset *commonstruct.FightSetting) {
	if attackset.TacticsID == 0 {
		attackset.TacticsID = attackTroops.TacticsID[0]
	}

	if defenseset.TacticsID == 0 {
		defenseset.TacticsID = defenseTroops.TacticsID[0]
	}

	//防守方不可放技能
	defenseset.SkillID = 0

	//自动施放  自动选择一个技能->
	//手动 1.有设置  放技能或者战法
	//     2.没设置 放技能->技能如果CD->放战法
	//
	rand.Seed(time.Now().Unix())
	selectnum := rand.Intn(2)

	dskillCfg := cfg.GetSkillCfg(defenseTroops.SkillID)
	var skillCD bool
	if dskillCfg == nil {
		log.Warn("没技能：", defenseTroops.SkillID, defenseTroops)
	} else {
		//技能是否有CD
		skillCD = dskillCfg.UseTime > defenseTroops.SkillUseNumber
	}

	if !skillCD {
		defenseset.SkillID = 0
	}

	//第一次对碰 守方可放技能
	if !attackTroops.RoundWins && !defenseTroops.RoundWins {
		if defenseset.AutoSelect {
			//主动技能
			if dskillCfg.Skilltype == 1 && skillCD {
				defenseset.SkillID = defenseTroops.SkillID
				defenseset.TacticsID = 0
			} else if len(defenseTroops.TacticsID) > 0 {
				defenseset.TacticsID = defenseTroops.TacticsID[0]
			}
		} else if defenseset.TacticsID == 0 && len(defenseTroops.TacticsID) > 0 {
			defenseset.TacticsID = defenseTroops.TacticsID[selectnum] //next .. 随机选一个战术
		}
	}

	askillCfg := cfg.GetSkillCfg(attackTroops.SkillID)

	//技能是否有CD
	skillCD = askillCfg.UseTime > attackTroops.SkillUseNumber
	if !skillCD {
		attackset.SkillID = 0
	}

	if attackset.AutoSelect {
		//主动技能
		if askillCfg.Skilltype == 1 && skillCD {
			attackset.SkillID = attackTroops.SkillID
			attackset.TacticsID = 0
		} else if len(attackTroops.TacticsID) > 0 {
			attackset.TacticsID = attackTroops.TacticsID[selectnum] //next ..选克制防守方战术
		}
	} else if attackset.TacticsID == 0 && attackset.SkillID == 0 {
		attackset.TacticsID = attackTroops.TacticsID[selectnum] //next .. 选一个克制防守方的战术
	}
}

//战术
func tacticsCalculation(attackTroops, defenseTroops *commonstruct.TroopsStruct, attacktacticsID, defensetacticsID int32) (Record *fight.P_FightRecord, fightover bool) {
	//log.Info("战术 tacticsCalculation")

	attackTroops.SelectTactics = attacktacticsID
	defenseTroops.SelectTactics = defensetacticsID

	record := &fight.P_FightRecord{TroopsID: attackTroops.TroopsID, Type: 2, Value: attacktacticsID}

	rowkey, DHP := getHPfun(defenseTroops.RowHP)
	hp := tacticsHurt(attackTroops, defenseTroops, attacktacticsID, defensetacticsID)
	if DHP-hp <= 0 {
		hp = defenseTroops.RowHP[rowkey]
		defenseTroops.RowHP[rowkey] = 0
	} else {
		defenseTroops.RowHP[rowkey] = DHP - hp
	}
	record.Defense = append(record.Defense, &fight.P_FightItem{TroopsID: defenseTroops.TroopsID, Rolw: rowkey, Loss: hp, Dead: defenseTroops.RowHP[rowkey] <= 0})

	//攻方受击
	hp2 := tacticsHurt(defenseTroops, attackTroops, defensetacticsID, attacktacticsID)
	rowkey2, DHP2 := getHPfun(attackTroops.RowHP)
	if DHP2-hp2 <= 0 {
		hp2 = attackTroops.RowHP[rowkey2]
		attackTroops.RowHP[rowkey2] = 0
	} else {
		attackTroops.RowHP[rowkey2] = DHP2 - hp2
	}
	record.Attack = append(record.Attack, &fight.P_FightItem{TroopsID: attackTroops.TroopsID, Rolw: rowkey2, Loss: hp2, Dead: attackTroops.RowHP[rowkey2] <= 0})

	return record, defenseTroops.RowHP[rowkey] <= 0 || attackTroops.RowHP[rowkey2] <= 0
}

func tacticsHurt(attackTroops, defenseTroops *commonstruct.TroopsStruct, attacktacticsID, defensetacticsID int32) int32 {
	_, attackHp := getHPfun(attackTroops.RowHP)
	attacktactics := cfg.GetSkillTacticsCfg(attacktacticsID)
	defensetactics := cfg.GetSkillTacticsCfg(defensetacticsID)

	atype := attacktactics.Type
	dtype := defensetactics.Type
	//克制关系
	//2-1-3
	if atype == dtype {
		if attacktactics.Lv > defensetactics.Lv {
			return int32(attacktactics.BaseDamage + int(0.15*float64(attackHp)))
		}
		return int32(attacktactics.BaseDamage + int(0.05*float64(attackHp)))

	} else if atype == 2 && dtype == 1 || atype == 1 && dtype == 3 || atype == 3 && dtype == 2 {
		//克制
		return int32(attacktactics.BaseDamage + int(0.3*float64(attackHp)))
	} else {
		//被克制
		return int32(0.02 * float64(attackHp))
	}

}

//普通攻击战斗计算
func attackHurt(atthp int32, attackTroops, defenseTroops *commonstruct.TroopsStruct) int32 {
	//不破防
	if float64(attackTroops.Attack) < float64(defenseTroops.Defensive)*1.1 {
		return int32(1 + attackTroops.AttackSuper)
	}

	value := (float64(attackTroops.Attack-defenseTroops.Defensive))*(float64(atthp/(2*attackTroops.MaxNumber))+0.5) + float64(attackTroops.AttackSuper) - float64(defenseTroops.DefensiveSuper)

	//&bigmap.P_FightItem{TroopsID: defenseTroops.TroopsID, Rolw: defenseKey, Loss: ghp, Dead: defenseHp <= 0}
	return int32(tool.Round(value))
}

//普通攻击战斗计算
func attacklogic(attackTroops, defenseTroops *commonstruct.TroopsStruct) (fightRecord []*fight.P_FightRecord) {
	//log.Info("平砍 attackCalculation2")
	attackKey, attackHp := getHPfun(attackTroops.RowHP)
	defenseKey, defenseHp := getHPfun(defenseTroops.RowHP)
	recordlist := []*fight.P_FightRecord{}

	var attackOneKill, defeseOnekill bool
	num := 0

	//平砍伤害
	for {
		record := &fight.P_FightRecord{
			Attack:  []*fight.P_FightItem{},
			Defense: []*fight.P_FightItem{},
		}

		//7必杀
		num++
		if num >= 7 {
			log.Infof("7杀: %v  %v ", attackTroops.Number, defenseTroops.Number)
			if attackTroops.Number > defenseTroops.Number {
				attackOneKill = true
			} else {
				defeseOnekill = true
			}
		}

		olddefenseHp := defenseHp
		//攻击方
		ghp := int32(attackHurt(attackHp, attackTroops, defenseTroops))
		defenseHp = defenseHp - ghp
		if defenseHp <= 0 || attackOneKill {
			if defenseHp > 0 {
				ghp = defenseHp
			}
			defenseHp = 0
		}
		defenseTroops.Number -= ghp
		defenseTroops.RowHP[defenseKey] = defenseHp
		record.Defense = append(record.Defense, &fight.P_FightItem{TroopsID: defenseTroops.TroopsID, Rolw: defenseKey, Loss: ghp, Dead: defenseHp <= 0})

		//防守方
		ahp := int32(attackHurt(olddefenseHp, defenseTroops, attackTroops))
		attackHp = attackHp - ahp
		if attackHp <= 0 || defeseOnekill {
			if attackHp > 0 {
				ahp = attackHp
			}
			attackHp = 0
		}
		attackTroops.Number -= ahp
		attackTroops.RowHP[attackKey] = attackHp

		record.Attack = append(record.Attack, &fight.P_FightItem{TroopsID: attackTroops.TroopsID, Rolw: attackKey, Loss: ahp, Dead: attackHp <= 0})
		recordlist = append(recordlist, record)

		if attackHp == 0 || defenseHp == 0 {
			attackTroops.RoundWins = attackHp != 0
			defenseTroops.RoundWins = defenseHp != 0
			break
		}
	}

	//log.Debug("平砍伤害：", recordlist)
	return recordlist
}

func getHPfun(list []int32) (int32, int32) {
	for k, v := range list {
		if v > 0 {
			return int32(k), v
		}
	}
	return 0, 0
}

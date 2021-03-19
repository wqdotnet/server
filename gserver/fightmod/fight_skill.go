package fightmod

import (
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"slgserver/msgproto/fight"
	"slgserver/tool"
)

//战法(技能)
func skillCalculation(attackTroops, defenseTroops, defenseTroops2 *commonstruct.TroopsStruct, attackSkillID, defenseSkillID int32) (Record []*fight.P_FightRecord, fightover bool) {
	//log.Info("skillCalculation 战法(技能) :", attackSkillID, defenseSkillID)
	attackTroops.SkillUseNumber++
	recordlist := make([]*fight.P_FightRecord, 0)

	//双方施法
	askillCfg := cfg.GetSkillCfg(attackSkillID)
	var isover bool = false

	switch askillCfg.EffectType {
	case 1:
		if record, over := skillEffectDirect(askillCfg, attackTroops, defenseTroops, defenseTroops2); record != nil {
			isover = over
			recordlist = append(recordlist, record)
		}
	}

	for _, addskillid := range askillCfg.AddSkillsID {
		list, over := skillCalculation(attackTroops, defenseTroops, defenseTroops2, addskillid, 0)
		for _, v := range list {
			recordlist = append(recordlist, v)
		}
		isover = isover || over
	}

	return recordlist, isover
}

//技能效果  1 直接效果
func skillEffectDirect(askillCfg *cfg.SkillCfg, attackTroops, defenseTroops, defenseTroops2 *commonstruct.TroopsStruct) (*fight.P_FightRecord, bool) {
	record := &fight.P_FightRecord{TroopsID: attackTroops.TroopsID, Type: 1, Value: askillCfg.ID}
	rowkey, _ := getHPfun(defenseTroops.RowHP)

	switch askillCfg.EffectSubtypes1 {
	case 1:
		var num int32 = 0
		//askillCfg.Range 打几排
		for k, v := range defenseTroops.RowHP {
			if v == 0 {
				continue
			}
			hp := int32(skillHurt(askillCfg, attackTroops, defenseTroops))

			if v-hp <= 0 {
				hp = defenseTroops.RowHP[k]
				defenseTroops.RowHP[k] = 0
			} else {
				defenseTroops.RowHP[k] = v - hp
			}

			record.Defense = append(record.Defense, &fight.P_FightItem{TroopsID: defenseTroops.TroopsID, Rolw: int32(k), Loss: hp, Dead: defenseTroops.RowHP[k] <= 0})
			num++
			//打几排
			if k+1 >= int(askillCfg.Range) {
				break
			}
		}

		//需要打 askillCfg.Range 排，实际只打了 num 排
		// 向第二个位置的队伍攻击
		if num < askillCfg.Range && defenseTroops2 != nil {
			for k, v := range defenseTroops2.RowHP {
				if v == 0 {
					continue
				}
				hp := int32(skillHurt(askillCfg, attackTroops, defenseTroops2))

				if v-hp <= 0 {
					hp = defenseTroops2.RowHP[k]
					defenseTroops2.RowHP[k] = 0
				} else {
					defenseTroops2.RowHP[k] = v - hp
				}

				record.Defense = append(record.Defense, &fight.P_FightItem{TroopsID: defenseTroops2.TroopsID, Rolw: int32(k), Loss: hp, Dead: defenseTroops2.RowHP[k] <= 0})
				//打几排
				if k+1 >= int(askillCfg.Range-num) {
					break
				}
			}
		}

	case 4:

		//肯定触发
		if tool.Random(askillCfg.Key1) {
			return nil, false
		}
	}

	//buff
	for _, buffid := range askillCfg.AddbuffID {
		skillEffectBuff(cfg.GetBuffCfg(buffid), attackTroops, defenseTroops)
		record.Defense = append(record.Defense, &fight.P_FightItem{Bufid: buffid})
	}

	//log.Info("技能施放结果：", record, HP, defenseTroops.RowHP[rowkey])
	return record, defenseTroops.RowHP[rowkey] == 0
}

//
func skillHurt(askillCfg *cfg.SkillCfg, attackTroops, defenseTroops *commonstruct.TroopsStruct) int32 {

	//战法破不破防
	pfattack := float64(attackTroops.Attack-defenseTroops.Defensive)*askillCfg.Value1*(1+0+0)+float64(attackTroops.Strong) < 1.1*float64(defenseTroops.Control)
	//普攻破不破防
	pfskill := float64(attackTroops.Attack) < float64(defenseTroops.Defensive)*1.1

	if pfattack && pfskill {
		value := float64(attackTroops.Attack-defenseTroops.Defensive)*askillCfg.Value1*(1+0+0) + float64(attackTroops.Strong-defenseTroops.Control)
		return int32(tool.Round(value))
	} else if !pfattack {
		value := 0.5 * float64(attackTroops.Strong-defenseTroops.Control) * float64(1+(attackTroops.Attack-defenseTroops.Defensive)) / 1.1 * float64(defenseTroops.Defensive)
		return int32(tool.Round(value))
	} else if !pfskill {
		value := 0.1 * float64(attackTroops.Attack-defenseTroops.Defensive) * askillCfg.Value1 * (1 + 0 + 0)
		return int32(tool.Round(value))
	}

	return 0
}

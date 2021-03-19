package clienconnect

import (
	"slgserver/gserver/bigmapmanage"
	"slgserver/msgproto/fight"

	log "github.com/sirupsen/logrus"
)

//module 用户登陆模块
func (c *Client) fightModule(method int32, buf []byte) {
	switch fight.MSG_FIGHT(method) {
	case fight.MSG_FIGHT_C2S_Auto_Select:
		autoSelect := &fight.C2S_AutoSelect{}
		if decode(autoSelect, buf) {
			c.autoSelect(autoSelect)
		}
	case fight.MSG_FIGHT_C2S_Select_Tactics:
		selectTactics := &fight.C2S_SelectTactics{}
		if decode(selectTactics, buf) {
			c.selectTactics(selectTactics)
		}
	default:
		log.Info("loginModule null methodID:", method)
	}
}

//设置自动选择战术
func (c *Client) autoSelect(autoTactics *fight.C2S_AutoSelect) {
	c.roleinfo.Settings.AutoSelectTactics = autoTactics.AutoSelect
	bigmapmanage.SendFightSetting(c.roleid, 0, 0, 0, autoTactics.AutoSelect)
	c.Send(int32(fight.MSG_FIGHT_Module_FIGHT),
		int32(fight.MSG_FIGHT_S2C_Auto_Select),
		&fight.S2C_AutoSelect{AutoSelect: autoTactics.AutoSelect})
}

//选择战术
func (c *Client) selectTactics(autoTactics *fight.C2S_SelectTactics) {
	bigmapmanage.SendFightSetting(c.roleid, autoTactics.TroopsID, autoTactics.SkillID, autoTactics.TacticsID, false)

	// c.Send(int32(fight.MSG_FIGHT_Module_FIGHT),
	// 	int32(fight.MSG_FIGHT_S2C_Select_Tactics),
	// 	&fight.S2C_SelectTactics{TacticsID: autoTactics.TacticsID, TroopsID: 0, Msg: ""})
}

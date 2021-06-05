package cfg

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func init() {
	InitViperConfig("../../config", "json")

	//viper.AddConfigPath("./config")
	//viper.SetConfigName("mapinfo")
	log.Info("err:", GameCfg.ErrorCode.CfgList)
	// log.Info("MapInfo :", len(GameCfg.MapInfo.Areas))
	// log.Infof("troops:%v", len(GameCfg.Troops.CfgList))
	// log.Infof("ErrorCode:%v", len(GameCfg.ErrorCode.CfgList))
	// log.Infof("role exp:%v", len(GameCfg.RoleExp.ExpList))
	// log.Infof("global: %v", GetGlobalInt("expRole"))
	// log.Infof("skill :%v", len(GameCfg.Skill.SkillList))
	// log.Infof("SkillLandform :%v", len(GameCfg.Skill.SkillLandform))
	// log.Infof("SkillBuff :%v", len(GameCfg.Skill.BuffList))
}

func TestMapCfg(t *testing.T) {

	//10000 50982 132201
	// level, exp := AddRoleExp(2, 60, 10000+50982+132201+50)
	// assert.Equal(t, int(level), 5)
	// assert.Equal(t, int(exp), 110)

	// assert.Equal(t, GetErrorCodeNumber("PARAMETER_EMPTY"), "参数不能为空")
	// assert.Equal(t, GetTroopsCfg(1).Name, "骑兵")

	// assert.Equal(t, CheckBigMapConfig(), true)
	//assert.Equal(t, AreasIsBeside(1319, 1316), false)

	// assert.Equal(t, AreasIsBeside(2427, 2725), true)
	// assert.Equal(t, AreasIsBeside(2427, 3966), false)
}
package cfg

import (
	"testing"

	"github.com/sirupsen/logrus"
	"go.uber.org/atomic"
)

func init() {
	InitViperConfig("../../config", "json")

	var atom atomic.Bool
	atom.Load()
	atom.Store(false)

	//viper.AddConfigPath("./config")
	//viper.SetConfigName("mapinfo")
	logrus.Info("err:", GetGameCfg().ErrorCode.CfgList)
	// logrus.Info("MapInfo :", len(GameCfg.MapInfo.Areas))
	// logrus.Infof("troops:%v", len(GameCfg.Troops.CfgList))
	// logrus.Infof("ErrorCode:%v", len(GameCfg.ErrorCode.CfgList))
	// logrus.Infof("role exp:%v", len(GameCfg.RoleExp.ExpList))
	// logrus.Infof("global: %v", GetGlobalInt("expRole"))
	// logrus.Infof("skill :%v", len(GameCfg.Skill.SkillList))
	// logrus.Infof("SkillLandform :%v", len(GameCfg.Skill.SkillLandform))
	// logrus.Infof("SkillBuff :%v", len(GameCfg.Skill.BuffList))
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

package cfg

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	InitViperConfig("../../config", "json")
	logrus.Info("ErrorCode:", GetGameCfg().ErrorCode.CfgList)
	logrus.Info("ExpXiufaInfo len :", len(GetGameCfg().ExpXiufaInfo))
	logrus.Info("item len :", len(GetGameCfg().ItemInfo))
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

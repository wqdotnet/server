package cfg

var (
	//GameCfg 全局配置
	GameCfg cfgCollection
)

type cfgCollection struct {
	//错误提示码
	ErrorCode struct {
		CfgList []ErrorCodeCfg
	}

	//全局配置
	Global struct {
		Initialize []GlobalCfg
	}

	//升级经验
	RoleExp struct {
		ExpList []ExpCfg
	}

	//大地图配置
	MapInfo MapCfgStruct
	//部队
	Troops struct {
		CfgList []TroopsCfg
	}

	Skill struct {
		SkillList     []SkillCfg
		SkillLandform []SillLandformCfg
		BuffList      []BuffCfg
	}
}

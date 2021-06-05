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
}

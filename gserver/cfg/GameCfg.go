package cfg

import (
	"go.uber.org/atomic"
)

var (
	//gameCfg 全局配置
	gameCfg atomic.Value
)

func GetGameCfg() *cfgCollection {
	cfg := gameCfg.Load().(*cfgCollection)
	return cfg
}

func saveCfg(cfg *cfgCollection) {
	gameCfg.Store(cfg)
}

type cfgCollection struct {

	//角色数据

	//错误提示码
	ErrorCode struct {
		CfgList []*ErrorCodeCfg
	}
}

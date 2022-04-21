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

	ExpXiufaInfo []*ExpXiufaInfo
}

type ExpXiufaInfo struct {
	Level            int     `json:"level"`
	NeedExp          int     `json:"needExp"`
	FailExp          int     `json:"failExp"`
	MaxExp           int     `json:"maxExp"`
	BigLevel         string  `json:"bigLevel"`
	HeadSize         int     `json:"headSize"`
	CycleEXP         int     `json:"cycleEXP"`
	ExpWeight        []int   `json:"expWeight"`
	ExpWeightValue   [][]int `json:"expWeightValue"`
	ExpUp            int     `json:"expUp"`
	PropertiesID     []int   `json:"propertiesId"`
	AttributeValues  []int   `json:"attributeValues"`
	Times            int     `json:"times"`
	Properties       []int   `json:"properties"`
	PropertiesWeight []int   `json:"propertiesWeight"`
	CycleProperties  []int   `json:"cycleProperties"`
	PropertiesMax    []int   `json:"propertiesMax"`
	TribulationID    int     `json:"tribulationId"`
	ShowNum          []int   `json:"showNum"`
	IDGroup          int     `json:"idGroup"`
	MailID           int     `json:"mailId"`
}

package cfg

import log "github.com/sirupsen/logrus"

//GlobalCfg 全局配置
type GlobalCfg struct {
	ID    int     `json:"ID"`
	Type  string  `json:"Type"`
	Intor string  `json:"Intor"`
	Param string  `json:"Param"`
	Value float32 `json:"Value"`
	Ps    string  `json:"ps"`
}

//GetGlobalInt int32
func GetGlobalInt(key string) int32 {
	for _, v := range GameCfg.Global.Initialize {
		if v.Param == key {
			return int32(v.Value)
		}
	}
	log.Warnf("config key [%v] is null", key)
	return 0
}

//GetGlobalInt64 int64
func GetGlobalInt64(key string) int64 {
	for _, v := range GameCfg.Global.Initialize {
		if v.Param == key {
			return int64(v.Value)
		}
	}
	log.Warnf("config key [%v] is null", key)
	return 0
}

//GetGlobalfloat float32
func GetGlobalfloat(key string) float32 {
	for _, v := range GameCfg.Global.Initialize {
		if v.Param == key {
			return v.Value
		}
	}
	log.Warnf("config key [%v] is null", key)
	return 0
}

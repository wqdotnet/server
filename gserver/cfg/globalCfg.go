package cfg

import (
	log "github.com/sirupsen/logrus"
)

var (
	//GlobalCfg 全局配置
	GlobalCfg global
)

type global struct {
	//大地图配置
	MapInfo mapinfo
}

//错误提示
var (
	ERROR_PARAMETER_EMPTY = "参数不能为空"
	ERROR_NOT_LOGIN       = "用户未登陆"
	ERROR_AccountNull     = "未找到账号 or 密码错误"
	ERROR_RoleNull        = "角色为空"
	ERROR_AccountExists   = "账号已存在"
	ERROR_RoleNameExists  = "角色名已存在"
)

//GetCountryAreasIndex 查找国家起启坐标
func GetCountryAreasIndex(countryid int32) int32 {
	for _, arecfg := range GlobalCfg.MapInfo.Areas {
		if arecfg.Type == int(countryid) {
			return int32(arecfg.Setindex)
		}
	}
	return 0
}

//AreasIsBeside 区域是否相连
func AreasIsBeside(areas, nextareas int32) bool {
	index := GlobalCfg.MapInfo.IndexCfg[areas]
	beside := GlobalCfg.MapInfo.Areas[index-1].Beside

	for _, v := range beside {
		if GlobalCfg.MapInfo.Areas[v-1].Setindex == int(nextareas) {
			return true
		}
	}

	return false
}

//CheckBigMapConfig 地图数据检查
func CheckBigMapConfig() bool {
	log.Info("CheckBigMapConfig Areas Number:", len(GlobalCfg.MapInfo.Areas))
	if &GlobalCfg.MapInfo == nil || len(GlobalCfg.MapInfo.Areas) == 0 {
		return false
	}
	areaslist := GlobalCfg.MapInfo.Areas

	for _, arecfg := range areaslist {
		index := GlobalCfg.MapInfo.IndexCfg[arecfg.Setindex]
		tmpareas := GlobalCfg.MapInfo.Areas[index-1]
		if tmpareas.Setindex != arecfg.Setindex {
			return false
		}

		log.Trace("checkout ", index, arecfg.Setindex, arecfg.Beside)

		for _, v := range arecfg.Beside {
			tmpbaside := GlobalCfg.MapInfo.Areas[v-1]

			log.Trace("		beside: ", v, tmpbaside.Setindex, tmpbaside.Beside)

			var tmpbool bool = false
			for _, k := range tmpbaside.Beside {
				if k == index {
					tmpbool = true
				}
			}

			if !tmpbool {
				return false
			}
		}

	}
	return true
}

package bigmapmanage

import (
	"server/gserver/cfg"

	log "github.com/sirupsen/logrus"
)

//-----------------------------地图区域-----------------------------

//AreasInfo 区域信息
type AreasInfo struct {
	AreasIndex int32 //区域ID
	Type       int32 //0 中立1-3:国家
	State      int32 //0 正常 1 战斗
	Occupy     int32 //占领信息  0 :无人占领  1-3国家KEY
}

//初始化大地图所有区域信息
func initBigmapAreasInfo() {
	//配置信息
	for _, ares := range cfg.GlobalCfg.MapInfo.Areas {
		index := int32(ares.Setindex)
		SetAreasInfo(index, AreasInfo{AreasIndex: index, Type: int32(ares.Type), State: 0, Occupy: int32(ares.Type)})
	}

	//缓存信息覆盖

}

//地图关闭时缓存所有区域占领信息
func storageArreasInfo() {
	AreasRange(func(areas AreasInfo) bool {

		return true
	})
}

//SetAreasInfo 设置区域
func SetAreasInfo(key int32, info AreasInfo) {
	if _, ok := areasSMap.Load(key); ok {
		log.Warnf("Areasindex key:[%v] is existx", key)
	} else {
		areasSMap.Store(key, info)
	}
}

//GetAreasInfo 获取区域信息
func GetAreasInfo(key int32) *AreasInfo {
	if info, ok := areasSMap.Load(key); ok {
		areas := info.(AreasInfo)
		return &areas
	}
	return nil
}

//GetAllAreas get all
func GetAllAreas() []AreasInfo {
	var areasList []AreasInfo
	areasSMap.Range(func(key, value interface{}) bool {
		areas := value.(AreasInfo)
		areasList = append(areasList, areas)
		return true
	})
	return areasList
}

//AreasRange func
func AreasRange(exefunc func(value AreasInfo) bool) {
	areasSMap.Range(func(key, value interface{}) bool {
		areas := value.(AreasInfo)
		return exefunc(areas)
	})
}

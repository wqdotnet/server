package bigmapmanage

//-----------------------------地图区域-----------------------------

//AreasInfo 区域信息
type AreasInfo struct {
	AreasIndex int32   //区域ID
	Type       int32   //0 中立1-3:国家
	State      int32   //0 正常 1 战斗
	Occupy     int32   //占领信息  0 :无人占领  1-3国家KEY
	troopsA    []int32 //部队A
	troopsB    []int32 //部队B
	troopsC    []int32 //部队C

	areasmsgchan chan string
}

//===========================================================================================
//saveAreasInfo save
func saveAreasInfo(info AreasInfo) {
	areasSMap.Store(info.AreasIndex, info)
}

//getAreasInfo 获取区域信息
func getAreasInfo(key int32) *AreasInfo {
	if info, ok := areasSMap.Load(key); ok {
		areas := info.(AreasInfo)
		return &areas
	}
	return nil
}

//getAllAreas get all
// func getAllAreas() []AreasInfo {
// 	var areasList []AreasInfo
// 	areasSMap.Range(func(key, value interface{}) bool {
// 		areas := value.(AreasInfo)
// 		areasList = append(areasList, areas)
// 		return true
// 	})
// 	return areasList
// }

//AreasRange func
func AreasRange(exefunc func(value AreasInfo) bool) {
	areasSMap.Range(func(key, value interface{}) bool {
		areas := value.(AreasInfo)
		return exefunc(areas)
	})
}

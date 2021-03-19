package cfg

import log "github.com/sirupsen/logrus"

//MapCfgStruct 地图配置结构
type MapCfgStruct struct {
	PointsAry  [][]int `json:"pointsAry"`
	ArrowAry1  [][]int `json:"arrowAry1"`
	ArrowAry2  [][]int `json:"arrowAry2"`
	MaxX       int     `json:"maxX"`
	IndexCfg   []int   `json:"indexCfg"`
	MinTileNum int     `json:"minTileNum"`
	MaxTileNum int     `json:"maxTileNum"`
	Mapsize    struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"mapsize"`
	Tilesize struct {
		Height int     `json:"height"`
		Width  float64 `json:"width"`
	} `json:"tilesize"`
	Movespeed int `json:"movespeed"`
	Metatype  int `json:"metatype"`
	MaxY      int `json:"maxY"`
	Areas     []struct {
		AreasIndex []int  `json:"areasIndex"`
		SegInfo    []int  `json:"segInfo"`
		Type       int    `json:"type"`
		Name       string `json:"name"`
		Setindex   int    `json:"setindex"`
		Beside     []int  `json:"beside"`
		SetList    []int  `json:"setList"`
	} `json:"areas"`
}

//GetCountryAreasIndex 查找国家起启坐标
func GetCountryAreasIndex(countryid int32) int32 {
	for _, arecfg := range GameCfg.MapInfo.Areas {
		if arecfg.Type == int(countryid) {
			return int32(arecfg.Setindex)
		}
	}
	return 0
}

//AreasIsBeside 区域是否相连
func AreasIsBeside(areas, nextareas int32) bool {
	if int(areas) > len(GameCfg.MapInfo.IndexCfg)-1 {
		log.Warnf("AreasIsBeside IndexCfg:  [%v]  [%v]", areas, len(GameCfg.MapInfo.IndexCfg))
		return false
	}

	index := GameCfg.MapInfo.IndexCfg[areas-1]
	if index-1 > len(GameCfg.MapInfo.Areas) {
		log.Warnf("AreasIsBeside Areas: [%v]  [%v]", index, len(GameCfg.MapInfo.Areas))
		return false
	}

	beside := GameCfg.MapInfo.Areas[index-1].Beside
	for _, v := range beside {
		if GameCfg.MapInfo.Areas[v-1].Setindex == int(nextareas) {
			return true
		}
	}

	log.Warnf("[%v ---  %v]  index:[%v]  beside:[%v] ", areas, nextareas, index, beside)
	return false
}

//GetBeside 获取相连
func GetBeside(areas int32) []int {
	index := GameCfg.MapInfo.IndexCfg[areas-1]
	var list = make([]int, 0)
	for _, v := range GameCfg.MapInfo.Areas[index-1].Beside {
		list = append(list, GameCfg.MapInfo.Areas[v-1].Setindex)

	}
	return list
}

//CheckBigMapConfig 地图数据检查
func CheckBigMapConfig() bool {
	log.Info("CheckBigMapConfig Areas Number:", len(GameCfg.MapInfo.Areas))
	if &GameCfg.MapInfo == nil || len(GameCfg.MapInfo.Areas) == 0 {
		return false
	}
	areaslist := GameCfg.MapInfo.Areas

	for _, arecfg := range areaslist {
		index := GameCfg.MapInfo.IndexCfg[arecfg.Setindex-1]
		tmpareas := GameCfg.MapInfo.Areas[index-1]

		log.Trace("checkout ", index, arecfg.Setindex, arecfg.Beside)

		if tmpareas.Setindex != arecfg.Setindex {
			log.Tracef("err:[%v] Setindex:[%v]   tmpareas:[%v]", index, arecfg.Setindex, tmpareas)
			return false
		}

		for _, v := range arecfg.Beside {
			tmpbaside := GameCfg.MapInfo.Areas[v-1]

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

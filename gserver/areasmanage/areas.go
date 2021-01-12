package areasmanage

// import (
// 	"encoding/json"
// 	"server/db"
// 	"server/gserver/cfg"

// 	log "github.com/sirupsen/logrus"
// )

// //StartAreasGoroutine 创建区域进程
// func StartAreasGoroutine(areasindex int) {
// 	go func() {

// 		// defer func() {
// 		// 	if err := recover(); err != nil {
// 		// 		log.Errorf("recover: %v", err)
// 		// 	}
// 		// }()

// 		for {
// 			select {}
// 		}

// 	}()
// }

// //初始化大地图所有区域信息
// func initBigmapAreasInfo() {
// 	//配置信息
// 	for _, ares := range cfg.GlobalCfg.MapInfo.Areas {
// 		index := int32(ares.Setindex)
// 		if _, ok := areasSMap.Load(index); ok {
// 			log.Warnf("Areasindex key:[%v] is existx", index)
// 		} else {
// 			areasSMap.Store(index, AreasInfo{AreasIndex: index, Type: int32(ares.Type), State: 0, Occupy: int32(ares.Type)})
// 		}
// 	}

// 	//缓存信息覆盖
// 	value, _ := db.HVALS("areasSMap")
// 	for _, v := range value {
// 		areas := &AreasInfo{}
// 		json.Unmarshal(v, areas)
// 		areasSMap.Store(areas.AreasIndex, *areas)

// 		log.Info("load areasSMap: ", areas)
// 	}

// 	db.RedisExec("del", "areasSMap")
// }

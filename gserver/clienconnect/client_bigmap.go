package clienconnect

import (
	"server/gserver/bigmapmanage"
	"server/msgproto/bigmap"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

// type rolemapinfo struct {
// 	areasIndex []int
// 	segInfo    []int
// }

func (c *Client) bigmapModule(method int32, buf []byte) {
	switch bigmap.MSG_BIGMAP_Module_BIGMAP {
	// case bigmap.MSG_BIGMAP_C2S_GetAreasInfo:
	// 	getareas := &bigmap.C2S_GetAreasInfo{}
	// 	if e := proto.Unmarshal(buf, getareas); e != nil {
	// 		log.Error(e)
	// 		return
	// 	}
	// 	c.getAreasInfo(getareas)
	// case bigmap.MSG_BIGMAP_C2S_GetAreasTroops:
	// 	gettroops := &bigmap.C2S_GetAreasTroops{}
	// 	if e := proto.Unmarshal(buf, gettroops); e != nil {
	// 		log.Error(e)
	// 		return
	// 	}
	// 	c.getAreasTroops(gettroops)
	case bigmap.MSG_BIGMAP_C2S_Move:
		move := &bigmap.C2S_Move{}
		if e := proto.Unmarshal(buf, move); e != nil {
			log.Error(e)
			return
		}
		c.move(move)
	case bigmap.MSG_BIGMAP_C2S_StopMoving:
		stopmove := &bigmap.C2S_StopMoving{}
		if e := proto.Unmarshal(buf, stopmove); e != nil {
			log.Error(e)
			return
		}
		c.stopMoving(stopmove)
	default:
		log.Info("loginModule null methodID:", method)
	}
}

//大地图区域信息
func (c *Client) sendAllArease() {
	areaslist := &bigmap.S2C_AreasInfo{}
	bigmapmanage.AreasRange(func(value bigmapmanage.AreasInfo) bool {

		//只发送被占领的，正在发生战斗的
		if value.Occupy > 0 || value.State > 0 {
			areaslist.AreasInfoList = append(areaslist.AreasInfoList,
				&bigmap.P_AreasInfo{AreasIndex: value.AreasIndex,
					Type:  value.Occupy, //占据国家
					State: value.State})
		}
		return true
	})

	log.Debug(areaslist)
	c.Send(int32(bigmap.MSG_BIGMAP_Module_BIGMAP), int32(bigmap.MSG_BIGMAP_S2C_AreasInfo), areaslist)
}

//移动
func (c *Client) move(move *bigmap.C2S_Move) {

	info, err := getTroopsinfo(c.roleid, move.TroopsID)
	if err != nil {
		log.Error(err)
	}

	//1. 数据验证->当前部队是否已经在大地图移动
	//next ....

	info.AreasList = move.AreasList
	info.State = 1
	info.ArrivalTime = int64(len(move.AreasList)*3) + time.Now().Unix()

	// var areasindex int32 = 0
	// for _, arecfg := range cfg.GlobalCfg.MapInfo.Areas {
	// 	if arecfg.Type == int(info.Country) {
	// 		areasindex = int32(arecfg.Setindex)
	// 	}
	// }

	//2. 路径验证 -> 判断是否合法
	//next ....

	//发送给客户端 启动坐标 状态更新 预计到达时间
	c.s2cMove(move.TroopsID, info.ArrivalTime)
	//队伍移动发送大地图处理
	bigmapmanage.SendTroopsMove(*info)

	updateTroopsInfo(info)
}

//发送部队移动信息至客户端
func (c *Client) s2cMove(troopsid int32, arrivaltime int64) {
	s2c := &bigmap.S2C_Move{}
	s2c.TroopsID = troopsid
	s2c.ArrivalTime = arrivaltime

	c.Send(int32(bigmap.MSG_BIGMAP_Module_BIGMAP), int32(bigmap.MSG_BIGMAP_S2C_Move), s2c)
}

//暂停移动
func (c *Client) stopMoving(stop *bigmap.C2S_StopMoving) {
	bigmapmanage.SendStopMove(stop.TroopsID)
}

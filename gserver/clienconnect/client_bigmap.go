package clienconnect

import (
	"server/gserver/bigmapmanage"
	"server/gserver/cfg"
	"server/gserver/commonstruct"
	"server/msgproto/bigmap"
	"server/msgproto/common"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

// type rolemapinfo struct {
// 	areasIndex []int
// 	segInfo    []int
// }

func (c *Client) bigmapModule(method int32, buf []byte) {
	switch bigmap.MSG_BIGMAP(method) {
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
		log.Info("bigmap null methodID:", method)
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

	c.Send(int32(bigmap.MSG_BIGMAP_Module_BIGMAP), int32(bigmap.MSG_BIGMAP_S2C_AreasInfo), areaslist)
}

//移动
func (c *Client) move(move *bigmap.C2S_Move) {
	//1. 数据从 c.troopslist 中读取
	//2. 判断部队状态 0 3 则为出发和 大地图驻扎出发,其它状态不可移动
	//3. 数据定时保存DB 统一处理
	if move.AreasList == nil || len(move.AreasList) == 0 {
		log.Warn("部队移动路径不能为空")
		return
	}
	var info *commonstruct.TroopsStruct

	info = c.troopslist[move.TroopsID]
	if info == nil {
		//无此部队
		log.Warnf("[%v] [%v] 无此部队  部队id:[%v]", c.account, c.rolename, move.TroopsID)
		return
	}

	//此支部队已在大地图 则从大地图中得到全新数据
	if bmtroops, ok := bigmapmanage.GetBigMapTroopsInfo(move.TroopsID); ok {
		info = &bmtroops
		//info.State == common.TroopsState_Move
		if info.State == common.TroopsState_fight {
			// 战斗中的无法接受命令
			log.Warn("大地图 移动 战斗中的无法接受命令 TroopsID:", move.TroopsID)
		}
	} else {
		//大地图中未找到该部队，部队位置启始为主城
		info.State = common.TroopsState_StandBy
		info.AreasIndex = cfg.GetCountryAreasIndex(info.Country)
	}

	info.AreasList = move.AreasList
	info.MoveNum = 0
	info.ArrivalTime = int64(len(move.AreasList)*3) + time.Now().Unix()

	//发送给客户端 启动坐标 状态更新 预计到达时间
	c.s2cMove(move.TroopsID, info.AreasIndex, int32(info.State), info.ArrivalTime)

	//队伍移动发送大地图处理
	info.State = common.TroopsState_Move
	bigmapmanage.SendTroopsMove(*info)
	//角色部队数据更新
	//updateTroopsInfo(info)
}

//发送部队移动信息至客户端
func (c *Client) s2cMove(troopsid, areasindex, state int32, arrivaltime int64) {
	s2c := &bigmap.S2C_Move{}
	s2c.TroopsID = troopsid
	s2c.ArrivalTime = arrivaltime
	s2c.AreasIndex = areasindex
	s2c.State = state
	c.Send(int32(bigmap.MSG_BIGMAP_Module_BIGMAP), int32(bigmap.MSG_BIGMAP_S2C_Move), s2c)
}

//暂停移动
func (c *Client) stopMoving(stop *bigmap.C2S_StopMoving) {
	bigmapmanage.SendStopMove(stop.TroopsID)
}

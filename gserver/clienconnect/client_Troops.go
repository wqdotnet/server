package clienconnect

import (
	"context"
	"slgserver/db"
	"slgserver/gserver/bigmapmanage"
	"slgserver/gserver/cfg"
	"slgserver/gserver/commonstruct"
	"slgserver/msgproto/common"
	"slgserver/msgproto/troops"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//module 用户登陆模块
func (c *Client) troopsModule(method int32, buf []byte) {
	switch troops.MSG_TROOPS(method) {
	case troops.MSG_TROOPS_C2S_Behavior:
		behavior := &troops.C2S_Behavior{}
		if decode(behavior, buf) {
			c.troopsbehavior(behavior)
		}

	default:
		log.Info("loginModule null methodID:", method)
	}
}

func (c *Client) troopsbehavior(behavior *troops.C2S_Behavior) {
	info := c.troopslist[behavior.TroopsID]
	if info == nil {
		c.Send(int32(troops.MSG_TROOPS_Module_TROOPS),
			int32(troops.MSG_TROOPS_S2C_Behavior),
			&troops.S2C_Behavior{Msg: cfg.GetErrorCodeNumber("TROOPS_TYPE_NULL")})
		return
	}

	switch behavior.BehaviorID {
	case troops.TroopsBehavior_AddExp:
		if len(behavior.ParValue) > 0 {
			c.troopsAddExp(behavior.TroopsID, behavior.ParValue[0])
		}
	case troops.TroopsBehavior_Recruit:
		if len(behavior.ParValue) > 0 {
			c.troopsRecruit(int32(behavior.ParValue[0]))
		}
	case troops.TroopsBehavior_OnStage:
		if len(behavior.ParValue) > 0 {
			c.troopsOnStage(behavior.TroopsID, int32(behavior.ParValue[0]))
		}
	case troops.TroopsBehavior_Exit:
		c.troopsExit(behavior.TroopsID)
	default:
		log.Warnf("unknown : %v", behavior.BehaviorID)
		return
	}

	c.Send(int32(troops.MSG_TROOPS_Module_TROOPS),
		int32(troops.MSG_TROOPS_S2C_UpdateTroopsInfo),
		&troops.S2C_UpdateTroopsInfo{TroopsInfo: info.ConvertTroopsProto()})
}

//部队添加经验测试接口
func (c *Client) troopsAddExp(troopsid int32, addExp int64) {
	troopsinfo := c.troopslist[troopsid]
	if troopsinfo != nil {
		troopsinfo.AddExp(addExp)
	}

	//更新至大地图
	bigmapmanage.SendUpdateTroopsInfo(troopsinfo)
	c.Send(int32(troops.MSG_TROOPS_Module_TROOPS),
		int32(troops.MSG_TROOPS_S2C_Behavior),
		&troops.S2C_Behavior{TroopsID: troopsid, BehaviorID: troops.TroopsBehavior_AddExp, Item: map[string]int64{"addExp": addExp}},
		//&troops.S2C_Behavior{Troopsid: info.TroopsID, AddExp: addexp.AddExp, Level: info.Level, Exp: info.Exp}
	)
}

//部队招募
func (c *Client) troopsRecruit(typeID int32) {
	//创建部队信息
	troopsinfo := commonstruct.NewTroops(c.rolename, db.GetAutoID(db.TroopsTable), 0, c.country, typeID)
	if troopsinfo == nil {
		c.Send(int32(troops.MSG_TROOPS_Module_TROOPS),
			int32(troops.MSG_TROOPS_S2C_Behavior),
			&troops.S2C_Behavior{BehaviorID: troops.TroopsBehavior_Recruit, Msg: cfg.GetErrorCodeNumber("TROOPS_TYPE_NULL")})
		return
	}

	troopsinfo.AreasIndex = cfg.GetCountryAreasIndex(c.country)
	troopsinfo.Roleid = c.roleid
	db.InsertOne(db.TroopsTable, troopsinfo)
	c.troopslist[troopsinfo.TroopsID] = troopsinfo

	c.Send(int32(troops.MSG_TROOPS_Module_TROOPS), int32(troops.MSG_TROOPS_S2C_Behavior),
		&troops.S2C_Behavior{
			TroopsID:   troopsinfo.TroopsID,
			BehaviorID: troops.TroopsBehavior_Recruit,
			Item:       map[string]int64{"TroopsID": int64(troopsinfo.TroopsID)},
		})
}

//上阵
func (c *Client) troopsOnStage(troopsid, stageNumber int32) {
	troopsinfo := c.troopslist[troopsid]

	sendError := func(strError string) {
		c.Send(int32(troops.MSG_TROOPS_Module_TROOPS), int32(troops.MSG_TROOPS_S2C_Behavior),
			&troops.S2C_Behavior{BehaviorID: troops.TroopsBehavior_OnStage, Msg: cfg.GetErrorCodeNumber(strError)})
	}

	//上阵位置错误
	if stageNumber == 0 || stageNumber > 5 {
		sendError("TROOPS_ONSTAGE_ERROR")
		return
	}

	//上阵位置错误
	for _, v := range c.troopslist {
		if v.StageNumber == stageNumber {
			sendError("TROOPS_ONSTAGE_ERROR")
			return
		}
	}

	//战斗状态
	if troopsinfo.State != common.TroopsState_StandBy {
		sendError("TROOPS_FIGHT")
		return
	}

	troopsinfo.StageNumber = stageNumber
	//更新至大地图
	bigmapmanage.SendUpdateTroopsInfo(troopsinfo)
	c.Send(int32(troops.MSG_TROOPS_Module_TROOPS), int32(troops.MSG_TROOPS_S2C_Behavior),
		&troops.S2C_Behavior{TroopsID: troopsinfo.TroopsID, BehaviorID: troops.TroopsBehavior_OnStage, Item: map[string]int64{"StageNumber": int64(stageNumber)}})
}

//下阵
func (c *Client) troopsExit(troopsid int32) {
	troopsinfo := c.troopslist[troopsid]
	sendError := func(strError string) {
		c.Send(int32(troops.MSG_TROOPS_Module_TROOPS), int32(troops.MSG_TROOPS_S2C_Behavior),
			&troops.S2C_Behavior{BehaviorID: troops.TroopsBehavior_Exit, Msg: cfg.GetErrorCodeNumber(strError)})
	}

	if troopsinfo.StageNumber == 0 {
		//未上阵
		sendError("TROOPS_ONSTAGE_ERROR")
		return
	}

	//战斗状态
	if troopsinfo.State == common.TroopsState_fight {
		sendError("TROOPS_FIGHT")
		return
	} else if troopsinfo.State != common.TroopsState_StandBy {
		//清除部队在大地图中状态
		troopsinfo.StageNumber = 0
		bigmapmanage.SendTroopsExitBigmap(&commonstruct.TroopsExitBigmap{TroopsID: troopsinfo.TroopsID, Type: 1})
	} else {
		troopsinfo.StageNumber = 0
		//更新至大地图
		bigmapmanage.SendUpdateTroopsInfo(troopsinfo)
	}

	c.Send(int32(troops.MSG_TROOPS_Module_TROOPS), int32(troops.MSG_TROOPS_S2C_Behavior),
		&troops.S2C_Behavior{TroopsID: troopsinfo.TroopsID, BehaviorID: troops.TroopsBehavior_Exit})
}

//======================================================================================================================
//获取部队信息
//后期整理 所有数据库相关操作移至db 模块
// func getTroopsinfo(roleid, troopsid int32) (*commonstruct.TroopsStruct, error) {
// 	filter := bson.D{primitive.E{Key: "roleid", Value: roleid}, primitive.E{Key: "troopsid", Value: troopsid}}
// 	troops := &commonstruct.TroopsStruct{}
// 	if err := db.FindOneBson(db.TroopsTable, &troops, filter); err != nil {
// 		return nil, err
// 	}
// 	return troops, nil
// }

func getRoleAlltroops(roleid int32) []*commonstruct.TroopsStruct {
	filter := bson.D{primitive.E{Key: "roleid", Value: roleid}}
	var results []*commonstruct.TroopsStruct
	cur, err := db.FindBson(db.TroopsTable, filter)
	if err != nil {
		return nil
	}
	for cur.Next(context.TODO()) {
		var elem commonstruct.TroopsStruct
		err := cur.Decode(&elem)
		if err != nil {
			log.Error(err)
			return nil
		}
		results = append(results, &elem)
	}
	return results
}

//部队信息
func (c *Client) sendTroopsList() {

	//角色上线部队信息初始化(用户登陆后 获取个人部队信息 和 大地图部队信息)
	//1. 从数据库/缓存 中查找数据
	//2. 和 troopsSMap 中的数据进行对比 更新同步
	//3. 推送给客户端

	var list []*troops.P_Troops
	for _, value := range getRoleAlltroops(c.roleid) {
		if v, ok := bigmapmanage.GetMapTroopsInfo(value.TroopsID); ok {
			log.Debug("大地图中的部队:", value.TroopsID, v)
			c.troopslist[v.TroopsID] = &v
			list = append(list, v.ConvertTroopsProto())
		} else {
			//log.Debug("角色未出征部队:", value)
			c.troopslist[value.TroopsID] = value
			list = append(list, value.ConvertTroopsProto())
		}
	}

	s2cmsg := &troops.S2C_TroopsList{TroopsList: list}
	c.Send(int32(troops.MSG_TROOPS_Module_TROOPS), int32(troops.MSG_TROOPS_S2C_TroopsList), s2cmsg)
}

//更新部队数据
//后期整理 所有数据库相关操作移至db 模块
func updateTroopsInfo(troops *commonstruct.TroopsStruct) {
	upinfo := bson.D{primitive.E{Key: "areasindex", Value: troops.AreasIndex},
		primitive.E{Key: "country", Value: troops.Country},
		primitive.E{Key: "areaslist", Value: troops.AreasList},
		primitive.E{Key: "movestamp", Value: troops.MoveStamp},
		primitive.E{Key: "state", Value: troops.State},
		primitive.E{Key: "type", Value: troops.Type},
		primitive.E{Key: "number", Value: troops.Number},
		primitive.E{Key: "level", Value: troops.Level},
		primitive.E{Key: "roleid", Value: troops.Roleid},
		primitive.E{Key: "arrivaltime", Value: troops.ArrivalTime},
		primitive.E{Key: "stagenumber", Value: troops.StageNumber},
	}

	db.Update(db.TroopsTable,
		bson.D{primitive.E{Key: "troopsid", Value: troops.TroopsID}},
		bson.D{primitive.E{Key: "$set", Value: upinfo}},
	)
}

func (c *Client) s2cUpdateTroopsInfo(troopsinfo *commonstruct.TroopsStruct) {
	c.Send(
		int32(troops.MSG_TROOPS_Module_TROOPS),
		int32(troops.MSG_TROOPS_S2C_UpdateTroopsInfo),
		&troops.S2C_UpdateTroopsInfo{TroopsInfo: troopsinfo.ConvertTroopsProto()},
	)
}

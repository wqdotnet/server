package clienconnect

import (
	"context"
	"server/db"
	"server/gserver/bigmapmanage"
	"server/gserver/commonstruct"
	"server/msgproto/bigmap"
	"server/msgproto/common"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

	var list []*bigmap.P_Troops
	for _, value := range getRoleAlltroops(c.roleid) {
		if v, ok := bigmapmanage.GetTroopsInfo(value.TroopsID); ok {
			list = append(list, convertTroopsProto(v))
		} else {
			list = append(list, convertTroopsProto(value))
		}
	}

	s2cmsg := &bigmap.S2C_TroopsList{TroopsList: list}
	log.Debug(s2cmsg)
	c.Send(int32(bigmap.MSG_BIGMAP_Module_BIGMAP), int32(bigmap.MSG_BIGMAP_S2C_AreasTroops), s2cmsg)
}

//获取部队信息
//后期整理 所有数据库相关操作移至db 模块
func getTroopsinfo(roleid, troopsid int32) (*commonstruct.TroopsStruct, error) {
	filter := bson.D{primitive.E{Key: "roleid", Value: roleid}, primitive.E{Key: "troopsid", Value: troopsid}}
	troops := &commonstruct.TroopsStruct{}
	if err := db.FindOneBson(db.TroopsTable, &troops, filter); err != nil {
		return nil, err
	}
	return troops, nil
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
		primitive.E{Key: "arrivaltime", Value: troops.ArrivalTime}}

	db.Update(db.TroopsTable,
		bson.D{primitive.E{Key: "troopsid", Value: troops.TroopsID}},
		bson.D{primitive.E{Key: "$set", Value: upinfo}},
	)
}

func (c *Client) s2cUpdateTroopsInfo(troops *commonstruct.TroopsStruct) {
	s2c := &bigmap.S2C_UpdateTroopsInfo{}
	s2c.TroopsInfo = convertTroopsProto(troops)
	c.Send(int32(bigmap.MSG_BIGMAP_Module_BIGMAP), int32(bigmap.MSG_BIGMAP_S2C_UpdateTroopsInfo), nil)
}

func convertTroopsProto(troops *commonstruct.TroopsStruct) *bigmap.P_Troops {
	return &bigmap.P_Troops{
		TroopsID:   troops.TroopsID,
		Country:    troops.Country,
		AreasList:  troops.AreasList,
		AreasIndex: troops.AreasIndex,

		State:  common.TroopsState(troops.State),
		Type:   troops.Type,
		Number: troops.Number,
		Level:  troops.Level,
		Roleid: troops.Roleid,
	}
}

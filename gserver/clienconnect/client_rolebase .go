package clienconnect

import (
	"server/gserver/cfg"
	pbrole "server/proto/role"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//练功修为同步

//练功突破 CMD_Role_Insight_CS

//雷劫
func (c *Client) upgrade(msg *pbrole.C2S_Upgrade) {
	retmsg := &pbrole.S2C_Upgrade{}
	rolebase := &c.roleData.RoleBase
	lvcfg := cfg.GetLvExpInfo(rolebase.Level)
	if rolebase.Exp < int64(lvcfg.NeedExp) {
		retmsg.Retcode = cfg.GetErrorCodeNumber("EXP_NOT_ENOUGH")
		c.SendToClient(int32(pbrole.MSG_ROLE_Module), int32(pbrole.MSG_ROLE_Upgrade), retmsg)
		return
	}

	rolebase.Exp -= int64(lvcfg.NeedExp)
	rolebase.Level++
	rolebase.SetDirtyData(primitive.E{Key: "exp", Value: rolebase.Exp},
		primitive.E{Key: "level", Value: rolebase.Level})
	rolebase.CalculationProperties()

	retmsg.Exp = rolebase.Exp
	retmsg.Level = rolebase.Level
	c.SendToClient(int32(pbrole.MSG_ROLE_Module), int32(pbrole.MSG_ROLE_Upgrade), retmsg)
	c.SendToClient(int32(pbrole.MSG_ROLE_Module),
		int32(pbrole.MSG_ROLE_AttributeChange),
		&pbrole.S2C_AttributeChange_S{
			AttributeList: rolebase.AttributeValue,
			CE:            rolebase.CE,
		})

}

//背包

//阵法

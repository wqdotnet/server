package commonstruct

//ProcessMsgType 进程消息
type ProcessMsgType string

const (
	//ProcessMsgSocket socket
	ProcessMsgSocket ProcessMsgType = "Socket"
	//ProcessMsgTimeInterval 不定时消息
	ProcessMsgTimeInterval ProcessMsgType = "TimeInterval"
	//ProcessMsgRoleLogin 角色重复登陆挤下线
	ProcessMsgRoleLogin ProcessMsgType = "RoleLogin"
	//ProcessMsgTroopsMove 部队大地图中移动
	ProcessMsgTroopsMove ProcessMsgType = "TroopsMove"
	//ProcessMsgOverMove 部队结束移动
	ProcessMsgOverMove ProcessMsgType = "OverMove"
	//ProcessMsgOnFitht 部队触发战斗
	ProcessMsgOnFitht ProcessMsgType = "OnFitht"
	//ProcessMsgOverFitht 部队结束战斗
	ProcessMsgOverFitht ProcessMsgType = "OverFitht"
	//ProcessMsgAreasState 区域状态变化
	ProcessMsgAreasState ProcessMsgType = "AreasState"
	//ProcessMsgUpdateTroopsInfo 更新部队信息
	ProcessMsgUpdateTroopsInfo ProcessMsgType = "UpdateTroopsInfo"
	//ProcessMsgAddExp 加经验
	ProcessMsgAddExp ProcessMsgType = "AddExp"
)

//ProcessMsg 进程间消息
type ProcessMsg struct {
	MsgType ProcessMsgType
	Module  int32
	Method  int32
	Data    interface{}
}

package commonstruct

import (
	"server/msgproto/common"
	"time"
)

//AccountInfoStruct 账号信息
type AccountInfoStruct struct {
	Account            string
	AccountID          int32 //账号id
	Password           string
	Equipment          string //设备信息
	RegistrationSource string //注册来源(平台)
	RegistrationTime   time.Time
	RoleID             int32 //角色id

}

// //TroopsState 部队状态
// type TroopsState int32

// const (
// 	//TSStandBy 等待命令
// 	TSStandBy TroopsState = 0
// 	//TSMove 移动中
// 	TSMove TroopsState = 1
// 	//TSPause 暂停
// 	TSPause TroopsState = 2
// 	//TSfight 战斗中
// 	TSfight TroopsState = 3
// )

//TroopsStruct 部队信息
type TroopsStruct struct {
	TroopsID   int32   //部队名/id
	Country    int32   //国家归属
	AreasList  []int32 //移动路径 区域id list
	AreasIndex int32   //当前区域ID

	MoveStamp int64 //上次移动时间戳

	State common.TroopsState //状态 0:未出动    1:移动  2:驻扎(暂停)  3:战斗
	//兵种组成
	Type   int32 //部队类型
	Number int32 //数量
	Level  int32 //等级
	Roleid int32

	ArrivalTime int64 //预计到达时间
}

//ProcessMsg 进程间消息
type ProcessMsg struct {
	MsgType string
	Data    interface{}
}

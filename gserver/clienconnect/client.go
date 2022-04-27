package clienconnect

import (
	"fmt"
	"runtime"
	"server/gserver/cfg"
	"server/gserver/commonstruct"
	"server/gserver/nodeManange"
	"server/proto/account"
	"server/proto/item"
	"server/proto/protocol_base"
	"server/proto/role"
	"server/tools"
	"time"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//Client 客户端连接
type Client struct {
	process         *gen.ServerProcess
	registerName    string
	sendChan        chan []byte
	infofunc        map[int32]func(buf []byte)
	genServerStatus gen.ServerStatus

	roleID       int32      //角色ID
	connectState userStatus //连接状态
}

type userStatus int32

const (
	//StatusSockert socker 连接状态
	StatusSockert userStatus = 0
	//StatusLogin 已登陆成功
	StatusLogin userStatus = 1
	//正常游戏中
	StatusGame userStatus = 2

	//StatusSqueezeOut 重复登陆 挤下线
	//StatusSqueezeOut userStatus = 2
)

//===========GateGenHanderInterface 接口实现===============
func NewClient() *Client {
	client := &Client{}
	client.initMsgRoute()
	return client
}

func (c *Client) initMsgRoute() {
	//消息注册
	c.infofunc = make(map[int32]func(buf []byte))
	//心跳
	c.infofunc[int32(protocol_base.MSG_BASE_HeartBeat)] = createRegisterFunc(c.heartBeat)

	//账号
	c.infofunc[int32(account.MSG_ACCOUNT_Login)] = createRegisterFunc(c.accountLogin)
	c.infofunc[int32(account.MSG_ACCOUNT_Register)] = createRegisterFunc(c.registerAccount)
	c.infofunc[int32(account.MSG_ACCOUNT_CreateRole)] = createRegisterFunc(c.accountCreateRole)
	//角色

	//道具
	c.infofunc[int32(item.MSG_Item_GetBackpackInfo)] = createRegisterFunc(c.getBackpackInfo)

	c.genServerStatus = gen.ServerStatusOK
}

func (c *Client) InitHander(process *gen.ServerProcess, sendChan chan []byte) {
	c.process = process
	c.sendChan = sendChan

}

func (c *Client) MsgHander(module, method int32, buf []byte) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error(err)
		}
	}()

	//禁用模块
	//next...

	if msgfunc := c.infofunc[method]; msgfunc != nil {
		if c.connectState == StatusGame || module == int32(account.MSG_ACCOUNT_Module) {
			msgfunc(buf)
		} else {
			logrus.Errorf("未登陆 [%v] [%v] [%v]", module, method, buf)
		}
	} else {
		logrus.Warnln("未注册的消息", module, method)
	}
}

func (c *Client) LoopHander() time.Duration {
	defer func() {
		if err := recover(); err != nil {
			var err string
			for i := 0; i < 10; i++ {
				pc, fn, line, _ := runtime.Caller(i)
				if line == 0 {
					break
				}
				err += fmt.Sprintf("funcname:[%v] fn:[%v] line:[%v] \n", runtime.FuncForPC(pc).Name(), fn, line)
			}
			logrus.Error("err: \n", err)
		}
	}()

	if c.connectState == StatusGame {
		roledata := commonstruct.GetRoleAllData(c.roleID)
		nowtime := time.Now().Unix()
		num := (nowtime - roledata.RoleBase.PracticeTimestamp) / 5
		if num > 0 {
			roledata.RoleBase.PracticeTimestamp = nowtime
			if expcfg := cfg.GetLvExpInfo(roledata.RoleBase.Level); expcfg != nil {
				expcfg := cfg.GetLvExpInfo(roledata.RoleBase.Level)
				addexp := int64(expcfg.CycleEXP) * num
				roledata.RoleBase.Exp += addexp
				//加经验通知
				c.SendToClient(int32(role.MSG_ROLE_Module),
					int32(role.MSG_ROLE_AddExp),
					&role.S2C_AddExp_S{Exp: roledata.RoleBase.Exp, Addexp: addexp})

			} else {
				logrus.Error("expcfg is nil:", roledata.RoleBase.Level, roledata.RoleBase)
			}
		}

		//logrus.Debug("loop role exp:", roledata.RoleBase.Exp, num)
	}

	return time.Second * 5
}

func (c *Client) HandleCall(message etf.Term) {

}

func (c *Client) HandleInfo(message etf.Term) {
	// switch info := message.(type) {
	// case etf.Tuple:
	// 	switch info[0].(string) {
	// 	case "BroadcastMsg":
	// 		module := info[1].(int32)
	// 		method := info[2].(int32)
	// 		buf := info[3].(proto.Message)
	// 		c.SendToClient(module, method, buf)
	// 	}
	// }
}

func (c *Client) GenServerStatus() gen.ServerStatus {
	return c.genServerStatus
}

func (c *Client) Terminate(reason string) {
	logrus.Debugf("client 关闭  [%v] roleid:[%v]  [%v]", reason, c.roleID, c.process.Name())
	switch reason {
	case "Extrusionline": //挤下线 不中断socket连接 ,只将状态设置成为登陆状态
		c.SendToClient(int32(protocol_base.MSG_BASE_Module),
			int32(protocol_base.MSG_BASE_NoticeMsg),
			&protocol_base.S2C_NoticeMsg{
				Retcode:   0,
				NoticeMsg: reason,
			})
		c.connectState = StatusSockert
		return
	}

	node := nodeManange.GetNode(nodeManange.GateNode)
	node.UnregisterName(c.registerName)
	c.registerName = ""
	c.process = nil
	c.sendChan = nil
	c.roleID = 0
	c.connectState = StatusSockert

}

//==========================

// //SendToClient 发送消息至客户端
func (c *Client) SendToClient(module int32, method int32, pb proto.Message) {
	//logrus.Debugf("client send msg [%v] [%v] [%v]", module, method, pb)
	data, err := proto.Marshal(pb)
	if err != nil {
		logrus.Errorf("proto encode error[%v] [%v][%v] [%v]", err.Error(), module, method, pb)
		return
	}
	// msginfo := &common.NetworkMsg{}
	// msginfo.Module = module
	// msginfo.Method = method
	// msginfo.MsgBytes = data
	// msgdata, err := proto.Marshal(msginfo)
	// if err != nil {
	// 	logrus.Errorf("msg encode error[%s]\n", err.Error())
	// }
	// gateGS.sendChan <- msgdata

	mldulebuf := tools.IntToBytes(module, 2)
	methodbuf := tools.IntToBytes(method, 2)
	c.sendChan <- tools.BytesCombine(mldulebuf, methodbuf, data)
}

//==========msg register =======
//消息注册
func createRegisterFunc[T any](execfunc func(*T)) func(buf []byte) {
	return func(buf []byte) {
		info := new(T)
		err := decodeProto(info, buf)
		if err != nil {
			logrus.Errorf("decode error[%v]", err.Error())
		} else {
			//logrus.Debugf("client msg:[%v] [%v]", info, tools.GoID())
			execfunc(info)
		}
	}
}

//protobuf 解码
func decodeProto(info interface{}, buf []byte) error {
	if data, ok := info.(protoreflect.ProtoMessage); ok {
		return proto.Unmarshal(buf, data)
	}
	return nil
}

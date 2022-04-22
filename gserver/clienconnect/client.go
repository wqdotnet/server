package clienconnect

import (
	"server/db"
	"server/gserver/commonstruct"
	"server/proto/account"
	"server/tools"
	"time"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//Client 客户端连接
type Client struct {
	process         *gen.ServerProcess
	sendChan        chan []byte
	infofunc        map[int32]func(buf []byte)
	genServerStatus gen.ServerStatus
}

type userStatus int32

const (
	//StatusSockert socker 连接状态
	StatusSockert userStatus = 0
	//StatusLogin 已登陆成功
	StatusLogin userStatus = 1
	//StatusSqueezeOut 重复登陆 挤下线
	StatusSqueezeOut userStatus = 2
)

//===========GateGenHanderInterface 接口实现===============
func NewClient() *Client {
	client := &Client{}
	client.initMsgRoute()
	return client
}

func (c *Client) initMsgRoute() {
	c.infofunc = make(map[int32]func(buf []byte))
	//消息注册
	c.infofunc[int32(account.MSG_ACCOUNT_Login)] = createRegisterFunc(c.accountLogin)
	c.infofunc[int32(account.MSG_ACCOUNT_Register)] = createRegisterFunc(c.registerAccount)
	c.infofunc[int32(account.MSG_ACCOUNT_CreateRole)] = createRegisterFunc(c.accountCreateRole)
	c.genServerStatus = gen.ServerStatusOK
}

func (c *Client) InitHander(process *gen.ServerProcess, sendChan chan []byte) {
	c.process = process
	c.sendChan = sendChan
}

func (c *Client) MsgHander(module, method int32, buf []byte) {
	if msgfunc := c.infofunc[method]; msgfunc != nil {
		msgfunc(buf)
	} else {
		logrus.Warnln("未注册的消息", module, method)
	}
}

func (c *Client) LoopHander() time.Duration {
	c.accountLogin(&account.C2S_Login{Account: "asdf", Password: "1234"})
	return time.Second
}

func (c *Client) HandleCall(message etf.Term) {

}
func (c *Client) HandleInfo(message etf.Term) {}

func (c *Client) GenServerStatus() gen.ServerStatus {
	return c.genServerStatus
}

func (c *Client) Terminate(reason string) {
	c.process = nil
	c.sendChan = nil
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

//================db get data=============
func GetAccountinfo(account, passwd string) (bool, *commonstruct.AccountInfo) {
	filter := bson.D{
		primitive.E{Key: "account", Value: account},
		primitive.E{Key: "password", Value: passwd},
	}
	accountinfo := &commonstruct.AccountInfo{}
	if err := db.FindOneBson(db.AccountTable, accountinfo, filter); err != nil {
		return false, nil
	}

	return true, accountinfo
}

func GetRoleAllData(roleid int32) *commonstruct.RoleData {

	return &commonstruct.RoleData{}
}

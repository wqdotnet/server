package clienconnect

import (
	"fmt"
	"server/tools"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// import (
// 	"net"
// 	"server/db"
// 	"server/gserver/commonstruct"
// 	"server/gserver/process"
// 	"server/msgproto/account"
// 	"server/msgproto/bigmap"
// 	"server/msgproto/protocol_base"
// 	"server/msgproto/troops"
// 	"server/network"
func init() {
	fmt.Print("client init --------------------------------------")
}

// 	"github.com/golang/protobuf/proto"
// 	log "github.com/sirupsen/logrus"
// )

//Client 客户端连接
type Client struct {
	sendChan chan []byte
	infofunc map[int32]func(buf []byte)

	//用户 连接状态 [0:连接] [1:已登陆] [2:下线]
	status userStatus
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

func (c *Client) InitHander(sendChan chan []byte) {
	c.sendChan = sendChan
	c.infofunc = make(map[int32]func(buf []byte))

	//消息注册
	c.infofunc[101] = createRegisterFunc(msgCreateRole)
}

func (c *Client) MsgHander(module, method int32, buf []byte) {
	c.infofunc[method](buf)
}

// //Send 发送消息
func (c *Client) Send(module int32, method int32, pb proto.Message) {
	//log.Debugf("client send msg [%v] [%v] [%v]", module, method, pb)
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Errorf("proto encode error[%v] [%v][%v] [%v]", err.Error(), module, method, pb)
		return
	}
	// msginfo := &common.NetworkMsg{}
	// msginfo.Module = module
	// msginfo.Method = method
	// msginfo.MsgBytes = data
	// msgdata, err := proto.Marshal(msginfo)
	// if err != nil {
	// 	log.Errorf("msg encode error[%s]\n", err.Error())
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
			log.Errorf("decode error[%v]", err.Error())
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

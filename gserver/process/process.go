package process

import (
	"slgserver/gserver/commonstruct"
	"sync"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var (
	//玩家连接进程注册
	gameProcess sync.Map
)

//Register 注册协程
func Register(key interface{}, info chan commonstruct.ProcessMsg) {
	if IsRegister(key) {
		log.Warnf("gameProcess [%v] is Register", key)
	}
	gameProcess.Store(key, info)
}

//IsRegister 进程是否注册
func IsRegister(key interface{}) bool {
	_, ok := gameProcess.Load(key)
	return ok
}

func writeChWithSelect(ch chan commonstruct.ProcessMsg, msg commonstruct.ProcessMsg) {
	// unixnano := time.Now().UnixNano()
	// k := rand.Intn(100000)
	// log.Infof("[%v]bigmap to role ", k)
	ch <- msg
	// nano := time.Now().UnixNano() - unixnano
	// log.Infof("[%v]bigmap to role end:[%v][%v][%v][%v]", k, msg.MsgType, nano/1e6, nano/1e9, nano)

	//timeout := time.NewTimer(time.Microsecond * 500)
	// select {
	// case ch <- msg:
	// case <-timeout.C:
	// 	log.Warn("消息超时:", msg)
	// }
}

//SendMsg send msg
func SendMsg(key interface{}, msg commonstruct.ProcessMsg) bool {
	info, ok := gameProcess.Load(key)
	if ok {
		go func() {
			info.(chan commonstruct.ProcessMsg) <- msg
		}()
		//writeChWithSelect(info.(chan commonstruct.ProcessMsg), msg)
		return true
	}

	log.Debug(key, msg)
	return false
}

//SendSocketMsg send protobuf to socket
func SendSocketMsg(key interface{}, module int32, method int32, pb proto.Message) bool {
	info, ok := gameProcess.Load(key)
	if ok {
		go func() {
			data, err := proto.Marshal(pb)
			if err != nil {
				log.Warn(err)
			} else {
				writeChWithSelect(info.(chan commonstruct.ProcessMsg), commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgSocket, Module: module, Method: method, Data: data})
			}
		}()
		return true
	}
	log.Debug(key, pb)
	return false
}

//SendAllSocketMsg send all roles
func SendAllSocketMsg(module int32, method int32, pb proto.Message) {
	go func() {
		gameProcess.Range(func(key, value interface{}) bool {
			data, err := proto.Marshal(pb)
			if err != nil {
				log.Warn(err)
			} else {
				msg := commonstruct.ProcessMsg{MsgType: commonstruct.ProcessMsgSocket,
					Module: module,
					Method: method,
					Data:   data,
				}
				//value.(chan commonstruct.ProcessMsg) <- msg
				writeChWithSelect(value.(chan commonstruct.ProcessMsg), msg)
			}
			return true
		})
	}()
}

//UnRegister 取消
func UnRegister(key interface{}) {
	// 删除
	gameProcess.Delete(key)
}

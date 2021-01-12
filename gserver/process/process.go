package process

import (
	"server/gserver/commonstruct"
	"sync"

	log "github.com/sirupsen/logrus"
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

//SendMsg send msg
func SendMsg(key interface{}, msg commonstruct.ProcessMsg) bool {
	info, ok := gameProcess.Load(key)
	if ok {
		go func() {
			info.(chan commonstruct.ProcessMsg) <- msg
		}()
		return true
	}

	log.Debug(key, msg)
	return false
}

//UnRegister 取消
func UnRegister(key interface{}) {
	// 删除
	gameProcess.Delete(key)
}

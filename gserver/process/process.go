package process

import (
	"server/gserver/commonstruct"
	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	gameProcess sync.Map
)

//Register 注册协程
func Register(key interface{}, info chan commonstruct.ProcessMsg) {
	if _, ok := gameProcess.Load(key); ok {
		log.Warnf("gameProcess [%v] is Register", key)
	} else {
		gameProcess.Store(key, info)
	}
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
		info.(chan commonstruct.ProcessMsg) <- msg
		return true
	}

	log.Debug(ok, info)
	return false
}

//UnRegister 取消
func UnRegister(key interface{}) {
	// 删除
	gameProcess.Delete(key)
}

// //全局协程注册服务
// var cslist sync.Map

// //CSInterface 协程接口
// type CSInterface interface {
// 	GetSPType() CSType
// }

// //CSType 协和类型
// type CSType int32

// const (
// 	//GameServer 游戏服务
// 	GameServer CSType = 0
// 	//ClientConnect 用户连接bind
// 	ClientConnect CSType = 1
// )

// type registerChan struct {
// 	msgchan chan interface{}
// 	msgfunc func(msgchan chan interface{})
// }

// //Register 注册协程
// func Register(key string, info CSInterface) {
// 	if _, ok := cslist.Load(key); ok {
// 		log.Warn("sp [%s] is Register\n", key)
// 	} else {
// 		cslist.Store(key, &registerChan{
// 			msgchan: make(chan interface{}),
// 		})
// 		//cslist.Store(key, info)
// 	}
// }

// //SendMsg send msg
// func SendMsg(key string, info interface{}) {
// 	if info, ok := cslist.Load(key); ok {
// 		info.(registerChan).msgchan <- info
// 	} else {
// 		log.Warn("no Register：", key)
// 	}
// }

// //GetCoroutine 获取协程
// func GetCoroutine(key string) CSInterface {
// 	if info, ok := cslist.Load(key); ok {
// 		return info.(CSInterface)
// 	}
// 	return nil
// }

// //UnRegister 取消
// func UnRegister(key string) {
// 	// 删除
// 	cslist.Delete(key)
// }

// //Range 循环遍历
// func Range(exec func(key string, value CSInterface) bool) {
// 	// 遍历
// 	cslist.Range(func(key, value interface{}) bool {
// 		switch value.(type) {
// 		case *CSInterface:
// 			return exec(key.(string), value.(CSInterface))
// 		default:
// 		}
// 		return true
// 	})
// }

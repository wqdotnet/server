package cservice

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

//全局协程注册服务
var cslist sync.Map

//CSInterface 协程接口
type CSInterface interface {
	GetSPType() CSType
}

//CSType 协和类型
type CSType int32

const (
	//GameServer 游戏服务
	GameServer CSType = 0
	//ClientConnect 用户连接bind
	ClientConnect CSType = 1
)

//Register 注册协程
func Register(key string, info CSInterface) {
	if _, ok := cslist.Load(key); ok {
		log.Warn("sp [%s] is Register\n", key)
	} else {
		cslist.Store(key, info)
	}
}

//GetCoroutine 获取协程
func GetCoroutine(key string) CSInterface {
	if info, ok := cslist.Load(key); ok {
		return info.(CSInterface)
	}
	return nil
}

//UnRegister 取消
func UnRegister(key string) {
	// 删除
	cslist.Delete(key)
}

//Range 循环遍历
func Range(exec func(key string, value CSInterface) bool) {
	// 遍历
	cslist.Range(func(key, value interface{}) bool {
		switch value.(type) {
		case *CSInterface:
			return exec(key.(string), value.(CSInterface))
		default:
		}
		return true
	})
}

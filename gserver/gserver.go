package gserver

import (
	"fmt"
	net "server/network"
	//msg "server/proto"
)

func Start() {
	nw := net.NewNetWorkX()
	nw.Start()
	fmt.Println("start game server")
}

// //UserLogin 用户登陆
// func UserLogin(buf []byte) {

// }

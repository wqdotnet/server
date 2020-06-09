package gserver

import (
	"fmt"
	net "server/network"
	//msg "server/proto"
)

//StartGServer 启动game server
func StartGServer() {
	fmt.Println("start game server [success]")
	//启动网络
	nw := net.NewNetWorkX()
	nw.Start()

	fmt.Println("Shut down the server")
}

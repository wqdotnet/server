package gserver

import (
	"fmt"
	net "server/network"
	"sync"
	//msg "server/proto"
)

//StartGServer 启动game server
func StartGServer() {
	fmt.Println("start game server ")

	//go web.Start()

	//启动网络
	nw := net.NewNetWorkX(&sync.Pool{
		New: func() interface{} {
			return new(client)
		},
	})
	nw.Start()

	fmt.Println("Shut down the server")
}

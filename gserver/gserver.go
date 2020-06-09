package gserver

import (
	"fmt"
	net "server/network"
	"server/web"
	"sync"
	//msg "server/proto"
)

//StartGServer 启动game server
func StartGServer() {

	web.Start()
	fmt.Println("start game server [success]")

	//启动网络
	nw := net.NewNetWorkX(&sync.Pool{
		New: func() interface{} {
			return new(client)
		},
	})
	nw.Start()

	fmt.Println("Shut down the server")
}

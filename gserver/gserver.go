package gserver

import (
	"fmt"
	net "server/network"
	"sync"
	//msg "server/proto"
)

// ServerConfig  server cfg
type ServerConfig struct {
	OpenHTTP string
	HTTPPort int

	NetWork string
	Port    int

	//proto_path=%s  --go_out

	ProtoPath string
	GoOut     string
}

// ServerCfg  Program overall configuration
var ServerCfg = ServerConfig{

	OpenHTTP: "localhost",
	HTTPPort: 8080,

	// #network : tcp/udp
	NetWork: "tcp",
	Port:    3344,

	// #protobuf path
	ProtoPath: "/proto",
	GoOut:     "/proto",
}

//StartGServer 启动game server
func StartGServer() {
	fmt.Println("start game server ")
	//ServerConfig
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

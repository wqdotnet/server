package gserver

import (
	"fmt"
	"server/network"
	"server/web"
	"sync"
	//msg "server/proto"
)

// ServerConfig  server cfg
type ServerConfig struct {
	OpenHTTP bool
	HTTPPort int32

	NetType string
	Port    int32
	Packet  int32

	ProtoPath string
	GoOut     string
}

// ServerCfg  Program overall configuration
var ServerCfg = ServerConfig{

	OpenHTTP: false,
	HTTPPort: 8080,

	// #network : tcp/udp
	NetType: "tcp",
	Port:    3344,
	Packet:  2,

	// #protobuf path
	ProtoPath: "./proto",
	GoOut:     "./proto",
}

//StartGServer 启动game server
//go run main.go start --config=E:/worke/server/cfg.yaml
func StartGServer() {
	fmt.Println("start game server ")
	//ServerConfig
	if ServerCfg.OpenHTTP == true {
		go web.Start(ServerCfg.HTTPPort)
	}

	// classname := reflect.TypeOf((*ServerConfig)(nil)).Elem()
	// fmt.Println("reflect:", classname)
	// msgt := reflect.New(classname).Interface()
	// switch msgt.(type) {
	// // 有新的连接
	// case *ServerConfig:
	// 	fmt.Println("格式解析")
	// }

	//启动网络
	nw := network.NewNetWorkX(&sync.Pool{
		New: func() interface{} {
			return new(client)
		},
	})

	nw.Port = ServerCfg.Port
	nw.Packet = ServerCfg.Packet
	nw.NetType = ServerCfg.NetType
	nw.Start()

	fmt.Println("Shut down the server")
}

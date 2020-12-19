package gserver

import (
	"server/db"
	"server/gserver/cfg"
	"server/gserver/clienconnect"
	"server/gserver/cservice"
	"server/logger"
	"server/network"
	"server/web"
	"sync"

	log "github.com/sirupsen/logrus"
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

	MongoConnStr string
	Mongodb      string

	RedisConnStr string

	CfgPath string
	CfgType string

	LogWrite bool
	Loglevel string
	LogPath  string
	LogName  string
}

// ServerCfg  Program overall configuration
var ServerCfg = ServerConfig{
	// http
	OpenHTTP: true,
	HTTPPort: 8080,

	// #network : tcp/udp
	NetType: "tcp",
	Port:    3344,
	Packet:  2,

	// #protobuf path
	ProtoPath: "./proto",
	GoOut:     "./proto",

	MongoConnStr: "mongodb://localhost:27017",
	Mongodb:      "mygame",

	RedisConnStr: "127.0.0.1:6379",

	CfgPath: "./config",
	CfgType: "",

	Loglevel: "info",
	LogPath:  "./log",
	LogName:  "log",
	LogWrite: false,
}

type gameServer struct {
	nw *network.NetWorkx
	//game config
	command chan string
}

//GameServerInfo game info
var GameServerInfo = gameServer{
	nw: network.NewNetWorkX(&sync.Pool{
		New: func() interface{} {
			return new(clienconnect.Client)
		},
	},
		ServerCfg.Port,
		ServerCfg.Packet,
		ServerCfg.NetType),
	command: make(chan string),
}

func (s *gameServer) GetSPType() cservice.CSType {
	return cservice.GameServer
}

//StartGServer 启动game server
//go run main.go start --config=E:/worke/server/cfg.yaml
func StartGServer() {
	if level, err := log.ParseLevel(ServerCfg.Loglevel); err == nil {
		logger.Init(level, ServerCfg.LogWrite, ServerCfg.LogName, ServerCfg.LogPath)
	} else {
		logger.Init(log.InfoLevel, ServerCfg.LogWrite, ServerCfg.LogName, ServerCfg.LogPath)
	}

	log.Info("start game server ")

	cfg.InitViperConfig(ServerCfg.CfgPath, ServerCfg.CfgType)
	db.StartMongodb(ServerCfg.Mongodb, ServerCfg.MongoConnStr)
	db.StartRedis(ServerCfg.RedisConnStr)
	clienconnect.InitAutoID()

	if ServerCfg.OpenHTTP {
		go web.Start(ServerCfg.HTTPPort)
	}

	//启动网络
	GameServerInfo.nw.Start()

	//登陆成功后注册进程
	cservice.Register("server", &GameServerInfo)

	for {
		select {
		case command := <-GameServerInfo.command:
			switch command {
			case "down":
				log.Warn("Shut down the game server")
			default:
				log.Warn("command:", command)
			}
		}
	}

}

package gserver

import (
	"os"
	"os/signal"
	"server/db"
	"server/gserver/bigmapmanage"
	"server/gserver/cfg"
	"server/gserver/clienconnect"
	"server/gserver/timedtasks"
	"server/logger"
	"server/network"
	"server/web"
	"sync"
	"syscall"

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
			return clienconnect.NewClient() //new(clienconnect.Client)
		},
	},
		ServerCfg.Port,
		ServerCfg.Packet,
		ServerCfg.NetType),
	command: make(chan string),
}

//StartGServer 启动game server
//go run main.go start --config=E:/worke/server/cfg.yaml
func StartGServer() {
	log.Info("start game server")
	if level, err := log.ParseLevel(ServerCfg.Loglevel); err == nil {
		logger.Init(level, ServerCfg.LogWrite, ServerCfg.LogName, ServerCfg.LogPath)
	} else {
		logger.Init(log.InfoLevel, ServerCfg.LogWrite, ServerCfg.LogName, ServerCfg.LogPath)
	}

	cfg.InitViperConfig(ServerCfg.CfgPath, ServerCfg.CfgType)
	db.StartMongodb(ServerCfg.Mongodb, ServerCfg.MongoConnStr)
	db.StartRedis(ServerCfg.RedisConnStr)
	clienconnect.InitAutoID()

	//ctx, cancelFunc := context.WithCancel(context.Background())
	bigmapmanage.StartBigmapGoroutine()
	defer bigmapmanage.CloneBigmap()

	//启动定时器
	timedtasks.StartCronTasks()
	//大地图loop
	timedtasks.AddTasks("bigmaploop", "* * * * * ?", func() {
		bigmapmanage.SendMsgBigMap("BigMapLoop_OneSecond")
	})
	defer timedtasks.RemoveTasks("bigmaploop")

	if ServerCfg.OpenHTTP {
		go web.Start(ServerCfg.HTTPPort)
	}

	//启动网络
	GameServerInfo.nw.Start()

	var exitChan = make(chan os.Signal)
	//signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	signal.Notify(exitChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1,
		syscall.SIGUSR2, syscall.SIGTSTP)

	for {
		select {
		case command := <-GameServerInfo.command:
			switch command {
			case "down":
				log.Warn("Shut down the game server")
			default:
				log.Warn("command:", command)
			}
		case s := <-exitChan:
			log.Info("收到退出信号", s)
			return
			//os.Exit(1) //如果ctrl+c 关不掉程序，使用os.Exit强行关掉
		}
	}
}

//SendGameServerMsg game system msg
func SendGameServerMsg(msg string) {
	GameServerInfo.command <- msg
}

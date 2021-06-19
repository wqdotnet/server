package gserver

import (
	"os"
	"os/signal"
	"runtime"
	"server/db"
	"server/gserver/cfg"
	"server/gserver/timedtasks"
	"server/logger"
	"server/network"
	"server/web"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
	//msg "server/proto"

	"github.com/halturin/ergo"
)

type gameServer struct {
	nw *network.NetWorkx
	//game config
	command chan string
	node    *ergo.Node
}

//GameServerInfo game info
var GameServerInfo gameServer

func (gs *gameServer) startOtp() {
	opts := ergo.NodeOptions{
		ListenRangeBegin: uint16(ServerCfg.ListenRangeBegin),
		ListenRangeEnd:   uint16(ServerCfg.ListenRangeEnd),
		EPMDPort:         uint16(ServerCfg.EPMDPort),
	}

	gs.node = ergo.CreateNode(ServerCfg.NodeName, ServerCfg.Cookie, opts)
	process, _ := gs.node.Spawn("serverSup", ergo.ProcessOptions{}, &serverSup{})
	process.Wait()
}

//StartGServer 启动game server
//go run main.go start --config=E:/worke/server/cfg.yaml
func StartGServer() {
	log.Infof("====================== Begin Start [%v][%v] ======================", ServerCfg.ServerName, ServerCfg.ServerID)
	if level, err := log.ParseLevel(ServerCfg.Loglevel); err == nil {
		logger.Init(level, ServerCfg.LogWrite, ServerCfg.LogName, ServerCfg.LogPath)
	} else {
		logger.Init(log.InfoLevel, ServerCfg.LogWrite, ServerCfg.LogName, ServerCfg.LogPath)
	}

	// if ServerCfg.Daemon {
	//https://github.com/takama/daemon
	// }

	cfg.InitViperConfig(ServerCfg.CfgPath, ServerCfg.CfgType)

	db.StartMongodb(ServerCfg.Mongodb, ServerCfg.MongoConnStr)
	db.StartRedis(ServerCfg.RedisConnStr, ServerCfg.RedisDB)

	//启动定时器
	timedtasks.StartCronTasks()
	// //定时器
	// timedtasks.AddTasks("loop", "* * * * * ?", func() {
	// 	log.Info("server time:", time.Now())
	// })
	//defer timedtasks.RemoveTasks("loop")

	if ServerCfg.OpenHTTP {
		go web.Start(ServerCfg.HTTPPort)
	}

	if ServerCfg.StatsView {
		go web.StartStatsView(ServerCfg.StatsViewPort)
	}

	GameServerInfo = gameServer{
		nw: network.NewNetWorkX(&sync.Pool{
			New: func() interface{} {
				return nil // clienconnect.NewClient() //new(clienconnect.Client)
			}},
			ServerCfg.Port,
			ServerCfg.Packet,
			ServerCfg.Readtimeout,
			ServerCfg.NetType,
			ServerCfg.MsgTime,
			ServerCfg.MsgNum,
			func() { SendGameServerMsg("StartSuccess") },
			func() { db.RedisExec("del", "ConnectNumber") },
			func() { log.Info("connect number: ", db.INCRBY("ConnectNumber", 1)) },
			func() { log.Info("connect number: ", db.INCRBY("ConnectNumber", -1)) },
		),
		command: make(chan string),
	}
	//启动网络
	//GameServerInfo.nw.Start()
	//defer GameServerInfo.nw.Close()
	//GameServerInfo.startOtp()

	//退出消息监控
	var exitChan = make(chan os.Signal)
	if runtime.GOOS == "linux" {
		//signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
		signal.Notify(exitChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTSTP)
	}

	for {
		select {
		case command := <-GameServerInfo.command:
			switch command {
			case "StartSuccess":
				log.Info("====================== Start Game Server Success =========================")
			case "down":
				log.Warn("Shut down the game server")
			default:
				log.Warn("command:", command)
			}
		case s := <-exitChan:
			log.Info("收到信号: ", s)
			if s.String() == "quit" || s.String() == "terminated" {
				//os.Exit(1)
				return
			}
			//case <-time.After(60 * time.Second):
			//log.Infof("time: [%v]  online:[%v]", time.Now().Format(tool.DateTimeFormat), db.RedisGetInt("ConnectNumber"))
		}
	}
}

//SendGameServerMsg game system msg
func SendGameServerMsg(msg string) {
	GameServerInfo.command <- msg
}

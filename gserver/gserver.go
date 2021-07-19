package gserver

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"server/db"
	"server/gserver/cfg"
	"server/logger"
	"server/network"
	"server/web"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
	//msg "server/proto"

	"github.com/halturin/ergo"
)

//GameServerInfo game info
var GameServerInfo *gameServer

type gameServer struct {
	nw       *network.NetWorkx
	serverid int32
	command  chan string

	//由于没有 erlang:nodes()  手动维护所有节点信息
	//三种节点类型 gate server db
	//gate、db 集群可能会有多个
	nodes map[string]*ergo.Node
}

func (g *gameServer) Start() {
	//启动网络
	g.nw.Start()
	gateNodeName := fmt.Sprintf("gatewayNode_%v@127.0.0.1", g.serverid)
	serverNodeName := fmt.Sprintf("serverNode_%v@127.0.0.1", g.serverid)
	dbNodeName := fmt.Sprintf("dbNode_%v@127.0.0.1", g.serverid)

	gateNode, _ := StartGateSupNode(gateNodeName)
	serverNode, _ := StartGameServerSupNode(serverNodeName)
	dbNode, _ := StartDataBaseSupSupNode(dbNodeName)
	g.nodes[gateNode.FullName] = gateNode
	g.nodes[serverNode.FullName] = serverNode
	g.nodes[dbNode.FullName] = dbNode
}

func (g *gameServer) Close() {
	g.nw.Close()
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

	//启动定时器
	//timedtasks.StartCronTasks()
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

	GameServerInfo = &gameServer{
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
		command:  make(chan string),
		nodes:    make(map[string]*ergo.Node),
		serverid: ServerCfg.ServerID,
	}
	GameServerInfo.Start()
	defer GameServerInfo.Close()

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

package gserver

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"server/db"
	"server/gserver/cfg"
	"server/gserver/clienconnect"
	"server/gserver/genServer"
	"server/gserver/nodeManange"
	"server/logger"
	"server/network"
	"server/web"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"
	"github.com/sirupsen/logrus"

	//msg "server/proto"

	"github.com/facebookgo/pidfile"
)

//GameServerInfo game info
var GameServerInfo *gameServer

type gameServer struct {
	nw      *network.NetWorkx
	command chan string
}

func (g *gameServer) Start() {
	nodeManange.Start(&ServerCfg, g.command)

	//启动网络
	gateNode := nodeManange.GetNode(fmt.Sprintf("gatewayNode_%v@127.0.0.1", ServerCfg.ServerID))
	if gateNode != nil {
		g.nw.Start(gateNode)
		//panic("节点启动失败")
	}
}

func (g *gameServer) Close() {
	for _, node := range nodeManange.GetNodes() {
		for _, process := range node.ProcessList() {
			process.Exit("server stop")
			process.Wait()
		}
		node.Stop()
		node.Wait()
	}
	g.nw.Close()
}

//StartGServer 启动game server
//go run main.go start --config=E:/worke/server/cfg.yaml
func StartGServer() {
	logrus.Infof("============================= Begin Start [%v][%v] ===============================", ServerCfg.ServerName, ServerCfg.ServerID)
	if level, err := logrus.ParseLevel(ServerCfg.Loglevel); err == nil {
		logger.Init(level, ServerCfg.LogWrite, ServerCfg.LogName, ServerCfg.LogPath)
	} else {
		logger.Init(logrus.InfoLevel, ServerCfg.LogWrite, ServerCfg.LogName, ServerCfg.LogPath)
	}

	//set pid file
	//file, _ := ioutil.TempFile("", fmt.Sprintf("pid_%v_%v_", ServerCfg.ServerName, ServerCfg.ServerID))
	filename := fmt.Sprintf("/tmp/pid_%v_%v", ServerCfg.ServerName, ServerCfg.ServerID)
	pidfile.SetPidfilePath(filename)
	if i, _ := pidfile.Read(); i != 0 {
		logrus.Warnf("服务已启动请检查或清除 进程id [%v] pidfile: [%v]  ", i, filename)
		return
	}

	// if ServerCfg.Daemon {
	//https://github.com/takama/daemon
	// }

	cfg.InitViperConfig(ServerCfg.CfgPath, ServerCfg.CfgType)
	if ServerCfg.WatchConfig {
		cfg.WatchConfig(ServerCfg.CfgPath, func(in fsnotify.Event) {
			logrus.Debug("Config file changed: [%v]  ", in.Name)
			cfg.InitViperConfig(ServerCfg.CfgPath, ServerCfg.CfgType)
		})
	}

	//启动定时器
	//timedtasks.StartCronTasks()
	// //定时器
	// timedtasks.AddTasks("loop", "* * * * * ?", func() {
	// 	logrus.Info("server time:", time.Now())
	// })
	//defer timedtasks.RemoveTasks("loop")

	db.StartMongodb(ServerCfg.Mongodb, ServerCfg.MongoConnStr)
	if ok, err := db.MongodbPing(); ok {
		logrus.Info("mongodb conn success")
	} else {
		panic(err)
	}

	db.StartRedis(ServerCfg.RedisConnStr, ServerCfg.RedisDB)
	if ok, err := db.RedisConn(); ok {
		logrus.Info("redis conn success")
	} else {
		panic(err)
	}

	if ServerCfg.OpenHTTP {
		go web.Start(ServerCfg.HTTPPort)
	}

	if ServerCfg.OpenPyroscope {
		profiler.Start(profiler.Config{
			ApplicationName: fmt.Sprintf("%v_%v", ServerCfg.ServerName, ServerCfg.ServerID),
			ServerAddress:   ServerCfg.PyroscopeHost,
		})
	}

	GameServerInfo = &gameServer{
		nw: network.NewNetWorkX(
			func() genServer.GateGenHanderInterface {
				return &clienconnect.Client{}
			},
			// &sync.Pool{
			// New: func() interface{} {
			// 	return &genServer.GateGenServer{}
			// }},
			ServerCfg.Port,
			ServerCfg.Packet,
			ServerCfg.Readtimeout,
			ServerCfg.NetType,
			ServerCfg.MsgTime,
			ServerCfg.MsgNum,
			func() { SendGameServerMsg("StartSuccess") },
			func() { db.RedisExec("del", "ConnectNumber") },
			func() { logrus.Info("connect number: ", db.RedisINCRBY("ConnectNumber", 1)) },
			func() { logrus.Info("connect number: ", db.RedisINCRBY("ConnectNumber", -1)) },
		),
		command: make(chan string),
	}

	GameServerInfo.Start()
	defer ClonseServer()
	defer GameServerInfo.Close()

	//退出消息监控
	var exitChan = make(chan os.Signal)

	if runtime.GOOS == "linux" {
		//signal.Notify(exitChan, os.Interrupt, os.Kill, syscall.SIGTERM)
		signal.Notify(exitChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGTSTP)
	} else {
		signal.Notify(exitChan, os.Interrupt)
	}

	for {
		select {
		case command := <-GameServerInfo.command:
			switch command {
			case "StartSuccess":
				pid := StartSuccess()
				logrus.Infof("====================== Start Game Server pid:[%v] Success =========================", pid)
			case "shutdown":
				logrus.Warn("Shut down the game server")
				return
			default:
				logrus.Warn("command:", command)
			}
		case s := <-exitChan:
			logrus.Info("收到信号: ", s)
			if runtime.GOOS == "linux" && s.String() == "quit" || s.String() == "terminated" {
				return
			} else if runtime.GOOS == "windows" && s.String() == "interrupt" {
				return
			}
			// case <-time.After(1 * time.Second):
			// 	logrus.Infof("time: [%v]  online:[%v]  [%v]", time.Now().Format(tools.DateTimeFormat), db.RedisGetInt("ConnectNumber"), GameServerInfo.nw.ConnectCount)
		}
	}

}

//SendGameServerMsg game system msg
func SendGameServerMsg(msg string) {
	GameServerInfo.command <- msg
}

//成功启动
func StartSuccess() int {
	pidfile.Write()
	logrus.Infof("pidfile :%v", pidfile.GetPidfilePath())
	i, _ := pidfile.Read()
	return i
}

//关闭服务
func ClonseServer() {
	logrus.Info("delete pidfile: ", pidfile.GetPidfilePath())
	os.Remove(pidfile.GetPidfilePath())
}

package genserver

import (
	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
	log "github.com/sirupsen/logrus"
)

//接收处理socket 发送过来的信息
// 处理玩家独立无交互的游戏逻辑
// 在socket中断后 此进程会保留一段时间以便于重新建立连接

type GateGenServer struct {
	ergo.GenServer
	process  *ergo.Process
	sendChan chan []byte
}

func (gateGS *GateGenServer) Unregister() {
	gateGS.process.Node.Unregister(gateGS.process.Name())
}

func (gateGS *GateGenServer) Register(name string) {
	gateGS.process.Node.Register(name, gateGS.process.Self())
}

type gateState struct {
}

// Init initializes process state using arbitrary arguments
// Init(...) -> state
func (gateGS *GateGenServer) Init(p *ergo.Process, args ...interface{}) interface{} {
	log.Infof("Init (%v): args %v ", p.Name(), args)
	gateGS.process = p
	gateGS.sendChan = args[0].(chan []byte)
	return gateState{}
}

// HandleCast serves incoming messages sending via gen_server:cast
// HandleCast -> ("noreply", state) - noreply
//		         ("stop", reason) - stop with reason
func (gateGS *GateGenServer) HandleCast(message etf.Term, state interface{}) (string, interface{}) {
	log.Infof("HandleCast (%v): %v", gateGS.process.Name(), message)

	switch info := message.(type) {
	case etf.Atom:
		switch info {
		case "stop":
			return "stop", "normal"
		case "SocketStop":
			return "stop", "normal"
		}
	case etf.Tuple:
		module := info[0]
		method := info[1]
		buf := info[2].([]byte)
		log.Debug("socket info ", module, method, buf)
		//gateGS.sendChan <- []byte("send msg test")
	case []byte:
		log.Debug("[]byte", info)

	}

	// switch message {
	// case etf.Atom("stop"):
	// 	return "stop", "normal"
	// case etf.List([]byte):
	// }
	return "noreply", state
}

// HandleCall serves incoming messages sending via gen_server:call
// HandleCall -> ("reply", message, state) - reply
//				 ("noreply", _, state) - noreply
//		         ("stop", reason, _) - normal stop
func (gateGS *GateGenServer) HandleCall(from etf.Tuple, message etf.Term, state interface{}) (string, etf.Term, interface{}) {
	log.Infof("HandleCall (%v): %v, From: %v", gateGS.process.Name(), message, from)

	reply := etf.Term(etf.Tuple{etf.Atom("error"), etf.Atom("unknown_request")})

	switch message {
	case etf.Atom("ping"):
		reply = etf.Term(etf.Atom("pong"))
	}
	return "reply", reply, state
}

// HandleInfo serves all another incoming messages (Pid ! message)
// HandleInfo -> ("noreply", state) - noreply
//		         ("stop", reason) - normal stop
func (gateGS *GateGenServer) HandleInfo(message etf.Term, state interface{}) (string, interface{}) {
	log.Infof("HandleInfo (%v): %v", gateGS.process.Name(), message)

	return "noreply", state
}

// Terminate called when process died
func (gateGS *GateGenServer) Terminate(reason string, state interface{}) {
	gateGS.Unregister()
	log.Infof("Terminate (%v): %v", gateGS.process.Name(), reason)
}

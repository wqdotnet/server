package genserver

import (
	"net"

	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

//接收处理socket 发送过来的信息
// 处理玩家独立无交互的游戏逻辑
// 在socket中断后 此进程会保留一段时间以便于重新建立连接

type GateGenServer struct {
	ergo.GenServer
	process *ergo.Process
}

type gateState struct {
}

func (dgs *GateGenServer) OnConnect(sendchan chan []byte, packet int32, msgchan chan []byte, addr net.Addr) {
}
func (dgs *GateGenServer) OnClose() {
}
func (dgs *GateGenServer) OnMessage(module int32, method int32, buf []byte) {
}
func (dgs *GateGenServer) Send(module int32, method int32, pb proto.Message) {
}
func (dgs *GateGenServer) ProcessMessage(msg []byte) bool {
	return false
}

// Init initializes process state using arbitrary arguments
// Init(...) -> state
func (dgs *GateGenServer) Init(p *ergo.Process, args ...interface{}) interface{} {
	log.Infof("Init (%v): args %v ", p.Name(), args)

	dgs.process = p
	return gateState{}
}

// HandleCast serves incoming messages sending via gen_server:cast
// HandleCast -> ("noreply", state) - noreply
//		         ("stop", reason) - stop with reason
func (dgs *GateGenServer) HandleCast(message etf.Term, state interface{}) (string, interface{}) {
	log.Infof("HandleCast (%v): %v", dgs.process.Name(), message)

	switch info := message.(type) {
	case etf.Atom:
		switch info {
		case "stop":
			return "stop", "normal"
		case "SocketStop":
			return "stop", "normal"
		}

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
func (dgs *GateGenServer) HandleCall(from etf.Tuple, message etf.Term, state interface{}) (string, etf.Term, interface{}) {
	log.Infof("HandleCall (%v): %v, From: %v", dgs.process.Name(), message, from)

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
func (dgs *GateGenServer) HandleInfo(message etf.Term, state interface{}) (string, interface{}) {
	log.Infof("HandleInfo (%v): %v", dgs.process.Name(), message)

	return "noreply", state
}

// Terminate called when process died
func (dgs *GateGenServer) Terminate(reason string, state interface{}) {
	log.Infof("Terminate (%v): %v", dgs.process.Name(), reason)
}

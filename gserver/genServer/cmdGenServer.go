package genserver

import (
	"fmt"
	"server/gserver/cfg"

	"github.com/facebookgo/pidfile"
	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
	log "github.com/sirupsen/logrus"
)

//命令服务 用于接收 外部发过来的服务命令

// GenServer implementation structure
type CmdGenServer struct {
	ergo.GenServer
	process       *ergo.Process
	CfgPath       string
	CfgType       string
	ServerCmdChan chan string
	ServerNmae    string
}

type cmdState struct {
}

// Init initializes process state using arbitrary arguments
// Init(...) -> state
func (dgs *CmdGenServer) Init(p *ergo.Process, args ...interface{}) interface{} {
	log.Infof("Init (%v): args %v ", p.Name(), args)
	dgs.process = p
	dgs.CfgPath = args[0].(string)
	dgs.CfgType = args[1].(string)
	dgs.ServerCmdChan = args[2].(chan string)
	dgs.ServerNmae = args[3].(string)

	return cmdState{}
}

// HandleCast serves incoming messages sending via gen_server:cast
// HandleCast -> ("noreply", state) - noreply
//		         ("stop", reason) - stop with reason
func (dgs *CmdGenServer) HandleCast(message etf.Term, state interface{}) (string, interface{}) {
	log.Infof("HandleCast (%v): %v", dgs.process.Name(), message)
	switch message {
	case etf.Atom("stop"):
		return "stop", "normal"
	}
	return "noreply", state
}

// HandleCall serves incoming messages sending via gen_server:call
// HandleCall -> ("reply", message, state) - reply
//				 ("noreply", _, state) - noreply
//		         ("stop", reason, _) - normal stop
func (dgs *CmdGenServer) HandleCall(from etf.Tuple, message etf.Term, state interface{}) (string, etf.Term, interface{}) {
	log.Infof("HandleCall (%v): %v ", dgs.process.Name(), message)
	reply := etf.Term(etf.Tuple{etf.Atom("error"), etf.Atom("unknown_request")})

	switch message {
	case etf.Atom("ping"):
		reply = etf.Atom(dgs.ServerNmae)
	case etf.Atom("ReloadCfg"):
		cfg.InitViperConfig(dgs.CfgPath, dgs.CfgType)
		reply = etf.Atom(dgs.ServerNmae)
	case etf.Atom("shutdown"):
		dgs.ServerCmdChan <- "shutdown"
		reply = etf.Atom(dgs.ServerNmae)
	case etf.Atom("state"):
		i, _ := pidfile.Read()
		reply = etf.Atom(fmt.Sprintf(" [%v] pid:[%v]", dgs.ServerNmae, i))
	}
	return "reply", reply, state
}

// HandleInfo serves all another incoming messages (Pid ! message)
// HandleInfo -> ("noreply", state) - noreply
//		         ("stop", reason) - normal stop
func (dgs *CmdGenServer) HandleInfo(message etf.Term, state interface{}) (string, interface{}) {
	log.Infof("HandleInfo (%v): %v", dgs.process.Name(), message)

	return "noreply", state
}

// Terminate called when process died
func (dgs *CmdGenServer) Terminate(reason string, state interface{}) {
	log.Infof("Terminate (%v): %v", dgs.process.Name(), reason)

}
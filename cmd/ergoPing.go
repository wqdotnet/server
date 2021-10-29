package cmd

import (
	"fmt"
	"server/gserver"

	"github.com/ergo-services/ergo"
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/ergo-services/ergo/node"
)

var (
	genServerName string
	gateNodeName  string
	process       *gen.ServerProcess
)

func call(cmd ...string) (etf.Term, error) {
	if len(cmd) == 1 {
		return process.Call(etf.Tuple{genServerName, gateNodeName}, etf.Atom(cmd[0]))
	} else {
		return process.Call(etf.Tuple{genServerName, gateNodeName}, cmd)
	}
}

// func cast(cmd ...string) {
// 	if len(cmd) == 1 {
// 		process.Cast(etf.Tuple{genServerName, gateNodeName}, etf.Atom(cmd[0]))
// 	} else {
// 		process.Cast(etf.Tuple{genServerName, gateNodeName}, cmd)
// 	}
// }

// func send(cmd ...string) {
// 	if len(cmd) == 1 {
// 		process.Send(etf.Tuple{genServerName, gateNodeName}, etf.Atom(cmd[0]))
// 	} else {
// 		process.Send(etf.Tuple{genServerName, gateNodeName}, cmd)
// 	}
// }

//"gatewayNode[serverid]@[ip]"
func ping(serverid, ip string) bool {
	_, process := startDebugGen("debug_server@127.0.0.1")
	genServerName = "cmdServer"
	gateNodeName = fmt.Sprintf("serverNode_%v@%v", serverid, ip)

	//process.Send(etf.Tuple{"gateServer", "demo@127.0.0.1"}, etf.Map{"abc": []byte("operation cwal")})

	if err := process.Send(etf.Tuple{genServerName, gateNodeName}, etf.Atom("ping")); err != nil {
		return false
	}
	return true

}

func startDebugGen(nodeName string) (node.Node, gen.Process) {
	opts := node.Options{
		ListenRangeBegin: uint16(gserver.ServerCfg.ListenRangeBegin),
		ListenRangeEnd:   uint16(gserver.ServerCfg.ListenRangeEnd),
		EPMDPort:         uint16(gserver.ServerCfg.EPMDPort),
	}
	node, _ := ergo.StartNode(nodeName, gserver.ServerCfg.Cookie, opts)
	// Spawn supervisor process
	process, _ := node.Spawn("deubg_gen", gen.ProcessOptions{}, &DebugGenServer{})

	return node, process
}

// GenServer implementation structure
type DebugGenServer struct {
	gen.Server
	process *gen.ServerProcess
}

type debugState struct {
}

// Init initializes process state using arbitrary arguments
// Init(...) -> state
func (dgs *DebugGenServer) Init(p *gen.ServerProcess, args ...interface{}) interface{} {
	dgs.process = p
	return debugState{}
}

// HandleCast serves incoming messages sending via gen_server:cast
// HandleCast -> ("noreply", state) - noreply
//		         ("stop", reason) - stop with reason
func (dgs *DebugGenServer) HandleCast(message etf.Term, state interface{}) (string, interface{}) {
	return "noreply", state
}

// HandleCall serves incoming messages sending via gen_server:call
// HandleCall -> ("reply", message, state) - reply
//				 ("noreply", _, state) - noreply
//		         ("stop", reason, _) - normal stop
func (dgs *DebugGenServer) HandleCall(from etf.Tuple, message etf.Term, state interface{}) (string, etf.Term, interface{}) {
	return "reply", etf.Term(etf.Atom("")), state
}

// HandleInfo serves all another incoming messages (Pid ! message)
// HandleInfo -> ("noreply", state) - noreply
//		         ("stop", reason) - normal stop
func (dgs *DebugGenServer) HandleInfo(message etf.Term, state interface{}) (string, interface{}) {
	return "noreply", state
}

// Terminate called when process died
func (dgs *DebugGenServer) Terminate(reason string, state interface{}) {
}

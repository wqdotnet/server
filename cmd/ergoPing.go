package cmd

import (
	"fmt"
	"server/gserver"

	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
)

var (
	genServerName string
	gateNodeName  string
	process       *ergo.Process
)

func call(cmd ...string) (etf.Term, error) {
	if len(cmd) == 1 {
		return process.Call(etf.Tuple{genServerName, gateNodeName}, etf.Atom(cmd[0]))
	} else {
		return process.Call(etf.Tuple{genServerName, gateNodeName}, cmd)
	}
}

func cast(cmd ...string) {
	if len(cmd) == 1 {
		process.Cast(etf.Tuple{genServerName, gateNodeName}, etf.Atom(cmd[0]))
	} else {
		process.Cast(etf.Tuple{genServerName, gateNodeName}, cmd)
	}
}

func send(cmd ...string) {
	if len(cmd) == 1 {
		process.Send(etf.Tuple{genServerName, gateNodeName}, etf.Atom(cmd[0]))
	} else {
		process.Send(etf.Tuple{genServerName, gateNodeName}, cmd)
	}
}

//"gatewayNode[serverid]@[ip]"
func connGenServer(serverid, ip string) bool {
	_, process = startDebugGen("debug_server@127.0.0.1")
	genServerName = "cmdServer"
	gateNodeName = fmt.Sprintf("serverNode_%v@%v", serverid, ip)

	fmt.Printf("---------- Console  connect  {'%v','%v'} ----------- \n", genServerName, gateNodeName)

	//process.Send(etf.Tuple{"gateServer", "demo@127.0.0.1"}, etf.Map{"abc": []byte("operation cwal")})
	if info, err := process.Call(etf.Tuple{genServerName, gateNodeName}, etf.Atom("ping")); err != nil {
		fmt.Printf("[pang]   err:[%v]\n", err)
		return false
	} else {
		fmt.Printf("[%v] \n", info)
		return true
	}
}

func startDebugGen(nodeName string) (*ergo.Node, *ergo.Process) {
	opts := ergo.NodeOptions{
		ListenRangeBegin: uint16(gserver.ServerCfg.ListenRangeBegin),
		ListenRangeEnd:   uint16(gserver.ServerCfg.ListenRangeEnd),
		EPMDPort:         uint16(gserver.ServerCfg.EPMDPort),
	}
	node := ergo.CreateNode(nodeName, gserver.ServerCfg.Cookie, opts)
	// Spawn supervisor process
	process, _ := node.Spawn("deubg_gen", ergo.ProcessOptions{}, &DebugGenServer{})
	return node, process
}

// GenServer implementation structure
type DebugGenServer struct {
	ergo.GenServer
	process *ergo.Process
}

type debugState struct {
}

// Init initializes process state using arbitrary arguments
// Init(...) -> state
func (dgs *DebugGenServer) Init(p *ergo.Process, args ...interface{}) interface{} {
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

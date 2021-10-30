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
	genServerName  string
	gateNodeName   string
	debugGenServer *DebugGenServer
)

func call(cmd ...string) (etf.Term, error) {
	if len(cmd) == 1 {

		return debugGenServer.process.Call(gen.ProcessID{Name: genServerName, Node: gateNodeName}, etf.Atom(cmd[0]))
	} else {
		return debugGenServer.process.Call(gen.ProcessID{Name: genServerName, Node: gateNodeName}, cmd)
	}
}

func ping(serverid, ip string) bool {
	_, process := startDebugGen("debug_server@127.0.0.1")
	genServerName = "cmdServer"
	gateNodeName = fmt.Sprintf("serverNode_%v@%v", serverid, ip)

	if err := process.Send(gen.ProcessID{Name: genServerName, Node: gateNodeName}, etf.Term("ping")); err != nil {
		fmt.Println(err)
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
	debugGenServer = &DebugGenServer{}
	// Spawn supervisor process
	process, _ := node.Spawn("deubg_gen", gen.ProcessOptions{}, debugGenServer)

	return node, process
}

// GenServer implementation structure
type DebugGenServer struct {
	gen.Server
	process *gen.ServerProcess
}

// Init initializes process state using arbitrary arguments
// Init(...) -> state
func (dgs *DebugGenServer) Init(process *gen.ServerProcess, args ...etf.Term) error {
	dgs.process = process
	return nil
}

// HandleCast serves incoming messages sending via gen_server:cast
// HandleCast -> ("noreply", state) - noreply
//		         ("stop", reason) - stop with reason
func (dgs *DebugGenServer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	return gen.ServerStatusOK
}

// HandleCall serves incoming messages sending via gen_server:call
// HandleCall -> ("reply", message, state) - reply
//				 ("noreply", _, state) - noreply
//		         ("stop", reason, _) - normal stop
func (dgs *DebugGenServer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	return etf.Term(""), gen.ServerStatusOK
}

// HandleInfo serves all another incoming messages (Pid ! message)
// HandleInfo -> ("noreply", state) - noreply
//		         ("stop", reason) - normal stop
func (dgs *DebugGenServer) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	return gen.ServerStatusOK
}

// Terminate called when process died
func (dgs *DebugGenServer) Terminate(process *gen.ServerProcess, reason string) {
}

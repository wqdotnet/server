package genserver

import (
	"fmt"

	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
)

// GenServer implementation structure
type GateWayGenServer struct {
	ergo.GenServer
	process *ergo.Process
}

type gateState struct {
}

// Init initializes process state using arbitrary arguments
// Init(...) -> state
func (dgs *GateWayGenServer) Init(p *ergo.Process, args ...interface{}) interface{} {
	fmt.Printf("Init (%s): args %v \n", p.Name(), args)
	dgs.process = p
	return dbState{}
}

// HandleCast serves incoming messages sending via gen_server:cast
// HandleCast -> ("noreply", state) - noreply
//		         ("stop", reason) - stop with reason
func (dgs *GateWayGenServer) HandleCast(message etf.Term, state interface{}) (string, interface{}) {
	fmt.Printf("HandleCast (%s): %#v\n", dgs.process.Name(), message)
	switch message {
	case etf.Atom("stop"):
		return "stop", "they said"
	}
	return "noreply", state
}

// HandleCall serves incoming messages sending via gen_server:call
// HandleCall -> ("reply", message, state) - reply
//				 ("noreply", _, state) - noreply
//		         ("stop", reason, _) - normal stop
func (dgs *GateWayGenServer) HandleCall(from etf.Tuple, message etf.Term, state interface{}) (string, etf.Term, interface{}) {
	fmt.Printf("HandleCall (%s): %#v, From: %#v\n", dgs.process.Name(), message, from)

	reply := etf.Term(etf.Tuple{etf.Atom("error"), etf.Atom("unknown_request")})

	switch message {
	case etf.Atom("hello"):
		reply = etf.Term(etf.Atom("hi"))
	}
	return "reply", reply, state
}

// HandleInfo serves all another incoming messages (Pid ! message)
// HandleInfo -> ("noreply", state) - noreply
//		         ("stop", reason) - normal stop
func (dgs *GateWayGenServer) HandleInfo(message etf.Term, state interface{}) (string, interface{}) {
	fmt.Printf("HandleInfo (%s): %#v\n", dgs.process.Name(), message)
	return "noreply", state
}

// Terminate called when process died
func (dgs *GateWayGenServer) Terminate(reason string, state interface{}) {
	fmt.Printf("Terminate (%s): %#v\n", dgs.process.Name(), reason)
}

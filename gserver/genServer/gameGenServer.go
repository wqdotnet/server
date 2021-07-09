package genserver

import (
	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
	log "github.com/sirupsen/logrus"
)

// GenServer implementation structure
type GameGenServer struct {
	ergo.GenServer
	process *ergo.Process
}

type gameGenState struct {
}

// Init initializes process state using arbitrary arguments
// Init(...) -> state
func (dgs *GameGenServer) Init(p *ergo.Process, args ...interface{}) interface{} {
	log.Info("Init (%s): args %v \n", p.Name(), args)
	dgs.process = p
	return gameGenState{}
}

// HandleCast serves incoming messages sending via gen_server:cast
// HandleCast -> ("noreply", state) - noreply
//		         ("stop", reason) - stop with reason
func (dgs *GameGenServer) HandleCast(message etf.Term, state interface{}) (string, interface{}) {
	log.Info("HandleCast (%s): %#v\n", dgs.process.Name(), message)
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
func (dgs *GameGenServer) HandleCall(from etf.Tuple, message etf.Term, state interface{}) (string, etf.Term, interface{}) {
	log.Info("HandleCall (%s): %#v, From: %#v\n", dgs.process.Name(), message, from)

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
func (dgs *GameGenServer) HandleInfo(message etf.Term, state interface{}) (string, interface{}) {
	log.Info("HandleInfo (%s): %#v\n", dgs.process.Name(), message)
	return "noreply", state
}

// Terminate called when process died
func (dgs *GameGenServer) Terminate(reason string, state interface{}) {
	log.Info("Terminate (%s): %#v\n", dgs.process.Name(), reason)
}

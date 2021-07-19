package genserver

import (
	"server/db"

	log "github.com/sirupsen/logrus"

	"github.com/halturin/ergo"
	"github.com/halturin/ergo/etf"
)

//数据落地服务

// GenServer implementation structure
type DbGenServer struct {
	ergo.GenServer
	process *ergo.Process
}

type dbState struct {
}

// Init initializes process state using arbitrary arguments
// Init(...) -> state
func (dgs *DbGenServer) Init(p *ergo.Process, args ...interface{}) interface{} {
	log.Infof("Init Gen_Server:(%v): args %v ", p.Name(), args)

	db.StartMongodb(args[0].(string), args[1].(string))
	db.StartRedis(args[2].(string), args[3].(int))

	dgs.process = p
	return dbState{}
}

// HandleCast serves incoming messages sending via gen_server:cast
// HandleCast -> ("noreply", state) - noreply
//		         ("stop", reason) - stop with reason
func (dgs *DbGenServer) HandleCast(message etf.Term, state interface{}) (string, interface{}) {
	log.Infof("HandleCast (%v): %v", dgs.process.Name(), message)
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
func (dgs *DbGenServer) HandleCall(from etf.Tuple, message etf.Term, state interface{}) (string, etf.Term, interface{}) {
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
func (dgs *DbGenServer) HandleInfo(message etf.Term, state interface{}) (string, interface{}) {
	log.Infof("HandleInfo (%v): %v", dgs.process.Name(), message)
	return "noreply", state
}

// Terminate called when process died
func (dgs *DbGenServer) Terminate(reason string, state interface{}) {
	log.Infof("Terminate (%v): %v", dgs.process.Name(), reason)
}

// //process.Send(etf.Tuple{GenServerName, NodeName}, etf.Atom("show me the money"))
// process.Send(etf.Tuple{"example", "demo@127.0.0.1"}, etf.Map{"abc": []byte("operation cwal")})

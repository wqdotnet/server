package genServer

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	log "github.com/sirupsen/logrus"
)

//游戏逻辑服务
type GameGenServer struct {
	gen.Server
	process *gen.ServerProcess
}

func (dgs *GameGenServer) Init(process *gen.ServerProcess, args ...etf.Term) error {
	log.Infof("Init (%v): args %v ", process.Name(), args)
	dgs.process = process
	return nil
}

func (dgs *GameGenServer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	log.Infof("HandleCast (%v): %v", dgs.process.Name(), message)

	return gen.ServerStatusOK
}

func (dgs *GameGenServer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	log.Infof("HandleCall (%v): %v, From: %v", dgs.process.Name(), message, from)

	reply := etf.Term(etf.Tuple{etf.Atom("error"), etf.Atom("unknown_request")})

	return reply, gen.ServerStatusOK
}

func (dgs *GameGenServer) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	log.Debugf("HandleInfo (%v): %v", dgs.process.Name(), message)

	return gen.ServerStatusOK
}

func (dgs *GameGenServer) Terminate(process *gen.ServerProcess, reason string) {
	log.Infof("Terminate (%v): %v", dgs.process.Name(), reason)
}

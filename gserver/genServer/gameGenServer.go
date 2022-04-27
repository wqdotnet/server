package genServer

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/sirupsen/logrus"
)

//中心服务
type GameGenServer struct {
	gen.Server
}

func (dgs *GameGenServer) Init(process *gen.ServerProcess, args ...etf.Term) error {
	logrus.Infof("Init (%v): args %v ", process.Name(), args)
	return nil
}

func (dgs *GameGenServer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	logrus.Infof("HandleCast (%v): %v", process.Name(), message)

	return gen.ServerStatusOK
}

func (dgs *GameGenServer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	logrus.Infof("HandleCall (%v): %v, From: %v", process.Name(), message, from)

	reply := etf.Term(etf.Tuple{etf.Atom("error"), etf.Atom("unknown_request")})

	return reply, gen.ServerStatusOK
}

func (dgs *GameGenServer) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	logrus.Debugf("HandleInfo (%v): %v", process.Name(), message)

	return gen.ServerStatusOK
}

func (dgs *GameGenServer) Terminate(process *gen.ServerProcess, reason string) {
	logrus.Infof("Terminate (%v): %v", process.Name(), reason)
}

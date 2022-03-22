package genServer

import (
	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/sirupsen/logrus"
)

//数据落地服务

type DbGenServer struct {
	gen.Server
	process *gen.ServerProcess
}

func (dgs *DbGenServer) Init(process *gen.ServerProcess, args ...etf.Term) error {
	logrus.Infof("Init (%v): args %v ", process.Name(), args)

	dgs.process = process
	return nil
}

func (dgs *DbGenServer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	logrus.Infof("HandleCast (%v): %v", dgs.process.Name(), message)

	return gen.ServerStatusOK
}

func (dgs *DbGenServer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	logrus.Infof("HandleCall (%v): %v, From: %v", dgs.process.Name(), message, from)

	reply := etf.Term(etf.Tuple{etf.Atom("error"), etf.Atom("unknown_request")})

	return reply, gen.ServerStatusOK
}

func (dgs *DbGenServer) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	logrus.Infof("HandleInfo (%v): %v", dgs.process.Name(), message)

	return gen.ServerStatusOK
}

func (dgs *DbGenServer) Terminate(process *gen.ServerProcess, reason string) {

	logrus.Infof("Terminate (%v): %v", dgs.process.Name(), reason)
}

package genServer

import (
	"fmt"
	"server/gserver/cfg"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/facebookgo/pidfile"
	"github.com/sirupsen/logrus"
)

//命令服务 用于接收 外部发过来的服务命令

// GenServer implementation structure
type CmdGenServer struct {
	gen.Server
	process       *gen.ServerProcess
	CfgPath       string
	CfgType       string
	ServerCmdChan chan string
	ServerNmae    string
}

func (dgs *CmdGenServer) Init(process *gen.ServerProcess, args ...etf.Term) error {

	logrus.Infof("Init (%v): args %v ", process.Name(), args)
	dgs.process = process
	dgs.CfgPath = args[0].(string)
	dgs.CfgType = args[1].(string)
	dgs.ServerCmdChan = args[2].(chan string)
	dgs.ServerNmae = args[3].(string)

	return nil
}

func (dgs *CmdGenServer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	logrus.Infof("HandleCast (%v): %v", dgs.process.Name(), message)
	// switch message {
	// case etf.Atom("stop"):
	// 	return gen.ServerStatusStopWithReason("stop normal")
	// }
	return gen.ServerStatusOK
}

func (gd *CmdGenServer) HandleDirect(process *gen.ServerProcess, message interface{}) (interface{}, error) {

	return nil, nil
}

func (dgs *CmdGenServer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	logrus.Infof("HandleCall (%v): %v ", dgs.process.Name(), message)
	reply := etf.Term(etf.Tuple{etf.Atom("error"), etf.Atom("unknown_request")})

	switch message {
	case etf.Atom("ping"):
		reply = etf.Atom(dgs.ServerNmae)
	case etf.Atom("ReloadCfg"):
		cfg.InitViperConfig(dgs.CfgPath, dgs.CfgType)
		reply = etf.Atom("ReloadCfg ok")
	case etf.Atom("shutdown"):
		dgs.ServerCmdChan <- "shutdown"
		reply = etf.Atom(dgs.ServerNmae)
	case etf.Atom("state"):
		i, _ := pidfile.Read()
		reply = etf.Atom(fmt.Sprintf(" [%v] pid:[%v]", dgs.ServerNmae, i))
	default:
		logrus.Debug("info:", message)
	}
	return reply, gen.ServerStatusOK
}

func (dgs *CmdGenServer) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	logrus.Infof("HandleInfo (%v): %v", dgs.process.Name(), message)

	return gen.ServerStatusOK
}

func (dgs *CmdGenServer) Terminate(process *gen.ServerProcess, reason string) {
	logrus.Infof("Terminate (%v): %v", dgs.process.Name(), reason)

}

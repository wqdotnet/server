package genServer

import (
	"server/db"
	"server/gserver/commonstruct"
	"time"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/sirupsen/logrus"
)

//数据落地服务

type DbGenServer struct {
	gen.Server
}

func (dgs *DbGenServer) Init(process *gen.ServerProcess, args ...etf.Term) error {
	logrus.Infof("Init (%v): args %v ", process.Name(), args)
	process.SendAfter(process.Self(), etf.Atom("loop"), time.Minute*10)
	return nil
}

func (dgs *DbGenServer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	logrus.Infof("HandleCast (%v): %v", process.Name(), message)

	return gen.ServerStatusOK
}

func (dgs *DbGenServer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	logrus.Infof("HandleCall (%v): %v, From: %v", process.Name(), message, from)

	reply := etf.Term(etf.Tuple{etf.Atom("error"), etf.Atom("unknown_request")})

	return reply, gen.ServerStatusOK
}

func (dgs *DbGenServer) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	logrus.Debugf("HandleInfo (%v): %v", process.Name(), message)
	switch info := message.(type) {
	case etf.Atom:
		switch info {
		case "loop":
			process.SendAfter(process.Self(), etf.Atom("loop"), time.Minute*10)
			loop()
		}
	}

	return gen.ServerStatusOK
}

func (dgs *DbGenServer) Terminate(process *gen.ServerProcess, reason string) {
	logrus.Infof("Terminate (%v): %v", process.Name(), reason)
}

func loop() {
	commonstruct.RangeAllData(func(rd *commonstruct.RoleData) (issave bool) {
		if rd.RoleBase.DirtyData {
			logrus.Debug("保存数据", rd.RoleBase.Name, rd.RoleBase.RoleID)
			if err := db.InsertOne(db.RoleBaseTable, rd.RoleBase); err == nil {
				rd.RoleBase.DirtyData = false
				issave = true
			}
		}

		if rd.RoleItems.DirtyData {
			logrus.Debug("保存数据", rd.RoleBase.Name, rd.RoleBase.RoleID)
			if err := db.InsertOne(db.RoleItemsTable, rd.RoleItems); err == nil {
				rd.RoleItems.DirtyData = false
				issave = true
			}
		}

		return issave
	})
}

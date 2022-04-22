package genServer

import (
	"runtime"
	"time"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/sirupsen/logrus"
)

//接收处理socket 发送过来的信息
// 处理玩家独立无交互的游戏逻辑
// 在socket中断后 此进程会保留一段时间以便于重新建立连接

type GateGenServer struct {
	gen.Server
	sendChan     chan []byte
	clientHander GateGenHanderInterface
}

func (gateGS *GateGenServer) Init(process *gen.ServerProcess, args ...etf.Term) error {
	logrus.Infof("Init (%v): args %v ", process.Name(), args)
	gateGS.sendChan = args[0].(chan []byte)
	gateGS.clientHander = args[1].(GateGenHanderInterface)
	gateGS.clientHander.InitHander(process, gateGS.sendChan)

	process.SendAfter(process.Self(), etf.Atom("loop"), time.Second)
	return nil
}

func (gateGS *GateGenServer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	//logrus.Infof("gateGen HandleCast (%v): %v", process.Name(), message)
	defer func() {
		if err := recover(); err != nil {
			pc, fn, line, _ := runtime.Caller(5)
			logrus.Errorf("process:[%v] funcname:[%v] fn:[%v] line:[%v]", process.Name(), runtime.FuncForPC(pc).Name(), fn, line)
		}
	}()

	switch info := message.(type) {
	case etf.Atom:
		switch info {
		case "SocketStop":
			return gen.ServerStatusStopWithReason("stop normal")
		case "timeloop":
			logrus.Debug("time loop")
		}
	case etf.Tuple:
		module := info[0].(int32)
		method := info[1].(int32)
		buf := info[2].([]byte)
		gateGS.clientHander.MsgHander(module, method, buf)
		//gateGS.sendChan <- []byte("send msg test")
	case []byte:
		logrus.Debug("[]byte:", info)
	}
	return gateGS.clientHander.GenServerStatus()
}

func (gateGS *GateGenServer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	logrus.Infof("HandleCall (%v): %v, From: %v", process.Name(), message, from)

	switch info := message.(type) {
	case etf.Atom:
		switch info {
		case "Extrusionline": //挤下线
			gateGS.clientHander.Terminate("Extrusionline")
			return etf.Term("ignore"), gen.ServerStatusStop
		}
	}
	gateGS.clientHander.HandleCall(message)

	reply := etf.Atom("ignore")
	return reply, gateGS.clientHander.GenServerStatus()
}

func (gateGS *GateGenServer) HandleInfo(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	switch info := message.(type) {
	case etf.Atom:
		switch info {
		case "loop":
			after := gateGS.clientHander.LoopHander()
			if after < time.Millisecond {
				after = time.Second
			}
			process.SendAfter(process.Self(), etf.Atom("loop"), after)
			return gateGS.clientHander.GenServerStatus()
		case "stop":
			return gen.ServerStatusStop
		}
	}

	gateGS.clientHander.HandleInfo(message)
	return gateGS.clientHander.GenServerStatus()
}

// Terminate called when process died
func (gateGS *GateGenServer) Terminate(process *gen.ServerProcess, reason string) {
	//logrus.Infof("Terminate (%v): %v", process.Name(), reason)
	gateGS.clientHander.Terminate(reason)
}

// // //Send 发送消息
// func (gateGS *GateGenServer) Send(module int32, method int32, pb proto.Message) {
// 	//logrus.Debugf("client send msg [%v] [%v] [%v]", module, method, pb)
// 	data, err := proto.Marshal(pb)
// 	if err != nil {
// 		logrus.Errorf("proto encode error[%v] [%v][%v] [%v]", err.Error(), module, method, pb)
// 		return
// 	}
// 	// msginfo := &common.NetworkMsg{}
// 	// msginfo.Module = module
// 	// msginfo.Method = method
// 	// msginfo.MsgBytes = data
// 	// msgdata, err := proto.Marshal(msginfo)
// 	// if err != nil {
// 	// 	logrus.Errorf("msg encode error[%s]\n", err.Error())
// 	// }
// 	// gateGS.sendChan <- msgdata

// 	mldulebuf := tools.IntToBytes(module, 2)
// 	methodbuf := tools.IntToBytes(method, 2)
// 	gateGS.sendChan <- tools.BytesCombine(mldulebuf, methodbuf, data)

// }

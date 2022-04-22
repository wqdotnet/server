package genServer

import (
	"time"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
)

type GateGenHanderInterface interface {
	InitHander(process *gen.ServerProcess, sendChan chan []byte)
	LoopHander() (nextloop time.Duration)
	MsgHander(module, method int32, buf []byte)
	HandleCall(message etf.Term)
	HandleInfo(message etf.Term)
	Terminate(reason string)
	GenServerStatus() gen.ServerStatus
}

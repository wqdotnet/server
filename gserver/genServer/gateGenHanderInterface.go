package genServer

import "time"

type GateGenHanderInterface interface {
	InitHander(sendChan chan []byte)
	LoopHander() (nextloop time.Duration)
	MsgHander(module, method int32, buf []byte)
}

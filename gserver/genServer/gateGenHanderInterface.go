package genServer

type GateGenHanderInterface interface {
	InitHander(sendChan chan []byte)
	MsgHander(module, method int32, buf []byte)
}

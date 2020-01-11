package netwrok

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

//NetworkInterface network
type NetworkInterface interface {
	Start()
	Stop()
	Send(msg *NetworkMsg)
}

//NetworkMsg   tcp udp send/receive msg
type NetworkMsg struct {
	// recvbuf []byte
	// bufptr  []byte
	Module int
	method int
	buf    []byte
}

//Send send msg
func Send(c *NetworkInterface, module int, method int, pb proto.Message) {

	// encode
	data, err := proto.Marshal(pb)
	if err != nil {
		fmt.Printf("proto encode error[%s]\n", err.Error())
		return
	}

	msg := &NetworkMsg{}
	msg.buf = data
	msg.Module = module
	msg.method = method

}

package network

import (
	"fmt"
)

//TCPNetwork tcp/ip
type TCPNetwork struct {
	laddr string
}

//Start NetworkInterface.Start
func (c TCPNetwork) Start() {
	fmt.Printf("TcpNetwork start")

}

//Stop NetworkInterface.Stop
func (c TCPNetwork) Stop() {
	fmt.Printf("TcpNetwork Stop ")
}

//Send network sendmsg
func (c TCPNetwork) Send(msg []byte) {
	fmt.Printf("TcpNetwork send ")
	//outpb proto.Message
	//EncodeSend(1,1,outpb)
}

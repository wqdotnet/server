package netwrok

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

//TcpNetwork tcp/ip
type TcpNetwork struct {
	laddr string
}

//Start NetworkInterface.Start
func Start(c *TcpNetwork) {
	fmt.Printf("TcpNetwork start")

}

//Stop NetworkInterface.Stop
func Stop(c *TcpNetwork) {
	fmt.Printf("TcpNetwork Stop ")
}

//Send network sendmsg
func Send(c *TcpNetwork, outpb proto.Message) {
	fmt.Printf("TcpNetwork send ")

	//EncodeSend(1,1,outpb)
}

package network

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

//TCPNetwork tcp/ip
type TCPNetwork struct {
	network string
	address string
}

//Start NetworkInterface.Start
func (c TCPNetwork) Start() {
	fmt.Printf("TcpNetwork start")
	tcpServer, _ := net.ResolveTCPAddr("tcp4", ":8080")
	listener, _ := net.ListenTCP("tcp", tcpServer)

	for {
		//当有新的客户端请求来的时候，拿到与客户端的连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		//处理逻辑
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//读取客户端传送的消息
	go func() {
		response, _ := ioutil.ReadAll(conn)
		fmt.Println(string(response))
	}()

	//向客户端发送消息
	time.Sleep(1 * time.Second)
	now := time.Now().String()
	conn.Write([]byte(now))
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

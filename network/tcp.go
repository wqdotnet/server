package network

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
)

//TCPNetwork tcp/ip
type TCPNetwork struct {
	network string
	address string
}

type innerBuffer []byte

func (in *innerBuffer) readN(n int) (buf []byte, err error) {
	if n <= 0 {
		return nil, errors.New("zero or negative length is invalid")
	} else if n > len(*in) {
		return nil, errors.New("exceeding buffer length")
	}
	buf = (*in)[:n]
	*in = (*in)[n:]
	return
}

//Start NetworkInterface.Start
func (c TCPNetwork) Start() {
	fmt.Println("TcpNetwork start")
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
	//defer conn.Close()
	fmt.Println("socket handle")
	//读取客户端传送的消息
	go func() {
		var (
			response innerBuffer
			header   []byte
			err      error
		)
		response, _ = ioutil.ReadAll(conn)
		fmt.Println(string(response))

		header, err = response.readN(2)
		fmt.Println(header, err)

	}()

	// //向客户端发送消息
	// time.Sleep(1 * time.Second)
	// now := time.Now().String()
	// conn.Write([]byte(now))
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

package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
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

func bytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.LittleEndian, &data)
	return int(data)
}

//Start start
func (c *TCPNetwork) Start(nw *NetWorkx) {
	fmt.Println(fmt.Sprintf("tcp run on localhost: [%v]", nw.Port))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", nw.Port))
	defer listener.Close()
	checkError(err)
	//go func() {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		//c := nw.UserPool.Get().(ClientInterface)
		go nw.HandleClient(conn)
		//go handleClient(conn, nw.Packet, c)
	}
	//}()
	//select {}
}

// func handleClient(conn net.Conn, packet int32, client ClientInterface) {
// 	// close connection on exit
// 	defer conn.Close()
// 	defer client.OnClose()

// 	client.OnConnect()
// 	// var buffer = make([]byte, 1024, 1024)
// 	//var oneRead innerBuffer
// 	// var n int
// 	// var e error
// 	for {
// 		_, buf, e := UnpackToBlockFromReader(conn, packet)
// 		if e != nil {
// 			fmt.Println("socket error:", e.Error())
// 			return
// 		}

// 		//[:n]
// 		client.OnMessage(1, 1, buf[packet:])

// 		// oneRead = buf
// 		// _, _ = oneRead.readN(int(packet))
// 		// client.OnMessage(1, 2, oneRead)

// 		//next 消息处理
// 		// _, err2 := conn.Write(NewByte(1, 2, 3, 4, 5, 6, 7, 8, 9))
// 		// if err2 != nil {
// 		// 	fmt.Println(err2.Error())
// 		// 	return
// 		// }
// 	}
// }

// func NewByte(byts ...byte) []byte {
// 	var rs = make([]byte, 0, 512)
// 	rs = append(rs, byts...)
// 	return rs
// }

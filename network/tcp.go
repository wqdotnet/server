package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/fwhezfwhez/errorx"
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

func readUntil(reader io.Reader, buf []byte) error {
	if len(buf) == 0 {
		return nil
	}
	var offset int
	for {
		n, e := reader.Read(buf[offset:])
		if e != nil {
			if e == io.EOF {
				return e
			}
			return errorx.Wrap(e)
		}
		offset += n
		if offset >= len(buf) {
			break
		}
	}
	return nil
}

//Start start
func (c *TCPNetwork) Start(nw *NetWorkx) {
	fmt.Println(fmt.Sprintf("tcp run on localhost: [%v]", nw.Port))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", nw.Port))
	defer listener.Close()
	checkError(err)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			c := nw.UserPool.Get().(ClientInterface)
			go handleClient(conn, c)
		}
	}()
	select {}
}

func handleClient(conn net.Conn, client ClientInterface) {
	// close connection on exit
	defer conn.Close()
	defer client.OnClose()

	client.OnConnect()
	var oneRead innerBuffer
	var e error
	for {
		oneRead, e = readOnce(conn)
		if e != nil {
			if e == io.EOF {
				break
			}
			fmt.Println(errorx.Wrap(e).Error())
			return
		}
		fmt.Println("socket []byte len: ", len(oneRead))

		buf, err := oneRead.readN(2)
		if err != nil {
			fmt.Println("socket error:", err)
		}

		head := binary.BigEndian.Uint16(buf)
		fmt.Println(fmt.Sprintf("head: [%v]  body： [%v]", int(head), len(oneRead)))
		// if int(blen) == len(buf) {

		// }

		client.OnMessage(1, 2, oneRead)
		fmt.Println(fmt.Sprintf("receive from client: %v", string(oneRead)))

		//next 消息处理

		// _, err2 := conn.Write(NewByte(1, 2, 3, 4, 5, 6, 7, 8, 9))
		// if err2 != nil {
		// 	fmt.Println(err2.Error())
		// 	return
		// }
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readOnce(reader io.Reader) ([]byte, error) {
	var buffer = make([]byte, 512, 512)
	var n int
	var e error

	n, e = reader.Read(buffer)
	if e != nil {
		return nil, e
	}

	return buffer[0:n], nil
}

// func NewByte(byts ...byte) []byte {
// 	var rs = make([]byte, 0, 512)
// 	rs = append(rs, byts...)
// 	return rs
// }

//Stop NetworkInterface.Stop
func (c *TCPNetwork) Stop() {
	fmt.Printf("TcpNetwork Stop ")
}

//Send network sendmsg
func (c *TCPNetwork) Send(msg []byte) {
	fmt.Printf("TcpNetwork send ")
	//outpb proto.Message
	//EncodeSend(1,1,outpb)
}

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
			go handleClient(conn, nw.Packet, c)
		}
	}()
	select {}
}

func handleClient(conn net.Conn, packet int32, client ClientInterface) {
	// close connection on exit
	defer conn.Close()
	defer client.OnClose()

	client.OnConnect()
	// var buffer = make([]byte, 1024, 1024)
	var oneRead innerBuffer
	// var n int
	// var e error
	for {
		buf, e := UnpackToBlockFromReader(conn, packet)
		if e != nil {
			fmt.Println("socket error:", e.Error())
			return
		}

		client.OnMessage(1, 1, buf)

		oneRead = buf
		_, _ = oneRead.readN(2)

		client.OnMessage(1, 2, oneRead)

		//next 消息处理
		// _, err2 := conn.Write(NewByte(1, 2, 3, 4, 5, 6, 7, 8, 9))
		// if err2 != nil {
		// 	fmt.Println(err2.Error())
		// 	return
		// }
	}
}

// UnpackToBlockFromReader -> unpack the first block from the reader.
// protocol is PackWithMarshaller().
// [4]byte -- length             fixed_size,binary big endian encode
// [4]byte -- messageID          fixed_size,binary big endian encode
// [4]byte -- headerLength       fixed_size,binary big endian encode
// [4]byte -- bodyLength         fixed_size,binary big endian encode
// []byte -- header              marshal by json
// []byte -- body                marshal by marshaller
// ussage:
// for {
//     blockBuf, e:= UnpackToBlockFromReader(reader)
// 	   go func(buf []byte){
//         // handle a message block apart
//     }(blockBuf)
//     continue
// }
func UnpackToBlockFromReader(reader io.Reader, packet int32) ([]byte, error) {
	if reader == nil {
		return nil, errors.New("reader is nil")
	}
	var info = make([]byte, packet, packet)
	if e := readUntil(reader, info); e != nil {
		if e == io.EOF {
			return nil, e
		}
		return nil, errorx.Wrap(e)
	}

	length, e := LengthOf(info, packet)
	if e != nil {
		return nil, e
	}
	var content = make([]byte, length, length)
	if e := readUntil(reader, content); e != nil {
		if e == io.EOF {
			return nil, e
		}
		return nil, errorx.Wrap(e)
	}

	return append(info, content...), nil
}

//LengthOf Length of the stream starting validly.
//Length doesn't include length flag itself, it refers to a valid message length after it.
func LengthOf(stream []byte, packet int32) (int32, error) {
	if len(stream) < int(packet) {
		return 0, errors.New(fmt.Sprint("stream lenth should be bigger than ", packet))
	}

	switch packet {
	case 2:
		return int32(binary.BigEndian.Uint16(stream[0:2])), nil
	case 4:
		return int32(binary.BigEndian.Uint32(stream[0:4])), nil
	default:
		errstr := fmt.Sprintf("stream lenth seting error  [packet: %v]", packet)
		return 0, errors.New(errstr)
	}

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

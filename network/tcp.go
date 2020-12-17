package network

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

//TCPNetwork tcp/ip
type TCPNetwork struct {
	// network string
	// address string
}

// type innerBuffer []byte

// func (in *innerBuffer) readN(n int) (buf []byte, err error) {
// 	if n <= 0 {
// 		return nil, errors.New("zero or negative length is invalid")
// 	} else if n > len(*in) {
// 		return nil, errors.New("exceeding buffer length")
// 	}
// 	buf = (*in)[:n]
// 	*in = (*in)[n:]
// 	return
// }

// func bytesToInt(bys []byte) int {
// 	bytebuff := bytes.NewBuffer(bys)
// 	var data int64
// 	binary.Read(bytebuff, binary.LittleEndian, &data)
// 	return int(data)
// }

//Start start
func (c *TCPNetwork) Start(nw *NetWorkx) {
	log.Info(fmt.Sprintf("tcp run on localhost: [%v]", nw.Port))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", nw.Port))
	defer listener.Close()
	checkError(err)
	//go func() {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error(err.Error())
			break
		}
		go nw.HandleClient(conn)
	}
	//}()
	//select {}
}

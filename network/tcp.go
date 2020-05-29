package network

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"github.com/panjf2000/gnet"
)

type echoServer struct {
	*gnet.EventServer
}

func (es *echoServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}
func (es *echoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// Echo synchronously.
	out = frame
	return

	/*
		// Echo asynchronously.
		data := append([]byte{}, frame...)
		go func() {
			time.Sleep(time.Second)
			c.AsyncWrite(data)
		}()
		return
	*/
}

func main() {
	var port int
	var multicore bool

	// Example command: go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()
	echo := new(echoServer)
	log.Fatal(gnet.Serve(echo, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore)))
}

//TCPNetwork tcp/ip
type TCPNetwork struct {
	network string
	address string
}

//Start NetworkInterface.Start
func (c TCPNetwork) Start() {
	fmt.Println("TcpNetwork start")

	// gnet.Serve(c, fmt.Sprintf("tcp://:%d", "port"), gnet.WithMulticore(true))

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

		response, _ := ioutil.ReadAll(conn)
		fmt.Println(string(response))
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

package network

import (
	"crypto/sha1"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
)

//KCPNetwork  kcp
type KCPNetwork struct {
}

//Start start
func (c *KCPNetwork) Start(nw *NetWorkx) {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)

	connstr := fmt.Sprintf("127.0.0.1:%v", nw.Port)
	if listener, err := kcp.ListenWithOptions(connstr, block, 10, 3); err == nil {
		if nw.StartHook != nil {
			nw.StartHook()
		}

		for {
			s, err := listener.AcceptKCP()
			if err != nil {
				logrus.Fatal(err)
			}
			go nw.HandleClient(s)
		}

	} else {
		logrus.Fatal(err)
	}
}

//Close 关闭
func (c *KCPNetwork) Close() {

}

//demo :
// func startDemo() {
// 	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
// 	block, _ := kcp.NewAESBlockCrypt(key)
// 	if listener, err := kcp.ListenWithOptions("127.0.0.1:12345", block, 10, 3); err == nil {
// 		// spin-up the client
// 		go client()
// 		for {
// 			s, err := listener.AcceptKCP()
// 			if err != nil {
// 				logrus.Fatal(err)
// 			}
// 			go handleEcho(s)
// 		}
// 	} else {
// 		logrus.Fatal(err)
// 	}
// }

// // handleEcho send back everything it received
// func handleEcho(conn *kcp.UDPSession) {
// 	buf := make([]byte, 4096)
// 	for {
// 		n, err := conn.Read(buf)
// 		if err != nil {
// 			logrus.Println(err)
// 			return
// 		}

// 		n, err = conn.Write(buf[:n])
// 		if err != nil {
// 			logrus.Println(err)
// 			return
// 		}
// 	}
// }

// func client() {
// 	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
// 	block, _ := kcp.NewAESBlockCrypt(key)

// 	// wait for server to become ready
// 	time.Sleep(time.Second)

// 	// dial to the echo server
// 	if sess, err := kcp.DialWithOptions("127.0.0.1:12345", block, 10, 3); err == nil {
// 		for {
// 			data := time.Now().String()
// 			buf := make([]byte, len(data))
// 			logrus.Println("sent:", data)
// 			if _, err := sess.Write([]byte(data)); err == nil {
// 				// read back the data
// 				if _, err := io.ReadFull(sess, buf); err == nil {
// 					logrus.Println("recv:", string(buf))
// 				} else {
// 					logrus.Fatal(err)
// 				}
// 			} else {
// 				logrus.Fatal(err)
// 			}
// 			time.Sleep(time.Second)
// 		}
// 	} else {
// 		logrus.Fatal(err)
// 	}
// }

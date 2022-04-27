// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net/http"
	"server/network"
	"server/tools"
	"sync/atomic"
	"time"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	// Time allowed to write a message to the peer.
	writeWait = 30 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// 客户端连接
func WsClient(hub *Hub, context *gin.Context, nw *network.NetWorkx) {
	upGrande := websocket.Upgrader{
		//设置允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		//设置请求协议
		Subprotocols: []string{context.GetHeader("Sec-WebSocket-Protocol")},
		// ReadBufferSize:  1024,
		// WriteBufferSize: 1024,
	}
	//创建连接
	conn, err := upGrande.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		context.JSON(51001, gin.H{
			"websocket connect error": context.Param("channel"),
		})
		return
	}

	process, clientHander, sendchan, err := nw.CreateProcess()
	defer nw.UserPool.Put(clientHander)

	wsclient := &wsClient{hub: hub, conn: conn, send: sendchan, testsend: make(chan []byte, 1)}
	wsclient.hub.register <- wsclient

	pongWait = time.Second * time.Duration(nw.Readtimeout)

	atomic.AddInt32(&nw.ConnectCount, 1)
	defer atomic.AddInt32(&nw.ConnectCount, -1)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go wsclient.writePump(nw.Packet)
	wsclient.readPump(process, nw.Packet)
}

// wsClient is a middleman between the websocket connection and the hub.
type wsClient struct {
	hub *Hub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte

	testsend chan []byte
}

func (c *wsClient) readPump(process gen.Process, packet int32) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(int64(maxMessageSize))
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	defer process.Send(process.Self(), etf.Term(etf.Tuple{etf.Atom("$gen_cast"), etf.Atom("SocketStop")}))

	for {
		messageType, messagebuf, err := unpackToBlockFromReader(c.conn, packet, []byte{})
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("WebSocket closed  error: %v", err.Error())
			}
			return
		}

		if messageType == 1 {
			messagebuf = bytes.TrimSpace(bytes.Replace(messagebuf, newline, space, -1))
			c.hub.broadcast <- messagebuf
		} else {
			if len(messagebuf) < 4 {
				logrus.Debug("buf len:", len(messagebuf))
				return
			}
			module := int32(binary.BigEndian.Uint16(messagebuf[packet : packet+2]))
			method := int32(binary.BigEndian.Uint16(messagebuf[packet+2 : packet+4]))
			process.Send(process.Self(), etf.Term(etf.Tuple{etf.Atom("$gen_cast"), etf.Tuple{module, method, messagebuf[packet+4:]}}))
		}
	}
}

func (c *wsClient) writePump(packet int32) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}

			le := tools.IntToBytes(int32(len(message)), packet)
			buf := tools.BytesCombine(le, message)
			w.Write(buf)

			if err := w.Close(); err != nil {
				return
			}
		case message, ok := <-c.testsend:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.testsend)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.testsend)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func unpackToBlockFromReader(conn *websocket.Conn, packet int32, lastbyte []byte) (int, []byte, error) {
	messageType, buf, err := conn.ReadMessage()
	buf = tools.BytesCombine(lastbyte, buf)
	if err != nil {
		return 0, nil, err
	}
	switch messageType {
	case 1:
		return messageType, buf, nil
	case 2:
		if len(buf) <= int(packet) {
			return messageType, nil, errors.New("packet size error")
		}
		lenght, err := LengthOf(buf, packet)
		if err != nil {
			return messageType, nil, err
		}

		if int(lenght) < len(buf[packet:]) {
			return unpackToBlockFromReader(conn, packet, buf[lenght:])
		} else {
			return messageType, buf, nil
		}
	}

	return 0, nil, nil
}

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

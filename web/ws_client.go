// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

import (
	"bytes"
	"encoding/binary"
	"net/http"
	"server/network"
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

	wsclient := &wsClient{hub: hub, conn: conn, send: sendchan}
	wsclient.hub.register <- wsclient

	pongWait = time.Second * time.Duration(nw.Readtimeout)

	atomic.AddInt32(&nw.ConnectCount, 1)
	defer atomic.AddInt32(&nw.ConnectCount, -1)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go wsclient.writePump()
	wsclient.readPump(process)
}

// wsClient is a middleman between the websocket connection and the hub.
type wsClient struct {
	hub *Hub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *wsClient) readPump(process gen.Process) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(int64(maxMessageSize))
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	defer process.Send(process.Self(), etf.Term(etf.Tuple{etf.Atom("$gen_cast"), etf.Atom("SocketStop")}))

	for {
		_, messagebuf, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Errorf("WebSocket closed  error: %v", err.Error())
			} else {
				logrus.Debugf("WebSocket closed: [%v]", err.Error())
			}
			return
		}
		messagebuf = bytes.TrimSpace(bytes.Replace(messagebuf, newline, space, -1))

		if len(messagebuf) < int(4) {
			logrus.Debug("buf len:", len(messagebuf))
			return
		}

		module := int32(binary.BigEndian.Uint16(messagebuf[:2]))
		method := int32(binary.BigEndian.Uint16(messagebuf[2:4]))
		process.Send(process.Self(), etf.Term(etf.Tuple{etf.Atom("$gen_cast"), etf.Tuple{module, method, messagebuf[4:]}}))

		//c.hub.broadcast <- messagebuf
	}
}

func (c *wsClient) writePump() {
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
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
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

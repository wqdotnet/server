// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package web

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*wsClient]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *wsClient

	// Unregister requests from clients.
	unregister chan *wsClient
}

var WSHub = &Hub{
	broadcast:  make(chan []byte),
	register:   make(chan *wsClient),
	unregister: make(chan *wsClient),
	clients:    make(map[*wsClient]bool),
}

// func newHub() *Hub {
// 	return &Hub{
// 		broadcast:  make(chan []byte),
// 		register:   make(chan *Client),
// 		unregister: make(chan *Client),
// 		clients:    make(map[*Client]bool),
// 	}
// }

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

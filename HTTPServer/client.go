package httpserver

import (
	"github.com/gorilla/websocket"
)

//client is one chat user
type client struct {
	//client is websocket
	socket *websocket.Conn
	//send is message channel
	send chan Message
	//chan is client chat room
	connect *connect
}

//client message json struct
type Message struct {
	Type string `json:"Type"`
	Data string `json:"Data"`
}

//method read
func (c *client) read() {
	var msg Message
	for {
		if err := c.socket.ReadJSON(&msg); err == nil {
			c.connect.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

//method write
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

package httpserver

import (
	"github.com/gorilla/websocket"
)

//client is one chat user
type client struct {
	//client is websocket
	socket *websocket.Conn
	//send is message channel
	send chan []byte
	//chan is client chat room
	connect *connect
}

//method read
func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
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
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

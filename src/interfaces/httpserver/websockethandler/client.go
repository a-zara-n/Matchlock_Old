package websockethandler

import "github.com/gorilla/websocket"

//client is one chat user
type client struct {
	//client is websocket
	socket *websocket.Conn
	//send is message channel
	send chan WebSocketRequest
	//chan is client chat room
	connect *webSocketHandler
}

//method read
func (c *client) read() {
	var msg WebSocketRequest
	for {
		if err := c.socket.ReadJSON(&msg); err == nil {
			switch msg.Type {
			case "Intercept", "HistoryCount":
				c.connect.forward <- msg

			}
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

package httpserver

import (
	"log"
	"net/http"

	"../channel"
	"../extractor"
	"github.com/gorilla/websocket"
)

type connect struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	forward chan Message
	// joinはチャットルームに参加しようとしているクライアントのためのチャネルです。
	join chan *client
	// leaveはチャットルームから退室しようとしているクライアントのためのチャネルです
	leave chan *client
	// clientsには在室しているすべてのクライアントが保持されます。
	clients map[*client]bool
	// tracerはチャットルーム上で行われた操作のログを受け取ります。
	channel *channel.Matchlock
	request *http.Request
}

func (c *connect) Run() {
	reqchan := c.channel.Request
	for {
		select {
		case client := <-c.join:
			//参加
			c.clients[client] = true
		case client := <-c.leave:
			//退室
			delete(c.clients, client)
			close(client.send)

		case msg := <-c.forward:
			c.request = extractor.GetRequestByString(msg.Data, c.request)
			reqchan.HMgToHsSignal <- c.request
			c.request = &http.Request{}
		case r := <-reqchan.HMgToHsSignal:
			c.request = r
			ret := extractor.GetStringByRequest(r)
			mes := Message{
				Type: "Request",
				Data: ret,
			}
			for client := range c.clients {
				select {
				case client.send <- mes:
				default:
					delete(c.clients, client)
					close(client.send)
				}
			}
		}
	}
}

var upgrader = &websocket.Upgrader{}

func (c *connect) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket:  socket,
		send:    make(chan Message),
		connect: c,
	}
	c.join <- client
	defer func() { c.leave <- client }()
	go client.write()
	client.read()
}

func newConnect(m *channel.Matchlock) *connect {
	return &connect{
		forward: make(chan Message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		channel: m,
	}
}

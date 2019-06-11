package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/WestEast1st/Matchlock/channel"
	"github.com/WestEast1st/Matchlock/extractor"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
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
	go func() {
		var count int
		hisdb := db.OpenDatabase()
		hisdb.Table("requests").Count(&count)
		historyCount = count
		time.Sleep(5 * time.Second)
		for {
			hisdb.Table("requests").Count(&count)
			if count != historyCount && historyCount < count {
				historys := getHistory(historyCount + 1)
				res, _ := json.Marshal(APIresponse{Data: historys})
				historyCount += len(historys)
				mes := Message{
					Type: "History",
					Data: string(res),
				}
				for client := range c.clients {
					select {
					case client.send <- mes:
					default:
						delete(c.clients, client)
						close(client.send)
					}
				}
				time.Sleep(50 * time.Millisecond)
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()
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

func (c *connect) ServeHTTP(cont echo.Context) error {
	var (
		w   = cont.Response()
		req = cont.Request()
	)
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return err
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
	return nil
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

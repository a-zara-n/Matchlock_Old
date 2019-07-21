package websockethandler

import (
	"log"

	"github.com/a-zara-n/Matchlock/src/config"

	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

//WebSocketHandler はWebSocketを用いたメッセージの処理を定義
type WebSocketHandler interface {
	ServeHTTP(cont echo.Context) error
	Run()
}
type webSocketHandler struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	forward chan WebSocketRequest
	// joinはチャットルームに参加しようとしているクライアントのためのチャネルです。
	join chan *client
	// leaveはチャットルームから退室しようとしているクライアントのためのチャネルです
	leave chan *client
	// clientsには在室しているすべてのクライアントが保持されます。
	clients map[*client]bool
	// tracerはチャットルーム上で行われた操作のログを受け取ります。
	channel *config.HTTPServerChannel
	request aggregate.Request
	usecase usecase.WebSocketUsecase
}

//WebSocketRequest json struct
type WebSocketRequest struct {
	Type string `json:"Type"`
	Data string `json:"Data"`
}

//NewWebSocketHandler はWebSocketのコネクションん管理を行います
func NewWebSocketHandler(c *config.HTTPServerChannel, ws usecase.WebSocketUsecase) WebSocketHandler {
	return &webSocketHandler{
		forward: make(chan WebSocketRequest),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		channel: c,
		usecase: ws,
		request: aggregate.Request{},
	}
}

var upgrader = &websocket.Upgrader{}

func (ws *webSocketHandler) ServeHTTP(cont echo.Context) error {
	socket, err := upgrader.Upgrade(cont.Response(), cont.Request(), nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return err
	}
	client := &client{
		socket:  socket,
		send:    make(chan WebSocketRequest),
		connect: ws,
	}
	ws.join <- client
	defer func() { ws.leave <- client }()
	go client.write()
	client.read()
	return nil
}

func (ws *webSocketHandler) Run() {
	for {
		select {
		case client := <-ws.join:
			//参加
			log.Println("正常にアクセスしました")
			ws.clients[client] = true
		case client := <-ws.leave:
			//退室
			delete(ws.clients, client)
			close(client.send)
		case msg := <-ws.forward:
			ws.channel.Response <- msg.Data
		case r := <-ws.channel.Request:
			log.Println("リクエストを受信しました")
			mes := WebSocketRequest{Type: "Request", Data: ws.usecase.GetHTTPRequestByString(r)}
			for client := range ws.clients {
				select {
				case client.send <- mes:
				default:
					delete(ws.clients, client)
					close(client.send)
				}
			}
		}
	}
}

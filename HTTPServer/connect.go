package httpserver

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"unsafe"

	"../channel"
	"github.com/gorilla/websocket"
)

type connect struct {
	// forwardは他のクライアントに転送するためのメッセージを保持するチャネルです。
	forward chan []byte
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

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

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
			c.SetRequest(msg)
			reqchan.HMgToHsSignal <- c.request
			c.request = &http.Request{}
		case r := <-reqchan.HMgToHsSignal:
			c.request = r
			ret := c.GetRequestByHeader(r)
			for client := range c.clients {
				select {
				case client.send <- []byte(ret):
				default:
					delete(c.clients, client)
					close(client.send)
				}
			}
		}
	}
}

func GetQuery(rq string) string {
	if rq != "" {
		return "?" + rq
	}
	return rq
}

func (c *connect) GetRequestByHeader(r *http.Request) string {
	headerKey := []string{}
	headerSlice := []string{}
	for k := range r.Header {
		headerKey = append(headerKey, k)
	}
	sort.Strings(headerKey)
	for _, v := range headerKey {
		h := strings.Join([]string{v, strings.Join(r.Header[v], ",")}, ": ")
		headerSlice = append(headerSlice, h)
	}
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	return strings.Join([]string{
		strings.Join([]string{r.Method, r.URL.Path, GetQuery(r.URL.RawQuery), r.Proto}, " "),
		strings.Join(headerSlice, "\n"),
		"",
		bufbody.String(),
	}, "\n")
}

func (c *connect) SetRequest(msg []byte) {
	editReq := strings.Split(*(*string)(unsafe.Pointer(&msg)), "\n")
	startLine := strings.Split(editReq[0], " ")
	host := c.request.Host
	c.request.Header = http.Header{}
	for _, v := range editReq[1:] {
		if v == "" {
			break
		}
		headL := strings.Split(v, ": ")
		if len(headL) <= 1 {
			headL = strings.Split(v, ":")
		}
		c.request.Header.Add(headL[0], strings.Join(headL[1:], ":"))
	}
	if s := c.request.Header.Get("Host"); s != "" {
		host = s
	}
	c.request.URL.Host = host
	c.request.URL.Path = startLine[1]
	c.request.Method = startLine[0]
	c.request.Proto = startLine[2]
	bodyStr := editReq[len(editReq)-1]
	c.request.ContentLength = int64(len(bodyStr))
	c.request.Body = ioutil.NopCloser(strings.NewReader(bodyStr))
}

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (c *connect) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket:  socket,
		send:    make(chan []byte, messageBufferSize),
		connect: c,
	}
	c.join <- client
	defer func() { c.leave <- client }()
	go client.write()
	client.read()
}

func newConnect(m *channel.Matchlock) *connect {
	return &connect{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		channel: m,
	}
}

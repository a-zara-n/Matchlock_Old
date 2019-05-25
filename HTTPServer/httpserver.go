package httpserver

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"../channel"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	t.templ.Execute(w, data)
}

type HttpServer interface {
	Run()
}
type httpServer struct {
	channels *channel.Matchlock
}

func (h *httpServer) Run() {
	c := newConnect(h.channels)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/connect", c)
	http.HandleFunc("/api/is/forward", func(w http.ResponseWriter, r *http.Request) {
		if h.channels.IsForward {
			h.channels.IsForward = false
		} else {
			h.channels.IsForward = true
		}
	})
	go c.Run()
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func NewHTTPServer(m *channel.Matchlock) HttpServer {
	return &httpServer{channels: m}
}

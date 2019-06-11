package httpserver

import (
	"html/template"
	"io"
	"net/http"

	"github.com/WestEast1st/Matchlock/channel"
	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type ServiceInfo struct {
	Title string
}

var serviceInfo = ServiceInfo{
	"Web dynamic testing",
}

type HttpServer interface {
	Run()
}
type httpServer struct {
	channels *channel.Matchlock
}

func (h *httpServer) Run() {
	conn := newConnect(h.channels)
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(99)
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}
	e.GET("/", h.index)
	e.GET("/connect", conn.ServeHTTP)
	e.GET("/api/is/forward", h.changeForward)
	e.GET("/api/history/all", GetHistryAll)
	e.GET("/api/request/:identifier", GetRequest)
	go conn.Run()
	e.Start(":8888")
}

func NewHTTPServer(m *channel.Matchlock) HttpServer {
	return &httpServer{channels: m}
}

func (h *httpServer) changeForward(c echo.Context) error {
	if h.channels.IsForward {
		h.channels.IsForward = false
	} else {
		h.channels.IsForward = true
	}
	return c.JSON(http.StatusOK, "{status: \"OK\"}")
}

func (h *httpServer) index(c echo.Context) error {
	data := struct {
		ServiceInfo
		Host string
	}{
		ServiceInfo: serviceInfo,
		Host:        c.Request().Host,
	}
	return c.Render(http.StatusOK, "index", data)
}

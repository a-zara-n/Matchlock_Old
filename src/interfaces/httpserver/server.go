package httpserver

import (
	"github.com/a-zara-n/MatchlockDDD/Matchlock/src/domain/entity"
	"github.com/labstack/echo"
)

type HttpServer interface {
	Run()
}
type httpServer struct {
	channels *entity.Channel
}

//NewHTTPServer は
func NewHTTPServer(c *entity.Channel) HttpServer {
	return &httpServer{channels: c}
}

func (h *httpServer) Run() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(99)
	e.Renderer = renders()
	e.Start(":8888")
}

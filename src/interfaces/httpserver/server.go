package httpserver

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver/htmlhandler"
	"github.com/labstack/echo"
)

type HttpServer interface {
	Run()
}
type httpServer struct {
	channels *entity.Channel
	htmlhandler.HTMLHandler
}

//NewHTTPServer は
func NewHTTPServer(c *entity.Channel, h usecase.HTMLUseCase) HttpServer {
	return &httpServer{c, htmlhandler.NewWarmupHandler(h)}
}

func (h *httpServer) Run() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(99)
	e.Renderer = renders()
	e = h.router(e)
	e.Start(":8888")
}

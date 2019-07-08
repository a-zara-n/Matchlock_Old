package httpserver

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver/apihandler"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver/htmlhandler"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver/websockethandler"
	"github.com/labstack/echo"
)

type HttpServer interface {
	Run()
}
type httpServer struct {
	htmlhandler.HTMLHandler
	apihandler.API
	websockethandler.WebSocketHandler
}

//NewHTTPServer は
func NewHTTPServer(c *entity.Channel, h usecase.HTMLUseCase, api *usecase.APIUsecase, ws usecase.WebSocketUsecase) HttpServer {
	return &httpServer{
		htmlhandler.NewWarmupHandler(h),
		apihandler.SetHandler(c, api),
		websockethandler.NewWebSocketHandler(c, ws),
	}
}

func (h *httpServer) Run() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(99)
	e.Renderer = renders()
	e = h.router(e)
	e.Start(":8888")
}

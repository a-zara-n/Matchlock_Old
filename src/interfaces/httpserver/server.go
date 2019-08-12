package httpserver

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/config"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver/apihandler"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver/htmlhandler"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver/middleware"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver/websockethandler"
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
)

type HttpServer interface {
	Run()
}
type httpServer struct {
	htmlhandler.HTMLHandler
	apihandler.API
	websockethandler.WebSocketHandler
	middleware.Middleware
}

//NewHTTPServer は
func NewHTTPServer(c *config.HTTPServerChannel, h usecase.HTMLUsecase, api *usecase.APIUsecase, ws usecase.WebSocketUsecase, m usecase.ManagerUsecase) HttpServer {
	return &httpServer{
		htmlhandler.NewHTMLHandler(h),
		apihandler.NewAPIHandler(api),
		websockethandler.NewWebSocketHandler(c, ws),
		middleware.NewMiddleware(m),
	}
}

func (h *httpServer) Run() {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(99)
	e.Renderer = renders()
	e.Use(em.CORS())
	e = h.router(e)
	go h.Middleware.Run()
	go h.WebSocketHandler.Run()
	e.Start(":8888")
}

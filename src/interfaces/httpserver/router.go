package httpserver

import (
	"github.com/labstack/echo"
)

//router はecho.Echoを排出します
func (h *httpServer) router(e *echo.Echo) *echo.Echo {
	e.GET("/", h.Index)
	e.GET("/connect", h.WebSocketHandler.ServeHTTP)
	e.GET("/api/is/forward", h.API.ChangeForward)
	e.GET("/api/history/:type", h.API.FetchHistory)
	e.GET("/api/message/:identifier", h.API.FetchMessage)
	return e
}

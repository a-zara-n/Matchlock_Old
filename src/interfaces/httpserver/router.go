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
	e.GET("/api/whitelist", h.API.FetchWhiteList)
	e.POST("/api/whitelist", h.API.AddWhiteList)
	e.PUT("/api/whitelist", h.API.UpdateWhiteList)
	e.DELETE("/api/whitelist", h.API.DeleteWhiteList)
	e.POST("/api/scan", h.API.RunScan)
	return e
}

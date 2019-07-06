package httpserver

import (
	"github.com/labstack/echo"
)

//router はecho.Echoを排出します
func (h *httpServer) router(e *echo.Echo) *echo.Echo {
	e.GET("/", h.Index)
	return e
}

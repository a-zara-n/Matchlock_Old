package usecase

import (
	"net/http"

	"github.com/labstack/echo"
)

type HTMLUsecase interface {
	GetIndex(c echo.Context) error
}

type htmlUsecase struct{}

type ServiceInfo struct {
	Title string
}

var serviceinfo = ServiceInfo{
	"Web dynamic testing",
}

func NewHTMLUsecase() HTMLUsecase {
	return &htmlUsecase{}
}

func (h *htmlUsecase) GetIndex(c echo.Context) error {
	data := struct {
		ServiceInfo
		Host string
	}{
		ServiceInfo: serviceinfo,
		Host:        c.Request().Host,
	}
	return c.Render(http.StatusOK, "index", data)
}

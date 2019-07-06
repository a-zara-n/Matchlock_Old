package usecase

import (
	"net/http"

	"github.com/labstack/echo"
)

type HTMLUseCase interface {
	GetIndex(c echo.Context) error
}

type htmlUseCase struct{}

type ServiceInfo struct {
	Title string
}

var serviceinfo = ServiceInfo{
	"Web dynamic testing",
}

func NewHTMLUseCase() HTMLUseCase {
	return &htmlUseCase{}
}

func (h *htmlUseCase) GetIndex(c echo.Context) error {
	data := struct {
		ServiceInfo
		Host string
	}{
		ServiceInfo: serviceinfo,
		Host:        c.Request().Host,
	}
	return c.Render(http.StatusOK, "index", data)
}

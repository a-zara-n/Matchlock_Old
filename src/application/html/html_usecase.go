package htmlusecase

import (
	"net/http"

	"github.com/labstack/echo"
)

type HTMLUseCase interface {
	GetIndex(c echo.Context) error
}

type htmlUseCase struct{}

type serviceInfo struct {
	Title string
}

var serviceinfo = serviceInfo{
	"Web dynamic testing",
}

func NewHTMLUseCase() HTMLUseCase {
	return &htmlUseCase{}
}

func (h *htmlUseCase) GetIndex(c echo.Context) error {
	data := struct {
		serviceInfo
		Host string
	}{
		serviceInfo: serviceinfo,
		Host:        c.Request().Host,
	}
	return c.Render(http.StatusOK, "index", data)
}

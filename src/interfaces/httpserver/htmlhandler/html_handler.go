package htmlhandler

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/labstack/echo"
)

type HTMLHandler interface {
	Index(c echo.Context) error
}

type htmlHandler struct {
	usecase.HTMLUsecase
}

func NewHTMLHandler(h usecase.HTMLUsecase) HTMLHandler {
	return &htmlHandler{h}
}

func (h *htmlHandler) Index(c echo.Context) error {
	return h.GetIndex(c)
}

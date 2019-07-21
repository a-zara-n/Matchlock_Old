package apihandler

import (
	"net/http"

	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/labstack/echo"
)

type API interface {
	ChangeForward(c echo.Context) error
}
type api struct {
	usecase *usecase.APIUsecase
}

func NewAPIHandler(a *usecase.APIUsecase) API {
	return &api{a}
}

func (api *api) ChangeForward(c echo.Context) error {
	retbool, _ := api.usecase.ChangeForward()
	c.JSON(http.StatusOK, map[string]interface{}{"status": retbool})
	return nil
}

package apihandler

import (
	"encoding/json"
	"net/http"

	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/labstack/echo"
)

//message json struct
type Message struct {
	Type string `json:"Type"`
	Data string `json:"Data"`
}

type APIresponse struct {
	Data interface{} `json:"Data"`
}
type API interface {
	ChangeForward(c echo.Context) error
	FetchHistory(c echo.Context) error
	FetchMessage(c echo.Context) error
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

func (api *api) FetchHistory(c echo.Context) error {
	var (
		historys interface{}
		response = c.Response()
	)
	switch c.Param("type") {
	case "all":
		historys = api.usecase.FetchAll()
	}
	responsedata, err := json.Marshal(APIresponse{Data: historys})
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return err
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(responsedata)
	return nil
}

func (api *api) FetchMessage(c echo.Context) error {
	var response = c.Response()
	responsedata, err := json.Marshal(APIresponse{Data: api.usecase.FetchMessage(c.Param("identifier"))})
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return err
	}
	response.Header().Set("Content-Type", "application/json")
	response.Write(responsedata)
	return nil
}

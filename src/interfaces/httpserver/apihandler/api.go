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

type ID struct {
	ID int `json:"id" form:"id" query:"id"`
}
type Regex struct {
	Regex string `json:"regex" form:"regex" query:"regex"`
}
type Host struct {
	Host string `json:"host" form:"host" query:"host"`
}

type request struct {
	ID
	Regex
}
type API interface {
	ChangeForward(c echo.Context) error
	FetchHistory(c echo.Context) error
	FetchMessage(c echo.Context) error
	FetchWhiteList(c echo.Context) error
	UpdateWhiteList(c echo.Context) error
	DeleteWhiteList(c echo.Context) error
	AddWhiteList(c echo.Context) error
	RunScan(c echo.Context) error
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
		historys = api.usecase.FetchHistoryAll()
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

func (api *api) FetchWhiteList(c echo.Context) error {
	c.JSON(http.StatusOK, APIresponse{Data: api.usecase.FetchWhiteList(0)})
	return nil
}

func (api *api) UpdateWhiteList(c echo.Context) error {
	req := new(request)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"status": false})
		return err
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": api.usecase.UpdateWhiteList(req.ID.ID, req.Regex.Regex)})
	return nil
}

func (api *api) DeleteWhiteList(c echo.Context) error {
	req := new(ID)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"status": false})
		return err
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": api.usecase.DelWhiteList(req.ID)})
	return nil
}

func (api *api) AddWhiteList(c echo.Context) error {
	req := new(Regex)
	if err := c.Bind(req); err != nil {
		c.JSON(http.StatusExpectationFailed, map[string]interface{}{"status": false})
		return err
	}
	if req.Regex == "" {
		c.JSON(http.StatusNoContent, map[string]interface{}{"status": false})
		return nil
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": api.usecase.AddWhiteList(req.Regex)})
	return nil
}

func (api *api) RunScan(c echo.Context) error {
	host := new(Host)
	if err := c.Bind(host); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"status": false})
		return err
	}
	c.JSON(http.StatusOK, map[string]interface{}{"status": api.usecase.RunScan(host.Host)})
	return nil
}

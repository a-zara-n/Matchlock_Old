package usecase

import (
	"net/http"

	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
)

type WebSocketUsecase interface {
	GetHTTPRequestByRequest(data string) *http.Request
	GetHTTPRequestByString(data *http.Request) string
}
type websocketusecase struct {
	request aggregate.Request
}

func NewWebSocketUsecase(memreq repository.RequestRepositry, memres repository.ResponseRepositry) WebSocketUsecase {
	return &websocketusecase{request: aggregate.Request{}}
}

func (wsu *websocketusecase) GetHTTPRequestByRequest(data string) *http.Request {
	wsu.request.SetHTTPRequestByString(data)
	return wsu.request.GetHTTPRequestByRequest()
}
func (wsu *websocketusecase) GetHTTPRequestByString(data *http.Request) string {
	wsu.request.SetHTTPRequestByRequest(data)
	return wsu.request.GetHTTPRequestByString()
}

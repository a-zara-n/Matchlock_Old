package usecase

import (
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
)

type WebSocketUsecase interface {
	GetHTTPRequestByRequest(data string) *aggregate.Request
	GetHTTPRequestByString(data *aggregate.Request) string
}
type websocketusecase struct {
	request aggregate.Request
}

func NewWebSocketUsecase(memreq repository.RequestRepositry, memres repository.ResponseRepositry) WebSocketUsecase {
	return &websocketusecase{request: aggregate.Request{}}
}

func (wsu *websocketusecase) GetHTTPRequestByRequest(data string) *aggregate.Request {
	wsu.request.SetHTTPRequestByString(data)
	return &wsu.request
}
func (wsu *websocketusecase) GetHTTPRequestByString(data *aggregate.Request) string {
	return data.GetHTTPRequestByString()
}

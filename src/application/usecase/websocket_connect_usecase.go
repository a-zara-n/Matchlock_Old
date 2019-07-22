package usecase

import (
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
)

type WebSocketUsecase interface {
	GetHTTPRequestByString(data *aggregate.Request) string
	DiffHistory(i int) interface{}
}
type websocketusecase struct {
	repository.HistoryRepository
}

func NewWebSocketUsecase(memreq repository.RequestRepositry, memres repository.ResponseRepositry, hh repository.HistoryRepository) WebSocketUsecase {
	return &websocketusecase{hh}
}

func (wsu *websocketusecase) GetHTTPRequestByString(data *aggregate.Request) string {
	return data.GetHTTPRequestByString()
}

func (wsu *websocketusecase) DiffHistory(i int) interface{} {
	return wsu.Fetch(i)
}

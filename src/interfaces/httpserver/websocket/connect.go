package websocket

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

type WebSocketHandler interface{}
type webSocketHandler struct{}

//NewWebSocketHandler はWebSocketのコネクションん管理を行います
func NewWebSocketHandler(c *entity.Channel, ws usecase.WebSocketUsecase) WebSocketHandler {
	return &webSocketHandler{}
}

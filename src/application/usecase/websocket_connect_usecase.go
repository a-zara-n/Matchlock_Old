package usecase

type WebSocketUsecase interface{}
type websocketusecase struct{}

func NewWebSocketUsecase() WebSocketUsecase {
	return &websocketusecase{}
}

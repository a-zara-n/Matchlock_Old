package registry

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

type Usecase interface {
	NewLogic() usecase.ProxyLogic
	NewHTMLUseCase() usecase.HTMLUseCase
	NewAPIUsecase() *usecase.APIUsecase
	NewCommandUsecase() usecase.CommandUsecase
	NewWebSocketUsecase() usecase.WebSocketUsecase
}

//NewLogic はusecase.ProxyLogicを取得
func NewLogic(white *entity.WhiteList) usecase.ProxyLogic {
	return usecase.NewLogic(white)
}

//NewHTMLUseCase はHTMLのレンダリングを行う
func NewHTMLUseCase() usecase.HTMLUseCase {
	return usecase.NewHTMLUseCase()
}

//NewAPIUsecase はAPIの処理を取得
func NewAPIUsecase() *usecase.APIUsecase {
	return usecase.NewAPIUsecase()
}

//NewWebSocketUsecase はWebSocket用のUseCaseを取得します
func NewWebSocketUsecase() usecase.WebSocketUsecase {
	return usecase.NewWebSocketUsecase()
}

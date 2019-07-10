package registry

import (
	"github.com/a-zara-n/Matchlock/src/application"
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
)

type Usecase interface {
	NewLogic(white *entity.WhiteList, c *entity.Channel) application.ProxyLogic
	NewHTMLUseCase() usecase.HTMLUsecase
	NewAPIUsecase(memreq repository.RequestRepositry, memres repository.ResponseRepositry) *usecase.APIUsecase
	NewCommandUsecase() usecase.CommandUsecase
	NewWebSocketUsecase(memreq repository.RequestRepositry, memres repository.ResponseRepositry) usecase.WebSocketUsecase
	NewManagerUsecase(channel *entity.Channel, memreq repository.RequestRepositry, memres repository.ResponseRepositry) usecase.ManagerUsecase
}

//NewLogic はusecase.ProxyLogicを取得
func NewLogic(white *entity.WhiteList, c *entity.Channel) application.ProxyLogic {
	return application.NewLogic(white, c)
}

//NewHTMLUseCase はHTMLのレンダリングを行う
func NewHTMLUseCase() usecase.HTMLUsecase {
	return usecase.NewHTMLUsecase()
}

//NewAPIUsecase はAPIの処理を取得
func NewAPIUsecase(memreq repository.RequestRepositry, memres repository.ResponseRepositry) *usecase.APIUsecase {
	return usecase.NewAPIUsecase(memreq, memres)
}

//NewWebSocketUsecase はWebSocket用のUseCaseを取得します
func NewWebSocketUsecase(memreq repository.RequestRepositry, memres repository.ResponseRepositry) usecase.WebSocketUsecase {
	return usecase.NewWebSocketUsecase(memreq, memres)
}

//NewManagerUsecase は管理を行うためのManagerUsecaseを取得します
func NewManagerUsecase(channel *entity.Channel, memreq repository.RequestRepositry, memres repository.ResponseRepositry) usecase.ManagerUsecase {
	return usecase.NewManagerUsecase(channel, memreq, memres)
}

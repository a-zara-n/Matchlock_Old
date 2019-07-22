package registry

import (
	"github.com/a-zara-n/Matchlock/src/application"
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/config"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/a-zara-n/Matchlock/src/domain/value"
)

type Usecase interface {
	NewLogic(white *entity.WhiteList, c *config.ProxyChannel) application.ProxyLogic
	NewHTMLUseCase() usecase.HTMLUsecase
	NewAPIUsecase(f *value.Forward, memreq repository.RequestRepositry, memres repository.ResponseRepositry, history repository.HistoryRepository, message repository.HTTPMessageRepository) *usecase.APIUsecase
	NewCommandUsecase() usecase.CommandUsecase
	NewWebSocketUsecase(memreq repository.RequestRepositry, memres repository.ResponseRepositry, hh repository.HistoryRepository) usecase.WebSocketUsecase
	NewManagerUsecase(channel config.Channel, memreq repository.RequestRepositry, memres repository.ResponseRepositry, history repository.HistoryRepository, f *value.Forward) usecase.ManagerUsecase
}

//NewLogic はusecase.ProxyLogicを取得
func (r *registry) NewLogic(white *entity.WhiteList, c *config.ProxyChannel) application.ProxyLogic {
	return application.NewLogic(white, c)
}

//NewHTMLUseCase はHTMLのレンダリングを行う
func (r *registry) NewHTMLUseCase() usecase.HTMLUsecase {
	return usecase.NewHTMLUsecase()
}

//NewAPIUsecase はAPIの処理を取得
func (r *registry) NewAPIUsecase(f *value.Forward, memreq repository.RequestRepositry, memres repository.ResponseRepositry, history repository.HistoryRepository, message repository.HTTPMessageRepository) *usecase.APIUsecase {
	return usecase.NewAPIUsecase(f, memreq, memres, history, message)
}
func (r *registry) NewCommandUsecase() usecase.CommandUsecase {
	return usecase.NewCommandUsecase()
}

//NewWebSocketUsecase はWebSocket用のUseCaseを取得します
func (r *registry) NewWebSocketUsecase(memreq repository.RequestRepositry, memres repository.ResponseRepositry, hh repository.HistoryRepository) usecase.WebSocketUsecase {
	return usecase.NewWebSocketUsecase(memreq, memres, hh)
}

//NewManagerUsecase は管理を行うためのManagerUsecaseを取得します
func (r *registry) NewManagerUsecase(channel config.Channel, memreq repository.RequestRepositry, memres repository.ResponseRepositry, history repository.HistoryRepository, f *value.Forward) usecase.ManagerUsecase {
	return usecase.NewManagerUsecase(channel, memreq, memres, history, f)
}

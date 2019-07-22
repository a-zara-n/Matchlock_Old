package usecase

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase/api"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/a-zara-n/Matchlock/src/domain/value"
)

//APIUsecase で利用されるUsecaseをまとめた構造体
type APIUsecase struct {
	api.ForwordInterface
	api.HistoryInterface
	api.MessageInterface
	api.WhiteListInterface
}

func NewAPIUsecase(f *value.Forward, whitelist *entity.WhiteList, history repository.HistoryRepository, message repository.HTTPMessageRepository) *APIUsecase {
	return &APIUsecase{
		api.NewForword(f), api.NewHistory(history), api.NewMessage(message), api.NewWhiteList(whitelist),
	}
}

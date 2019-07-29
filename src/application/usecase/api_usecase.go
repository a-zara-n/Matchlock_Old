package usecase

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase/api"
)

//APIUsecase で利用されるUsecaseをまとめた構造体
type APIUsecase struct {
	api.ForwordInterface
	api.HistoryInterface
	api.MessageInterface
	api.WhiteListInterface
	api.ScanInterface
}

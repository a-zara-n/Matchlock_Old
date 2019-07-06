package usecase

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase/api"
)

//APIusecase で利用されるUsecaseをまとめた構造体
type APIUsecase struct {
	api.Forward
}

func NewAPIUsecase() *APIUsecase {
	return &APIUsecase{
		api.NewForward(),
	}
}

package usecase

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase/api"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
)

//APIusecase で利用されるUsecaseをまとめた構造体
type APIUsecase struct {
	api.Forward
}

func NewAPIUsecase(memreq repository.RequestRepositry, memres repository.ResponseRepositry) *APIUsecase {
	return &APIUsecase{
		api.NewForward(),
	}
}

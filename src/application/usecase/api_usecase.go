package usecase

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase/api"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/a-zara-n/Matchlock/src/domain/value"
)

//APIUsecase で利用されるUsecaseをまとめた構造体
type APIUsecase struct {
	api.Forward
}

func NewAPIUsecase(f *value.Forward, memreq repository.RequestRepositry, memres repository.ResponseRepositry) *APIUsecase {
	return &APIUsecase{
		api.NewForward(f),
	}
}

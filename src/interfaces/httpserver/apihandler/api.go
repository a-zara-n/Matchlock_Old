package apihandler

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
)

type API interface {
}
type api struct {
}

func NewAPIHandler(a *usecase.APIUsecase) API {
	return &api{}
}

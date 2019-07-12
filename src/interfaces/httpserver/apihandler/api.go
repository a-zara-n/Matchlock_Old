package apihandler

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

type API interface {
}
type api struct {
}

func NewAPIHandler(c *entity.Channel, a *usecase.APIUsecase) API {
	return &api{}
}

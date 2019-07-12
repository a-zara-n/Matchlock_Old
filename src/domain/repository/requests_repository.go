package repository

import (
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

type RequestRepositry interface {
	Insert(Identifier string, IsEdit bool, a *aggregate.Request) bool
	GetRequest(Identifier string, IsEdit bool) *aggregate.Request
	Select(Identifier string, IsEdit bool) *entity.RequestInfo
}

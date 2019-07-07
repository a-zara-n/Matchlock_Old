package repository

import (
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

type RequestRepositry interface {
	Insert(a *aggregate.Request) bool
	GetRequest() *aggregate.Request
	Select() *entity.RequestInfo
	HistryCommon
}

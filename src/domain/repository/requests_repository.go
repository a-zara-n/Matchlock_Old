package repository

import (
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//RequestRepositry は
type RequestRepositry interface {
	Insert(Identifier string, IsEdit bool, a *aggregate.Request) bool
	Fetch(Identifier string, IsEdit bool) *aggregate.Request
}

//RequestInfoRepositry は
type RequestInfoRepositry interface {
	Insert(Identifier string, IsEdit bool, e *entity.RequestInfo) bool
	Fetch(Identifier string, IsEdit bool) *entity.RequestInfo
}

//RequestHeaderRepositry はDBを操作する関数を定義したinterfaceを記述しています
type RequestHeaderRepositry interface {
	Insert(Identifier string, IsEdit bool, e *entity.HTTPHeader) bool
	Fetch(Identifier string, IsEdit bool) *entity.HTTPHeader
}

//RequestDataRepositry はDBを操作する関数を定義したinterfaceを記述しています
type RequestDataRepositry interface {
	Insert(Identifier string, IsEdit bool, e *entity.Data) bool
	Fetch(Identifier string, IsEdit bool) *entity.Data
}

package repository

import (
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//ResponseRepositry はDBを操作する関数を定義したinterfaceを記述しています
type ResponseRepositry interface {
	Insert(Identifier string, a *aggregate.Response) bool
	GetRequest(Identifier string) *aggregate.Response
	Select(Identifier string) *entity.ResponseInfo
}

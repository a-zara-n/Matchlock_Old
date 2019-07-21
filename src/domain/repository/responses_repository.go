package repository

import (
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//ResponseRepositry はDBを操作する関数を定義したinterfaceを記述しています
type ResponseRepositry interface {
	Insert(Identifier string, a *aggregate.Response) bool
	Fetch(Identifier string) *aggregate.Response
}

//ResponseInfoRepositry は
type ResponseInfoRepositry interface {
	Insert(Identifier string, e *entity.ResponseInfo) bool
	Fetch(Identifier string) *entity.ResponseInfo
}

//ResponseHeaderRepositry はDBを操作する関数を定義したinterfaceを記述しています
type ResponseHeaderRepositry interface {
	Insert(Identifier string, e *entity.HTTPHeader) bool
	Fetch(Identifier string) *entity.HTTPHeader
}

//ResponseBodyRepositry はDBを操作する関数を定義したinterfaceを記述しています
type ResponseBodyRepositry interface {
	Insert(Identifier string, e *entity.Body) bool
	Fetch(Identifier string) *entity.Body
}

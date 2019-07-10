package repository

import "github.com/a-zara-n/Matchlock/src/domain/entity"

//RequestHeaderRepositry はDBを操作する関数を定義したinterfaceを記述しています
type RequestHeaderRepositry interface {
	Insert(Identifier string, IsEdit bool, e *entity.HTTPHeader) bool
	Select(Identifier string, IsEdit bool) *entity.HTTPHeader
}

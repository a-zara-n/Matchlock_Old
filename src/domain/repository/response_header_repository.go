package repository

import "github.com/a-zara-n/Matchlock/src/domain/entity"

//ResponseHeaderRepositry はDBを操作する関数を定義したinterfaceを記述しています
type ResponseHeaderRepositry interface {
	Insert(Identifier string, e *entity.HTTPHeader) bool
	Select(Identifier string) *entity.HTTPHeader
}

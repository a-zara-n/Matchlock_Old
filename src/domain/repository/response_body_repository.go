package repository

import "github.com/a-zara-n/Matchlock/src/domain/entity"

//ResponseBodyRepositry はDBを操作する関数を定義したinterfaceを記述しています
type ResponseBodyRepositry interface {
	Insert(Identifier string, e *entity.Body) bool
	Select(Identifier string) *entity.Body
}

package repository

import (
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//RequestDataRepositry はDBを操作する関数を定義したinterfaceを記述しています
type RequestDataRepositry interface {
	Insert(Identifier string, IsEdit bool, e *entity.Data) bool
	Select(Identifier string, IsEdit bool) *entity.Data
}

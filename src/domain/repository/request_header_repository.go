package repository

import "github.com/a-zara-n/Matchlock/src/domain/entity"

//RequestHeaderRepositry はDBを操作する関数を定義したinterfaceを記述しています
type RequestHeaderRepositry interface {
	Insert(e *entity.HTTPHeader) bool
	Select() *entity.HTTPHeader
	HistryCommon
}

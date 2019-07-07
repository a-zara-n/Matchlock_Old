package repository

import "github.com/a-zara-n/Matchlock/src/domain/entity"

//ResponseHeaderRepositry はDBを操作する関数を定義したinterfaceを記述しています
type ResponseHeaderRepositry interface {
	Insert(e *entity.HTTPHeader) bool
	Select() *entity.HTTPHeader
	HistryCommon
}

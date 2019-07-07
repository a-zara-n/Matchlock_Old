package datastore

import (
	"strings"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/jinzhu/gorm"
)

//ResponseHeaderRepositry は保存用のRepositryです
type ResponseHeaderRepositry struct {
	historyCommon
}

//ResponseHeader は保存用のschemaです
type ResponseHeader struct {
	gorm.Model
	Identifier string
	Name       string
	Value      string
}

//NewResponseHeaderRepositry はResponseHeaderRepositryを取得する
func NewResponseHeaderRepositry(Identifier string, IsEdit bool, db *gorm.DB) repository.ResponseHeaderRepositry {
	return &ResponseHeaderRepositry{
		historyCommon{Identifier, IsEdit, db},
	}
}

//SetIsEdit は編集のフラグを書き換えることができます
func (r *ResponseHeaderRepositry) SetIsEdit(flag bool) { r.IsEdit = flag }

//SetIdentifier はIdentifierを書き換えることができます
func (r *ResponseHeaderRepositry) SetIdentifier(id string) { r.Identifier = id }

//Insert はResponseHeaderを保存します
func (r *ResponseHeaderRepositry) Insert(e *entity.HTTPHeader) bool {
	insertHeader := &ResponseHeader{
		Identifier: r.Identifier,
	}
	for _, key := range e.GetKeys() {
		insertHeader.Name = key
		insertHeader.Value = strings.Join(e.Header[key], ",")
		r.DB.Create(insertHeader)
	}
	return true
}

//Select はentity.HTTPHeader を取得できる
func (r *ResponseHeaderRepositry) Select() *entity.HTTPHeader {
	rets := []*ResponseHeader{}
	r.DB.Where("Identifier = ?", r.Identifier).Find(rets)
	retentity := &entity.HTTPHeader{}
	for _, data := range rets {
		retentity.Header.Add(data.Name, data.Value)
	}
	return retentity
}

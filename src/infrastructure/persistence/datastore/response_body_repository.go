package datastore

import (
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/jinzhu/gorm"
)

//ResponseBodyRepositry は保存用のRepositryです
type ResponseBodyRepositry struct {
	historyCommon
}

//ResponseBody は保存用のschemaです
type ResponseBody struct {
	gorm.Model
	Identifier string
	Body       string
	Encodetype string
	Length     int64
}

//NewResponseBodyRepositry はResponseBodyRepositryを取得する
func NewResponseBodyRepositry(Identifier string, IsEdit bool, db *gorm.DB) repository.ResponseBodyRepositry {
	return &ResponseBodyRepositry{
		historyCommon{Identifier, IsEdit, db},
	}
}

//SetIsEdit は編集のフラグを書き換えることができます
func (r *ResponseBodyRepositry) SetIsEdit(flag bool) { r.IsEdit = flag }

//SetIdentifier はIdentifierを書き換えることができます
func (r *ResponseBodyRepositry) SetIdentifier(id string) { r.Identifier = id }

//Insert はBodyを保存します
func (r *ResponseBodyRepositry) Insert(e *entity.Body) bool {
	insertData := &ResponseBody{
		Identifier: r.Identifier,
		Body:       e.Body,
		Encodetype: e.Encodetype,
		Length:     e.Length,
	}
	r.DB.Create(insertData)
	return true
}

//Select はentity.Bodyを取得出来る
func (r *ResponseBodyRepositry) Select() *entity.Body {
	rets := []*ResponseBody{}
	r.DB.Where("Identifier = ?", r.Identifier).Find(rets)
	retentity := &entity.Body{
		Body:       rets[0].Body,
		Encodetype: rets[0].Encodetype,
		Length:     rets[0].Length,
	}
	return retentity
}

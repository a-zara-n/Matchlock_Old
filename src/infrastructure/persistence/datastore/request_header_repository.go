package datastore

import (
	"strings"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/jinzhu/gorm"
)

//RequestHeaderRepositry は保存用のRepositryです
type RequestHeaderRepositry struct {
	historyCommon
}

//RequestHeader は保存用のschemaです
type RequestHeader struct {
	gorm.Model
	Identifier string
	IsEdit     bool
	Name       string
	Value      string
}

//NewRequestHeaderRepositry はRequestHeaderを取得する
func NewRequestHeaderRepositry(db *gorm.DB) repository.RequestHeaderRepositry {
	return &RequestHeaderRepositry{
		historyCommon{DB: db},
	}
}

//Insert はHTTPHeaderを保存します
func (r *RequestHeaderRepositry) Insert(Identifier string, IsEdit bool, e *entity.HTTPHeader) bool {
	insertHeader := &RequestHeader{
		Identifier: Identifier,
		IsEdit:     IsEdit,
	}
	for _, key := range e.GetKeys() {
		insertHeader.Name = key
		insertHeader.Value = strings.Join(e.Header[key], ",")
		r.DB.Create(insertHeader)
	}
	return true
}

//Select はentity.HTTPHeaderを取得します
func (r *RequestHeaderRepositry) Select(Identifier string, IsEdit bool) *entity.HTTPHeader {
	rets := []*RequestHeader{}
	r.DB.Where("Identifier = ? AND IsEdit = ?", Identifier, IsEdit).Find(rets)
	retentity := &entity.HTTPHeader{}
	for _, data := range rets {
		retentity.Header.Add(data.Name, data.Value)
	}
	return retentity
}

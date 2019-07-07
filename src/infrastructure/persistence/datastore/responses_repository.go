package datastore

import (
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/jinzhu/gorm"
)

//ResponseRepositry は保存用のRepositryです
type ResponseRepositry struct {
	historyCommon
	Header repository.ResponseHeaderRepositry
	Body   repository.ResponseBodyRepositry
}

//Response は保存用のschemaです
type Response struct {
	gorm.Model
	Identifier string
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
}

//NewResponseRepositry はResponseRepositryを取得する
func NewResponseRepositry(Identifier string, IsEdit bool, db *gorm.DB) repository.ResponseRepositry {
	header := NewResponseHeaderRepositry(Identifier, IsEdit, db)
	body := NewResponseBodyRepositry(Identifier, IsEdit, db)
	return &ResponseRepositry{
		historyCommon{Identifier, IsEdit, db},
		header,
		body,
	}
}

//SetIsEdit は編集のフラグを書き換えることができます
func (r *ResponseRepositry) SetIsEdit(flag bool) {
	go r.Header.SetIsEdit(flag)
	go r.Body.SetIsEdit(flag)
	r.IsEdit = flag
}

//SetIdentifier はIdentifierを書き換えることができます
func (r *ResponseRepositry) SetIdentifier(id string) {
	go r.Header.SetIdentifier(id)
	go r.Body.SetIdentifier(id)
	r.Identifier = id
}

//Insert はResponseを保存します
func (r *ResponseRepositry) Insert(a *aggregate.Response) bool {
	go r.insert(&a.Info)
	go r.Header.Insert(&a.Header)
	go r.Body.Insert(&a.Body)
	return true
}

//Insert はRequestInfoを保存します
func (r *ResponseRepositry) insert(e *entity.ResponseInfo) bool {
	insertRequestInfo := &Response{
		Identifier: r.Identifier,
		Status:     e.Status,
		StatusCode: e.StatusCode,
		Proto:      e.Proto,
		ProtoMajor: e.ProtoMajor,
		ProtoMinor: e.ProtoMinor,
	}
	r.DB.Create(insertRequestInfo)
	return true
}

//GetRequest はaggregate.Responseを取得します
func (r *ResponseRepositry) GetRequest() *aggregate.Response {
	retentity := &aggregate.Response{
		Info:   *r.Select(),
		Header: *r.Header.Select(),
		Body:   *r.Body.Select(),
	}
	return retentity
}

//Select はentity.ResponseInfoを取得します
func (r *ResponseRepositry) Select() *entity.ResponseInfo {
	rets := []*Response{}
	r.DB.Where("Identifier = ?", r.Identifier).Find(rets)
	retentity := &entity.ResponseInfo{
		Status:     rets[0].Status,
		StatusCode: rets[0].StatusCode,
		Proto:      rets[0].Proto,
		ProtoMajor: rets[0].ProtoMajor,
		ProtoMinor: rets[0].ProtoMinor,
	}
	return retentity
}

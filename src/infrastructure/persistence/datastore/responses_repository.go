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
func NewResponseRepositry(db *gorm.DB) repository.ResponseRepositry {
	header := NewResponseHeaderRepositry(db)
	body := NewResponseBodyRepositry(db)
	return &ResponseRepositry{
		historyCommon{DB: db},
		header,
		body,
	}
}

//Insert はResponseを保存します
func (r *ResponseRepositry) Insert(Identifier string, a *aggregate.Response) bool {
	go r.insert(Identifier, a.Info)
	go r.Header.Insert(Identifier, a.Header)
	go r.Body.Insert(Identifier, a.Body)
	return true
}

//Insert はRequestInfoを保存します
func (r *ResponseRepositry) insert(Identifier string, e *entity.ResponseInfo) bool {
	insertRequestInfo := &Response{
		Identifier: Identifier,
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
func (r *ResponseRepositry) GetRequest(Identifier string) *aggregate.Response {
	retentity := &aggregate.Response{
		Info:   r.Select(Identifier),
		Header: r.Header.Select(Identifier),
		Body:   r.Body.Select(Identifier),
	}
	return retentity
}

//Select はentity.ResponseInfoを取得します
func (r *ResponseRepositry) Select(Identifier string) *entity.ResponseInfo {
	rets := []*Response{}
	r.DB.Where("Identifier = ?", Identifier).Find(rets)
	retentity := &entity.ResponseInfo{
		Status:     rets[0].Status,
		StatusCode: rets[0].StatusCode,
		Proto:      rets[0].Proto,
		ProtoMajor: rets[0].ProtoMajor,
		ProtoMinor: rets[0].ProtoMinor,
	}
	return retentity
}

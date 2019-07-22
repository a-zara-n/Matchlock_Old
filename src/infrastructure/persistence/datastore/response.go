package datastore

import (
	"strings"

	"github.com/a-zara-n/Matchlock/src/config"
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
)

//Response は保存用のRepositryです
type Response struct {
	Info   repository.ResponseInfoRepositry
	Header repository.ResponseHeaderRepositry
	Body   repository.ResponseBodyRepositry
}

//NewResponse は
func NewResponse(dbconfig config.DatabaseConfig) repository.ResponseRepositry {
	return &Response{NewResponseInfo(dbconfig), NewResponseHeader(dbconfig), NewResponseBody(dbconfig)}
}

//Insert はResponseを保存します
func (r *Response) Insert(Identifier string, a *aggregate.Response) bool {
	go r.Info.Insert(Identifier, a.Info)
	go r.Header.Insert(Identifier, a.Header)
	go r.Body.Insert(Identifier, a.Body)
	return true
}

//Fetch はaggregate.Responseを取得します
func (r *Response) Fetch(Identifier string) *aggregate.Response {
	return &aggregate.Response{
		Info:   r.Info.Fetch(Identifier),
		Header: r.Header.Fetch(Identifier),
		Body:   r.Body.Fetch(Identifier),
	}
}

//ResponseInfo は保存用のRepositryです
type ResponseInfo struct {
	historyCommon
}

//NewResponseInfo は
func NewResponseInfo(dbconfig config.DatabaseConfig) repository.ResponseInfoRepositry {
	return &ResponseInfo{historyCommon{DBconfig: dbconfig}}
}

//Insert はRequestInfoを保存します
func (r *ResponseInfo) Insert(Identifier string, e *entity.ResponseInfo) bool {
	db := r.OpenDB()
	defer db.Close()
	insertRequestInfo := &ResponseInfoSchema{
		Identifier: Identifier,
		Status:     e.Status,
		StatusCode: e.StatusCode,
		Proto:      e.Proto,
		ProtoMajor: e.ProtoMajor,
		ProtoMinor: e.ProtoMinor,
	}
	db.Create(insertRequestInfo)
	return true
}

//Fetch はentity.ResponseInfoを取得します
func (r *ResponseInfo) Fetch(Identifier string) *entity.ResponseInfo {
	rets := []*ResponseInfoSchema{}
	db := r.OpenDB()
	defer db.Close()
	db.Where("Identifier = ?", Identifier).Find(rets)
	retentity := &entity.ResponseInfo{
		Status:     rets[0].Status,
		StatusCode: rets[0].StatusCode,
		Proto:      rets[0].Proto,
		ProtoMajor: rets[0].ProtoMajor,
		ProtoMinor: rets[0].ProtoMinor,
	}
	return retentity
}

//ResponseHeader は保存用のRepositryです
type ResponseHeader struct {
	historyCommon
}

//NewResponseHeader は
func NewResponseHeader(dbconfig config.DatabaseConfig) repository.ResponseHeaderRepositry {
	return &ResponseHeader{historyCommon{DBconfig: dbconfig}}
}

//Insert はResponseHeaderを保存します
func (r *ResponseHeader) Insert(Identifier string, e *entity.HTTPHeader) bool {
	db := r.OpenDB()
	defer db.Close()
	for _, key := range e.GetKeys() {
		insertHeader := &ResponseHeaderSchema{
			Identifier: Identifier,
			Name:       key,
			Value:      strings.Join(e.Header[key], ","),
		}
		db.Create(insertHeader)
	}
	return true
}

//Fetch はentity.HTTPHeader を取得できる
func (r *ResponseHeader) Fetch(Identifier string) *entity.HTTPHeader {
	rets := []*ResponseHeaderSchema{}
	db := r.OpenDB()
	defer db.Close()
	db.Where("Identifier = ?", Identifier).Find(rets)
	retentity := &entity.HTTPHeader{}
	for _, data := range rets {
		retentity.Header.Add(data.Name, data.Value)
	}
	return retentity
}

//ResponseBody は保存用のRepositryです
type ResponseBody struct {
	historyCommon
}

//NewResponseBody は
func NewResponseBody(dbconfig config.DatabaseConfig) repository.ResponseBodyRepositry {
	return &ResponseBody{historyCommon{DBconfig: dbconfig}}
}

//Insert はBodyを保存します
func (r *ResponseBody) Insert(Identifier string, e *entity.Body) bool {
	db := r.OpenDB()
	defer db.Close()
	insertData := &ResponseBodySchema{
		Identifier: Identifier,
		Body:       e.Body,
		Encodetype: strings.Join(e.Encodetype, ","),
		Length:     e.Length,
	}
	db.Create(insertData)
	return true
}

//Fetch はentity.Bodyを取得出来る
func (r *ResponseBody) Fetch(Identifier string) *entity.Body {
	rets := []*ResponseBodySchema{}
	db := r.OpenDB()
	defer db.Close()
	db.Where("Identifier = ?", Identifier).Find(rets)
	retentity := &entity.Body{
		Body:       rets[0].Body,
		Encodetype: strings.Split(rets[0].Encodetype, ","),
		Length:     rets[0].Length,
	}
	return retentity
}

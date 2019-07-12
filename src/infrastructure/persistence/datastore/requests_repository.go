package datastore

import (
	"net/url"

	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/jinzhu/gorm"
)

//RequestRepositry は保存用のRepositryです
type RequestRepositry struct {
	historyCommon
	Header repository.RequestHeaderRepositry
	Data   repository.RequestDataRepositry
}

//Request は保存用のschemaです
type Request struct {
	gorm.Model
	Identifier string
	IsEdit     bool
	Host       string
	Method     string
	URL        string
	Path       string
	Proto      string
}

//NewRequestRepositry はRequestRepositryを取得する
func NewRequestRepositry(db *gorm.DB) repository.RequestRepositry {
	datarepo := NewRequestDataRepositry(db)
	headrepo := NewRequestHeaderRepositry(db)
	return &RequestRepositry{
		historyCommon{DB: db},
		headrepo,
		datarepo,
	}
}

//Insert はaggregate.Requestに保存されたデータをDBに格納する
func (r *RequestRepositry) Insert(Identifier string, IsEdit bool, a *aggregate.Request) bool {
	go r.insert(Identifier, IsEdit, a.Info)
	go r.Data.Insert(Identifier, IsEdit, a.Data)
	go r.Header.Insert(Identifier, IsEdit, a.Header)
	return true
}

//Insert はRequestInfoを保存します
func (r *RequestRepositry) insert(Identifier string, IsEdit bool, e *entity.RequestInfo) bool {
	insertRequestInfo := &Request{
		Identifier: Identifier,
		IsEdit:     IsEdit,
		Host:       e.Host,
		Method:     e.Method,
		URL:        e.URL.String(),
		Path:       e.Path,
		Proto:      e.Proto,
	}
	r.DB.Create(insertRequestInfo)
	return true
}

//GetRequest は全てのデータを保持した状態のRequestを取得できる
func (r *RequestRepositry) GetRequest(Identifier string, IsEdit bool) *aggregate.Request {
	retentity := &aggregate.Request{
		Info:   r.Select(Identifier, IsEdit),
		Header: r.Header.Select(Identifier, IsEdit),
		Data:   r.Data.Select(Identifier, IsEdit),
	}
	return retentity
}

//Select はRequest Infoを取得出来る
func (r *RequestRepositry) Select(Identifier string, IsEdit bool) *entity.RequestInfo {
	rets := []*Request{}
	r.DB.Where("Identifier = ? AND IsEdit = ?", Identifier, IsEdit).Find(rets)
	u, _ := url.Parse(rets[0].URL)
	retentity := &entity.RequestInfo{
		Host:   rets[0].Host,
		Method: rets[0].Method,
		URL:    u,
		Path:   rets[0].Path,
		Proto:  rets[0].Proto,
	}
	return retentity
}

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
func NewRequestRepositry(Identifier string, IsEdit bool, db *gorm.DB) repository.RequestRepositry {
	datarepo := NewRequestDataRepositry(Identifier, IsEdit, db)
	headrepo := NewRequestHeaderRepositry(Identifier, IsEdit, db)
	return &RequestRepositry{
		historyCommon{Identifier, IsEdit, db},
		headrepo,
		datarepo,
	}
}

//SetIsEdit は編集のフラグを書き換えることができます
func (r *RequestRepositry) SetIsEdit(flag bool) {
	r.Data.SetIsEdit(flag)
	r.Header.SetIsEdit(flag)
	r.IsEdit = flag
}

//SetIdentifier はIdentifierを書き換えることができます
func (r *RequestRepositry) SetIdentifier(id string) {
	r.Data.SetIdentifier(id)
	r.Header.SetIdentifier(id)
	r.Identifier = id
}

//Insert はaggregate.Requestに保存されたデータをDBに格納する
func (r *RequestRepositry) Insert(a *aggregate.Request) bool {
	go r.insert(a.Info)
	go r.Data.Insert(a.Data)
	go r.Header.Insert(a.Header)
	return true
}

//Insert はRequestInfoを保存します
func (r *RequestRepositry) insert(e *entity.RequestInfo) bool {
	insertRequestInfo := &Request{
		Identifier: r.Identifier,
		IsEdit:     r.IsEdit,
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
func (r *RequestRepositry) GetRequest() *aggregate.Request {
	retentity := &aggregate.Request{
		Info:   r.Select(),
		Header: r.Header.Select(),
		Data:   r.Data.Select(),
	}
	return retentity
}

//Select はRequest Infoを取得出来る
func (r *RequestRepositry) Select() *entity.RequestInfo {
	rets := []*Request{}
	r.DB.Where("Identifier = ? AND IsEdit = ?", r.Identifier, r.IsEdit).Find(rets)
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

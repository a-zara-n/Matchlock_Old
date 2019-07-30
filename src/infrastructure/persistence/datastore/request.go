package datastore

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/a-zara-n/Matchlock/src/config"
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
)

//Request は保存用のリポジトリを設定します
type Request struct {
	Info   repository.RequestInfoRepositry
	Header repository.RequestHeaderRepositry
	Data   repository.RequestDataRepositry
}

func NewRequest(dbconfig config.DatabaseConfig) repository.RequestRepositry {
	return &Request{NewRequestInfo(dbconfig), NewRequestHeader(dbconfig), NewRequestData(dbconfig)}
}

func (req *Request) Insert(Identifier string, IsEdit bool, a *aggregate.Request) bool {
	go req.Info.Insert(Identifier, IsEdit, a.Info)
	go req.Data.Insert(Identifier, IsEdit, a.Data)
	go req.Header.Insert(Identifier, IsEdit, a.Header)
	return true
}

func (req *Request) Fetch(Identifier string, IsEdit bool) *aggregate.Request {
	retentity := &aggregate.Request{
		Info:   req.Info.Fetch(Identifier, IsEdit),
		Header: req.Header.Fetch(Identifier, IsEdit),
		Data:   req.Data.Fetch(Identifier, IsEdit),
	}
	return retentity
}

//FetchHostRequests は第一引数で渡したホスト名のリクエストの一覧を探索します
func (req *Request) FetchHostRequests(host string) []*aggregate.Request {
	hosts := req.Info.FetchInfo(host)
	return req.requestFactry(hosts)
}

//FetchHostRequests は第一引数で渡したホスト名のリクエストの一覧を探索します
func (req *Request) requestFactry(infolist []*entity.RequestInfo) []*aggregate.Request {
	var reqestlist []*aggregate.Request
	if len(infolist) > 1 {
		reqestlist = append(reqestlist, req.requestFactry(infolist[1:])...)
	}
	info := infolist[0]
	header := req.Header.FetchHeader(info)
	data := req.Data.FetchData(info)
	return append(reqestlist, &aggregate.Request{
		Info: info, Header: header, Data: data,
	})
}

//RequestInfo は
type RequestInfo struct {
	Common
}

func NewRequestInfo(dbconfig config.DatabaseConfig) repository.RequestInfoRepositry {
	return &RequestInfo{Common: Common{DBconfig: dbconfig}}
}

//Insert はRequestInfoを保存します
func (r *RequestInfo) Insert(Identifier string, IsEdit bool, e *entity.RequestInfo) bool {
	db := r.OpenDB()
	defer db.Close()
	e.URL.RawQuery = ""
	insertRequestInfo := &RequestInfoSchema{
		Identifier: Identifier,
		IsEdit:     IsEdit,
		Host:       e.Host,
		Method:     e.Method,
		URL:        e.URL.String(),
		Path:       e.Path,
		Proto:      e.Proto,
	}
	db.Create(insertRequestInfo)
	return true
}

//Fetch はRequest Infoを取得出来る
func (r *RequestInfo) Fetch(Identifier string, IsEdit bool) *entity.RequestInfo {
	rets := []*RequestInfoSchema{}
	db := r.OpenDB()
	defer db.Close()
	db.Where("Identifier = ? AND IsEdit = ?", Identifier, IsEdit).Find(rets)
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

func (r *RequestInfo) FetchInfo(host string) []*entity.RequestInfo {
	var (
		requestinfos = []*entity.RequestInfo{}
		rets         = []*RequestInfoSchema{}
	)
	db := r.OpenDB()
	db.Debug().Select("Distinct method,url,proto,host,path").
		Where("host LIKE ?", "%"+host+"%").
		Find(&rets)
	for _, info := range rets {
		u, _ := url.Parse(info.URL)
		requestinfos = append(requestinfos, &entity.RequestInfo{
			Host:   info.Host,
			Method: info.Method,
			URL:    u,
			Path:   info.Path,
			Proto:  info.Proto,
		})
	}
	return requestinfos
}

//RequestHeader は
type RequestHeader struct {
	Common
}

//NewRequestHeader はRequestHeaderを取得する
func NewRequestHeader(dbconfig config.DatabaseConfig) repository.RequestHeaderRepositry {
	return &RequestHeader{Common: Common{DBconfig: dbconfig}}
}

//Insert はHTTPHeaderを保存します
func (r *RequestHeader) Insert(Identifier string, IsEdit bool, e *entity.HTTPHeader) bool {
	db := r.OpenDB()
	defer db.Close()
	for _, key := range e.GetKeys() {
		insertHeader := &RequestHeaderSchema{
			Identifier: Identifier,
			IsEdit:     IsEdit,
			Name:       key,
			Value:      strings.Join(e.Header[key], ","),
		}
		db.Create(insertHeader)
	}
	return true
}

//Fetch はentity.HTTPHeaderを取得します
func (r *RequestHeader) Fetch(Identifier string, IsEdit bool) *entity.HTTPHeader {
	rets := []*RequestHeaderSchema{}
	db := r.OpenDB()
	defer db.Close()
	db.Where("Identifier = ? AND IsEdit = ?", Identifier, IsEdit).Find(rets)
	retentity := &entity.HTTPHeader{}
	for _, data := range rets {
		retentity.Header.Add(data.Name, data.Value)
	}
	return retentity
}
func (r *RequestHeader) FetchHeader(info *entity.RequestInfo) *entity.HTTPHeader {
	rets := []RequestHeaderSchema{}
	db := r.OpenDB()
	defer db.Close()
	db.Table("request_header_schemas").Select("name, value, request_header_schemas.is_edit AS is_edit").
		Joins("LEFT JOIN request_info_schemas ON request_info_schemas.identifier = request_header_schemas.identifier").
		Where("host = ? AND path = ? AND method = ?", info.Host, info.Path, info.Method).
		Group("name").Find(&rets)
	header := http.Header{}
	for _, data := range rets {
		header.Add(data.Name, data.Value)
	}
	return &entity.HTTPHeader{header}
}

//RequestData は
type RequestData struct {
	Common
}

func NewRequestData(dbconfig config.DatabaseConfig) repository.RequestDataRepositry {
	return &RequestData{Common: Common{DBconfig: dbconfig}}
}

//Insert はRequestDataを保存します
func (r *RequestData) Insert(Identifier string, IsEdit bool, e *entity.Data) bool {
	db := r.OpenDB()
	defer db.Close()
	for _, key := range e.GetKeys() {
		insertData := &RequestDataSchema{
			Identifier: Identifier,
			IsEdit:     IsEdit,
			Style:      e.Type,
			Name:       key,
			Value:      e.Data[key].(string),
		}

		db.Create(insertData)
	}
	return true
}

func (r *RequestData) insert(parents string, name interface{}, value interface{}, data RequestDataSchema) {
	var pstr []string
	if parents != "" {
		pstr = append(pstr, parents)
	}
	typeof := reflect.TypeOf(value).String()
	ins := func(ret string) {
		data.Parents = strings.Join(pstr, ",")
		data.Name = name.(string)
		data.Value = ret
	}
	switch typeof {
	case "string":
		ins(value.(string))
	case "float64":
		ins(strconv.FormatFloat(value.(float64), 'f', 3, 64))
	case "[]interface {}":
		ss := slicestring(value.([]interface{}))
		ins(strings.Join(ss, ","))
	case "map[interface {}]interface {}":
		for p, v := range value.(map[interface{}]interface{}) {
			r.insert(strings.Join(append(pstr, name.(string)), ","), p, v, data)
		}

	}
}
func slicestring(slice []interface{}) []string {
	ret := []string{}
	for _, val := range slice {
		ret = append(ret, retstring(val)...)
	}
	return ret
}

func retstring(in interface{}) []string {
	switch reflect.TypeOf(in).String() {
	case "string":
		return []string{in.(string)}
	case "float64":
		return []string{strconv.FormatFloat(in.(float64), 'f', 3, 64)}
	case "[]interface {}":
		ss := slicestring(in.([]interface{}))
		return []string{"[" + strings.Join(ss, ",") + "]"}
	}
	return []string{}
}

//Fetch はentity.Dataを取得します
func (r *RequestData) Fetch(Identifier string, IsEdit bool) *entity.Data {
	db := r.OpenDB()
	defer db.Close()
	rets := []*RequestDataSchema{}
	db.Where("Identifier = ? AND IsEdit = ?", Identifier, IsEdit).Find(rets)
	retentity := &entity.Data{}
	retentity.Type = rets[0].Style
	for _, data := range rets {
		retentity.Keys = append(retentity.Keys, data.Name)
		retentity.Data[data.Name] = data.Value
	}
	return retentity
}
func (r *RequestData) FetchData(info *entity.RequestInfo) *entity.Data {
	db := r.OpenDB()
	defer db.Close()
	rets := []RequestDataSchema{}
	db.Select("name, value, style, request_data_schemas.is_edit AS is_edit").
		Joins("LEFT JOIN request_info_schemas ON request_info_schemas.identifier = request_data_schemas.identifier").
		Where("request_info_schemas.host = ? AND request_info_schemas.path = ? AND request_info_schemas.method = ?", info.Host, info.Path, info.Method).
		Group("name").Find(&rets)
	retentity := entity.Data{}
	if len(rets) < 1 {
		return &retentity
	}
	retentity.Type = rets[0].Style
	retentity.Data = map[string]interface{}{}
	for _, data := range rets {
		retentity.Keys = append(retentity.Keys, data.Name)
		retentity.Data[data.Name] = data.Value
	}
	return &retentity
}

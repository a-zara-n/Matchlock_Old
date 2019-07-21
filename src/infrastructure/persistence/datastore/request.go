package datastore

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/jinzhu/gorm"
)

//Request は保存用のリポジトリを設定します
type Request struct {
	Info   repository.RequestInfoRepositry
	Header repository.RequestHeaderRepositry
	Data   repository.RequestDataRepositry
}

func NewRequest(db *gorm.DB) repository.RequestRepositry {
	return &Request{NewRequestInfo(db), NewRequestHeader(db), NewRequestData(db)}
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

//RequestInfo は
type RequestInfo struct {
	historyCommon
}

func NewRequestInfo(db *gorm.DB) repository.RequestInfoRepositry {
	return &RequestInfo{historyCommon{DB: db}}
}

//Insert はRequestInfoを保存します
func (r *RequestInfo) Insert(Identifier string, IsEdit bool, e *entity.RequestInfo) bool {
	insertRequestInfo := &RequestInfoSchema{
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

//Fetch はRequest Infoを取得出来る
func (r *RequestInfo) Fetch(Identifier string, IsEdit bool) *entity.RequestInfo {
	rets := []*RequestInfoSchema{}
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

//RequestHeader は
type RequestHeader struct {
	historyCommon
}

//NewRequestHeader はRequestHeaderを取得する
func NewRequestHeader(db *gorm.DB) repository.RequestHeaderRepositry {
	return &RequestHeader{historyCommon{DB: db}}
}

//Insert はHTTPHeaderを保存します
func (r *RequestHeader) Insert(Identifier string, IsEdit bool, e *entity.HTTPHeader) bool {

	for _, key := range e.GetKeys() {
		insertHeader := &RequestHeaderSchema{
			Identifier: Identifier,
			IsEdit:     IsEdit,
			Name:       key,
			Value:      strings.Join(e.Header[key], ","),
		}
		r.DB.Create(insertHeader)
	}
	return true
}

//Fetch はentity.HTTPHeaderを取得します
func (r *RequestHeader) Fetch(Identifier string, IsEdit bool) *entity.HTTPHeader {
	rets := []*RequestHeaderSchema{}
	r.DB.Where("Identifier = ? AND IsEdit = ?", Identifier, IsEdit).Find(rets)
	retentity := &entity.HTTPHeader{}
	for _, data := range rets {
		retentity.Header.Add(data.Name, data.Value)
	}
	return retentity
}

//RequestData は
type RequestData struct {
	historyCommon
}

func NewRequestData(db *gorm.DB) repository.RequestDataRepositry {
	return &RequestData{historyCommon{DB: db}}
}

//Insert はRequestDataを保存します
func (r *RequestData) Insert(Identifier string, IsEdit bool, e *entity.Data) bool {
	for _, key := range e.GetKeys() {
		insertData := &RequestDataSchema{
			Identifier: Identifier,
			IsEdit:     IsEdit,
			Style:      e.Type,
			Name:       key,
			Value:      e.Data[key].(string),
		}

		r.DB.Create(insertData)
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
	rets := []*RequestDataSchema{}
	r.DB.Where("Identifier = ? AND IsEdit = ?", Identifier, IsEdit).Find(rets)
	retentity := &entity.Data{}
	retentity.Type = rets[0].Type
	for _, data := range rets {
		retentity.Keys = append(retentity.Keys, data.Name)
		retentity.Data[data.Name] = data.Value
	}
	return retentity
}

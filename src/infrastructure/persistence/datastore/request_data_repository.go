package datastore

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/jinzhu/gorm"
)

//RequestDataRepositry は
type RequestDataRepositry struct {
	historyCommon
}

//RequestData はDBを定義した構造体
type RequestData struct {
	gorm.Model
	Identifier string
	IsEdit     bool
	Parents    string
	Name       string
	Value      string
	Style      string //JSON / FORM
	Type       string // ex. int string float
}

//NewRequestDataRepositry はRequestDataを取得する
func NewRequestDataRepositry(db *gorm.DB) repository.RequestDataRepositry {
	return &RequestDataRepositry{
		historyCommon{DB: db},
	}
}

//Insert はRequestDataを保存します
func (r *RequestDataRepositry) Insert(Identifier string, IsEdit bool, e *entity.Data) bool {
	insertData := &RequestData{
		Identifier: Identifier,
		IsEdit:     IsEdit,
		Style:      e.Type,
	}
	for _, key := range e.GetKeys() {
		insertData.Name = key
		insertData.Value = e.Data[key].(string)
		r.DB.Create(insertData)
	}
	return true
}

func (r *RequestDataRepositry) insert(parents string, name interface{}, value interface{}, data RequestData) {
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

//Select はentity.Dataを取得します
func (r *RequestDataRepositry) Select(Identifier string, IsEdit bool) *entity.Data {
	rets := []*RequestData{}
	r.DB.Where("Identifier = ? AND IsEdit = ?", Identifier, IsEdit).Find(rets)
	retentity := &entity.Data{}
	retentity.Type = rets[0].Type
	for _, data := range rets {
		retentity.Keys = append(retentity.Keys, data.Name)
		retentity.Data[data.Name] = data.Value
	}
	return retentity
}

package datastore

import (
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
	Name       string
	Value      string
	Type       string
}

//NewRequestDataRepositry はRequestDataを取得する
func NewRequestDataRepositry(Identifier string, IsEdit bool, db *gorm.DB) repository.RequestDataRepositry {
	return &RequestDataRepositry{
		historyCommon{Identifier, IsEdit, db},
	}
}

//SetIsEdit は編集のフラグを書き換えることができます
func (r *RequestDataRepositry) SetIsEdit(flag bool) { r.IsEdit = flag }

//SetIdentifier はIdentifierを書き換えることができます
func (r *RequestDataRepositry) SetIdentifier(id string) { r.Identifier = id }

//Insert はRequestDataを保存します
func (r *RequestDataRepositry) Insert(e *entity.Data) bool {
	insertData := &RequestData{
		Identifier: r.Identifier,
		IsEdit:     r.IsEdit,
		Type:       e.Type,
	}
	for _, key := range e.GetKeys() {
		insertData.Name = key
		insertData.Value = e.Data[key].(string)
		r.DB.Create(insertData)
	}
	return true
}

//Select はentity.Dataを取得します
func (r *RequestDataRepositry) Select() *entity.Data {
	rets := []*RequestData{}
	r.DB.Where("Identifier = ? AND IsEdit = ?", r.Identifier, r.IsEdit).Find(rets)
	retentity := &entity.Data{}
	retentity.Type = rets[0].Type
	for _, data := range rets {
		retentity.Keys = append(retentity.Keys, data.Name)
		retentity.Data[data.Name] = data.Value
	}
	return retentity
}

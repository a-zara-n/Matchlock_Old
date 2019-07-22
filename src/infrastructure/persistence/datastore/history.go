package datastore

import (
	"fmt"

	"github.com/a-zara-n/Matchlock/src/config"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/jinzhu/gorm"
)

//Common はDB操作を含む全てのstructに含まれる共通構造体です
type Common struct {
	DBconfig config.DatabaseConfig
}

//OpenDB は設定されているコンフィグからDBを呼び出します
func (h *Common) OpenDB() *gorm.DB {
	return h.DBconfig.OpenDB(h.DBconfig.GetConnect())
}

//HTTPHistory はAPIで利用されるJSONのschemaを定義しています
type HTTPHistory struct {
	Common
	JSON []repository.HTTPHistoryDefinitionJSON
}

//NewHTTPHistory は
func NewHTTPHistory(dbconfig config.DatabaseConfig) repository.HistoryRepository {
	return &HTTPHistory{Common: Common{DBconfig: dbconfig}}
}

//Count は
func (hh *HTTPHistory) Count() int {
	var count int
	db := hh.OpenDB()
	defer db.Close()
	db.Table("history_schemas").Count(&count)
	return count
}

//Fetch は
func (hh *HTTPHistory) Fetch(i int) []repository.HTTPHistoryDefinitionJSON {
	historys := hh.JSON
	db := hh.OpenDB()
	defer db.Close()
	db.Table("request_info_schemas").
		Select(" id,request_info_schemas.identifier as identifier,method,host,path,url,param").
		Joins(requestDataTable).
		Where("id >= ?", i).
		Find(&historys)
	fmt.Println(historys)
	return historys
}

//Insert は
func (hh *HTTPHistory) Insert(Identifier string, IsEdit bool) {
	db := hh.OpenDB()
	defer db.Close()
	insertHistory := &HistorySchema{
		Identifier: Identifier,
		IsEdit:     IsEdit,
	}
	db.Create(insertHistory)
}

//Update は
func (hh *HTTPHistory) Update(Identifier string, IsEdit bool) {
	db := hh.OpenDB()
	defer db.Close()
	db.Model(&HistorySchema{}).Where("identifier = ?", Identifier).Update("is_edit", IsEdit)
}

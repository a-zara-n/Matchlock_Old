package datastore

import (
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type historyCommon struct {
	*gorm.DB
}

//HTTPHistoryDefinitionJSON はAPIで利用されるJSONのschemaを定義しています
type HTTPHistoryDefinitionJSON struct {
	ID         int64  `json:"ID"`
	Identifier string `json:"Identifier"`
	Method     string `json:"Method"`
	Host       string `json:"Host"`
	Path       string `json:"Path"`
	URL        string `json:"URL"`
	Param      string `json:"Param"`
}

//NewHTTPHistory はRequestDataを取得する
func NewHTTPHistory(db *gorm.DB) repository.RequestDataRepositry {
	return &RequestDataRepositry{
		historyCommon{DB: db},
	}
}

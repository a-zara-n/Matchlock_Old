package datastore

import (
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

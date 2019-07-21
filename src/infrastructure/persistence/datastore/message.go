package datastore

import (
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/jinzhu/gorm"
)

type HTTPMessage struct {
	historyCommon
}

//NewNewHTTPMessage はRequestDataを取得する
func NewHTTPMessage(db *gorm.DB) repository.HTTPMessageRepository {
	return &HTTPMessage{historyCommon{DB: db}}
}

func (hh *HTTPMessage) Fetch(Identifier string, IsEdit bool) *aggregate.HTTPMessages {
	return &aggregate.HTTPMessages{}
}

package datastore

import (
	"github.com/a-zara-n/Matchlock/src/config"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
)

type HTTPMessage struct {
	Common
	JSON []repository.HTTPMessageDefinitionJSON
}

//NewHTTPMessage はRequestDataを取得する
func NewHTTPMessage(dbconfig config.DatabaseConfig) repository.HTTPMessageRepository {
	return &HTTPMessage{Common: Common{DBconfig: dbconfig}}
}

//Fetch はhttpのrequest + edit request + responseを返します
func (hm *HTTPMessage) Fetch(Identifier string) []repository.HTTPMessageDefinitionJSON {
	httpmessage := hm.JSON
	db := hm.OpenDB()
	defer db.Close()
	db.Table("history_schemas").
		Select(httpData["SELECT"]).
		Joins(httpData["NoEditReq"]).
		Joins(httpData["EditReq"]).
		Joins(httpData["Response"]).
		Where("history_schemas.identifier = ?", Identifier).
		Find(&httpmessage)
	return httpmessage
}

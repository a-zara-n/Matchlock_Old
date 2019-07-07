package aggregate

import "github.com/a-zara-n/Matchlock/src/domain/value"

//HTTPPair はHTTPリクエストとレスポンスをまとめたものになります
type HTTPPair struct {
	value.Identifier
	IsEdit      bool
	Request     Request
	EditRequest Request
	Response    Response
}

//HTTPDataDefinitionByJSON はWSで利用されるJSONのデータ定義
type HTTPDataDefinitionByJSON struct {
	Identifier         string `json:"Identifier"`
	RequestMethod      string `json:"Method"`
	RequestPath        string `json:"Path"`
	RequestProto       string `json:"Proto"`
	RequestHost        string `json:"Host"`
	RequestHeaders     string `json:"Header"`
	RequestParam       string `json:"Param"`
	RequestEditMethod  string `json:"EditMethod"`
	RequestEditPath    string `json:"EditPath"`
	RequestEditProto   string `json:"EditProto"`
	RequestEditHost    string `json:"EditHost"`
	RequestEditHeaders string `json:"EditHeader"`
	RequestEditParam   string `json:"EditParam"`
	ResponseHeaders    string `json:"ResHeader"`
	Body               string `json:"ReqBody"`
}

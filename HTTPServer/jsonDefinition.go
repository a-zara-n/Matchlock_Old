package httpserver

//message json struct
type Message struct {
	Type string `json:"Type"`
	Data string `json:"Data"`
}

type APIresponse struct {
	Data interface{} `json:"Data"`
}

type historyData struct {
	ID         int64  `json:"ID"`
	Identifier string `json:"Identifier"`
	Method     string `json:"Method"`
	Host       string `json:"Host"`
	Path       string `json:"Path"`
	URL        string `json:"URL"`
	Param      string `json:"Param"`
}

type httpdata struct {
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

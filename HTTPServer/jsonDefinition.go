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

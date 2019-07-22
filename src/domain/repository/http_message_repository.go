package repository

type HTTPMessageRepository interface {
	Fetch(Identifier string) []HTTPMessageDefinitionJSON
}
type HTTPMessageDefinitionJSON struct {
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

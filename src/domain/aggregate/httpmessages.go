package aggregate

import (
	"log"
	"net/http"
	"reflect"

	"github.com/a-zara-n/Matchlock/src/domain/value"
)

//HTTPMessages はHTTPリクエストとレスポンスをまとめたものになります
type HTTPMessages struct {
	value.Identifier
	IsEdit      bool
	Request     *Request
	EditRequest *Request
	Response    *Response
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

//NewHTTPMessage は新規でHTTPMessageを生成します
func NewHTTPMessage() HTTPMessages {
	return HTTPMessages{
		Identifier:  value.Identifier{},
		IsEdit:      false,
		Request:     &Request{},
		EditRequest: &Request{},
		Response:    &Response{},
	}
}

//IsEdited はリクエストが編集されたかを確認するためのmethod
func (h *HTTPMessages) IsEdited() bool {
	var flag bool
	for _, key := range h.EditRequest.Header.GetKeys() {
		if h.Request.Header.Get(key) != h.EditRequest.Header.Header.Get(key) {
			flag = true
		}
	}
	if flag || !reflect.DeepEqual(h.Request.Info, h.EditRequest.Info) || !reflect.DeepEqual(h.Request.Data, h.EditRequest.Data) {
		log.Println("変更が発生しました")
		h.IsEdit = true
		return true
	}
	return false
}

//SetRequest は
func (h *HTTPMessages) SetRequest(req *http.Request) {
	h.Identifier.Set(req.URL.String())
	h.Request = NewHTTPRequestByRequest(req)
}

//SetEditedRequest は
func (h *HTTPMessages) SetEditedRequest(req *http.Request) {
	h.EditRequest = NewHTTPRequestByRequest(req)
}

//SetResponse は
func (h *HTTPMessages) SetResponse(resp *http.Response) {
	h.Response = NewHTTPResponseByResponse(resp)
}

//FetchRequest は
func (h *HTTPMessages) FetchRequest() *http.Request {
	return h.Request.GetHTTPRequestByRequest()
}

//FetchEditRequest は
func (h *HTTPMessages) FetchEditRequest() *http.Request {
	return h.EditRequest.GetHTTPRequestByRequest()
}

//FetchResponse は
func (h *HTTPMessages) FetchResponse() *http.Response {
	return h.Response.GetHTTPRequestByResponse()

}

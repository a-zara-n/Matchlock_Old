package aggregate

import (
	"net/http"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//Request はHTTPを
type Request struct {
	Info   *entity.RequestInfo
	Header *entity.HTTPHeader
	Data   *entity.Data
}

//NewHTTPRequestByRequest はhttp.Requestを利用しaggregate.Requestを取得できます
func NewHTTPRequestByRequest(req *http.Request) *Request {
	request := &Request{
		Info:   &entity.RequestInfo{},
		Header: &entity.HTTPHeader{},
		Data:   &entity.Data{},
	}
	request.SetHTTPRequestByRequest(req)
	return request
}

//SetHTTPRequestByRequest はhttp.Requestを元に設定をします
func (r *Request) SetHTTPRequestByRequest(req *http.Request) {
	r.Info.SetRequestINFO(req)
	r.Header.SetHTTPHeader(req.Header)
	if req.Method == "POST" {
		req.Body = r.Data.SetDataByHTTPBody(req.Body)
		return
	}
	r.Data.SetData(req.URL.RawQuery)
}

//SetHTTPRequestByString はstringを元に設定をします
func (r *Request) SetHTTPRequestByString(req string) {

}

//GetHTTPRequestByRequest はhttp.Requestを生成します
func (r *Request) GetHTTPRequestByRequest() *http.Request {
	return &http.Request{}
}

//GetHTTPRequestByString は文字列のhttp requestを生成します
func (r *Request) GetHTTPRequestByString() string {
	return ""
}

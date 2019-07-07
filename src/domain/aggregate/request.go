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

//SetHTTPRequestByRequest はhttp.Requestを元に設定をします
func (r *Request) SetHTTPRequestByRequest(req *http.Request) {
	go r.Info.SetRequestINFO(req)
	go r.Header.SetHTTPHeader(req.Header)
	req.Body = r.Data.SetDataByHTTPBody(req.Body)
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

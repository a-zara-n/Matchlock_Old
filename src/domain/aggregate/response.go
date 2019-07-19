package aggregate

import (
	"net/http"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//Response はHTTPのレスポンスに関するエンティティを集約した構造体
type Response struct {
	Info    *entity.ResponseInfo
	Header  *entity.HTTPHeader
	Body    *entity.Body
	Request *Request
}

//NewHTTPResponseByResponse はhttp.Responseを利用してaggregate.Responseを返します
func NewHTTPResponseByResponse(resp *http.Response) *Response {
	response := &Response{
		Info:    &entity.ResponseInfo{},
		Header:  &entity.HTTPHeader{},
		Body:    &entity.Body{},
		Request: &Request{},
	}
	response.SetHTTPResponseByResponse(resp)
	return response
}

//SetHTTPResponseByResponse は
func (resp *Response) SetHTTPResponseByResponse(response *http.Response) {
	resp.Info.Set(response)
	resp.Header.SetHTTPHeader(response.Header)
	response.Body = resp.Body.Set(response.Body, response.TransferEncoding)
	resp.Request = NewHTTPRequestByRequest(response.Request)
}

//GetHTTPRequestByResponse は
func (resp *Response) GetHTTPRequestByResponse() *http.Response {
	retresponse := resp.Info.Fetch()
	retresponse.Header = resp.Header.Header
	retresponse.Body = resp.Body.Get()
	retresponse.ContentLength = resp.Body.GetLength()
	retresponse.TransferEncoding = resp.Body.Encodetype
	retresponse.Request = resp.Request.GetHTTPRequestByRequest()
	return retresponse
}

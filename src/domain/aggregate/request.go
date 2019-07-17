package aggregate

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//Request はHTTPを
type Request struct {
	Info   *entity.RequestInfo
	Header *entity.HTTPHeader
	Data   *entity.Data
}

//NewHTTPRequestByRequest はhttp.Requestを利用しaggregate.Requestを取得できます
func NewHTTPRequestByRequest(request *http.Request) *Request {
	requestuest := &Request{
		Info:   &entity.RequestInfo{},
		Header: &entity.HTTPHeader{},
		Data:   &entity.Data{},
	}
	requestuest.SetHTTPRequestByRequest(request)
	return requestuest
}

//SetHTTPRequestByRequest はhttp.Requestを元に設定をします
func (req *Request) SetHTTPRequestByRequest(request *http.Request) {
	req.Info.SetRequestINFO(request)
	req.Header.SetHTTPHeader(request.Header)
	if request.Method == "POST" {
		request.Body = req.Data.SetDataByHTTPBody(request.Body)
		return
	}
	req.Data.SetData(request.URL.RawQuery)
}

//SetHTTPRequestByString はstringを元に設定をします
func (req *Request) SetHTTPRequestByString(request string) {
	slicehttprequest := strings.Split(request, "\n")
	req.Info.SetStatusLine(slicehttprequest[0])
	req.Header.SetStringHeader(strings.Join(slicehttprequest[1:len(slicehttprequest)-2], "\n"))
	req.Data.SetData(slicehttprequest[len(slicehttprequest)-1])
}

//GetHTTPRequestByRequest はhttp.Requestを生成します
func (req *Request) GetHTTPRequestByRequest() *http.Request {
	proto := strings.Split(strings.Split(req.Info.Proto, "/")[1], ".")
	major, err := strconv.Atoi(proto[0])
	if err != nil {
		major = 1
	}
	minor, err := strconv.Atoi(proto[1])
	if err != nil {
		minor = 0
	}
	return &http.Request{
		Host:          req.Info.Host,
		Method:        req.Info.Method,
		URL:           req.Info.URL,
		Proto:         req.Info.Proto,
		ProtoMajor:    major,
		ProtoMinor:    minor,
		Header:        req.Header.Header,
		Body:          req.Data.GetIoReadCloser(),
		ContentLength: int64(len(req.Data.FetchData())),
	}
}

//GetHTTPRequestByString は文字列のhttp requestを生成します
func (req *Request) GetHTTPRequestByString() string {
	return req.Info.GetStatusLine() + "\n" + req.Header.GetStringHeader() + "\n\n" + req.Data.FetchData()
}

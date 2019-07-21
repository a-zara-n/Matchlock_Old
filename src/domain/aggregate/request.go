package aggregate

import (
	"net/http"
	"reflect"
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

//DiffUpdate はaggrigateリクエスの差分を反映させます
func (req *Request) DiffUpdate(request *Request) {
	if !reflect.DeepEqual(req.Info, request.Info) {
		req.Info = request.Info
	}
	if !reflect.DeepEqual(req.Header, request.Header) {
		req.Header = request.Header
	}
	if !reflect.DeepEqual(req.Data, request.Data) {
		req.Data = request.Data
	}
}

//SetHTTPRequestByRequest はhttp.Requestを元に設定をします
func (req *Request) SetHTTPRequestByRequest(request *http.Request) {
	req.Info.SetRequestINFO(request)
	req.Header.SetHTTPHeader(request.Header)
	switch request.Method {
	case "POST", "DELETE", "PUT":
		request.Body = req.Data.SetDataByHTTPBody(request.Body)
	default:
		req.Data.SetData(request.URL.RawQuery)
	}
}

//SetHTTPRequestByString はstringを元に設定をします
func (req *Request) SetHTTPRequestByString(request string) {
	slicehttprequest := strings.Split(request, "\n")
	req.Info.SetStatusLine(slicehttprequest[0])
	req.Header.SetStringHeader(strings.Join(slicehttprequest[1:len(slicehttprequest)-2], "\n"))
	switch strings.Split(slicehttprequest[0], " ")[0] {
	case "POST", "DELETE", "PUT":
		req.Data.SetData(slicehttprequest[len(slicehttprequest)-1])
	default:
		pathquery := strings.Split(slicehttprequest[0], " ")[1]
		if query := strings.Split(pathquery, "?"); len(query) > 1 {
			req.Data.SetData(query[1])
		} else {
			req.Data.SetData("")
		}
	}
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
	var stringrequest string
	switch req.Info.Method {
	case "POST", "DELETE", "PUT":
		stringrequest = req.Info.GetStatusLine() + "\n" + req.Header.GetStringHeader() + "\n\n" + req.Data.FetchData()
	default:
		stringrequest = req.Info.GetStatusLine(req.Data.FetchData()) + "\n" + req.Header.GetStringHeader() + "\n"
	}
	return stringrequest
}

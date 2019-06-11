package scanner

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/WestEast1st/Matchlock/extractor"
)

type Request struct {
	Method string
	URL    string
	Proto  string
}

func (r *Request) GetRequest(host string) []http.Request {
	likehost := "%" + host + "%"
	db.Table = Request{}
	reqdb := db.OpenDatabase()
	var request []Request
	reqdb.
		Table("Requests").
		Select("Distinct method,url,proto").
		Where("host LIKE ?", likehost).
		Find(&request)
	retReq := []http.Request{}
	for _, r := range request {
		u, _ := url.Parse(r.URL)
		tr := http.Request{
			Method: r.Method,
			Host:   host,
			URL:    u,
			Proto:  r.Proto,
			Header: http.Header{},
		} //tmprequest
		rh, rd, strs := RequestHeader{}, RequestData{}, []string{}
		for _, h := range rh.GetHeader(tr.URL.Host, tr.URL.Path, r.Method) {
			tr.Header.Add(h.Name, h.Value)
		}

		for _, d := range rd.GetData(tr.URL.Host, tr.URL.Path, r.Method) {
			strs = append(strs, d.Name+"="+d.Value)
		}
		tr.Body = extractor.GetIOReadCloser(strings.Join(strs, "&"))
		retReq = append(retReq, tr)
	}
	return retReq
}

type RequestHeader struct {
	Name   string
	Value  string
	IsEdit bool
}

func (r *RequestHeader) GetHeader(host string, path string, method string) []RequestHeader {
	reqdb := db.OpenDatabase()
	var requestHeader []RequestHeader
	reqdb.
		Table("request_headers").
		Select("name, value, request_headers.is_edit AS is_edit").
		Joins("LEFT JOIN requests ON requests.identifier = request_headers.identifier").
		Where("host = ? AND path = ? AND method = ?", host, path, method).
		Group("name").
		Find(&requestHeader)
	return requestHeader
}

type RequestData struct {
	Name   string
	Value  string
	IsEdit bool
}

func (r *RequestData) GetData(host string, path string, method string) []RequestData {
	reqdb := db.OpenDatabase()
	var requestData []RequestData
	reqdb.
		Table("request_data").
		Select("name, value, request_data.is_edit AS is_edit").
		Joins("LEFT JOIN requests ON requests.identifier = request_data.identifier").
		Where("host = ? AND path = ? AND method = ?", host, path, method).
		Group("name").
		Find(&requestData)
	return requestData
}
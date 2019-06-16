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
	return r.getRequests(host, request)
}

func (r *Request) getRequests(host string, requests []Request) []http.Request {
	u, _ := url.Parse(requests[0].URL)
	rh, rd := RequestHeader{}, RequestData{}
	tr := []http.Request{{
		Method: requests[0].Method,
		Host:   host,
		URL:    u,
		Proto:  requests[0].Proto,
		Header: setHeader(http.Header{}, rh.GetHeader(u.Host, u.Path, requests[0].Method)),
		Body: extractor.GetIOReadCloser(strings.Join(
			getBodySlice(rd.GetData(u.Host, u.Path, requests[0].Method)), "&")),
	}}
	if len(requests) > 1 {
		return append(tr, r.getRequests(host, requests[1:])...)
	}
	return tr
}

func setHeader(header http.Header, hs []RequestHeader) http.Header {
	for _, h := range hs {
		header.Add(h.Name, h.Value)
	}
	return header
}

func getBodySlice(d []RequestData) []string {
	data := []string{d[0].Name + "=" + d[0].Value}
	if len(d) > 1 {
		return append(data, getBodySlice(d[1:])...)
	}
	return data
}

func merge(m1, m2 map[string]string) map[string]string {
	ans := map[string]string{}
	for k, v := range m1 {
		ans[k] = v
	}
	for k, v := range m2 {
		ans[k] = v
	}
	return (ans)
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

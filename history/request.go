package history

import (
	"net/http"
	"net/url"
	"sort"
	"strings"

	"../extractor"
	"github.com/jinzhu/gorm"
)

type Request struct {
	gorm.Model
	Identifier       string
	IsEdit           bool
	Host             string
	Method           string
	URL              string
	Proto            string
	ContentLength    int64
	TransferEncoding string
}

func (r *Request) SetRequest(req *http.Request, bstr string) {
	headerKey := []string{}
	for k := range req.Header {
		headerKey = append(headerKey, k)
	}
	sort.Strings(headerKey)
	for _, v := range headerKey {
		heade :=
			RequestHeader{
				Identifier: r.Identifier,
				IsEdit:     r.IsEdit,
				Method:     req.Method,
				Host:       req.Host,
				Path:       req.URL.Path,
			}
		heade.SetHeader(v, strings.Join(req.Header[v], ","))
	}
	fordata := func(datas []string) {
		for _, ds := range datas {
			data :=
				RequestData{
					Identifier: r.Identifier,
					IsEdit:     r.IsEdit,
					Method:     req.Method,
					Host:       req.Host,
					Path:       req.URL.Path,
				}
			d := strings.Split(ds, "=")
			if len(d) > 0 {
				if len(d) > 1 {
					data.SetData(d[0], d[1])
				}
			}
		}
	}

	for _, datas := range [][]string{strings.Split(req.URL.RawQuery, "&"), strings.Split(bstr, "&")} {
		fordata(datas)
	}
	db.Table = Request{}
	r.Host, r.Method, r.Proto, r.URL, r.ContentLength, r.TransferEncoding =
		req.Host, req.Method, req.Proto, req.URL.String(), req.ContentLength, strings.Join(req.TransferEncoding, ",")
	db.Insert(r)
}

func (r *Request) GetHost() []Request {
	db.Table = Request{}
	reqdb := db.OpenDatabase()
	var request []Request
	reqdb.
		Select("DISTINCT host,host").
		Find(&request)
	return request
}

func (r *Request) GetRequest(host string) []http.Request {
	host = "%" + host + "%"
	db.Table = Request{}
	reqdb := db.OpenDatabase()
	var request []Request
	reqdb.
		Select("Distinct method,url,proto").
		Where("host LIKE ?", host).
		Find(&request)
	retReq := []http.Request{}
	for _, r := range request {
		u, _ := url.Parse(r.URL)
		tr := http.Request{
			Method: r.Method,
			Host:   r.Host,
			URL:    u,
			Proto:  r.Proto,
			Header: http.Header{},
		} //tmprequest
		rh, rd, strs := RequestHeader{}, RequestData{}, []string{}
		for _, h := range rh.GetHeader(tr.URL.Host, tr.URL.Path, r.Method) {
			if h.Name == "Cookie" {
				continue
			}
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
	gorm.Model
	Identifier string
	Method     string
	Host       string
	Path       string
	Name       string
	Value      string
	IsEdit     bool
}

func (r *RequestHeader) SetHeader(name string, value string) {
	r.Name = name
	r.Value = value
	db.Table = RequestHeader{}
	db.Insert(r)
}

func (r *RequestHeader) GetHeader(host string, path string, method string) []RequestHeader {
	db.Table = RequestHeader{}
	reqdb := db.OpenDatabase()
	var requestHeader []RequestHeader
	reqdb.
		Select("*").
		Where("host = ? AND path = ? AND method = ?", host, path, method).
		Group("name").
		Find(&requestHeader)
	return requestHeader
}

type RequestData struct {
	gorm.Model
	Identifier string
	Method     string
	Host       string
	Path       string
	Name       string
	Value      string
	IsEdit     bool
}

func (r *RequestData) SetData(name string, value string) {
	r.Name = name
	r.Value = value
	db.Table = RequestData{}
	db.Insert(r)
}

func (r *RequestData) GetData(host string, path string, method string) []RequestData {
	db.Table = RequestData{}
	reqdb := db.OpenDatabase()
	var requestData []RequestData
	reqdb.
		Select("*").
		Where("host = ? AND path = ? AND method = ?", host, path, method).
		Group("name").
		Find(&requestData)
	return requestData
}

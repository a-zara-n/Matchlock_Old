package main

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"./datastore"
	"github.com/jinzhu/gorm"
)

var db = datastore.Database{Database: "./test.db"}

type History struct {
	gorm.Model
	Identifier string
}

func (h *History) SetIdentifier(id string) {
	h.Identifier = id
}

func (h *History) MemoryRequest(r *http.Request, isEdit bool, bstr string) {
	fmt.Println(bstr)
	var req = Request{Identifier: h.Identifier, IsEdit: isEdit}
	req.SetRequest(r, bstr)
	db.Table = History{}
	db.Insert(h)
}

func (h *History) MemoryResponse(r *http.Response, bstr string) {
	var res = Response{Identifier: h.Identifier}
	res.SetResponse(r)
}

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
	r.Host = req.Host
	r.Method = req.Method
	r.Proto = req.Proto
	r.URL = req.URL.String()
	r.ContentLength = req.ContentLength
	r.TransferEncoding = strings.Join(req.TransferEncoding, ",")
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

func (r *Request) GetRequest(host string) []Request {
	host = "%" + host + "%"
	db.Table = Request{}
	reqdb := db.OpenDatabase()
	var request []Request
	reqdb.
		Select("DISTINCT method, url").
		Where("host LIKE ?", host).
		Find(&request)
	return request
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
func (r *RequestHeader) GetHeader(host string) []RequestHeader {
	host = "%" + host + "%"
	db.Table = RequestHeader{}
	reqdb := db.OpenDatabase()
	var requestHeader []RequestHeader
	reqdb.
		Select("DISTINCT method, url").
		Where("host LIKE ?", host).
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

func (r *RequestData) GetData() {

}

type Response struct {
	gorm.Model
	Identifier string
}

func (r *Response) SetResponse(res *http.Response) {

}
func (r *Response) GetResponse() {

}

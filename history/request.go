package history

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

type Request struct {
	gorm.Model
	Identifier string
	IsEdit     bool
	Host       string
	Method     string
	URL        string
	Path       string
	Proto      string
}

func (r *Request) SetRequest(req *http.Request, bstr string) {
	db.Table = Request{}
	r.Host, r.Method, r.Proto, r.URL, r.Path =
		req.Host, req.Method, req.Proto, req.URL.String(), req.URL.Path
	db.Insert(r)
}

type RequestHeader struct {
	gorm.Model
	Identifier string
	Name       string
	Value      string
	IsEdit     bool
}

func (r *Request) SetHeader(header http.Header) {
	ise := r.IsEdit
	id := r.Identifier
	db.Table = RequestHeader{}
	for name, value := range header {
		insHeader := &RequestHeader{
			Identifier: id,
			Name:       name,
			Value:      strings.Join(value, ","),
			IsEdit:     ise,
		}
		db.Insert(insHeader)
	}
}

type RequestData struct {
	gorm.Model
	Identifier       string
	Name             string
	Value            string
	TransferEncoding string
	IsEdit           bool
}

func (r *Request) SetData(bstr string, length int64, enctype []string) {
	id := r.Identifier
	ise := r.IsEdit
	db.Table = RequestData{}
	for _, params := range strings.Split(bstr, "&") {
		param := strings.Split(params, "=")
		data := &RequestData{
			Identifier:       id,
			Name:             param[0],
			Value:            strings.Join(param[1:], "="),
			TransferEncoding: strings.Join(enctype, ","),
			IsEdit:           ise,
		}
		db.Insert(data)

	}
}

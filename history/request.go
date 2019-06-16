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

func (r *Request) SetRequest(req *http.Request) {
	db.Table = Request{}
	req.URL.RawQuery = ""
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
	var insertHeader func(headerKeys []string)
	db.Table = RequestHeader{}
	insertHeader = func(hkeys []string) {
		recursiveExec(hkeys, insertHeader)
		db.Insert(&RequestHeader{
			Identifier: r.Identifier,
			Name:       hkeys[0],
			Value:      quoteEscape(strings.Join(header[hkeys[0]], ",")),
			IsEdit:     r.IsEdit,
		})
	}
	insertHeader(getKeys(header))
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
	if bstr == "" {
		return
	}
	var innsertData func(params []string)
	db.Table = RequestData{}
	innsertData = func(params []string) {
		recursiveExec(params, innsertData)
		param := strings.Split(params[0], "=")
		db.Insert(&RequestData{
			Identifier:       r.Identifier,
			Name:             param[0],
			Value:            quoteEscape(strings.Join(param[1:], "=")),
			TransferEncoding: strings.Join(enctype, ","),
			IsEdit:           r.IsEdit,
		})
	}
	innsertData(strings.Split(bstr, "&"))
}

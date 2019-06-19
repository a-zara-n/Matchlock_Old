package history

import (
	"net/http"
	"strings"

	"github.com/a-zara-n/Matchlock/datastore"
	"github.com/a-zara-n/Matchlock/shared"
	"github.com/jinzhu/gorm"
)

type Response struct {
	gorm.Model
	Identifier string
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
}

func (r *Response) SetResponse(res *http.Response) {
	datastore.DB.Insert(&Response{
		Identifier: r.Identifier,
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Proto:      res.Proto,
		ProtoMajor: res.ProtoMajor,
		ProtoMinor: res.ProtoMinor,
	})
}

type ResponseHeader struct {
	gorm.Model
	Identifier string
	Name       string
	Value      string
}

func (r *Response) SetResponseHeader(header http.Header) {
	var insertHeader func(headerKeys []string)
	insertHeader = func(hkeys []string) {
		shared.RecursiveExec(hkeys, insertHeader)
		datastore.DB.Insert(&ResponseHeader{
			Identifier: r.Identifier,
			Name:       hkeys[0],
			Value:      strings.Join(header[hkeys[0]], ","),
		})
	}
	insertHeader(shared.GetKeys(header))
}

type ResponseBody struct {
	gorm.Model
	Identifier string
	Body       string
	Encodetype string
	Length     int64
}

func (r *Response) SetResponseBody(body string, length int64, tenc []string) {
	datastore.DB.Insert(&ResponseBody{
		Identifier: r.Identifier,
		Body:       body,
		Encodetype: strings.Join(tenc, ","),
		Length:     length,
	})
}

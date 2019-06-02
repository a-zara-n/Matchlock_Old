package history

import (
	"net/http"
	"strings"

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

	resp := &Response{
		Identifier: r.Identifier,
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Proto:      res.Proto,
		ProtoMajor: res.ProtoMajor,
		ProtoMinor: res.ProtoMinor,
	}
	db.Table = Request{}
	db.Insert(resp)
}

type ResponseHeader struct {
	gorm.Model
	Identifier string
	Name       string
	Value      string
}

func (r *Response) SetResponseHeader(header http.Header) {

	db.Table = &ResponseHeader{}
	for name, data := range header {
		respH := &ResponseHeader{
			Identifier: r.Identifier,
			Name:       name,
			Value:      strings.Join(data, ","),
		}
		db.Insert(respH)
	}
}

type ResponseBody struct {
	gorm.Model
	Identifier string
	Body       string
	Encodetype string
	Length     int64
}

func (r *Response) SetResponseBody(body string, length int64, tenc []string) {
	respB := &ResponseBody{
		Identifier: r.Identifier,
		Body:       body,
		Encodetype: strings.Join(tenc, ","),
		Length:     length,
	}
	db.Table = ResponseBody{}
	db.Insert(respB)
}

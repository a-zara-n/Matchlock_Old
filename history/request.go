package history

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/WestEast1st/Matchlock/datastore"
	"github.com/WestEast1st/Matchlock/shared"
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
	req.URL.RawQuery = ""
	r.Host, r.Method, r.Proto, r.URL, r.Path =
		req.Host, req.Method, req.Proto, req.URL.String(), req.URL.Path
	datastore.DB.Insert(r)
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
	insertHeader = func(hkeys []string) {
		shared.RecursiveExec(hkeys, insertHeader)
		datastore.DB.Insert(&RequestHeader{
			Identifier: r.Identifier,
			Name:       hkeys[0],
			Value:      shared.QuoteEscape(strings.Join(header[hkeys[0]], ",")),
			IsEdit:     r.IsEdit,
		})
	}
	insertHeader(shared.GetKeys(header))
}

type RequestData struct {
	gorm.Model
	Identifier       string
	Name             string
	Value            string
	Type             string
	TransferEncoding string
	IsEdit           bool
}

func (r *Request) SetData(bstr string, length int64, enctype []string) {
	if bstr == "" {
		return
	}
	var typestr string
	if shared.CheckRegexp(`^{(\".*\":\"?.*\"?,?)+[^,]}$`, bstr) {
		typestr = "JSON"
		ret := []string{}
		bbyte := []byte(bstr)
		var list interface{}
		json.Unmarshal(bbyte, &list)
		hoge := list.(map[string]interface{})
		for key := range hoge {
			ret = append(ret, key+"="+hoge[key].(string))
		}
		bstr = strings.Join(ret, "&")
	}
	var innsertData func(params []string)
	innsertData = func(params []string) {
		shared.RecursiveExec(params, innsertData)
		param := strings.Split(params[0], "=")
		datastore.DB.Insert(&RequestData{
			Identifier:       r.Identifier,
			Name:             strings.Replace(param[0], ":", "\x3a", -1),
			Value:            shared.QuoteEscape(strings.Join(param[1:], "=")),
			Type:             typestr,
			TransferEncoding: strings.Join(enctype, ","),
			IsEdit:           r.IsEdit,
		})
	}
	innsertData(strings.Split(bstr, "&"))
}

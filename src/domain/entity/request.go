package entity

import (
	"net/http"
	"net/url"
)

//RequestInfo は
type RequestInfo struct {
	Host   string
	Method string
	URL    *url.URL
	Path   string
	Proto  string
}

//SetRequestINFO はリクエストの情報を設定します
func (ri *RequestInfo) SetRequestINFO(r *http.Request) {
	ri.Host, ri.Method, ri.URL, ri.Path, ri.Proto =
		r.Host, r.Method, r.URL, ri.URL.Path, r.Proto
}

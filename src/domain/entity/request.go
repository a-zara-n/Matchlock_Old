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
	ri.Host = r.Host
	ri.Method = r.Method
	ri.URL = r.URL
	ri.Path = ri.URL.Path
	ri.Proto = r.Proto
}

func (ri *RequestInfo) GetStartLine() []string {
	return []string{ri.Method, ri.Path, ri.Proto}
}

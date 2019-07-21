package entity

import (
	"net/http"
	"net/url"
	"strings"
)

//RequestInfo は
type RequestInfo struct {
	Host   string
	Method string
	URL    *url.URL
	Query  Data
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
	ri.Query = Data{}
	if r.URL.RawQuery != "" {
		ri.Query.SetData(r.URL.RawQuery)
	}
}

func (ri *RequestInfo) GetStatusLine(query ...string) string {
	if len(query) > 0 {
		ri.Query.SetData(query[0])
	}
	statusline := []string{ri.Method, ri.Path, ri.Proto}
	if len(ri.Query.GetKeys()) != 0 {
		statusline[1] += "?" + ri.Query.FetchData()
	}
	return strings.Join(statusline, " ")
}

func (ri *RequestInfo) SetStatusLine(startline string) {
	sline := strings.Split(startline, " ")
	pathandquery := strings.Split(sline[1], "?")
	ri.Method = sline[0]
	ri.Proto = sline[2]
	ri.Path = pathandquery[0]
	if len(pathandquery) == 2 && len(pathandquery[1]) > 0 {
		ri.URL.RawQuery = pathandquery[1]
		ri.Query.SetData(pathandquery[1])
	}
}

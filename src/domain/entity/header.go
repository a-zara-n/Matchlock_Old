package entity

import (
	"net/http"
	"sort"
	"strings"
)

//HTTPHeader はhttp.Headerを設定します
type HTTPHeader struct {
	http.Header
}

//SetHTTPHeader はHTTPheaderを設定することが可能です
func (h *HTTPHeader) SetHTTPHeader(header http.Header) {
	h.Header = header
}

//GetKeys はHeaderのnamesを取得します
func (h *HTTPHeader) GetKeys() []string {
	keys := []string{}
	for k := range h.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (h *HTTPHeader) SetStringHeader(header string) {
	h.Header = map[string][]string{}
	for _, value := range strings.Split(header, "\n") {
		head := strings.Split(value, ": ")
		for _, v := range strings.Split(head[1], ",") {
			h.Header.Add(head[0], v)
		}
	}
}

func (h *HTTPHeader) GetStringHeader() string {
	header := []string{}
	keys := h.GetKeys()
	sort.Strings(keys)
	for _, key := range keys {
		header = append(header, key+": "+strings.Join(h.Header[key], ","))
	}
	return strings.Join(header, "\n")
}

package entity

import "net/http"

//HTTPHeader はhttp.Headerを設定します
type HTTPHeader struct {
	http.Header
}

//GetKeys はHeaderのnamesを取得します
func (h *HTTPHeader) GetKeys() []string {
	keys := []string{}
	for k := range h.Header {
		keys = append(keys, k)
	}
	return keys
}

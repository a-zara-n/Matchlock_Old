package extractor

import (
	"bytes"
	"io"
	"net/http"
	"sort"
	"strings"
)

func GetStringByRequest(r *http.Request) string {
	return strings.Join([]string{
		getStartLine(r),
		strings.Join(GetHeader(r.Header), "\n"),
		"",
		GetStringBody(r.Body),
	}, "\n")
}

func GetStringBody(b io.ReadCloser) string {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(b)
	return bufbody.String()
}

func GetHeader(h http.Header) []string {
	headerKey, headerSlice := []string{}, []string{}
	for k := range h {
		headerKey = append(headerKey, k)
	}
	sort.Strings(headerKey)
	for _, v := range headerKey {
		headerSlice =
			append(
				headerSlice,
				strings.Join([]string{v, strings.Join(h[v], ",")}, ": "),
			)
	}
	return headerSlice
}

func getQuery(rq string) string {
	if rq != "" {
		return "?" + rq
	}
	return rq
}

func getStartLine(r *http.Request) string {
	return strings.Join(
		[]string{
			r.Method,
			r.URL.Path,
			getQuery(r.URL.RawQuery),
			r.Proto,
		},
		" ")
}

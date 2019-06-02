package extractor

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetRequestByString(msg string, req *http.Request) *http.Request {
	editReq := strings.Split(msg, "\n")
	req.Header = setHTTPHeader(editReq[1:])

	startLine := strings.Split(editReq[0], " ")
	pathAndQuery := strings.Split(startLine[1], "?")
	req.URL.Path, req.Method, req.Proto =
		pathAndQuery[0], startLine[0], startLine[2]
	if len(pathAndQuery) > 1 {
		req.URL.ForceQuery = true
		req.URL.RawQuery = pathAndQuery[1]
	} else {
		req.URL.ForceQuery = false
	}

	bodyStr := editReq[len(editReq)-1]
	req.ContentLength, req.Body =
		int64(len(bodyStr)), GetIOReadCloser(bodyStr)

	host := req.Host
	if s := req.Header.Get("Host"); s != "" {
		host = s
	}
	req.URL.Host = host
	return req
}

func GetIOReadCloser(b string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(b))
}

func setHTTPHeader(h []string) http.Header {
	head := http.Header{}
	for _, v := range h {
		if v == "" {
			break
		}
		headL := strings.Split(v, ": ")
		if len(headL) <= 1 {
			headL = strings.Split(v, ":")
		}
		head.Add(headL[0], strings.Join(headL[1:], ":"))
	}
	return head
}

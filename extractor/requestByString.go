package extractor

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//GetRequestByString はメッセージ(string)からhttp.requestを生成するための関数です。
func GetRequestByString(msg string, req *http.Request) *http.Request {
	var (
		ers   = strings.Split(msg, "\n")     //editedRequestString
		sline = strings.Split(ers[0], " ")   //startLine
		pAndQ = strings.Split(sline[1], "?") //pathAndQuery
		bstr  = ers[len(ers)-1]              //String by body
		host  = req.Host
	)
	req.Header = setHTTPHeader(ers[1:])
	if s := req.Header.Get("Host"); s != "" {
		host = s
	}
	if len(pAndQ) > 1 {
		req.URL.ForceQuery, req.URL.RawQuery =
			true,
			pAndQ[1]
	} else {
		req.URL.ForceQuery = false
	}

	req.URL.Host, req.URL.Path, req.Method,
		req.Proto, req.ContentLength, req.Body =
		host,
		pAndQ[0], //path
		sline[0], //Method
		sline[2], //Protocol
		int64(len(bstr)), //ContentLength
		GetIOReadCloser(bstr) //Body

	return req
}

//GetIOReadCloser 文字列をio.ReadCloserにする抽象化関数
func GetIOReadCloser(b string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(b))
}

//setHTTPHeader は文字列のhttpheaderをhttp.headerと同じ方に変更する関数です
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

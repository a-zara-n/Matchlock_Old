package extractor

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

var (
	testRequest1 = []http.Request{
		{Method: "POST", URL: u1, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: head, Body: ioutil.NopCloser(strings.NewReader(bodystring)), ContentLength: int64(len(bodystring)), Host: "localhost"},
		{Method: "GET", URL: u2, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: head, Body: ioutil.NopCloser(strings.NewReader("")), Host: "localhost"},
	}
	testRequest2 = []http.Request{
		{Method: "POST", URL: u1, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: head, Body: ioutil.NopCloser(strings.NewReader(bodystring)), ContentLength: int64(len(bodystring)), Host: "localhost"},
		{Method: "GET", URL: u2, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: head, Body: ioutil.NopCloser(strings.NewReader("")), Host: "localhost"},
	}
)

func Test_GetRequestByString(t *testing.T) {
	for index := 0; index < 2; index++ {
		req := GetRequestByString(success[index], &testRequest1[index])
		bufbody1 := new(bytes.Buffer)
		bufbody2 := new(bytes.Buffer)
		bufbody1.ReadFrom(req.Body)
		bufbody2.ReadFrom(testRequest2[index].Body)
		if bufbody1.String() != bufbody2.String() {
			t.Error("It is not the expected request string. loop is", index+1)
		}
	}
}

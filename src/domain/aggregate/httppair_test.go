package aggregate

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

var (
	httppair = &HTTPPair{}
	u1, _    = url.Parse("http://localhost/testing/")
	u2, _    = url.Parse("http://localhost/testing/?input=usa")
	head     = http.Header{
		"Host":            {"loacalhost"},
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Foo":             {"Bar", "two"},
	}
	headSlice  = []string{"Accept-Encoding: gzip, deflate", "Accept-Language: en-us", "Foo: Bar,two", "Host: loacalhost"}
	bodystring = "submit=admin&message=Hi+my+name+is+admin."
	success    = []string{
		"POST /testing/  HTTP/1.0\nAccept-Encoding: gzip, deflate\nAccept-Language: en-us\nFoo: Bar,two\nHost: loacalhost\n\nsubmit=admin&message=Hi+my+name+is+admin.",
		"GET /testing/ ?input=usa HTTP/1.0\nAccept-Encoding: gzip, deflate\nAccept-Language: en-us\nFoo: Bar,two\nHost: loacalhost\n\n",
	}
	testRequest1 = []http.Request{
		{Method: "POST", URL: u1, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: head, Body: ioutil.NopCloser(strings.NewReader(bodystring)), ContentLength: int64(len(bodystring)), Host: "localhost"},
		{Method: "GET", URL: u2, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: head, Body: ioutil.NopCloser(strings.NewReader("")), Host: "localhost"},
	}
	testRequest2 = []http.Request{
		{Method: "POST", URL: u1, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: head, Body: ioutil.NopCloser(strings.NewReader(bodystring + "A")), ContentLength: int64(len(bodystring + "A")), Host: "localhost"},
		{Method: "GET", URL: u2, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: head, Body: ioutil.NopCloser(strings.NewReader("")), Host: "localhost"},
	}
)

func TestIsEdit(t *testing.T) {
	httppair.Request = NewHTTPRequestByRequest(&testRequest1[0])
	httppair.EditRequest = NewHTTPRequestByRequest(&testRequest1[0])
	if httppair.IsEdited() {
		t.Error("正しい値が帰ってきていません")
	}
	httppair.Request = NewHTTPRequestByRequest(&testRequest1[0])
	httppair.EditRequest = NewHTTPRequestByRequest(&testRequest2[0])
	if !httppair.IsEdited() {
		t.Error("正しい値が帰ってきていません")
	}
}

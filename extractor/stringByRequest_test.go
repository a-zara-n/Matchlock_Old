package extractor

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

var (
	u1, _ = url.Parse("http://localhost/testing/")
	u2, _ = url.Parse("http://localhost/testing/?input=usa")
	head  = http.Header{
		"Host":            {"loacalhost"},
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Foo":             {"Bar", "two"},
	}
	headSlice   = []string{"Accept-Encoding: gzip, deflate", "Accept-Language: en-us", "Foo: Bar,two", "Host: loacalhost"}
	bodystring  = "submit=admin&message=Hi+my+name+is+admin."
	testRequest = []http.Request{
		{Method: "POST", URL: u1, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: head, Body: ioutil.NopCloser(strings.NewReader(bodystring)), ContentLength: int64(len(bodystring)), Host: "localhost"},
		{Method: "GET", URL: u2, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: head, Body: ioutil.NopCloser(strings.NewReader("")), Host: "localhost"},
	}
	success = []string{
		"POST /testing/  HTTP/1.0\nAccept-Encoding: gzip, deflate\nAccept-Language: en-us\nFoo: Bar,two\nHost: loacalhost\n\nsubmit=admin&message=Hi+my+name+is+admin.",
		"GET /testing/ ?input=usa HTTP/1.0\nAccept-Encoding: gzip, deflate\nAccept-Language: en-us\nFoo: Bar,two\nHost: loacalhost\n\n",
	}
)

func Test_getQuery(t *testing.T) {
	if getQuery(u2.RawQuery) != "?input=usa" {
		t.Error("It is not a required query")
	}
	if getQuery(u1.RawQuery) != "" {
		t.Error("It is not a required query")
	}
}
func Test_GetHeader(t *testing.T) {
	if !reflect.DeepEqual(GetHeader(head), headSlice) {
		t.Error("Expected Headerslice could not be obtained")
	}
}
func Test_GetStringBody(t *testing.T) {
	if GetStringBody(ioutil.NopCloser(strings.NewReader(bodystring))) != bodystring {
		t.Error("It is not an expected body")
	}
	if GetStringBody(ioutil.NopCloser(strings.NewReader(""))) != "" {
		t.Error("It is not an expected body")
	}
}

func Test_GetStringByRequest(t *testing.T) {
	for index := 0; index < 2; index++ {
		if success[index] != GetStringByRequest(&testRequest[index]) {
			t.Error("It is not the expected request string. loop is", index+1)
		}
	}
}

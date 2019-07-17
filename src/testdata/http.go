package testdata

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//HTTPTestData はHTTP系のテストデータをまとめたものです
type HTTPTestData struct {
	Request      map[string]*Request
	RequestCase  []string
	Response     map[string]*Response
	ResponseCase []string
}

func (h *HTTPTestData) FetchTestRequest(i int) *Request {
	return h.Request[h.RequestCase[i]]
}
func (h *HTTPTestData) GetRequestCount() int {
	return len(h.RequestCase)
}
func (h *HTTPTestData) FetchTestResponse(i int) *Response {
	return h.Response[h.RequestCase[i]]
}
func (h *HTTPTestData) GetResponseCount() int {
	return len(h.ResponseCase)
}
func NewHTTPTestData() HTTPTestData {
	return HTTPTestData{
		Request:      TestRequests,
		RequestCase:  TestCaseSlice,
		Response:     TestResponses,
		ResponseCase: TestCaseSlice,
	}
}

//Request は
type Request struct {
	HTTP   *http.Request
	String string
	header
	Query query
}
type query struct {
	Raw    string
	Keys   []string
	Result map[string]interface{}
	Fetch  string
}
type header struct {
	Header      http.Header
	HeaderSlice []string
}

//Response は
type Response struct {
	HTTP   *http.Response
	String string
	header
	Body body
}

type body struct {
	Raw string
	IO  io.ReadCloser
}

var (
	host = "localhost"
	//TestURLPackageURL はテスト用のURLデータ
	TestURLPackageURL, _ = url.Parse("http://localhost/testing/")
	//TestURLPackageURLPlusQuery はテスト用のクエリ付きのデータです
	TestURLPackageURLPlusQuery, _ = url.Parse("http://localhost/testing/?input=usa")
	//TestHeader はテスト用のhttpHeader
	TestHeader = http.Header{
		"Host":            {"loacalhost"},
		"Accept-Encoding": {"gzip", "deflate"},
		"Accept-Language": {"en-us"},
		"Foo":             {"Bar", "two"},
	}
	TestHeaderKeys = []string{"Accept-Encoding", "Accept-Language", "Foo", "Host"}
	TestFailHeader = http.Header{}
	//TestHeaderSlice はHeaderの情報をstring sliceでまとめたもの
	TestHeaderSlice  = []string{"Accept-Encoding: gzip,deflate", "Accept-Language: en-us", "Foo: Bar,two", "Host: loacalhost"}
	Testheaderstruct = header{TestHeader, TestHeaderSlice}

	TestCaseSlice = []string{"FORM_success", "FORM_ADD_success", "JSON_success", "JSON_fail", "GET", "GET_Query"}
	//TestQuery はFORM形式とJSON形式のクエリ
	TestQuery = map[string]query{
		TestCaseSlice[0]: {
			Raw:    `name=hoge&age=20&like=mikan`,
			Keys:   []string{"name", "age", "like"},
			Result: map[string]interface{}{"name": "hoge", "age": "20", "like": "mikan"},
			Fetch:  `age=20&like=mikan&name=hoge`,
		},
		TestCaseSlice[1]: {
			Raw:    `name=hoge&age=20&like=mikanA`,
			Keys:   []string{"name", "age", "like"},
			Result: map[string]interface{}{"name": "hoge", "age": "20", "like": "mikanA"},
			Fetch:  `age=20&like=mikanA&name=hoge`,
		},
		TestCaseSlice[2]: {
			Raw:    `{"age":20,"lang":["jp","en","fr"],"like":"mikan","name":"hoge"}`,
			Keys:   []string{"name", "age", "like", "lang"},
			Result: map[string]interface{}{"name": "hoge", "age": float64(20), "like": "mikan", "lang": []interface{}{"jp", "en", "fr"}},
			Fetch:  `{"age":20,"lang":["jp","en","fr"],"like":"mikan","name":"hoge"}`,
		},
		TestCaseSlice[3]: {
			Raw:    `{"name":hoge}`,
			Keys:   []string{},
			Result: map[string]interface{}{},
			Fetch:  `{}`,
		},
		TestCaseSlice[4]: {
			Raw:    ``,
			Keys:   []string{},
			Result: map[string]interface{}{},
			Fetch:  ``,
		},
		TestCaseSlice[5]: {
			Raw:    `input=usa`,
			Keys:   []string{},
			Result: map[string]interface{}{},
			Fetch:  `input=usa`,
		},
	}
	TestRequests = map[string]*Request{
		TestCaseSlice[0]: {
			HTTP: &http.Request{Method: "POST", URL: TestURLPackageURL, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
				Body: ioutil.NopCloser(strings.NewReader(TestQuery[TestCaseSlice[0]].Raw)), ContentLength: int64(len(TestQuery[TestCaseSlice[0]].Fetch)), Host: host},
			String: "POST /testing/ HTTP/1.0\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestQuery[TestCaseSlice[0]].Raw,
			header: Testheaderstruct,
			Query:  TestQuery[TestCaseSlice[0]],
		},
		TestCaseSlice[1]: {
			HTTP: &http.Request{Method: "POST", URL: TestURLPackageURL, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
				Body: ioutil.NopCloser(strings.NewReader(TestQuery[TestCaseSlice[1]].Raw)), ContentLength: int64(len(TestQuery[TestCaseSlice[1]].Fetch)), Host: host},
			String: "POST /testing/ HTTP/1.0\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestQuery[TestCaseSlice[1]].Raw,
			header: Testheaderstruct,
			Query:  TestQuery[TestCaseSlice[1]],
		},
		TestCaseSlice[2]: {
			HTTP: &http.Request{Method: "POST", URL: TestURLPackageURL, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
				Body: ioutil.NopCloser(strings.NewReader(TestQuery[TestCaseSlice[2]].Raw)), ContentLength: int64(len(TestQuery[TestCaseSlice[2]].Fetch)), Host: host},
			String: "POST /testing/ HTTP/1.0\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestQuery[TestCaseSlice[2]].Raw,
			header: Testheaderstruct,
			Query:  TestQuery[TestCaseSlice[2]],
		},
		TestCaseSlice[3]: {
			HTTP: &http.Request{Method: "POST", URL: TestURLPackageURL, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
				Body: ioutil.NopCloser(strings.NewReader(TestQuery[TestCaseSlice[3]].Raw)), ContentLength: int64(len(TestQuery[TestCaseSlice[3]].Fetch)), Host: host},
			String: "POST /testing/ HTTP/1.0\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestQuery[TestCaseSlice[3]].Raw,
			header: Testheaderstruct,
			Query:  TestQuery[TestCaseSlice[3]],
		},
		TestCaseSlice[4]: {
			HTTP: &http.Request{Method: "GET", URL: TestURLPackageURL, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
				Body: ioutil.NopCloser(strings.NewReader(TestQuery[TestCaseSlice[4]].Raw)), ContentLength: int64(len(TestQuery[TestCaseSlice[4]].Raw)), Host: host},
			String: "GET /testing/ HTTP/1.0\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestQuery[TestCaseSlice[4]].Raw,
			header: Testheaderstruct,
			Query:  TestQuery[TestCaseSlice[4]],
		},
		TestCaseSlice[5]: {
			HTTP: &http.Request{Method: "GET", URL: TestURLPackageURLPlusQuery, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
				Body: ioutil.NopCloser(strings.NewReader(TestQuery[TestCaseSlice[5]].Raw)), ContentLength: int64(len(TestQuery[TestCaseSlice[5]].Raw)), Host: host},
			String: "GET /testing/?input=usa HTTP/1.0\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestQuery[TestCaseSlice[5]].Raw,
			header: Testheaderstruct,
			Query:  TestQuery[TestCaseSlice[5]],
		},
	}

	TestResponse = http.Response{
		Status:           "200 OK",
		StatusCode:       200,
		Proto:            "HTTP/1.0",
		ProtoMajor:       1,
		ProtoMinor:       0,
		Header:           TestHeader,
		TransferEncoding: []string{"chunked"},
	}
	TestHTML = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"
 "http://www.w3.org/TR/html4/strict.dtd">

<html>

 <head>
  <title>タイトルを指定する</title>
 </head>

 <body>
  ここに内容を書く
 </body>

</html>`
	TestBodys = map[string]body{
		TestCaseSlice[0]: {
			Raw: TestHTML,
			IO:  ioutil.NopCloser(strings.NewReader(TestHTML)),
		},
		TestCaseSlice[1]: {
			Raw: TestHTML,
			IO:  ioutil.NopCloser(strings.NewReader(TestHTML)),
		},
		TestCaseSlice[2]: {
			Raw: TestHTML,
			IO:  ioutil.NopCloser(strings.NewReader(TestHTML)),
		},
		TestCaseSlice[3]: {
			Raw: TestHTML,
			IO:  ioutil.NopCloser(strings.NewReader(TestHTML)),
		},
		TestCaseSlice[4]: {
			Raw: TestHTML,
			IO:  ioutil.NopCloser(strings.NewReader(TestHTML)),
		},
		TestCaseSlice[5]: {
			Raw: TestHTML,
			IO:  ioutil.NopCloser(strings.NewReader(TestHTML)),
		},
	}
	testHTTPSetFunc = func(casestr string) *http.Response {
		TestResponse.Body = TestBodys[casestr].IO
		TestResponse.ContentLength = int64(len(TestBodys[casestr].Raw))
		return &TestResponse
	}
	TestResponses = map[string]*Response{
		TestCaseSlice[0]: {
			HTTP:   testHTTPSetFunc(TestCaseSlice[0]),
			String: TestResponse.Proto + " " + TestResponse.Status + "\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestBodys[TestCaseSlice[0]].Raw,
			header: Testheaderstruct,
			Body:   TestBodys[TestCaseSlice[0]],
		},
		TestCaseSlice[1]: {
			HTTP:   testHTTPSetFunc(TestCaseSlice[1]),
			String: TestResponse.Proto + " " + TestResponse.Status + "\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestBodys[TestCaseSlice[1]].Raw,
			header: Testheaderstruct,
			Body:   TestBodys[TestCaseSlice[1]],
		},
		TestCaseSlice[2]: {
			HTTP:   testHTTPSetFunc(TestCaseSlice[2]),
			String: TestResponse.Proto + " " + TestResponse.Status + "\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestBodys[TestCaseSlice[2]].Raw,
			header: Testheaderstruct,
			Body:   TestBodys[TestCaseSlice[2]],
		},
		TestCaseSlice[3]: {
			HTTP:   testHTTPSetFunc(TestCaseSlice[3]),
			String: TestResponse.Proto + " " + TestResponse.Status + "\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestBodys[TestCaseSlice[3]].Raw,
			header: Testheaderstruct,
			Body:   TestBodys[TestCaseSlice[3]],
		},
		TestCaseSlice[4]: {
			HTTP:   testHTTPSetFunc(TestCaseSlice[4]),
			String: TestResponse.Proto + " " + TestResponse.Status + "\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestBodys[TestCaseSlice[4]].Raw,
			header: Testheaderstruct,
			Body:   TestBodys[TestCaseSlice[4]],
		},
		TestCaseSlice[5]: {
			HTTP:   testHTTPSetFunc(TestCaseSlice[5]),
			String: TestResponse.Proto + " " + TestResponse.Status + "\n" + strings.Join(TestHeaderSlice, "\n") + "\n\n" + TestBodys[TestCaseSlice[5]].Raw,
			header: Testheaderstruct,
			Body:   TestBodys[TestCaseSlice[5]],
		},
	}
)

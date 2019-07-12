package testdata

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	//TestURLPackageURL はテスト用のURLデータ
	TestURLPackageURL, _ = url.Parse("http://localhost/testing/")
	//TestURLPackageURLPlusQuery はテスト用のクエリ付きのデータです
	TestURLPackageURLPlusQuery, _ = url.Parse("http://localhost/testing/?input=usa")
	//TestHeader はテスト用のhttpHeader
	TestHeader = http.Header{
		"Host":            {"loacalhost"},
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Foo":             {"Bar", "two"},
	}
	//TestHeaderSlice はHeaderの情報をstring sliceでまとめたもの
	TestHeaderSlice = []string{"Accept-Encoding: gzip, deflate", "Accept-Language: en-us", "Foo: Bar,two", "Host: loacalhost"}
	//TestQuery はFORM形式とJSON形式のクエリ
	TestQuery = map[string]map[string]interface{}{
		"FORM": {
			"Raw":    `name=hoge&age=20&like=mikan`,
			"Keys":   []string{"name", "age", "like"},
			"Result": map[string]interface{}{"name": "hoge", "age": "20", "like": "mikan"},
			"Fetch":  `age=20&like=mikan&name=hoge`,
		},
		"FORM_ADD": {
			"Raw":    `name=hoge&age=20&like=mikanA`,
			"Keys":   []string{"name", "age", "like"},
			"Result": map[string]interface{}{"name": "hoge", "age": "20", "like": "mikanA"},
			"Fetch":  `age=20&like=mikanA&name=hoge`,
		},
		"JSON_Success": {
			"Raw":    `{"age":20,"lang":["jp","en","fr"],"like":"mikan","name":"hoge"}`,
			"Keys":   []string{"name", "age", "like", "lang"},
			"Result": map[string]interface{}{"name": "hoge", "age": float64(20), "like": "mikan", "lang": []interface{}{"jp", "en", "fr"}},
			"Fetch":  `{"age":20,"lang":["jp","en","fr"],"like":"mikan","name":"hoge"}`,
		},
		"JSON_Fail": {
			"Raw":    `{"name":hoge}`,
			"Keys":   []string{},
			"Result": map[string]interface{}{},
			"Fetch":  `{}`,
		},
	}
	//TestDataNames は
	TestDataNames = []string{"StandardPOST", "AddStringPOST", "JSONPOST", "JSONPOST_Fail", "StandardGET"}
	//TestDataKeys は
	TestDataKeys = map[string]string{
		"StandardPOST":  "FORM",
		"AddStringPOST": "FORM_ADD",
		"JSONPOST":      "JSON_Success",
		"JSONPOST_Fail": "JSON_Fail",
		"StandardGET":   "",
	}

	//TestSuccessReturn はパース成功後の値
	TestSuccessReturn = map[string]string{
		"StandardPOST":  "POST /testing/  HTTP/1.0\nAccept-Encoding: gzip, deflate\nAccept-Language: en-us\nFoo: Bar,two\nHost: loacalhost\n\n" + TestQuery["FORM"]["Raw"].(string),
		"AddStringPOST": "POST /testing/  HTTP/1.0\nAccept-Encoding: gzip, deflate\nAccept-Language: en-us\nFoo: Bar,two\nHost: loacalhost\n\n" + TestQuery["FORM_Add"]["Raw"].(string),
		"JSONPOST":      "POST /testing/  HTTP/1.0\nAccept-Encoding: gzip, deflate\nAccept-Language: en-us\nFoo: Bar,two\nHost: loacalhost\n\n" + TestQuery["JSON_Success"]["Raw"].(string),
		"JSONPOST_Fail": "POST /testing/  HTTP/1.0\nAccept-Encoding: gzip, deflate\nAccept-Language: en-us\nFoo: Bar,two\nHost: loacalhost\n\n" + TestQuery["JSON_Fail"]["Raw"].(string),
		"StandardGET":   "GET /testing/ ?input=usa HTTP/1.0\nAccept-Encoding: gzip, deflate\nAccept-Language: en-us\nFoo: Bar,two\nHost: loacalhost\n\n",
	}
	//TestRequest はhttprequestの配列
	TestRequest = map[string]http.Request{
		"StandardPOST": {Method: "POST", URL: TestURLPackageURL, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
			Body: ioutil.NopCloser(strings.NewReader(TestQuery["FORM"]["Raw"].(string))), ContentLength: int64(len(TestQuery["FORM"]["Raw"].(string))), Host: "localhost"},
		"AddStringPOST": {Method: "POST", URL: TestURLPackageURL, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
			Body: ioutil.NopCloser(strings.NewReader(TestQuery["FORM_Add"]["Raw"].(string))), ContentLength: int64(len(TestQuery["FORM_Add"]["Raw"].(string))), Host: "localhost"},
		"JSONPOST": {Method: "POST", URL: TestURLPackageURL, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
			Body: ioutil.NopCloser(strings.NewReader(TestQuery["JSON_Success"]["Raw"].(string))), ContentLength: int64(len(TestQuery["JSON_Success"]["Raw"].(string))), Host: "localhost"},
		"JSONPOST_Fail": {Method: "POST", URL: TestURLPackageURL, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
			Body: ioutil.NopCloser(strings.NewReader(TestQuery["JSON_Success"]["Raw"].(string))), ContentLength: int64(len(TestQuery["JSON_Fail"]["Raw"].(string))), Host: "localhost"},
		"StandardGET": {Method: "GET", URL: TestURLPackageURLPlusQuery, Proto: "HTTP/1.0", ProtoMajor: 1, ProtoMinor: 0, Header: TestHeader,
			Body: ioutil.NopCloser(strings.NewReader("")), Host: "localhost"},
	}
)

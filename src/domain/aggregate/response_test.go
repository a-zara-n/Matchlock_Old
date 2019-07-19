package aggregate

import (
	"reflect"
	"testing"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

var testresponse = &Response{
	Info:    &entity.ResponseInfo{},
	Header:  &entity.HTTPHeader{},
	Body:    &entity.Body{},
	Request: &Request{},
}

func TestSetResponse(t *testing.T) {
	for i := 0; i < test.GetRequestCount(); i++ {
		testingdata := test.FetchTestResponse(i)
		testresponse.SetHTTPResponseByResponse(testingdata.HTTP)
		if testresponse.Info.Proto != testingdata.HTTP.Proto || testresponse.Info.ProtoMajor != testingdata.HTTP.ProtoMajor || testresponse.Info.ProtoMinor != testingdata.HTTP.ProtoMinor {
			t.Errorf("プロトコル関連のセッティングが不適切です \n Proto : %v != %v \n ProtoMajor : %v != %v \n ProtoMiner : %v != %v",
				testresponse.Info.Proto, testingdata.HTTP.Proto, testresponse.Info.ProtoMajor, testingdata.HTTP.ProtoMajor,
				testresponse.Info.ProtoMinor, testingdata.HTTP.ProtoMinor)
		}
		if testresponse.Info.Status != testingdata.HTTP.Status || testresponse.Info.StatusCode != testingdata.HTTP.StatusCode {
			t.Errorf("ステータス処理周辺のセッティングが不適切です\n Status : %v != %v \n StatusCode : %v != %v",
				testresponse.Info.Status, testingdata.HTTP.Status, testresponse.Info.StatusCode, testingdata.HTTP.StatusCode)
		}
		if !reflect.DeepEqual(testresponse.Header.Header, testingdata.Header) {
			t.Errorf("http.Headerのセッティングが不適切です \n Header : %v != %v", testresponse.Header.Header, testingdata.Header)
		}
		if testresponse.Body.Body != testingdata.Body.Raw {
			t.Errorf("bodyの設定が不適切です \n Body : %v != %v", testresponse.Body.Body, testingdata.Body.Raw)
		}
		if !reflect.DeepEqual(testresponse.Request.GetHTTPRequestByRequest(), testingdata.HTTP.Request) {
			t.Error("http.Requestの設定が不適切です")
		}
		resp := testresponse.GetHTTPRequestByResponse()
		if !reflect.DeepEqual(resp, testingdata.HTTP) {
			t.Error("http.Requestの出力が不適切です\nこのテストの前のテストがRedの場合は、そのテストの改善をしてください")
			t.Error(resp)
			t.Error(testingdata.HTTP)
		}
	}
}

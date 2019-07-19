package aggregate

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/testdata"
)

func TestSetHTTPRequestByRequest(t *testing.T) {
	testrequestuest := &Request{
		Info:   &entity.RequestInfo{},
		Header: &entity.HTTPHeader{},
		Data:   &entity.Data{},
	}
	for i := 0; i < test.GetRequestCount(); i++ {
		testingdata := test.FetchTestRequest(i)
		testrequestuest.SetHTTPRequestByRequest(testingdata.HTTP)
		statusline := strings.Split(testingdata.String, "\n")[0]
		if statusline != testrequestuest.Info.GetStatusLine() {
			t.Errorf("取得されたステータスラインが異なります \n TestData : %v \n Return : %v", statusline, testrequestuest.Info.GetStatusLine())
		}
		if testrequestuest.Header.GetStringHeader() != strings.Join(testdata.TestHeaderSlice, "\n") {
			t.Error("値が異なります")
		}
		if testrequestuest.Data.FetchData() != testingdata.Query.Fetch {
			t.Errorf("出力が異なります \n RetData :%v \n TestData : %v", testrequestuest.Data.FetchData(), testingdata.Query.Fetch)
		}
	}
}

func TestSetHTTPRequestByString(t *testing.T) {
	testrequestuest := &Request{
		Info:   &entity.RequestInfo{},
		Header: &entity.HTTPHeader{},
		Data:   &entity.Data{},
	}
	for i := 0; i < test.GetRequestCount(); i++ {
		testingdata := test.FetchTestRequest(i)
		testrequestuest.SetHTTPRequestByString(testingdata.String)
		statusline := strings.Split(testingdata.String, "\n")[0]
		if statusline != testrequestuest.Info.GetStatusLine() {
			t.Errorf("取得されたステータスラインが異なります \n TestData : %v \n Return : %v", statusline, testrequestuest.Info.GetStatusLine())
		}
		if testrequestuest.Header.GetStringHeader() != strings.Join(testdata.TestHeaderSlice, "\n") {
			t.Error("値が異なります")
		}
		if testrequestuest.Data.FetchData() != testingdata.Query.Fetch {
			t.Errorf("出力が異なります \n RetData :%v \n TestData : %v", testrequestuest.Data.FetchData(), testingdata.Query.Fetch)
		}
	}
}

func TestGetHTTPRequestByRequest(t *testing.T) {
	testrequestuest := &Request{
		Info:   &entity.RequestInfo{},
		Header: &entity.HTTPHeader{},
		Data:   &entity.Data{},
	}
	for i := 0; i < test.GetRequestCount(); i++ {
		testingdata := test.FetchTestRequest(i)
		testrequestuest.SetHTTPRequestByRequest(testingdata.HTTP)
		diffrequest := testrequestuest.GetHTTPRequestByRequest()
		if diffrequest.Method != testingdata.HTTP.Method || !reflect.DeepEqual(diffrequest.URL, testingdata.HTTP.URL) ||
			diffrequest.Proto != testingdata.HTTP.Proto || diffrequest.ProtoMajor != testingdata.HTTP.ProtoMajor ||
			diffrequest.ProtoMinor != testingdata.HTTP.ProtoMinor {
			t.Error("RequestInfoの排出作業にエラーが発生しています")
		}
		if !reflect.DeepEqual(diffrequest.Method, testingdata.HTTP.Method) {
			t.Error("Headerにエラーが発生しました")
		}
		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(diffrequest.Body)
		if bufbody.String() != testingdata.Query.Fetch || diffrequest.ContentLength != testingdata.HTTP.ContentLength {
			t.Log(diffrequest.ContentLength, testingdata.HTTP.ContentLength)
			t.Log(bufbody.String(), testingdata.Query.Fetch)
			t.Error("Dataにエラーが発生しました")
		}
	}
}

func TestGetHTTPRequestByString(t *testing.T) {
	testrequestuest := &Request{
		Info:   &entity.RequestInfo{},
		Header: &entity.HTTPHeader{},
		Data:   &entity.Data{},
	}
	for i := 0; i < test.GetRequestCount(); i++ {
		testingdata := test.FetchTestRequest(i)
		testrequestuest.SetHTTPRequestByRequest(testingdata.HTTP)
		retstring := testrequestuest.GetHTTPRequestByString()
		if retstring != testingdata.String {
			t.Error("正確な返答が返ってきていません")
			t.Error("\n" + retstring)
			t.Error("\n" + testingdata.String)
		}
	}
}

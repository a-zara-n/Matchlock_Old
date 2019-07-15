package entity

import (
	"strings"
	"testing"
)

var testRequestInfo = &RequestInfo{}

func TestRequestInfo(t *testing.T) {
	for i := 0; i < test.GetRequestCount(); i++ {
		testRequestInfo.SetRequestINFO(test.FetchTestRequest(i).HTTP)
		statusline := strings.Split(test.FetchTestRequest(i).String, "\n")[0]
		if statusline != testRequestInfo.GetStatusLine() {
			t.Errorf("取得されたステータスラインが異なります \n TestData : %v \n Return : %v", statusline, testRequestInfo.GetStatusLine())
		}
	}
	testRequestInfo.SetRequestINFO(test.FetchTestRequest(0).HTTP)
	testRequestInfo.SetStatusLine("GET /avredsc?nam=893goro HTTP/0.9")
	if "nam=893goro" != testRequestInfo.Query.FetchData() && "GET" != testRequestInfo.Method && "/avredsc" != testRequestInfo.Path && "HTTP/0.9" != testRequestInfo.Proto {
		t.Error("適正な代入が行われておりません")
	}
}

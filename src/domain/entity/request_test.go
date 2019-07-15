package entity

import (
	"strings"
	"testing"
)

var testRequestInfo = &RequestInfo{}

func TestGetStartLine(t *testing.T) {
	for i := 0; i < test.GetRequestCount(); i++ {
		testRequestInfo.SetRequestINFO(test.FetchTestRequest(i).HTTP)
		statusline := strings.Split(test.FetchTestRequest(i).String, "\n")[0]
		if statusline != testRequestInfo.GetStatusLine() {
			t.Errorf("取得されたステータスラインが異なります \n TestData : %v \n Return : %v", statusline, testRequestInfo.GetStatusLine())
		}
	}
}

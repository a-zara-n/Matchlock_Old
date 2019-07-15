package aggregate

import (
	"testing"

	"github.com/a-zara-n/Matchlock/src/testdata"
)

var test = testdata.NewHTTPTestData()
var httppair = &HTTPPair{}

func TestIsEdit(t *testing.T) {
	httppair.Request = NewHTTPRequestByRequest(test.FetchTestRequest(0).HTTP)
	httppair.EditRequest = NewHTTPRequestByRequest(test.FetchTestRequest(0).HTTP)
	if httppair.IsEdited() {
		t.Error("正しい値が帰ってきていません")
	}
	httppair.Request = NewHTTPRequestByRequest(test.FetchTestRequest(0).HTTP)
	httppair.EditRequest = NewHTTPRequestByRequest(test.FetchTestRequest(1).HTTP)
	if !httppair.IsEdited() {
		t.Error("正しい値が帰ってきていません")
	}
}

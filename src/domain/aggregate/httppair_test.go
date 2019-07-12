package aggregate

import (
	"reflect"
	"testing"

	"github.com/a-zara-n/Matchlock/src/testdata"
)

var httppair = &HTTPPair{}

func TestIsEdit(t *testing.T) {
	httppair.Request = NewHTTPRequestByRequest(&testdata.TestRequest[0])
	httppair.EditRequest = NewHTTPRequestByRequest(&testdata.TestRequest[0])
	if httppair.IsEdited() {
		t.Error("正しい値が帰ってきていません")
	}
	httppair.Request = NewHTTPRequestByRequest(&testdata.TestRequest[0])
	httppair.EditRequest = NewHTTPRequestByRequest(&testdata.TestRequest[1])
	if !httppair.IsEdited() {
		t.Error("正しい値が帰ってきていません")
	}
}

func TestRequestStartLine(t *testing.T) {
	httppair.Request = NewHTTPRequestByRequest(&testdata.TestRequest[0])
	if !reflect.DeepEqual(httppair.Request.Info.GetStartLine(), []string{"POST", "/testing/", "HTTP/1.0"}) {
		t.Error("帰ってきた値が異なります")
	}
}

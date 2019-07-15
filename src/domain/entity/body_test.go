package entity

import (
	"bytes"
	"testing"
)

var testbody = &Body{}

func TestBody(t *testing.T) {
	for i := 0; i < test.GetResponseCount(); i++ {
		testbody.Set(test.FetchTestResponse(i).Body.IO)
		if testbody.Body != test.FetchTestResponse(i).Body.Raw {
			t.Error("適正に文字列化されておりません")
		}
		if testbody.GetLength() != int64(len(test.FetchTestResponse(i).Body.Raw)) {
			t.Error("正しい文字数が帰ってきません")
		}

		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(testbody.Get())
		if bufbody.String() != test.FetchTestResponse(i).Body.Raw {
			t.Error("適正な値が返されていません")
		}

	}
}

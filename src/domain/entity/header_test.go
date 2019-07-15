package entity

import (
	"reflect"
	"testing"

	"github.com/a-zara-n/Matchlock/src/testdata"
)

var testHeader = &HTTPHeader{}

func TestHeaderGetKeys(t *testing.T) {
	testHeader.SetHTTPHeader(testdata.TestHeader)
	if !reflect.DeepEqual(testHeader.Header, testdata.TestHeader) {
		t.Error("入力されたヘッダーが異なります")
	}
	if !reflect.DeepEqual(testHeader.GetKeys(), testdata.TestHeaderKeys) {
		t.Error("出力されたヘッダーのkeyが異なります")
	}
}

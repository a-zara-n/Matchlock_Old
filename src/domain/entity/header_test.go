package entity

import (
	"fmt"
	"reflect"
	"strings"
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
		fmt.Println(testHeader.GetKeys(), testdata.TestHeaderKeys)
		t.Error("出力されたヘッダーのkeyが異なります")
	}
	testHeader = &HTTPHeader{}
}

func TestSetStringHeader(t *testing.T) {
	testHeader.SetStringHeader(strings.Join(testdata.TestHeaderSlice, "\n"))
	if !reflect.DeepEqual(testHeader.Header, testdata.TestHeader) {
		t.Error("値が異なります")
	}
}

func TestGetStringHeader(t *testing.T) {
	testHeader.Header = testdata.TestHeader
	if testHeader.GetStringHeader() != strings.Join(testdata.TestHeaderSlice, "\n") {
		t.Error("値が異なります")
	}
}

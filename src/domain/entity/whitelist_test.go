package entity

import (
	"testing"

	"github.com/a-zara-n/Matchlock/src/testdata"
)

var testWhitelist = WhiteList{}

func TestWhiteList(t *testing.T) {
	testWhitelist.Add(`^\w*\.?(localhost)(\.+\w+)*$`)
	if !testWhitelist.Check(testdata.TestURLPackageURL.Host) {
		t.Error("チェックが失敗しています")
	}
	testdata.TestURLPackageURL.Host = "localhot"
	if testWhitelist.Check(testdata.TestURLPackageURL.Host) {
		t.Error("チェックが失敗しています")
	}
	testdata.TestURLPackageURL.Host = "localhost"
	testWhitelist.Del(0)
	if len(testWhitelist.List) != 0 {
		t.Error("削除が完了していません")
	}
	if testWhitelist.Del(20) {
		t.Error("チェックに失敗しています")
	}
	tr := `^\w*\.?(localhost`
	for index := 0; index < 5; index++ {
		testWhitelist.Add(tr + `)(\.+\w+)*$`)
		tr += "A"
	}
	if !testWhitelist.Check(testdata.TestURLPackageURL.Host + "AAAA") {
		t.Error("チェックが失敗しています")
	}
	testWhitelist.Del(4)
	if testWhitelist.Check(testdata.TestURLPackageURL.Host + "AAAA") {
		t.Error("チェックが失敗しています")
	}
	if !testWhitelist.Check(testdata.TestURLPackageURL.Host + "A") {
		t.Error("チェックが失敗しています")
	}
	testWhitelist.Del(1)
	if testWhitelist.Check(testdata.TestURLPackageURL.Host + "A") {
		t.Error("チェックが失敗しています")
	}
	if !testWhitelist.Check(testdata.TestURLPackageURL.Host) {
		t.Error("チェックが失敗しています")
	}
	testWhitelist.Del(0)
	if testWhitelist.Check(testdata.TestURLPackageURL.Host) {
		t.Error("チェックが失敗しています")
	}
}

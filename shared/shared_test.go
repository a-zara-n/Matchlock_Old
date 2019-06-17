package shared

import (
	"io/ioutil"
	"reflect"
	"sort"
	"strings"
	"sync"
	"testing"
)

func Test_Merge(t *testing.T) {
	test1 := map[string]string{"hoge": "123", "fuga": "456"}
	test2 := map[string]string{"piyo": "789"}
	ret := Merge(test1, test2)
	mapKeys := []string{}
	for key := range ret {
		mapKeys = append(mapKeys, key)
	}
	sort.Strings(mapKeys)
	if !reflect.DeepEqual(mapKeys, []string{"fuga", "hoge", "piyo"}) {
		t.Error("all keys are not merged")
	}
}

func Test_GetKeys(t *testing.T) {
	test := map[string][]string{"hoge": {"123"}, "fuga": {"456"}, "piyo": {"789"}}
	mapKeys := GetKeys(test)
	if !reflect.DeepEqual(mapKeys, []string{"fuga", "hoge", "piyo"}) {
		t.Error("all keys are not merged")
	}
}

func Test_QuoteEscape(t *testing.T) {
	if `\\\"` != QuoteEscape(`"`) {
		t.Error("Not escaped")
	}
	if `\\\'` != QuoteEscape(`'`) {
		t.Error("Not escaped")
	}
	if "`" != QuoteEscape("`") {
		t.Error("Not escaped")
	}
	if `i\\\'am` != QuoteEscape("i'am") {
		t.Error("Not escaped")
	}
}

func Test_RecursiveExec(t *testing.T) {
	var testFunc func(strs []string)
	wg := sync.WaitGroup{}
	out := []int{}
	testFunc = func(strs []string) {
		out = append(out, RecursiveExec(strs, testFunc))
		wg.Done()
	}
	wg.Add(1)
	testFunc([]string{"hoge"})
	wg.Wait()
	if len(out) != 1 {
		t.Error("Not Recursived")
	}
	out = []int{}
	wg.Add(3)
	testFunc([]string{"hoge", "hoge", "hoge"})
	wg.Wait()
	if len(out) != 3 {
		t.Error("Not Recursived")
	}
	out = []int{}
	wg.Add(12)
	testFunc([]string{"hoge", "hoge", "hoge", "hoge", "hoge", "hoge", "hoge", "hoge", "hoge", "hoge", "hoge", "hoge"})
	wg.Wait()
	if len(out) != 12 {
		t.Error("Not Recursived")
	}
}
func Test_CheckRegexp(t *testing.T) {
	if !CheckRegexp(`.*`, "WestEast1st") {
		t.Error("Not Check Regexp")
	}
	if CheckRegexp(`w`, "WestEast1st") {
		t.Error("Not Check Regexp")
	}
}
func Test_SeparationOfIOReadCloser(t *testing.T) {
	b := "http://example.com"
	str, iorc := SeparationOfIOReadCloser(ioutil.NopCloser(strings.NewReader(b)))
	if str != b {
		t.Error("Not Separation of io.ReadCloser")
	}
	i := ioutil.NopCloser(strings.NewReader(b))
	if reflect.TypeOf(iorc) != reflect.TypeOf(i) {
		t.Error("Not Separation of io.ReadCloser")
	}
}

package shared

import (
	"io"
	"regexp"
	"sort"
	"strings"

	"github.com/a-zara-n/Matchlock/extractor"
)

func Merge(m1, m2 map[string]string) map[string]string {
	ans := map[string]string{}
	for k, v := range m1 {
		ans[k] = v
	}
	for k, v := range m2 {
		ans[k] = v
	}
	return (ans)
}

func SeparationOfIOReadCloser(b io.ReadCloser) (string, io.ReadCloser) {
	bodyOfStr := extractor.GetStringBody(b)
	b = extractor.GetIOReadCloser(bodyOfStr)
	return bodyOfStr, b
}

func CheckRegexp(reg, str string) bool {
	r := regexp.MustCompile(reg)
	return r.MatchString(str)
}

func GetKeys(maps map[string][]string) []string {
	ret := []string{}
	for key := range maps {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret
}

func QuoteEscape(str string) string {
	str = strings.Replace(str, `"`, `\\\"`, -1)
	str = strings.Replace(str, "'", `\\\'`, -1)
	return str
}

func RecursiveExec(slice []string, fun func(slice []string)) int {
	if len(slice) > 1 {
		go fun(slice[1:])
		return len(slice)
	}
	return 1
}

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

/*
section\n
	slice[0] = form query
	slice[1] = json
*/
var (
	querySectString       = []string{"&", ","}
	keyandvalueSectString = []string{"=", "\":\""}
	joinFunc              = map[string]func(data []string, strs []string) string{
		"":     func(data []string, strs []string) string { return strings.Join(data, strs[0]) },
		"JSON": func(data []string, strs []string) string { return strings.Join(data, strs[1]) },
	}
	splitFunc = map[string]func(rawQuery string, strs []string) []string{
		"":     func(rawQuery string, strs []string) []string { return strings.Split(rawQuery, strs[0]) },
		"JSON": func(rawQuery string, strs []string) []string { return strings.Split(rawQuery, strs[1]) },
	}
	convParam = map[string]func(rawQuery string) string{
		"":     func(rawQuery string) string { return rawQuery },
		"JSON": func(rawQuery string) string { return "\"" + rawQuery + "\"" },
	}
	convQuery = map[string]func(rawQuery string) string{
		"":     func(rawQuery string) string { return rawQuery },
		"JSON": func(rawQuery string) string { return "{" + rawQuery + "}" },
	}
	deconv = map[string]func(decString string) string{
		"":     func(decString string) string { return decString },
		"JSON": func(decString string) string { return decString[1 : len(decString)-1] },
	}
)

func QueryConverter(typ string, data [][]string) string {
	var recursion func(data [][]string) []string
	recursion = func(data [][]string) []string {
		if len(data) > 1 {
			return append(
				[]string{convParam[typ](joinFunc[typ](data[0], keyandvalueSectString))},
				recursion(data[1:])...,
			)
		}
		return []string{convParam[typ](joinFunc[typ](data[0], keyandvalueSectString))}
	}
	return convQuery[typ](joinFunc[typ](recursion(data), querySectString))
}

func QueryDeconverter(typ string, rawQuery string) [][]string {
	var recursion func(data []string) [][]string
	params := splitFunc[typ](deconv[typ](rawQuery), querySectString)
	recursion = func(data []string) [][]string {
		if len(data) > 1 {
			return append(
				[][]string{splitFunc[typ](deconv[typ](data[0]), keyandvalueSectString)},
				recursion(data[1:])...,
			)
		}
		return [][]string{splitFunc[typ](deconv[typ](data[0]), keyandvalueSectString)}
	}
	return recursion(params)
}

package history

import (
	"sort"
	"strings"

	"github.com/WestEast1st/Matchlock/datastore"
)

var db = datastore.Database{Database: "./test.db"}

func getKeys(maps map[string][]string) []string {
	ret := []string{}
	for key := range maps {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret
}

func quoteEscape(str string) string {
	str = strings.Replace(str, `"`, `\\\"`, -1)
	str = strings.Replace(str, "'", `\\\'`, -1)
	return str
}

func recursiveExec(slice []string, fun func(slice []string)) {
	if len(slice) > 1 {
		go fun(slice[1:])
	}
}

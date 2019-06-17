package scanner

import (
	"strconv"
	"strings"
)

func getParamType(param string) string {
	if isInt(param) {
		return "INT"
	} else if isBool(param) {
		return "BOOL"
	}
	return "STRING"
}

func isInt(param string) bool {
	convI, _ := strconv.ParseInt(param, 10, 64)
	if convI == 0 {
		if strconv.FormatInt(convI, 10) != param {
			return false
		}
	}
	return true
}

func isBool(param string) bool {
	param = strings.ToLower(param)
	convB, _ := strconv.ParseBool(param)
	if !convB {
		if strconv.FormatBool(convB) != param {
			return false
		}
	}
	return true
}

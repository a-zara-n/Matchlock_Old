package datastore

import (
	"io/ioutil"
	"os"
)

var (
	historyCount int
	sqlpath      = "./_sql/"
	files        = []string{
		sqlpath + "requestDataTable.sql",
		sqlpath + "httpHistoryColumn.sql",
		sqlpath + "noEditRequest.sql",
		sqlpath + "editRequest.sql",
		sqlpath + "response.sql",
	}
	getSQL = func() []string {
		ret := []string{}
		for _, path := range files {
			f, _ := os.Open(path)
			defer f.Close()
			b, _ := ioutil.ReadAll(f)
			ret = append(ret, string(b))
		}
		return ret
	}()
	requestDataTable = getSQL[0]
	httpData         = map[string]string{
		"SELECT":    getSQL[1],
		"NoEditReq": getSQL[2],
		"EditReq":   getSQL[3],
		"Response":  getSQL[4],
	}
)

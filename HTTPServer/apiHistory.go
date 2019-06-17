package httpserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/WestEast1st/Matchlock/datastore"
	"github.com/labstack/echo"
)

var (
	historyCount int
	sqlpath      = "./sql/"
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

func getHistory(i int) []historyData {
	historys := []historyData{}
	db := datastore.DB.OpenDatabase()
	db.Table("requests").
		Select(" id,requests.identifier as identifier,method,host,path,url,param").
		Joins(requestDataTable).
		Where("id >= ?", i).
		Find(&historys)
	return historys
}

func GetHistryAll(c echo.Context) error {
	var w = c.Response()
	historyCount = 1
	historys := getHistory(historyCount)
	historyCount = len(historys)
	res, err := json.Marshal(APIresponse{Data: historys})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return nil
}

func GetRequest(c echo.Context) error {
	var w = c.Response()
	ret := httpdata{}
	db := datastore.DB.OpenDatabase()
	db.Table("histories").
		Select(httpData["SELECT"]).
		Joins(httpData["NoEditReq"]).
		Joins(httpData["EditReq"]).
		Joins(httpData["Response"]).
		Where("histories.identifier = ?", c.Param("identifier")).
		Find(&ret)
	res, err := json.Marshal(APIresponse{Data: ret})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return nil
}

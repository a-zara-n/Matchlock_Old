package scanner

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/WestEast1st/Matchlock/datastore"
	"github.com/WestEast1st/Matchlock/scanner/attacker"
)

type scanner struct {
	ScanTargets []http.Request
}

/*
tmp
func (s *scanner) () {

}
*/
var db = datastore.Database{Database: "./test.db"}

type getdata struct {
	Name  string
	Value string
	Count int
}

func (s *scanner) Scan() { //tmpname いずれ変える
	payloads := [][]string{}
	for _, name := range []string{"payload", "xss"} {
		f, err := os.Open("./payload/" + name + ".txt")
		if err != nil {
			fmt.Println("error")
		}
		defer f.Close()
		b, err := ioutil.ReadAll(f)
		payloads = append(payloads, strings.Split(string(b), "\n"))
	}
	for _, httpReq := range s.ScanTargets {
		paramdata := []attacker.ParamData{}
		for _, paramAndValue := range strings.Split(attacker.GetStringBody(httpReq.Body), "&") {
			var (
				maxTypeName  string
				maxTypeCount int
			)
			ss := strings.Split(paramAndValue, "=")
			name, reqdb, getdatas :=
				ss[0], db.OpenDatabase(), []getdata{}
			reqdb.
				Table("request_data").Select("Name,Value, count(Value) AS count").
				Where("host = ? AND path = ? AND method = ? AND name = ?",
					httpReq.URL.Host, httpReq.URL.Path, httpReq.Method, name).
				Group("value").Order("count Desc").
				Find(&getdatas)
			voting := map[string]int{"STRING": 0, "INT": 0, "BOOL": 0}
			for _, data := range getdatas {
				voting[s.getParamType(data.Value)] += data.Count
				if maxTypeCount < voting[s.getParamType(data.Value)] {
					maxTypeCount, maxTypeName = voting[s.getParamType(data.Value)], s.getParamType(data.Value)
				}
			}
			paramdata = append(paramdata, attacker.ParamData{
				Name:     name,
				Type:     maxTypeName,
				DefaultV: getdatas[0].Value,
			})
		}
		for _, payload := range payloads {
			attacker.Attack(httpReq, paramdata, payload)
		}
	}
}

func (s *scanner) s() {

}

func (s *scanner) getParamType(param string) string {
	if s.isInt(param) {
		return "INT"
	} else if s.isBool(param) {
		return "BOOL"
	}
	return "STRING"
}

func (s *scanner) isInt(param string) bool {
	convI, _ := strconv.ParseInt(param, 10, 64)
	if convI == 0 {
		convS := strconv.FormatInt(convI, 10)
		if convS != param {
			return false
		}
	}
	return true
}

func (s *scanner) isBool(param string) bool {
	param = strings.ToLower(param)
	convB, _ := strconv.ParseBool(param)
	if !convB {
		convS := strconv.FormatBool(convB)
		if convS != param {
			return false
		}
	}
	return true
}
func (s *scanner) C() {

}
func New(scanTargets []http.Request) scanner {
	s := scanner{
		ScanTargets: scanTargets,
	}
	return s
}

package scanner

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/WestEast1st/Matchlock/datastore"
	"github.com/WestEast1st/Matchlock/extractor"
	"github.com/WestEast1st/Matchlock/scanner/attacker"
	"github.com/WestEast1st/Matchlock/scanner/attacker/payload"
)

type scanner struct {
	ScanTargets []http.Request
}

const Inspection = "inspection"

type getdata struct {
	Name  string
	Value string
	Count int
}

func (s *scanner) setParamData(req http.Request, paramAndValues []string) []attacker.ParamData {
	var (
		maxTypeName  string
		maxTypeCount int
		t            string //type
		voting       = map[string]int{"STRING": 0, "INT": 0, "BOOL": 0}
	)
	getdatas := s.getDatas(req, paramAndValues[0])
	for _, data := range getdatas {
		t = s.getParamType(data.Value)
		voting[t] += data.Count
		if maxTypeCount < voting[t] {
			maxTypeCount, maxTypeName = voting[t], t
		}
	}
	paramData := []attacker.ParamData{{
		Name:     strings.Split(paramAndValues[0], "=")[0],
		Type:     maxTypeName,
		DefaultV: getdatas[0].Value,
	}}
	if len(paramAndValues) < 2 {
		return paramData
	}
	return append(paramData, s.setParamData(req, paramAndValues[1:])...)
}

func (s *scanner) attackRun(reqs []http.Request, ps map[string]map[string][]string) {
	/*
		I think that you can remove go in this function, but the inspection efficiency drops a little.
	*/
	if len(reqs) > 1 {
		go s.attackRun(reqs[1:], ps)
	}
	paramdata := []attacker.ParamData{}
	requestBody := extractor.GetStringBody(reqs[0].Body)
	if len(requestBody) < 1 {
		return
	}
	paramdata = s.setParamData(reqs[0], strings.Split(requestBody, "&"))
	go attacker.Attack(reqs[0], paramdata, ps)
}

func (s *scanner) Scan(typeString string) { //tmpname いずれ変える
	switch typeString {
	case Inspection:
		p := payload.Payload{}
		var ps = map[string]map[string][]string{}
		for _, ts := range p.GetTypeKeys(Inspection) {
			ps[ts] = map[string][]string{}
			for _, name := range p.GetFileName(ts) {
				ps[ts][name] = []string{}
				ps[ts][name] = p.GetPayload(ts, name)
			}
		}
		fmt.Println("# [INFO] The scan target is the following URL")
		for _, req := range s.ScanTargets {
			fmt.Println("	- ", req.URL.String())
		}
		fmt.Println("=============================================\n")
		s.attackRun(s.ScanTargets, ps)
	}
}

func (s *scanner) getDatas(httpReq http.Request, paramAndValue string) []getdata {
	ss := strings.Split(paramAndValue, "=")
	name, db, getdatas :=
		ss[0], datastore.DB.OpenDatabase(), []getdata{}
	db.Table("request_data").Select("Name, Value, count(Value) AS count").
		Joins("LEFT JOIN requests ON requests.identifier = request_data.identifier").
		Where("host = ? AND path = ? AND method = ? AND name = ?",
			httpReq.URL.Host, httpReq.URL.Path, httpReq.Method, name).
		Group("value").Order("count Desc").
		Find(&getdatas)
	return getdatas
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
		if strconv.FormatInt(convI, 10) != param {
			return false
		}
	}
	return true
}

func (s *scanner) isBool(param string) bool {
	param = strings.ToLower(param)
	convB, _ := strconv.ParseBool(param)
	if !convB {
		if strconv.FormatBool(convB) != param {
			return false
		}
	}
	return true
}

func New(scanTargets []http.Request) scanner {
	return scanner{ScanTargets: scanTargets}
}

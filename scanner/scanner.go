package scanner

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/a-zara-n/Matchlock/datastore"
	"github.com/a-zara-n/Matchlock/extractor"
	"github.com/a-zara-n/Matchlock/scanner/attacker"
	"github.com/a-zara-n/Matchlock/scanner/attacker/payload"
	"github.com/a-zara-n/Matchlock/shared"
)

type scanner struct {
	ScanTargets []http.Request
}

const Inspection = "inspection"

type getdata struct {
	Name  string
	Value string
	Type  string
	Count int
}

func (s *scanner) setParamData(req http.Request, paramAndValues [][]string) []attacker.ParamData {
	var (
		maxTypeName  string
		maxTypeCount int
		t            string //type
		voting       = map[string]int{"STRING": 0, "INT": 0, "BOOL": 0}
	)
	getdatas := s.getDatas(req, strings.Join(paramAndValues[0], "="))
	for _, data := range getdatas {
		t = getParamType(data.Value)
		voting[t] += data.Count
		if maxTypeCount < voting[t] {
			maxTypeCount, maxTypeName = voting[t], t
		}
	}
	paramData := []attacker.ParamData{{
		Name:     paramAndValues[0][0],
		TypeOf:   maxTypeName,
		Type:     getdatas[0].Type,
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
	var typ string
	if shared.CheckRegexp(`^{(\".*\":\"?.*\"?,?)+[^,]}$`, requestBody) {
		typ = "JSON"
	}
	paramdata = s.setParamData(reqs[0], shared.QueryDeconverter(typ, requestBody))
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
	db.Table("request_data").Select("Name, Value, Type, count(Value) AS count").
		Joins("LEFT JOIN requests ON requests.identifier = request_data.identifier").
		Where("host = ? AND path = ? AND method = ? AND name = ?",
			httpReq.URL.Host, httpReq.URL.Path, httpReq.Method, name).
		Group("value").Order("count Desc").
		Find(&getdatas)
	return getdatas
}

func New(scanTargets []http.Request) scanner {
	return scanner{ScanTargets: scanTargets}
}

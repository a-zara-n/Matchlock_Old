package attacker

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/cookiejar"

	"github.com/a-zara-n/Matchlock/scanner/attacker/payload"
	"github.com/a-zara-n/Matchlock/shared"
)

func Attack(req http.Request, paramdata []ParamData, ps map[string]map[string][]string) {
	bodys, names, defaultVs := setParamData(paramdata)
	jar, _ := cookiejar.New(nil)
	c := http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := c.Do(&req)
	if err != nil {
		fmt.Println("hoge")
	}
	b := shared.QueryConverter(paramdata[0].Type, bodys)
	var str string
	str, res.Body = shared.SeparationOfIOReadCloser(res.Body)
	attack := attacker{
		Request:      &req,
		Response:     res,
		ResponseBody: str,
		client:       &c,
		paramtmplate: template.Must(template.New("").Parse(b)),
	}
	/*
		You will lose some speed if you lose goroutines here.
		If necessary, remove it.
	*/
	go func(at attacker) {
		for div, datas := range ps {
			for types, data := range datas {
				go at.SimpleList(names, defaultVs, payload.Payload{
					Division: div,
					Type:     types,
					Data:     data,
				})
			}
		}
	}(attack)
}

type attacker struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody string
	client       *http.Client
	paramtmplate *template.Template
}

func (a attacker) AllChange(name []string, defaultV map[string]string, payloadData payload.Payload) {
	m := map[string]string{}
	for _, nm := range name {
		m[nm] = defaultV[nm]
	}
	for _, d := range payloadData.Data {
		for _, nm := range name {
			m[nm] = d
		}
		a.scanClientRun(m, payloadData)
	}
}

func (a attacker) SimpleList(name []string, defaultV map[string]string, payloadData payload.Payload) {
	m := map[string]string{}
	for _, nm := range name {
		m[nm] = defaultV[nm]
	}
	for _, nm := range name {
		tmp := m[nm]
		for _, d := range payloadData.Data {
			m[nm] = d
			a.scanClientRun(m, payloadData)
		}
		m[nm] = tmp
	}
}

func (a attacker) Cluster(name []string, defaultV map[string]string, payloadData payload.Payload) {
	m := map[string]string{}
	for _, nm := range name {
		m[nm] = defaultV[nm]
	}
	var function func(length int, i int, m map[string]string)
	function = func(length int, i int, m map[string]string) {
		a.Request.Close = true
		if length > i {
			for _, p := range payloadData.Data {
				m[name[i]] = p
				function(length, i+1, m)
			}
		} else {
			a.scanClientRun(m, payloadData)
		}
	}
	function(len(name), 0, m)
}

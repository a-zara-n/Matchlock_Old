package attacker

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/WestEast1st/Matchlock/extractor"
	"github.com/WestEast1st/Matchlock/scanner/attacker/decid"
	"github.com/WestEast1st/Matchlock/scanner/attacker/payload"
	"github.com/WestEast1st/Matchlock/shared"
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
	var str string
	str, res.Body = shared.SeparationOfIOReadCloser(res.Body)
	attack := attacker{
		Request:      &req,
		Response:     res,
		ResponseBody: str,
		client:       &c,
		paramtmplate: template.Must(template.New("").Parse(strings.Join(bodys, "&"))),
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

func (a attacker) scanClientRun(submitValues map[string]string, payloadData payload.Payload) {
	var buf bytes.Buffer
	a.paramtmplate.Execute(&buf, submitValues)
	a.setSubmitValue(html.UnescapeString(buf.String()))
	resp := a.sender()
	a.decider(resp.Body, payloadData, buf.String())
	resp.Body.Close()
}

func (a attacker) sender() *http.Response {
	resp, err := a.client.Do(a.Request)
	if err != nil {
		panic(err)
	}
	return resp
}

func (a attacker) setSubmitValue(submitValue string) {
	if a.Request.Method == "POST" {
		a.Request.Body = extractor.GetIOReadCloser(submitValue)
	} else {
		a.Request.URL.RawQuery = submitValue
	}
}

func (a attacker) decider(resp io.ReadCloser, payloadData payload.Payload, input string) {
	go decid.Decider(
		lineDiff(a.ResponseBody, extractor.GetStringBody(resp)), payloadData, *a.Request, input,
	)
}

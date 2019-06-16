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
	"github.com/sergi/go-diff/diffmatchpatch"
)

type ParamData struct {
	Name     string
	Type     string
	DefaultV string
}

/*
 Attack はattacker.goの関数を動かす仮の関数
*/

func setParamData(pdata []ParamData) ([]string, []string, map[string]string) {
	body, name, defvlue :=
		[]string{pdata[0].Name + "={{." + pdata[0].Name + "}}"},
		[]string{pdata[0].Name},
		map[string]string{pdata[0].Name: pdata[0].DefaultV}
	if len(pdata) > 1 {
		bodys, names, defvalues := setParamData(pdata[1:])
		return append(body, bodys...), append(name, names...), merge(defvlue, defvalues)
	}
	return body, name, defvlue
}

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
	str, res.Body = SeparationOfIOReadCloser(res.Body)
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
		var c int
		for div, datas := range ps {
			for types, data := range datas {
				c += len(data)
				p := payload.Payload{
					Division: div,
					Type:     types,
					Data:     data,
				}
				go at.SimpleList(names, defaultVs, p)
			}
		}
		//fmt.Println(int(math.Pow(float64(c), float64(len(names)))))
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
		var buf bytes.Buffer
		for _, nm := range name {
			m[nm] = d
		}
		a.paramtmplate.Execute(&buf, m)
		if a.Request.Method == "POST" {
			a.Request.Body = GetIOReadCloser(html.UnescapeString(buf.String()))
		} else {
			a.Request.URL.RawQuery = html.UnescapeString(buf.String())
		}
		resp, _ := a.client.Do(a.Request)
		fmt.Println(html.UnescapeString(buf.String()))
		resp.Body.Close()
		//fmt.Println(GetStringBody(resp.Body))
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
			var buf bytes.Buffer
			m[nm] = d
			a.paramtmplate.Execute(&buf, m)
			if a.Request.Method == "POST" {
				a.Request.Body = GetIOReadCloser(html.UnescapeString(buf.String()))
			} else {
				a.Request.URL.RawQuery = html.UnescapeString(buf.String())
			}
			resp, _ := a.client.Do(a.Request)
			body := GetStringBody(resp.Body)
			res := lineDiff(a.ResponseBody, body)
			go decid.Decider(res, payloadData, *a.Request, buf.String())
			resp.Body.Close()

			//fmt.Println(GetStringBody(resp.Body))
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
			var buf bytes.Buffer
			a.paramtmplate.Execute(&buf, m)
			if a.Request.Method == "POST" {
				a.Request.Body = GetIOReadCloser(html.UnescapeString(buf.String()))
			} else {
				a.Request.URL.RawQuery = html.UnescapeString(buf.String())
			}
			resp, err := a.client.Do(a.Request)
			if err != nil {
				panic(err)
			}
			//fmt.Println(resp.Status)
			body := GetStringBody(resp.Body)
			res := lineDiff(a.ResponseBody, body)
			go decid.Decider(res, payloadData, *a.Request, buf.String())
			resp.Body.Close()
			//time.Sleep(10 * time.Millisecond)
		}
	}
	function(len(name), 0, m)
}

func lineDiff(src1, src2 string) []diffmatchpatch.Diff {
	dmp := diffmatchpatch.New()
	a, b, c := dmp.DiffLinesToChars(src1, src2)
	diffs := dmp.DiffMain(a, b, false)
	result := dmp.DiffCharsToLines(diffs, c)
	return result
}

func SeparationOfIOReadCloser(b io.ReadCloser) (string, io.ReadCloser) {
	bodyOfStr := extractor.GetStringBody(b)
	b = extractor.GetIOReadCloser(bodyOfStr)
	return bodyOfStr, b
}

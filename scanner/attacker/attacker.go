package attacker

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

type ParamData struct {
	Name     string
	Type     string
	DefaultV string
}

/*
 Attack はattacker.goの関数を動かす仮の関数
*/
func Attack(req http.Request, paramdata []ParamData, ps map[string]map[string][]string) {
	var (
		body      []string
		names     []string
		defaultVs = map[string]string{}
	)
	jar, _ := cookiejar.New(nil)
	for _, pd := range paramdata {
		body = append(body, pd.Name+"={{."+pd.Name+"}}")
		names = append(names, pd.Name)
		defaultVs[pd.Name] = pd.DefaultV
	}
	attack := attacker{
		Request: &req,
		client: &http.Client{
			Jar: jar,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		paramtmplate: template.Must(template.New("").Parse(strings.Join(body, "&"))),
	}
	/*
		You will lose some speed if you lose goroutines here.
		If necessary, remove it.
	*/
	go func(at attacker) {
		for _, datas := range ps {
			for _, data := range datas {
				go at.Cluster(names, defaultVs, data)
			}
		}
	}(attack)

}

type attacker struct {
	Request      *http.Request
	client       *http.Client
	paramtmplate *template.Template
}

func (a attacker) AllChange(name []string, defaultV map[string]string, payload []string) {
	m := map[string]string{}
	for _, nm := range name {
		m[nm] = defaultV[nm]
	}
	for _, d := range payload {
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
		resp = resp
		//fmt.Println(GetStringBody(resp.Body))
	}

}

func (a attacker) SimpleList(name []string, defaultV map[string]string, payload []string) {
	m := map[string]string{}
	for _, nm := range name {
		m[nm] = defaultV[nm]
	}
	for _, nm := range name {
		tmp := m[nm]
		for _, d := range payload {
			var buf bytes.Buffer
			m[nm] = d
			a.paramtmplate.Execute(&buf, m)
			if a.Request.Method == "POST" {
				a.Request.Body = GetIOReadCloser(html.UnescapeString(buf.String()))
			} else {
				a.Request.URL.RawQuery = html.UnescapeString(buf.String())
			}
			resp, _ := a.client.Do(a.Request)
			fmt.Println(a.Request.URL)
			fmt.Println(GetStringBody(a.Request.Body))
			resp = resp
			//fmt.Println(GetStringBody(resp.Body))
		}
		m[nm] = tmp
	}
}

func (a attacker) Cluster(name []string, defaultV map[string]string, payload []string) {
	m := map[string]string{}
	for _, nm := range name {
		m[nm] = defaultV[nm]
	}
	var function func(length int, i int, m map[string]string)
	function = func(length int, i int, m map[string]string) {
		a.Request.Close = true

		if length > i {
			for _, p := range payload {
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
			resp.Body.Close()
			//time.Sleep(10 * time.Millisecond)
		}
	}
	function(len(name), 0, m)
}

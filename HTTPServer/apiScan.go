package httpserver

import (
	"bytes"
	"html/template"
	"net/http"
	"regexp"

	"github.com/WestEast1st/Matchlock/scanner"
	"github.com/labstack/echo"
)

func Scan(c echo.Context) error {
	/*
		//Domainはその文字列が含まれているホストを探すための文字列
		domain 	= c.Param("domain")
		//typeはAllやFuzzing 脆弱性名を利用する
		type 	= c.Param("type")
	*/
	d := struct {
		Host string
	}{
		Host: c.Param("domain"),
	}
	regexpString := `^[0-9a-zA-Z]*\.?({{.Host}})(\.+[0-9a-zA-Z]+)*$`
	tpl, _ := template.New("").Parse(regexpString)
	var buf bytes.Buffer
	tpl.Execute(&buf, d)
	rehost := buf.String()
	r := scanner.Request{}
	hosts := []struct {
		Host string
	}{}
	go func() {
		hostdb := db.OpenDatabase()
		hostdb.Table("requests").
			Select("DISTINCT host").
			Find(&hosts)
		for _, host := range hosts {

			if check_regexp(rehost, host.Host) {
				req := r.GetRequest(host.Host)
				s := scanner.New(req)
				s.Scan(c.Param("type"))
			}
		}
	}()
	res := c.Response()
	res.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, "{\"OK\":\"ok\"}")
	return nil
}

func check_regexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

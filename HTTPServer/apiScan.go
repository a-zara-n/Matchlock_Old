package httpserver

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/a-zara-n/Matchlock/shared"

	"github.com/a-zara-n/Matchlock/datastore"
	"github.com/a-zara-n/Matchlock/scanner"
	"github.com/labstack/echo"
)

func Scan(c echo.Context) error {
	/*
		//Domainはその文字列が含まれているホストを探すための文字列
		domain 	= c.Param("domain")
		//typeはAllやFuzzing 脆弱性名を利用する
		type 	= c.Param("type")
	*/
	var buf bytes.Buffer
	regexpString := `^[0-9a-zA-Z]*\.?({{.Host}})(\.+[0-9a-zA-Z]+)*$`
	tpl, _ := template.New("").Parse(regexpString)
	tpl.Execute(&buf, struct{ Host string }{Host: c.Param("domain")})
	rehost := buf.String()
	r := scanner.Request{}
	hosts := []struct {
		Host string
	}{}
	go func() {
		db := datastore.DB.OpenDatabase()
		db.Table("requests").
			Select("DISTINCT host").
			Find(&hosts)
		for _, host := range hosts {
			if shared.CheckRegexp(rehost, host.Host) {
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

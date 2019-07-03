package httpserver

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

//render はecho.Echoを排出します
func renders() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("../_view/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

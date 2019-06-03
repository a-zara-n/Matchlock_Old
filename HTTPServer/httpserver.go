package httpserver

import (
	"html/template"
	"io"
	"net/http"

	"../channel"
	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

/*
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data := map[string]interface{}{
		"Host": r.Host,
	}
	t.templ.Execute(w, data)
}
*/
type ServiceInfo struct {
	Title string
}

var serviceInfo = ServiceInfo{
	"Web dynamic testing",
}

type HttpServer interface {
	Run()
}
type httpServer struct {
	channels *channel.Matchlock
}

func (h *httpServer) Run() {
	t := &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}
	conn := newConnect(h.channels)
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(99)
	e.Renderer = t
	e.GET("/", func(c echo.Context) error {
		data := struct {
			ServiceInfo
			Host string
		}{
			ServiceInfo: serviceInfo,
			Host:        c.Request().Host,
		}
		return c.Render(http.StatusOK, "index", data)
	})
	e.GET("/connect", func(c echo.Context) error {
		conn.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	e.GET("/api/is/forward", func(c echo.Context) error {
		h.changeForward(c)
		return c.JSON(http.StatusOK, "{status: \"OK\"}")
	})
	e.GET("/api/history/all", func(c echo.Context) error {
		GetHistry(c.Response(), c.Request())
		return nil
	})
	go conn.Run()

	e.Start(":8888")
	/*
		http.Handle("/", &templateHandler{filename: "proto.html"})
		http.Handle("/connect", c)
		http.HandleFunc("/api/is/forward", h.changeForward)
		http.HandleFunc("/api/history/all", GetHistry)
		go c.Run()
		if err := http.ListenAndServe(":8888", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	*/
}

func NewHTTPServer(m *channel.Matchlock) HttpServer {
	return &httpServer{channels: m}
}

func (h *httpServer) changeForward(c echo.Context) {
	if h.channels.IsForward {
		h.channels.IsForward = false
	} else {
		h.channels.IsForward = true
	}
}

package service

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/a-zara-n/Matchlock/src/domain/value"
	"golang.org/x/net/html"
)

//CrawlerInterface は
type CrawlerInterface interface {
	Init(r io.Reader, method, urlstring string)
	Run(depth int)
}

//Crawler は
type Crawler struct {
	crawlingHost *regexp.Regexp
	data         crawlingdata
	memreq       repository.RequestRepositry
	memresp      repository.ResponseRepositry
}

//NewCrawler は
func NewCrawler(request repository.RequestRepositry, response repository.ResponseRepositry) CrawlerInterface {
	return &Crawler{
		memreq:  request,
		memresp: response,
	}
}

func (craw *Crawler) Init(r io.Reader, method, urlstring string) {
	u, _ := url.Parse(urlstring)
	craw.crawlingHost = regexp.MustCompile(`^(http(s)?|file|data|javascript)?://` + u.Host + `/.*`)
	craw.data = crawlingdata{
		Reader: r,
		Method: method,
		URL:    u,
		Parent: nil,
		Chiled: map[string]map[string]crawlingdata{},
	}
}

func (craw *Crawler) Run(depth int) {
	crawlinglist = value.NewCheckList()
	craw.data.Run(depth)
}

var (
	reqHost *regexp.Regexp
	r1      = regexp.MustCompile(`^(http(s)?|file|data|javascript)?://.*`)
	r2      = regexp.MustCompile(`^\.(\.)?.*?`)
	r3      = regexp.MustCompile(`^\/.*?`)
	r4      = regexp.MustCompile(`^.*\..*$`)
	r5      = regexp.MustCompile(`.*#.*$`)
	urlRoot = map[bool]func(val string, ustr string) string{
		false: func(val string, ustr string) string { return nomalizationURLPath(ustr + val) },
		true:  func(val string, ustr string) string { return nomalizationURLPath(ustr, val) },
	}
	ustrTraverse = map[bool]func(val string, ustr string) string{
		false: func(val string, ustr string) string { return urlRoot[r3.MatchString(val)](val, ustr) },
		true:  func(val string, ustr string) string { return nomalizationURLPath(ustr + "/" + val) },
	}
	// schema
	urlCOrrection = map[bool]func(val string, ustr string) string{
		false: func(val string, ustr string) string { return ustrTraverse[r2.MatchString(val)](val, ustr) },
		true:  func(val string, ustr string) string { return val },
	}
)

func nomalizationURLPath(ustr string, pathstr ...string) string {
	u, _ := url.Parse(ustr)
	p := map[bool]string{true: pathstr[0], false: u.Path}[len(pathstr) > 0]
	u.Path = path.Clean(p)
	return u.String()
}

var crawlinglist value.CheckListInterface

type crawlingdata struct {
	Reader    io.Reader
	Request   *http.Request
	Title     string
	Method    string
	URL       *url.URL
	Parent    *crawlingdata
	CrawlData []value.HTMLElement
	Chiled    map[string]map[string]crawlingdata //URL : Method : io.Reader
}

func (c *crawlingdata) Run(depth int) {
	crawlinglist.Check(c.URL.String())
	if depth == 0 {
		return
	}
	depth--
	doc, err := html.Parse(c.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Chiled = map[string]map[string]crawlingdata{}
	c.clowl(doc)
}

func (c *crawlingdata) clowl(doc *html.Node) {
	c.strageElementData(doc, []string{"form"}, c.fetchForm)
	c.strageElementData(doc, []string{"a", "img"}, c.resstr)
	c.strageElementData(doc, []string{"title"}, c.setTitle)
}

func (c *crawlingdata) elementFillter(key string, filltringstrings []string) bool {
	if filltringstrings[0] == key {
		return true
	}
	recursion := map[bool]func() bool{
		true:  func() bool { return c.elementFillter(key, filltringstrings[1:]) },
		false: func() bool { return false },
	}
	return recursion[len(filltringstrings) > 1]()
}

func (c *crawlingdata) strageElementData(n *html.Node, elements []string, stragefunc func(d *html.Node)) {
	h := map[bool]func(d *html.Node){true: stragefunc, false: func(d *html.Node) {}}
	h[n.Type == html.ElementNode && c.elementFillter(n.Data, elements)](n)
	for chiled := n.FirstChild; chiled != nil; chiled = chiled.NextSibling {
		c.strageElementData(chiled, elements, stragefunc)
	}
}

func (c *crawlingdata) resstr(d *html.Node) {
	for _, v := range d.Attr {
		u, _ := url.Parse(c.URL.String())
		u.Fragment, u.RawQuery = "", ""
		ps := strings.Split(u.Path, "/")
		if !r5.MatchString(v.Val) && c.elementFillter(v.Key, []string{"href", "src"}) {
			if r4.MatchString(ps[len(ps)-1]) {
				ps[len(ps)-1] = ""
			}
			u.Path = strings.Join(ps, "/")
			fmt.Println(u.String(), "  :   ", v.Val)
			urlstring, _ := url.QueryUnescape(urlCOrrection[r1.MatchString(v.Val)](v.Val, u.String()))
			u, _ = url.Parse(urlstring)
			if !crawlinglist.Find("GET "+urlstring) && reqHost.MatchString(urlstring) {
				c.Chiled[urlstring] = map[string]crawlingdata{"GET": {Method: "GET", URL: u, Parent: c}}
				crawlinglist.Check("GET " + urlstring)
			}
		}
	}
}

func (c *crawlingdata) setTitle(d *html.Node) {
	c.Title = d.FirstChild.Data
}

func (c *crawlingdata) fetchFormInput(d *html.Node) []value.HTMLElement {
	ret := []value.HTMLElement{}
	for chiled := d.FirstChild; chiled != nil; chiled = chiled.NextSibling {
		if c.elementFillter(chiled.Data, []string{"input"}) {
			ret = append(ret, value.HTMLElement{Tag: chiled.Data, Attr: chiled.Attr})
		}
		ret = append(ret, c.fetchFormInput(chiled)...)
	}
	return ret
}

func (c *crawlingdata) fetchForm(d *html.Node) {
	form := value.HTMLElement{Tag: d.Data, Attr: d.Attr, ParamData: c.fetchFormInput(d)}
	c.CrawlData = append(c.CrawlData, form)
}

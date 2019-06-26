package crawler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Crawler struct {
	Reader    io.Reader
	Request   *http.Request
	Title     string
	Method    string
	URL       string
	Host      string
	Parent    *Crawler
	CrawlData []Result
	Chiled    map[string]map[string]Crawler //URL : Method : io.Reader
}

type Result struct {
	Tag       string
	Attr      []html.Attribute
	ParamData []Result
}

var crawlURLList = map[string]bool{}

func (crawler *Crawler) strageElementData(n *html.Node, elements []string, stragefunc func(d *html.Node)) {
	h := map[bool]func(d *html.Node){true: stragefunc, false: func(d *html.Node) {}}
	h[n.Type == html.ElementNode && crawler.elementFillter(n.Data, elements)](n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		crawler.strageElementData(c, elements, stragefunc)
	}
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

func (crawler *Crawler) resstr(d *html.Node) {
	for _, v := range d.Attr {
		u, _ := url.Parse(crawler.URL)
		u.Fragment, u.RawQuery = "", ""
		ps := strings.Split(u.Path, "/")
		if !r5.MatchString(v.Val) && crawler.elementFillter(v.Key, []string{"href", "src"}) {
			if r4.MatchString(ps[len(ps)-1]) {
				ps[len(ps)-1] = ""
			}
			u.Path = strings.Join(ps, "/")
			fmt.Println(u.String(), "  :   ", v.Val)
			urlstring, _ := url.QueryUnescape(urlCOrrection[r1.MatchString(v.Val)](v.Val, u.String()))

			if !crawlURLList["GET "+urlstring] && reqHost.MatchString(urlstring) {
				crawler.Chiled[urlstring] = map[string]Crawler{"GET": {Method: "GET", URL: urlstring, Parent: crawler}}
				crawlURLList["GET "+urlstring] = true
			}
		}
	}
}

func (crawler *Crawler) setTitle(d *html.Node) {
	crawler.Title = d.FirstChild.Data
}

func (crawler *Crawler) fetchFormInput(d *html.Node) []Result {
	ret := []Result{}
	for c := d.FirstChild; c != nil; c = c.NextSibling {
		if crawler.elementFillter(c.Data, []string{"input"}) {
			ret = append(ret, Result{Tag: c.Data, Attr: c.Attr})
		}
		ret = append(ret, crawler.fetchFormInput(c)...)
	}
	return ret
}

func (crawler *Crawler) fetchForm(d *html.Node) {
	form := Result{Tag: d.Data, Attr: d.Attr, ParamData: crawler.fetchFormInput(d)}
	crawler.CrawlData = append(crawler.CrawlData, form)
}

func (crawler *Crawler) clow(depth int, doc *html.Node) {
	crawler.strageElementData(doc, []string{"form"}, crawler.fetchForm)
	crawler.strageElementData(doc, []string{"a", "img"}, crawler.resstr)
	crawler.strageElementData(doc, []string{"title"}, crawler.setTitle)
}

func (crawler *Crawler) Run(depth int) {
	if len(crawlURLList) == 0 {
		crawlURLList[crawler.URL] = true
	}
	if depth == 0 {
		return
	}
	depth--
	doc, err := html.Parse(crawler.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	crawler.Chiled = map[string]map[string]Crawler{}
	crawler.clow(depth, doc)
	for k, cm := range crawler.Chiled {
		for me, c := range cm {
			req, _ := http.NewRequest(me, k, strings.NewReader(""))
			client := http.Client{}
			res, _ := client.Do(req)
			if res.StatusCode < 400 {
				c.Host, c.Request, c.Reader = crawler.Host, req, res.Body
				c.Run(depth)
				crawler.Chiled[k][me] = c
			}
		}
	}
}

func (crawler *Crawler) elementFillter(tag string, hoge []string) bool {
	if hoge[0] == tag {
		return true
	}
	recursion := map[bool]func() bool{
		true:  func() bool { return crawler.elementFillter(tag, hoge[1:]) },
		false: func() bool { return false },
	}
	return recursion[len(hoge) > 1]()
}

func (c *Crawler) SetRequest(req *http.Request) {
	c.Request = req
}

func New(r io.Reader, method string, urlstring string) *Crawler {
	u, _ := url.Parse(urlstring)
	reqHost = regexp.MustCompile(`^(http(s)?|file|data|javascript)?://` + u.Host + `/.*`)
	return &Crawler{
		Reader: r,
		Method: method,
		URL:    urlstring,
		Host:   u.Host,
		Parent: nil,
	}
}

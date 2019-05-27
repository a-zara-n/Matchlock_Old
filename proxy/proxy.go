package proxy

import (
	"fmt"
	"net/http"
	"regexp"

	"../channel"
	"github.com/elazarl/goproxy"
)

type Proxy interface {
	Run()
}
type proxyInfo struct {
	proxy    *goproxy.ProxyHttpServer
	channels *channel.Matchlock
}

func (p *proxyInfo) Run() {
	var isHost bool
	p.proxy = goproxy.NewProxyHttpServer()
	p.proxy.Verbose = false
	//テストで配置
	//proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*$"))).
	//	HandleConnect(goproxy.AlwaysMitm)
	//を利用したいがうまく削除ができるか不明なので考慮する
	AddWhiteList(`^[0-9a-zA-Z]*\.?(localhost)(\.+[0-9a-zA-Z]+)*$`)
	AddWhiteList(`^[0-9a-zA-Z]*\.?(google)(\.+[0-9a-zA-Z]+)*$`)
	p.proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	reqchan := p.channels.Request
	reschan := p.channels.Response
	p.proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			isHost = whitelistMatch(r.Host)
			if isHost {
				fmt.Println("hoge")
				reqchan.ProxToHMgSignal <- r
				r = <-reqchan.ProxToHMgSignal
			}
			return r, nil
		})
	p.proxy.OnResponse().DoFunc(
		func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			if isHost {
				fmt.Println("foge")
				reschan.ProxToHMgSignal <- r
				r = <-reschan.ProxToHMgSignal
			}
			return r
		})

	http.ListenAndServe(":10080", p.proxy)
}

func requestHandler(r *http.Request) (*http.Request, error) {
	return nil, nil
}

func NewProxy(m *channel.Matchlock) Proxy {
	return &proxyInfo{channels: m}
}

var whitelist = []string{}

func AddWhiteList(domain string) bool {
	whitelist = append(whitelist, domain)
	return true
}
func RemoveWhiteList(i int) bool {
	whitelist = append(whitelist[:i], whitelist[i+1:]...)
	return true
}
func GetALLWhiteList(domain string) []string {
	return whitelist
}
func UpdataWhiteList(domains []string) bool {
	whitelist = domains
	return true
}

func whitelistMatch(host string) bool {
	fmt.Println(host)
	for _, d := range whitelist {
		if check_regexp(d, host) {
			return true
		}
	}
	return false
}
func check_regexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

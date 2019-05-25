package proxy

import (
	"net/http"

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
	p.proxy = goproxy.NewProxyHttpServer()
	//p.proxy.Verbose = true
	reqchan := p.channels.Request
	reschan := p.channels.Response
	p.proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			if p.channels.IsForward {
				reqchan.ProxToHMgSignal <- r
				r = <-reqchan.ProxToHMgSignal
			}
			return r, nil
		})
	p.proxy.OnResponse().DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		reschan.ProxToHMgSignal <- r
		r = <-reschan.ProxToHMgSignal
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

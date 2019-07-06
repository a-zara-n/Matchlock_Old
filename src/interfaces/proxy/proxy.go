package proxy

import (
	"net/http"

	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/elazarl/goproxy"
)

//Proxy はLocalProxyの起動に必要なmehtodを定義します
type Proxy interface {
	Run()
}
type proxy struct {
	proxy   *goproxy.ProxyHttpServer
	usecase usecase.ProxyLogic
}

//NewProxy は新規でProxtの設定を提供します
func NewProxy(usecase usecase.ProxyLogic) Proxy {
	return &proxy{usecase: usecase}
}

func (p *proxy) Run() {
	p.proxy = goproxy.NewProxyHttpServer()
	p.proxy.Verbose = false
	p.proxy.OnRequest().DoFunc(p.usecase.MatchlockLogic)
	http.ListenAndServe(":10080", p.proxy)
}

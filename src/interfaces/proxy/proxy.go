package proxy

import (
	"net/http"

	"github.com/a-zara-n/Matchlock/src/application"
	"github.com/a-zara-n/Matchlock/src/domain/entity"

	"github.com/elazarl/goproxy"
)

//Proxy はLocalProxyの起動に必要なmehtodを定義します
type Proxy interface {
	Run()
}
type proxy struct {
	proxy       *goproxy.ProxyHttpServer
	application application.ProxyLogic
}

//NewProxy は新規でProxtの設定を提供します
func NewProxy(c *entity.Channel, application application.ProxyLogic) Proxy {
	return &proxy{application: application}
}

func (p *proxy) Run() {
	p.proxy = goproxy.NewProxyHttpServer()
	p.proxy.Verbose = false
	p.proxy.OnRequest().DoFunc(p.application.MatchlockLogic)
	http.ListenAndServe(":10080", p.proxy)
}

package application

import (
	"net/http"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/elazarl/goproxy"
)

//ProxyLogic は動作するProxyの中枢機能になる
type ProxyLogic interface {
	MatchlockLogic(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response)
}

type proxylogic struct {
	WhiteList *entity.WhiteList
	channel   *entity.Channel
}

//NewLogic はProxyを利用する際にMatchlock側で定義するLogicをさしています
func NewLogic(white *entity.WhiteList, channel *entity.Channel) ProxyLogic {
	return &proxylogic{WhiteList: white, channel: channel}
}

func (l *proxylogic) MatchlockLogic(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	if l.WhiteList.Check(r.Host) {
		l.channel.Request.ProxToHMgSignal <- r
		resp := <-l.channel.Response.ProxToHMgSignal
		return nil, resp
	}
	return r, nil
}

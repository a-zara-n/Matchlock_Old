package application

import (
	"log"
	"net/http"

	"github.com/a-zara-n/Matchlock/src/config"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/elazarl/goproxy"
)

//ProxyLogic は動作するProxyの中枢機能になる
type ProxyLogic interface {
	MatchlockLogic(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response)
}

type proxylogic struct {
	WhiteList *entity.WhiteList
	channel   *config.ProxyChannel
}

//NewLogic はProxyを利用する際にMatchlock側で定義するLogicをさしています
func NewLogic(white *entity.WhiteList, channel *config.ProxyChannel) ProxyLogic {
	return &proxylogic{WhiteList: white, channel: channel}
}

func (l *proxylogic) MatchlockLogic(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	if l.WhiteList.Check(r.Host) {
		log.Println("ホワイトリストにマッチしました")
		l.channel.Request <- r
		resp := <-l.channel.Response
		return nil, resp
	}
	return r, nil
}

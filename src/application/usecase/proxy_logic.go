package usecase

import (
	"fmt"
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
}

//NewLogic はProxyを利用する際にMatchlock側で定義するLogicをさしています
func NewLogic(white *entity.WhiteList) ProxyLogic {
	return &proxylogic{white}
}

func (l *proxylogic) MatchlockLogic(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	if l.WhiteList.Check(r.Host) {
		return nil, nil
	}
	fmt.Println(r.URL)
	return r, nil
}

package registry

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/interfaces/proxy"
)

//Registry はNewで生成されるものを定義しています
type Registry interface {
	//entity
	NewChannel() *entity.Channel
	NewWhiteList() *entity.WhiteList
	//usecase
	NewLogic() usecase.ProxyLogic
	//infrastructure
	//interfase
	NewProxy(usecase usecase.ProxyLogic) proxy.Proxy
	//総合的なランディング
	Run()
}

//NewChannel はentity.Channelを取得
func NewChannel() *entity.Channel {
	return entity.NewMatchChannel()
}

//NewWhiteList はentity.WhiteListを取得
func NewWhiteList() *entity.WhiteList {
	return &entity.WhiteList{}
}

//NewProxy はproxy.Proxyを取得
func NewProxy(usecase usecase.ProxyLogic) proxy.Proxy {
	return proxy.NewProxy(usecase)
}

//NewLogic はusecase.ProxyLogicを取得
func NewLogic(white *entity.WhiteList) usecase.ProxyLogic {
	return usecase.NewLogic(white)
}

//Run はサーバー関連の起動をする
func Run() {
	//Entity
	whitelist := NewWhiteList()
	//UseCase
	proxylogic := NewLogic(whitelist)
	//Interface
	p := NewProxy(proxylogic)
	//Runding
	p.Run()
}

package registry

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver"
	"github.com/a-zara-n/Matchlock/src/interfaces/proxy"
)

//NewProxy はproxy.Proxyを取得
func NewProxy(usecase usecase.ProxyLogic) proxy.Proxy {
	return proxy.NewProxy(usecase)
}

//NewHTTPServer はサーバー起動に必要なhttpServerを取得します
func NewHTTPServer(c *entity.Channel, h usecase.HTMLUseCase) httpserver.HttpServer {
	return httpserver.NewHTTPServer(c, h)
}

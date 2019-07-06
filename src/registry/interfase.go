package registry

import (
	"github.com/a-zara-n/Matchlock/src/application"
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/interfaces/command"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver"
	"github.com/a-zara-n/Matchlock/src/interfaces/proxy"
)

//Interface はinterfase層で必要なものをまとめています
type Interface interface {
	NewProxy(usecase application.ProxyLogic) proxy.Proxy
	NewHTTPServer(c *entity.Channel, h usecase.HTMLUseCase, api *usecase.APIUsecase) httpserver.HttpServer
	NewCommand() command.Command
}

//NewProxy はproxy.Proxyを取得
func NewProxy(c *entity.Channel, application application.ProxyLogic) proxy.Proxy {
	return proxy.NewProxy(c, application)
}

//NewHTTPServer はサーバー起動に必要なhttpServerを取得します
func NewHTTPServer(c *entity.Channel, h usecase.HTMLUseCase, api *usecase.APIUsecase, ws usecase.WebSocketUsecase) httpserver.HttpServer {
	return httpserver.NewHTTPServer(c, h, api, ws)
}

//NewCommand はコマンドに必要な処理を永続化します
func NewCommand() command.Command {
	return command.NewCommand()
}

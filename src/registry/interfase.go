package registry

import (
	"github.com/a-zara-n/Matchlock/src/application"
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/config"
	"github.com/a-zara-n/Matchlock/src/interfaces/command"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver"
	"github.com/a-zara-n/Matchlock/src/interfaces/proxy"
)

//Interface はinterfase層で必要なものをまとめています
type Interface interface {
	NewProxy(application application.ProxyLogic) proxy.Proxy
	NewHTTPServer(httpserverchannel *config.HTTPServerChannel, html usecase.HTMLUsecase, api *usecase.APIUsecase, websocket usecase.WebSocketUsecase, manager usecase.ManagerUsecase) httpserver.HttpServer
	NewCommand() command.Command
}

//NewProxy はproxy.Proxyを取得
func (r *registry) NewProxy(application application.ProxyLogic) proxy.Proxy {
	return proxy.NewProxy(application)
}

//NewHTTPServer はサーバー起動に必要なhttpServerを取得します
func (r *registry) NewHTTPServer(httpserverchannel *config.HTTPServerChannel, html usecase.HTMLUsecase, api *usecase.APIUsecase, websocket usecase.WebSocketUsecase, manager usecase.ManagerUsecase) httpserver.HttpServer {
	return httpserver.NewHTTPServer(httpserverchannel, html, api, websocket, manager)
}

//NewCommand はコマンドに必要な処理を永続化します
func (r *registry) NewCommand() command.Command {
	return command.NewCommand()
}

package registry

import (
	"os"
	"os/signal"

	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver"
	"github.com/a-zara-n/Matchlock/src/interfaces/proxy"

	"github.com/a-zara-n/Matchlock/src/interfaces/command"

	"github.com/a-zara-n/Matchlock/src/config"
)

//Registry はNewで生成されるものを定義しています
type Registry interface {
	Config
	Entity
	Usecase
	Interface
	Repository
	//総合的なランディング
	Run()
}
type registry struct {
	Proxy   proxy.Proxy
	HTTP    httpserver.HttpServer
	Command command.Command
	Channel config.Channel
}

func NewRegistry() Registry {
	registry := &registry{}
	var (
		//config
		dbconf  = registry.NewDatabaseConfig()
		db      = dbconf.OpenDB(dbconf.GetConnect())
		channel = registry.NewMatchlockChannel()
		//Entity
		forward = registry.NewForward()
		//Repository
		reqrepo = registry.NewRequestRepositry(db)
		resrepo = registry.NewResponseRepositry(db)
		//UseCase
		html      = registry.NewHTMLUseCase()
		api       = registry.NewAPIUsecase(forward, reqrepo, resrepo)
		websocket = registry.NewWebSocketUsecase(reqrepo, resrepo)
		manager   = registry.NewManagerUsecase(&channel, reqrepo, resrepo, forward)
	)
	//Interface
	registry.Channel = channel
	registry.Proxy = registry.NewProxy(registry.NewLogic(registry.NewWhiteList(), channel.Proxy))
	registry.HTTP = registry.NewHTTPServer(channel.Server, html, api, websocket, manager)
	registry.Command = registry.NewCommand()
	return registry
}

//Run はサーバー関連の起動をする
func (r *registry) Run() {

	//Runding
	go r.Proxy.Run()
	go r.HTTP.Run()
	go r.Command.Run()
	sigClose(r.Channel)
}

// ctrl + c用の
func sigClose(m config.Channel) {
	// シグナル用のチャネル定義
	quit := make(chan os.Signal)
	// 受け取るシグナルを設定
	signal.Notify(quit, os.Interrupt)
	//<-m.ExitSignal
	<-quit // ここでシグナルを受け取るまで以降の処理はされない
}

package registry

import (
	"os"
	"os/signal"

	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/interfaces/httpserver"
	"github.com/a-zara-n/Matchlock/src/interfaces/proxy"
)

//Registry はNewで生成されるものを定義しています
type Registry interface {
	//entity
	NewChannel() *entity.Channel
	NewWhiteList() *entity.WhiteList
	//usecase
	NewLogic() usecase.ProxyLogic
	NewHTMLUseCase() usecase.HTMLUseCase
	//infrastructure
	//interfase
	NewProxy(usecase usecase.ProxyLogic) proxy.Proxy
	NewHTTPServer(c *entity.Channel, h usecase.HTMLUseCase) httpserver.HttpServer
	//総合的なランディング
	Run()
}

//Run はサーバー関連の起動をする
func Run() {
	//Entity
	whitelist := NewWhiteList()
	channel := NewChannel()
	//UseCase
	proxylogic := NewLogic(whitelist)
	html := NewHTMLUseCase()
	//Interface
	proxy := NewProxy(proxylogic)
	http := NewHTTPServer(channel, html)
	//Runding
	go proxy.Run()
	go http.Run()
	sigClose(channel)
}

// ctrl + c用の
func sigClose(m *entity.Channel) {
	// シグナル用のチャネル定義
	quit := make(chan os.Signal)
	// 受け取るシグナルを設定
	signal.Notify(quit, os.Interrupt)
	//<-m.ExitSignal
	<-quit // ここでシグナルを受け取るまで以降の処理はされない
}

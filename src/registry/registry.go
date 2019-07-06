package registry

import (
	"os"
	"os/signal"

	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//Registry はNewで生成されるものを定義しています
type Registry interface {
	Config
	Entity
	Usecase
	Interface
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
	apis := NewAPIUsecase()
	ws := NewWebSocketUsecase()
	//Interface
	proxy := NewProxy(channel, proxylogic)
	http := NewHTTPServer(channel, html, apis, ws)
	command := NewCommand()
	//Runding
	go proxy.Run()
	go http.Run()
	command.Run()
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

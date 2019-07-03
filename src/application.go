package main

import (
	"os"
	"os/signal"

	"github.com/a-zara-n/Matchlock/channel"
)

// ctrl + c用の
func sigClose(m *channel.Matchlock) {
	// シグナル用のチャネル定義
	quit := make(chan os.Signal)
	// 受け取るシグナルを設定
	signal.Notify(quit, os.Interrupt)
	//<-m.ExitSignal
	<-quit // ここでシグナルを受け取るまで以降の処理はされない
}

package main

import (
	"log"
	"os"
	"os/signal"
	"sort"

	cli "github.com/urfave/cli"
)

// server port
var proxyPort = "8888"

func main() {
	// 基本情報の設定
	app := cli.NewApp()
	app.Name = "VulunScaner"
	app.Usage = "It is a tool that can be used for vulnerability inspection when doing DevSecOps."
	app.Version = "β 0.0.1"
	// オプション
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "proxy-port, pp",
			Value: proxyPort,
			Usage: "Proxy port that works with this tool.",
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	// アプリケーションの起動
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	sigClose()
}

// ctrl + c用の
func sigClose() {
	// シグナル用のチャネル定義
	quit := make(chan os.Signal)
	// 受け取るシグナルを設定
	signal.Notify(quit, os.Interrupt)
	<-quit // ここでシグナルを受け取るまで以降の処理はされない
}

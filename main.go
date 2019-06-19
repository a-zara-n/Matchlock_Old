package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sort"

	hs "github.com/a-zara-n/Matchlock/HTTPServer"
	"github.com/a-zara-n/Matchlock/history"

	"github.com/a-zara-n/Matchlock/channel"
	"github.com/a-zara-n/Matchlock/datastore"
	"github.com/a-zara-n/Matchlock/proxy"
	cli "github.com/urfave/cli"
)

// server port
var proxyPort = "8888"
var title = `
   __  ___     __      __   __         __     ====
  /  |/  /__ _/ /_____/ /  / /__  ____/ /__  ====   Matchlock
 / /|_/ / _ '/ __/ __/ _ \/ / _ \/ __/  '_/ ====	〔 https://github.com/a-zara-n/Matchlock 〕
/_/  /_/\_,_/\__/\__/_//_/_/\___/\__/_/\_\ ====   
=============================================
`
var db = datastore.Database{Database: "./test.db"}

//init db
var dbschema = []interface{}{
	history.History{},
	history.Request{},
	history.RequestHeader{},
	history.RequestData{},
	history.Response{},
	history.ResponseHeader{},
	history.ResponseBody{},
	history.Cookie{},
}

func main() {
	// DB setting
	go func() {
		for _, v := range dbschema {
			db.Table = v
			db.InitMigration()
		}
	}()
	// 基本情報の設定
	app := cli.NewApp()
	app.Name = title
	app.Usage = "Matchlock is a web application vulnerability scanner"
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

	m := channel.NewMatchChannel()
	pr := proxy.NewProxy(m)
	hs := hs.NewHTTPServer(m)
	hh := NewHTTPmanager(m)

	app.Action = func(c *cli.Context) {
		fmt.Println(title)
		go pr.Run()
		go hs.Run()
		go hh.Run()
		sigClose(m)
	}
	// アプリケーションの起動
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

// ctrl + c用の
func sigClose(m *channel.Matchlock) {
	// シグナル用のチャネル定義
	quit := make(chan os.Signal)
	// 受け取るシグナルを設定
	signal.Notify(quit, os.Interrupt)
	//<-m.ExitSignal
	<-quit // ここでシグナルを受け取るまで以降の処理はされない
}

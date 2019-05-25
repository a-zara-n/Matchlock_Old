package main

import (
	"log"
	"os"
	"os/signal"
	"sort"

	hs "./HTTPServer"

	"./channel"
	"./datastore"
	"./proxy"
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

	m := channel.NewMatchChannel()
	pr := proxy.NewProxy(m)
	hs := hs.NewHTTPServer(m)
	hh := NewHHTTPmanager(m)

	app.Action = func(c *cli.Context) {
		db := datastore.Database{Database: "./test.db"}
		dbschema := []interface{}{
			History{},
			Request{},
			RequestHeader{},
			RequestData{},
			Response{},
		}
		go func() {
			for _, v := range dbschema {
				db.Table = v
				db.InitMigration()
			}
		}()
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

package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/a-zara-n/Matchlock/src/registry"
	"github.com/urfave/cli"
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

func main() {
	// 基本情報の設定
	app := cli.NewApp()
	app.Name = title
	app.Usage = "Matchlock is a web application vulnerability scanner"
	app.Version = "β 0.0.1"
	// オプション
	app.Flags = options()
	sort.Sort(cli.FlagsByName(app.Flags))
	app.Action = func(c *cli.Context) {
		fmt.Println(title)
		registry.Run()
	}
	// アプリケーションの起動
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

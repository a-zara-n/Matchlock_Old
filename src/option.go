package main

import "github.com/urfave/cli"

func options() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "proxy-port, pp",
			Value: proxyPort,
			Usage: "Proxy port that works with this tool.",
		},
	}
}

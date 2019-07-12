package config

/*
channel
このチャンネルパッケージはこのこのツールで利用される基幹チャンネルをまとめるて管理するためのチャンネルです。
*/
import (
	"net/http"
	"os"

	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
)

//ProxyChannel はProxyとmanagerとの間で利用される非同期チャンネルです
type ProxyChannel struct {
	Request  chan *http.Request
	Response chan *http.Response
	Common   *CommonChannel
}

//NewProxyChannel はチャンネルの生成を行います
func NewProxyChannel(c *CommonChannel) *ProxyChannel {
	return &ProxyChannel{
		Request:  make(chan *http.Request, 10000),
		Response: make(chan *http.Response, 10000),
		Common:   c,
	}
}

//HTTPServerChannel はhttpserverとmanagerとの間で利用される非同期チャンネルです
type HTTPServerChannel struct {
	Request  chan *aggregate.Request
	Response chan *aggregate.Request
	Common   *CommonChannel
}

//NewHTTPServerChannel はチャンネルの生成を行います
func NewHTTPServerChannel(c *CommonChannel) *HTTPServerChannel {
	return &HTTPServerChannel{
		Request:  make(chan *aggregate.Request, 10000),
		Response: make(chan *aggregate.Request, 10000),
		Common:   c,
	}
}

//CommonChannel は共通で必要なチャンネルです
type CommonChannel struct {
	ExitSignal chan os.Signal //終了のシグナルを送る
}

//NewCommonChannel は共通チャンネルを設定します
func NewCommonChannel() *CommonChannel {
	return &CommonChannel{
		ExitSignal: make(chan os.Signal),
	}
}

//Channel は全てのチャンネルを設定して返却します
type Channel struct {
	Proxy  *ProxyChannel
	Server *HTTPServerChannel
}

//NewMatchlockChannel はgoroutine間の通信を管理するMatchlockを返却する
func NewMatchlockChannel() Channel {
	commonchannel := NewCommonChannel()
	return Channel{
		Proxy:  NewProxyChannel(commonchannel),
		Server: NewHTTPServerChannel(commonchannel),
	}
}

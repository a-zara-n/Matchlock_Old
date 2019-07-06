package entity

/*
channel
このチャンネルパッケージはこのこのツールで利用される基幹チャンネルをまとめるて管理するためのチャンネルです。
*/
import (
	"net/http"
	"os"
)

//Matchlock はmain.goが管理するgoroutin間の通信を管理する部位である
type Channel struct {
	ExitSignal chan os.Signal //終了のシグナルを送る
	Request    RequestChan
	Response   ResponseChan
	IsForward  bool
}

//RequestChan はrequestをHTTPServerとプロキシのやり取りをするためのchannel
type RequestChan struct {
	ProxToHMgSignal chan *http.Request //ProxyとHTTPHQとの間のやり取り
	HMgToHsSignal   chan *http.Request //HTTPHQと
}

//ResponseChan はresponseをHTTPServerとプロキシのやり取りをするためのchannel
type ResponseChan struct {
	ProxToHMgSignal chan *http.Response //ProxyとHTTPHQとの間のやり取り
	HMgToHsSignal   chan *http.Response //HTTPHQと
}

//GetMatchlock はchannelをひとまとめにしたMatchlockを返却する
func (c *Channel) GetMatchlock() *Channel {
	return c
}

//NewMatchChannel はgoroutine間の通信を管理するMatchlockを返却する
func NewMatchChannel() *Channel {
	return &Channel{
		ExitSignal: make(chan os.Signal),
		Request: RequestChan{
			ProxToHMgSignal: make(chan *http.Request, 10000),
			HMgToHsSignal:   make(chan *http.Request),
		},
		Response: ResponseChan{
			ProxToHMgSignal: make(chan *http.Response, 10000),
			HMgToHsSignal:   make(chan *http.Response),
		},
		IsForward: false,
	}
}
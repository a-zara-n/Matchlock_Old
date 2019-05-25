package channel

import (
	"net/http"
	"os"
)

//Matchlock はmain.goが管理するgoroutin間の通信を管理する部位である
type Matchlock struct {
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
func (m *Matchlock) GetMatchlock() *Matchlock {
	return m
}

//NewMatchChannel はgoroutine間の通信を管理するMatchlockを返却する
func NewMatchChannel() *Matchlock {
	return &Matchlock{
		ExitSignal: make(chan os.Signal),
		Request: RequestChan{
			ProxToHMgSignal: make(chan *http.Request),
			HMgToHsSignal:   make(chan *http.Request),
		},
		Response: ResponseChan{
			ProxToHMgSignal: make(chan *http.Response),
			HMgToHsSignal:   make(chan *http.Response),
		},
		IsForward: false,
	}
}

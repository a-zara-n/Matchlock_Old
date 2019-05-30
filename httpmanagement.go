package main

import (
	"fmt"
	"io"
	"reflect"

	"./channel"
	"./extractor"
	"./history"
)

type HTTPmanager struct {
	channels *channel.Matchlock
}

func (h *HTTPmanager) Run() {
	var (
		bstr    string
		resH    = []history.History{}
		reqchan = h.channels.Request
		reschan = h.channels.Response
		sepIO   = SeparationOfIOReadCloser
	)
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			requestHistory := history.History{}
			resH = append(resH, requestHistory)
			bstr, req.Body = sepIO(req.Body)
			fmt.Println(req.URL.String())
			if h.channels.IsForward {
				reqchan.HMgToHsSignal <- req
				requestHistory.SetIdentifier(history.GetSha1(req.URL.String()))
				go requestHistory.MemoryRequest(req, false, bstr)
				creq := <-reqchan.HMgToHsSignal
				bstr, creq.Body = sepIO(req.Body)
				reqchan.ProxToHMgSignal <- creq
				if reflect.DeepEqual(req, creq) != true {
					go requestHistory.MemoryRequest(creq, true, bstr)
				}
			} else {
				requestHistory.SetIdentifier(history.GetSha1(req.URL.String()))
				go requestHistory.MemoryRequest(req, false, bstr)
				reqchan.ProxToHMgSignal <- req
			}
		case res := <-reschan.ProxToHMgSignal:
			bstr, res.Body = sepIO(res.Body)
			reschan.ProxToHMgSignal <- res
			go resH[0].MemoryResponse(res, bstr)
			if len(resH[1:]) > 0 {
				resH = resH[1:]
			}
		}
	}
}

func SeparationOfIOReadCloser(b io.ReadCloser) (string, io.ReadCloser) {
	bstr := extractor.GetStringBody(b)
	b = extractor.GetIOReadCloser(bstr)
	return bstr, b
}

func NewHHTTPmanager(m *channel.Matchlock) *HTTPmanager {
	return &HTTPmanager{channels: m}
}

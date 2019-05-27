package main

import (
	"fmt"
	"io"
	"reflect"

	"./extractor"

	"./channel"
)

type HTTPmanager struct {
	channels *channel.Matchlock
}

func (h *HTTPmanager) Run() {
	var (
		bstr    string
		resH    = []History{}
		reqchan = h.channels.Request
		reschan = h.channels.Response
		sepIO   = SeparationOfIOReadCloser
	)
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			history := History{}
			resH = append(resH, history)
			bstr, req.Body = sepIO(req.Body)
			fmt.Println(req.URL.String())
			if h.channels.IsForward {
				reqchan.HMgToHsSignal <- req
				history.SetIdentifier(GetSha1(req.URL.String()))
				go history.MemoryRequest(req, false, bstr)
				creq := <-reqchan.HMgToHsSignal
				bstr, creq.Body = sepIO(req.Body)
				reqchan.ProxToHMgSignal <- creq
				if reflect.DeepEqual(req, creq) != true {
					go history.MemoryRequest(creq, true, bstr)
				}
			} else {
				history.SetIdentifier(GetSha1(req.URL.String()))
				go history.MemoryRequest(req, false, bstr)
				reqchan.ProxToHMgSignal <- req
			}
		case res := <-reschan.ProxToHMgSignal:
			history := resH[0]
			if len(resH[1:]) > 0 {
				resH = resH[1:]
			}
			fmt.Println(res.Status)
			bstr, res.Body = sepIO(res.Body)
			reschan.ProxToHMgSignal <- res
			history.MemoryResponse(res, bstr)
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

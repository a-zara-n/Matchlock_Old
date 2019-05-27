package main

import (
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
		reqH    History
		resH    History
		reqchan = h.channels.Request
		reschan = h.channels.Response
		sepIO   = SeparationOfIOReadCloser
	)
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			reqH = History{}
			bstr, req.Body = sepIO(req.Body)
			if h.channels.IsForward {
				reqchan.HMgToHsSignal <- req
				reqH.SetIdentifier(GetSha1(req.URL.String()))
				go reqH.MemoryRequest(req, false, bstr)
				creq := <-reqchan.HMgToHsSignal
				bstr, creq.Body = sepIO(req.Body)
				resH = reqH
				reqchan.ProxToHMgSignal <- creq
				if reflect.DeepEqual(req, creq) != true {
					go reqH.MemoryRequest(creq, true, bstr)
				}
			} else {
				reqH.SetIdentifier(GetSha1(req.URL.String()))
				go reqH.MemoryRequest(req, false, bstr)
				resH = reqH
				reqchan.ProxToHMgSignal <- req
			}
		case res := <-reschan.ProxToHMgSignal:
			bstr, res.Body = sepIO(res.Body)
			reschan.ProxToHMgSignal <- res
			resH.MemoryResponse(res, bstr)
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

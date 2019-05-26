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
		reqchan = h.channels.Request
		reschan = h.channels.Response
		histry  = History{}
		sepIO   = SeparationOfIOReadCloser
	)
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			bstr, req.Body = sepIO(req.Body)
			if h.channels.IsForward {
				reqchan.HMgToHsSignal <- req
				histry.SetIdentifier(GetSha1(req.URL.String()))
				go histry.MemoryRequest(req, false, bstr)
				creq := <-reqchan.HMgToHsSignal
				bstr, creq.Body = sepIO(req.Body)
				reqchan.ProxToHMgSignal <- creq
				if reflect.DeepEqual(req, creq) != true {
					go histry.MemoryRequest(creq, true, bstr)
				}
			} else {
				histry.SetIdentifier(GetSha1(req.URL.String()))
				go histry.MemoryRequest(req, false, bstr)
				reqchan.ProxToHMgSignal <- req
			}
		case res := <-reschan.ProxToHMgSignal:
			bstr, res.Body = sepIO(res.Body)
			reschan.ProxToHMgSignal <- res
			histry.MemoryResponse(res, bstr)
			histry = History{}
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

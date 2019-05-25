package main

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"strings"

	"./channel"
)

type HTTPmanager struct {
	channels *channel.Matchlock
}

func (h *HTTPmanager) Run() {
	reqchan := h.channels.Request
	reschan := h.channels.Response
	histry := History{}
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			bufbody := new(bytes.Buffer)
			bufbody.ReadFrom(req.Body)
			bstr := bufbody.String()
			req.Body = ioutil.NopCloser(strings.NewReader(bstr))

			if h.channels.IsForward {
				reqchan.HMgToHsSignal <- req
				histry.SetIdentifier(GetSha1(req.URL.String()))
				go histry.MemoryRequest(req, false, bstr)
				creq := <-reqchan.HMgToHsSignal

				bufbody = new(bytes.Buffer)
				bufbody.ReadFrom(creq.Body)
				bstr, creq.Body = bufbody.String(), ioutil.NopCloser(strings.NewReader(bstr))

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
			bufbody := new(bytes.Buffer)
			bufbody.ReadFrom(res.Body)
			bstr := bufbody.String()
			res.Body = ioutil.NopCloser(strings.NewReader(bstr))
			reschan.ProxToHMgSignal <- res
			histry.MemoryResponse(res, bstr)
			histry = History{}
		}
	}
}
func NewHHTTPmanager(m *channel.Matchlock) *HTTPmanager {
	return &HTTPmanager{channels: m}
}

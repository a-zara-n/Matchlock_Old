package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"./channel"
	"./extractor"
	"./history"
)

type HTTPmanager struct {
	channels *channel.Matchlock
}

func (h *HTTPmanager) Run() {
	var (
		bstrreq   []string
		bstr      string
		b         string
		resH      = []history.History{}
		reqHeader = []http.Header{}
		reqchan   = h.channels.Request
		reschan   = h.channels.Response
		sepIO     = SeparationOfIOReadCloser
	)
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			bstrreq = []string{}
			requestHistory := history.History{
				IsEdit: false,
			}
			if h.channels.IsForward {
				requestHistory.IsEdit = true
			}
			b, req.Body = sepIO(req.Body)
			bstrreq = append(bstrreq, b)
			reqHeader = append(reqHeader, req.Header)
			requestHistory.SetIdentifier(history.GetSha1(req.URL.String()))
			go requestHistory.MemoryRequest(req, false, b)
			resH = append(resH, requestHistory)
			fmt.Println(req.URL.String())
			if h.channels.IsForward {
				reqchan.HMgToHsSignal <- req
				req = <-reqchan.HMgToHsSignal
				b, req.Body = sepIO(req.Body)
				bstrreq = append(bstrreq, b)
				reqHeader = append(reqHeader, req.Header)
				req.ContentLength = int64(len(b))
			}
			reqchan.ProxToHMgSignal <- req
			if h.channels.IsForward {
				var isEdit bool
				if bstrreq[0] != bstrreq[1] {
					isEdit = true
				}
				for name, data := range reqHeader[1] {
					if reqHeader[0][name] == nil {
						isEdit = true
						break
					}
					if strings.Join(reqHeader[0][name], ",") != strings.Join(data, ",") {
						isEdit = true
						break
					}
				}
				if isEdit {
					go requestHistory.MemoryRequest(req, true, bstrreq[1])
				}
			}
		case res := <-reschan.ProxToHMgSignal:
			bstr, res.Body = sepIO(res.Body)
			reschan.ProxToHMgSignal <- res
			go resH[0].MemoryResponse(res, bstr)
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

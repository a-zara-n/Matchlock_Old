package main

import (
	"io"
	"net/http"
	"time"

	"./channel"
	"./extractor"
	"./history"
)

//HTTPmanager is controls HTTP acquired by proxy etc.
type HTTPmanager struct {
	channels    *channel.Matchlock
	Information exchangeInformationOfHTTP
}

//Run is control of HTTP.
func (h *HTTPmanager) Run() {
	h.Information = newExchangeInformationOfHTTP()
	var (
		reqchan      = h.channels.Request
		reschan      = h.channels.Response
		sepIO        = SeparationOfIOReadCloser
		requests     = h.Information.Request
		editRequests = h.Information.EditRequest
	)
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			httphistory := history.History{
				IsEdit: h.channels.IsForward,
			}
			identifier := history.GetSha1(req.URL.String())
			httphistory.SetIdentifier(identifier)
			bodyOfStr := requests.SetRequest(req)
			go httphistory.MemoryRequest(req, false, bodyOfStr)
			if h.channels.IsForward {
				reqchan.HMgToHsSignal <- req
				req = <-reqchan.HMgToHsSignal
				bodyOfStr := editRequests.SetRequest(req)
				req.ContentLength = int64(len(bodyOfStr))
			}
			client := &http.Client{Timeout: time.Duration(10) * time.Second}
			req.RequestURI = ""
			resp, _ := client.Do(req)
			bodyOfStr, resp.Body = sepIO(resp.Body)
			go httphistory.MemoryResponse(resp, bodyOfStr)
			reschan.ProxToHMgSignal <- resp
			if h.channels.IsForward {
				isEdit, bodys, _ := h.Information.IsEdit()
				if isEdit {
					go httphistory.MemoryRequest(req, true, bodys[1])
				}
			}
		}
	}
}

//SeparationOfIOReadCloser io.ReadCloser => string & io.ReadCloser
func SeparationOfIOReadCloser(b io.ReadCloser) (string, io.ReadCloser) {
	bodyOfStr := extractor.GetStringBody(b)
	b = extractor.GetIOReadCloser(bodyOfStr)
	return bodyOfStr, b
}

//NewHTTPmanager is HTTPmanager structure return.
func NewHTTPmanager(m *channel.Matchlock) *HTTPmanager {
	return &HTTPmanager{channels: m}
}

package main

import (
	"io"
	"net/http"
	"time"

	"github.com/WestEast1st/Matchlock/channel"
	"github.com/WestEast1st/Matchlock/extractor"
	"github.com/WestEast1st/Matchlock/history"
	"github.com/WestEast1st/Matchlock/shared"
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
		reqchan   = h.channels.Request
		reschan   = h.channels.Response
		sepIO     = SeparationOfIOReadCloser
		bodyOfStr string
	)
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			httphistory := history.History{
				IsEdit: h.channels.IsForward,
			}
			httphistory.SetIdentifier(shared.GetSha1(req.URL.String()))
			go httphistory.MemoryRequest(req, false, h.Information.Request.SetRequest(req))
			if h.channels.IsForward {
				reqchan.HMgToHsSignal <- req
				req = <-reqchan.HMgToHsSignal
				req.ContentLength = int64(len(h.Information.EditRequest.SetRequest(req)))
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
			} else {
				h.Information.Request.DequeueBody()
				h.Information.Request.DequeueHeader()
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

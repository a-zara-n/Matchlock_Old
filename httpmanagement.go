package main

import (
	"io"
	"net/http"
	"strings"

	"./channel"
	"./extractor"
	"./history"
)

type HTTPmanager struct {
	channels    *channel.Matchlock
	Information []exchangeInformationOfHttp
}

type exchangeInformationOfHttp struct {
	Historys             []history.History
	Request, EditRequest []requestInfo
}

func (e *exchangeInformationOfHttp) IsEdit() (bool, []string, []http.Header) {
	bodyStrs := []string{e.Request[0].DequeueBody(), e.EditRequest[0].DequeueBody()}
	headers := []http.Header{e.Request[0].DequeueHeader(), e.EditRequest[0].DequeueHeader()}
	if bodyStrs[0] != bodyStrs[1] {
		return true, bodyStrs, headers
	}
	for name, data := range headers[0] {
		if headers[1][name] == nil {
			return true, bodyStrs, headers
		}
		if strings.Join(headers[1][name], ",") != strings.Join(data, ",") {
			return true, bodyStrs, headers
		}
	}
	return false, bodyStrs, headers
}

type requestInfo struct {
	QueueForBody    []string
	QueueForHeaders []http.Header
}

func (r *requestInfo) EnqueueBody(body string) {
	r.QueueForBody = append(r.QueueForBody, body)
}

func (r *requestInfo) EnqueueHeader(header http.Header) {
	r.QueueForHeaders = append(r.QueueForHeaders, header)
}

func (r *requestInfo) DequeueBody() string {
	data := r.QueueForBody[0]
	r.QueueForBody = r.QueueForBody[1:]
	return data
}

func (r *requestInfo) DequeueHeader() http.Header {
	data := r.QueueForHeaders[0]
	r.QueueForHeaders = r.QueueForHeaders[1:]
	return data
}

func (r *requestInfo) SetRequest(req http.Request) string {
	var bodyOfStr string
	bodyOfStr, req.Body = SeparationOfIOReadCloser(req.Body)
	r.EnqueueBody(bodyOfStr)
	r.EnqueueHeader(req.Header)
	return bodyOfStr
}

func (h *HTTPmanager) Run() {
	var (
		q4requestBody []string //Queue for request body
		bodyOfStr     string   //The body of the string
		historys      = []history.History{}
		headers       = []http.Header{} // request headers
		reqchan       = h.channels.Request
		reschan       = h.channels.Response
		sepIO         = SeparationOfIOReadCloser
	)
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			q4requestBody = []string{}

			httphistory := history.History{
				IsEdit: h.channels.IsForward,
			}

			httphistory.SetIdentifier(history.GetSha1(req.URL.String()))

			bodyOfStr, req.Body = sepIO(req.Body)
			q4requestBody, headers =
				append(q4requestBody, bodyOfStr), append(headers, req.Header)

			go httphistory.MemoryRequest(req, false, bodyOfStr)
			historys = append(historys, httphistory)

			if h.channels.IsForward {
				reqchan.HMgToHsSignal <- req
				req = <-reqchan.HMgToHsSignal

				bodyOfStr, req.Body = sepIO(req.Body)
				q4requestBody, headers =
					append(q4requestBody, bodyOfStr), append(headers, req.Header)

				req.ContentLength = int64(len(bodyOfStr))
			}

			reqchan.ProxToHMgSignal <- req

			if h.channels.IsForward {
				var isEdit bool
				if q4requestBody[0] != q4requestBody[1] {
					isEdit = true
				}
				for name, data := range headers[1] {
					if headers[0][name] == nil {
						isEdit = true
						break
					}
					if strings.Join(headers[0][name], ",") != strings.Join(data, ",") {
						isEdit = true
						break
					}
				}
				if isEdit {
					go httphistory.MemoryRequest(req, true, q4requestBody[1])
				}
			}
		case res := <-reschan.ProxToHMgSignal:
			bodyOfStr, res.Body = sepIO(res.Body)
			reschan.ProxToHMgSignal <- res
			go historys[0].MemoryResponse(res, bodyOfStr)
			historys = historys[1:]
		}
	}
}

func inspectionOfEdited() {

}

func SeparationOfIOReadCloser(b io.ReadCloser) (string, io.ReadCloser) {
	bodyOfStr := extractor.GetStringBody(b)
	b = extractor.GetIOReadCloser(bodyOfStr)
	return bodyOfStr, b
}

func NewHHTTPmanager(m *channel.Matchlock) *HTTPmanager {
	return &HTTPmanager{channels: m}
}

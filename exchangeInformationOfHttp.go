package main

import (
	"net/http"
	"strings"

	"github.com/WestEast1st/Matchlock/history"
)

type exchangeInformationOfHTTP struct {
	Historys             []history.History
	Request, EditRequest requestInfo
}

func (e *exchangeInformationOfHTTP) IsEdit() (bool, []string, []http.Header) {
	bodyStrs := []string{e.Request.DequeueBody(), e.EditRequest.DequeueBody()}
	headers := []http.Header{e.Request.DequeueHeader(), e.EditRequest.DequeueHeader()}
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

func (r *requestInfo) SetRequest(req *http.Request) string {
	var bodyOfStr string
	bodyOfStr, req.Body = SeparationOfIOReadCloser(req.Body)
	r.EnqueueBody(bodyOfStr)
	r.EnqueueHeader(req.Header)
	return bodyOfStr
}

func newExchangeInformationOfHTTP() exchangeInformationOfHTTP {
	return exchangeInformationOfHTTP{
		Historys:    []history.History{},
		Request:     requestInfo{},
		EditRequest: requestInfo{},
	}
}

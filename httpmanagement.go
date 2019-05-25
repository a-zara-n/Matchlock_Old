package main

import (
	"./channel"
)

type HTTPmanager struct {
	channels *channel.Matchlock
}

func (h *HTTPmanager) Run() {
	reqchan := h.channels.Request
	reschan := h.channels.Response
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			reqchan.HMgToHsSignal <- req
			req = <-reqchan.HMgToHsSignal
			reqchan.ProxToHMgSignal <- req
		case res := <-reschan.ProxToHMgSignal:
			reschan.ProxToHMgSignal <- res
		}
	}
}
func NewHHTTPmanager(m *channel.Matchlock) *HTTPmanager {
	return &HTTPmanager{channels: m}
}

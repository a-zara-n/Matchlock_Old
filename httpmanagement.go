package main

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"./channel"
	"./extractor"
	"./history"
)

type HTTPmanager struct {
	channels *channel.Matchlock
}

func (h *HTTPmanager) Run() {
	var (
		bstr    string
		resH    = []history.History{}
		reqchan = h.channels.Request
		reschan = h.channels.Response
		sepIO   = SeparationOfIOReadCloser
	)
	for {
		select {
		case req := <-reqchan.ProxToHMgSignal:
			requestHistory := history.History{
				IsEdit: false,
			}
			if h.channels.IsForward {
				requestHistory.IsEdit = true
			}
			bstr, req.Body = sepIO(req.Body)
			requestHistory.SetIdentifier(history.GetSha1(req.URL.String()))
			go requestHistory.MemoryRequest(req, false, bstr)
			resH = append(resH, requestHistory)
			fmt.Println(req.URL.String())
			if h.channels.IsForward {
				reqchan.HMgToHsSignal <- req
				req = <-reqchan.HMgToHsSignal
			}
			reqchan.ProxToHMgSignal <- req
		case res := <-reschan.ProxToHMgSignal:

			bstr, res.Body = sepIO(res.Body)
			reschan.ProxToHMgSignal <- res
			go resH[0].MemoryResponse(res, bstr)
			if resH[0].IsEdit {
				wg := &sync.WaitGroup{}
				go func() {
					isEdited := false
					wg.Add(2)
					go func() {
						defer wg.Done()
						headers := []history.RequestHeader{}
						db.Table = history.RequestHeader{}
						openDB := db.OpenDatabase()
						openDB.Where("Identifier = ?", resH[0].Identifier).Find(&headers)
						for _, header := range headers {
							if strings.Join(res.Request.Header[header.Name], ",") != header.Value {
								isEdited = true
								break
							}
						}
					}()
					go func() {
						defer wg.Done()
						var convstr []string
						datas := []history.RequestData{}
						db.Table = history.RequestData{}
						openDB := db.OpenDatabase()
						openDB.Where("Identifier = ?", resH[0].Identifier).Find(&datas)
						bstr, res.Request.Body = SeparationOfIOReadCloser(res.Request.Body)
						for _, data := range datas {
							convstr = append(convstr, data.Name+"="+data.Value)
						}
						if bstr != strings.Join(convstr, "&") {
							isEdited = true
						}
					}()
					wg.Wait()
					fmt.Println(isEdited)
				}()
			}
			if len(resH[1:]) > 0 {
				resH = resH[1:]
			}
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

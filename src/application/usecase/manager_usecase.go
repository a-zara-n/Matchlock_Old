package usecase

import (
	"net/http"
	"time"

	"github.com/a-zara-n/Matchlock/src/domain/repository"

	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//ManagerUsecase is Manage various centers in the software
type ManagerUsecase interface {
	InternalCommunication()
}

type managerUsecase struct {
	channel        *entity.Channel
	MemoryRequest  repository.RequestRepositry
	MemoryResponse repository.ResponseRepositry
}

//NewManagerUsecase はManagerUsecaseを返します
func NewManagerUsecase(channel *entity.Channel, memreq repository.RequestRepositry, memres repository.ResponseRepositry) ManagerUsecase {
	return &managerUsecase{channel, memreq, memres}
}
func (m *managerUsecase) InternalCommunication() {
	for {
		select {
		case req := <-m.channel.Request.ProxToHMgSignal:
			httppair := &aggregate.HTTPPair{}
			httppair.Request = aggregate.NewHTTPRequestByRequest(req)
			httppair.EditRequest = httppair.Request
			httppair.SetIdentifi(req.URL.String())
			go m.MemoryRequest.Insert(httppair.GetIdentifi(), false, httppair.Request)
			if m.channel.IsForward {
				m.channel.Request.HMgToHsSignal <- req
				req = <-m.channel.Request.HMgToHsSignal
				httppair.EditRequest = aggregate.NewHTTPRequestByRequest(req)
				req.ContentLength = int64(len(httppair.EditRequest.Data.FetchData()))
			}
			client := &http.Client{Timeout: time.Duration(10) * time.Second}
			req.RequestURI = ""
			resp, _ := client.Do(req)
			m.channel.Response.ProxToHMgSignal <- resp

			//保存methodを追加
			if m.channel.IsForward && httppair.IsEdited() {
				m.MemoryRequest.Insert(httppair.GetIdentifi(), true, httppair.Request)
			}
		}
	}
}

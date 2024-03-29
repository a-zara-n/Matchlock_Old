package usecase

import (
	"log"
	"net/http"
	"time"

	"github.com/a-zara-n/Matchlock/src/domain/value"

	"github.com/a-zara-n/Matchlock/src/config"
	"github.com/a-zara-n/Matchlock/src/domain/repository"

	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
)

//ManagerUsecase is Manage various centers in the software
type ManagerUsecase interface {
	InternalCommunication()
}

type managerUsecase struct {
	channel        config.Channel
	MemoryRequest  repository.RequestRepositry
	MemoryResponse repository.ResponseRepositry
	MemoryHistory  repository.HistoryRepository
	flag           *value.Forward
}

//NewManagerUsecase はManagerUsecaseを返します
func NewManagerUsecase(channel config.Channel, repuestrepo repository.RequestRepositry, responserepo repository.ResponseRepositry, history repository.HistoryRepository, forward *value.Forward) ManagerUsecase {

	return &managerUsecase{channel, repuestrepo, responserepo, history, forward}
}
func (m *managerUsecase) InternalCommunication() {
	for {
		select {
		case req := <-m.channel.Proxy.Request:
			log.Println("リクエストを受信しました")
			httpmessage := aggregate.NewHTTPMessage()
			httpmessage.SetRequest(req)
			httpmessage.SetEditedRequest(req)
			log.Println("リクエストを保存を開始します")
			go m.MemoryRequest.Insert(httpmessage.Get(), false, httpmessage.Request)
			go m.MemoryHistory.Insert(httpmessage.Get(), false)
			if m.flag.Get() {
				log.Println("forwardがONです")
				m.channel.Server.Request <- httpmessage.Request
				httpmessage.EditRequest.SetHTTPRequestByString(<-m.channel.Server.Response)
			}
			log.Println("クライアントを準備しています")
			client := &http.Client{Timeout: time.Duration(10) * time.Second}
			req.RequestURI = httpmessage.EditRequest.Info.URL.RequestURI()
			resp, _ := client.Do(httpmessage.EditRequest.GetHTTPRequestByRequest())

			httpmessage.SetResponse(resp)

			m.channel.Proxy.Response <- resp

			//保存methodを追加
			go m.MemoryResponse.Insert(httpmessage.Get(), httpmessage.Response)
			if m.flag.Get() && httpmessage.IsEdited() {
				go m.MemoryHistory.Update(httpmessage.Get(), true)
				go m.MemoryRequest.Insert(httpmessage.Get(), true, httpmessage.EditRequest)
			}
		}
	}
}

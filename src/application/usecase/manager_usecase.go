package usecase

import (
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
	channel        *config.Channel
	MemoryRequest  repository.RequestRepositry
	MemoryResponse repository.ResponseRepositry
	flag           *value.Forward
}

//NewManagerUsecase はManagerUsecaseを返します
func NewManagerUsecase(channel *config.Channel, repuestrepo repository.RequestRepositry, responserepo repository.ResponseRepositry, forward *value.Forward) ManagerUsecase {
	return &managerUsecase{channel, repuestrepo, responserepo, forward}
}
func (m *managerUsecase) InternalCommunication() {
	for {
		select {
		case req := <-m.channel.Proxy.Request:

			if m.flag.Get() {
				m.channel.Server.Request <- &aggregate.Request{}
				<-m.channel.Server.Response
			}
			client := &http.Client{Timeout: time.Duration(10) * time.Second}
			req.RequestURI = ""
			resp, _ := client.Do(req)
			m.channel.Proxy.Response <- resp

			//保存methodを追加
			if m.flag.Get() {

			}
		}
	}
}

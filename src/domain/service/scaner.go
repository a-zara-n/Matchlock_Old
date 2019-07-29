package service

import (
	"fmt"

	"github.com/a-zara-n/Matchlock/src/domain/entity"

	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/a-zara-n/Matchlock/src/domain/value"
)

//Inspection is Scanner Type name
const Inspection = "inspection"

//ScannerInterface はスキャン機能のインターフェースを定義しています
type ScannerInterface interface {
	Run(tys string)
	Listup(host, tys string)
}

//Scanner は必要な情報を保持します
type Scanner struct {
	Targets []*aggregate.Request
	Payload value.Payload
	request repository.RequestRepositry
}

//NewScanner はScannerを定義します
func NewScanner(req repository.RequestRepositry) ScannerInterface {
	return &Scanner{request: req}
}

func (scan *Scanner) Listup(host, tys string) {
	scan.Targets = scan.request.FetchHostRequests(host)
}

func (scan *Scanner) Run(tys string) {
	//仮置き
	switch tys {
	case "all":
		for _, target := range scan.Targets {
			data := target.Data
			fmt.Println("===ALL===")
			for _, key := range scan.Payload.GetTypeKeys("inspection") {
				for _, name := range scan.Payload.GetFileName(key) {
					scan.Payload.SetInfo(key, name)
					scan.AllChange(data.Type, data.GetKeys(), data.Data, scan.Payload)
				}
			}
		}
	case "simple":
		for _, target := range scan.Targets {
			data := target.Data
			fmt.Println("===Simple===")
			for _, key := range scan.Payload.GetTypeKeys("inspection") {
				for _, name := range scan.Payload.GetFileName(key) {
					scan.Payload.SetInfo(key, name)
					scan.SimpleList(data.Type, data.GetKeys(), data.Data, scan.Payload)
				}
			}
		}
	case "cluster":
		for _, target := range scan.Targets {
			data := target.Data
			fmt.Println("===Cluster===")
			for _, key := range scan.Payload.GetTypeKeys("inspection") {
				for _, name := range scan.Payload.GetFileName(key) {
					scan.Payload.SetInfo(key, name)
					scan.Cluster(data.Type, data.GetKeys(), data.Data, scan.Payload)
				}
			}
		}
	}
}

func (scan *Scanner) AllChange(style string, name []string, defaultV map[string]interface{}, payloads value.Payload) {
	data := entity.Data{
		Type: style,
		Data: defaultV,
	}
	for _, d := range payloads.GetPayload() {
		for _, nm := range name {
			data.Data[nm] = d
		}
		fmt.Println(data.FetchData())
	}
}

func (scan *Scanner) SimpleList(style string, name []string, defaultV map[string]interface{}, payloads value.Payload) {
	data := entity.Data{
		Type: style,
		Data: defaultV,
	}
	for _, nm := range name {
		tmp := data.Data[nm]
		for _, d := range payloads.GetPayload() {
			data.Data[nm] = d
			fmt.Println(data.FetchData())
		}
		data.Data[nm] = tmp
	}
}

func (scan *Scanner) Cluster(style string, name []string, defaultV map[string]interface{}, payloads value.Payload) {
	data := entity.Data{
		Type: style,
		Data: defaultV,
	}
	var recursive func(length int, i int, m entity.Data)
	recursive = func(length int, i int, m entity.Data) {
		//a.Request.Close = true
		if length > i {
			for _, p := range payloads.GetPayload() {
				m.Data[name[i]] = p
				recursive(length, i+1, m)
			}
		} else {
			fmt.Println(data.FetchData())
		}
	}
	recursive(len(name), 0, data)
}

func (scan *Scanner) scanClientRun(submitquery entity.Data, payloadData value.Payload) {

}

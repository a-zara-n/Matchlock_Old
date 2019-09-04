package service

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"

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
	client  http.Client
}

//NewScanner はScannerを定義します
func NewScanner(req repository.RequestRepositry) ScannerInterface {
	jar, _ := cookiejar.New(nil)
	c := http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return &Scanner{
		request: req,
		client:  c,
	}
}

func (scan *Scanner) Listup(host, tys string) {
	scan.Targets = scan.request.FetchHostRequests(host)
}

/*
Run はScannerのスキャン機能を走らせる際に利用するmethodです。
引数は scan typeのstritng型を引数に与えることで動作します。

scan types
| name  | discribe
|:-----:|:--------:
|all    |パラメータを同時で変更する
|simple |パラメータを順次変更する。変更をしない箇所はデフォルトの値にする
|cluster|パラメータとpayloadの組み合わせを全て試す
*/
func (scan *Scanner) Run(tys string) {
	//仮置き
	var modefunc func(target *aggregate.Request, name []string, defaultV map[string]interface{}, payloads value.Payload)
	switch tys {
	case "all":
		modefunc = scan.AllChange
	case "simple":
		modefunc = scan.SimpleList
	case "cluster":
		modefunc = scan.Cluster
	default:
		return
	}
	for _, target := range scan.Targets {
		data := target.Data
		for _, key := range scan.Payload.GetTypeKeys("inspection") {
			for _, name := range scan.Payload.GetFileName(key) {
				scan.Payload.SetInfo(key, name)
				modefunc(target, data.GetKeys(), data.Data, scan.Payload)
			}
		}
	}
}

//scan types methods
/*
この区分に属するmethodの引数は統一されるべきである
引数について
target :　スキャン対象を定義したaggrigate request
names : Data(QueryString)に含まれるnameを列挙した配列
defaultV : 現状のデフォルトvalueを渡す
Payload : 名前の通り取得したpayload valuesをentityで渡す
*/
//All はscan types all の動作を定義した関数
func (scan *Scanner) AllChange(target *aggregate.Request, names []string, defaultV map[string]interface{}, payloads value.Payload) {
	data := entity.Data{
		Type: target.Data.Type,
		Data: defaultV,
	}
	for _, d := range payloads.GetPayload() {
		for _, nm := range names {
			data.Data[nm] = d
			target.Data = &data
			//scan.clientRun(target)
		}
	}
}

//SimpleList はAllと同じくscan types のsimplelistを定義した関数
func (scan *Scanner) SimpleList(target *aggregate.Request, names []string, defaultV map[string]interface{}, payloads value.Payload) {
	data := entity.Data{
		Type: target.Data.Type,
		Data: defaultV,
	}
	for _, nm := range names {
		tmp := data.Data[nm]
		for _, d := range payloads.GetPayload() {
			data.Data[nm] = d
			target.Data = &data
			//scan.clientRun(target)
		}
		data.Data[nm] = tmp
	}
}

//Cluster は
func (scan *Scanner) Cluster(target *aggregate.Request, names []string, defaultV map[string]interface{}, payloads value.Payload) {
	data := entity.Data{
		Type: target.Data.Type,
		Data: defaultV,
	}
	var recursive func(length int, i int, m entity.Data)
	recursive = func(length int, i int, m entity.Data) {
		//a.Request.Close = true
		if length > i {
			for _, p := range payloads.GetPayload() {
				m.Data[names[i]] = p
				recursive(length, i+1, m)
			}
		} else {
			target.Data = &data
			//scan.clientRun(target)
		}
	}
	recursive(len(names), 0, data)
}

func (scan *Scanner) clientRun(target *aggregate.Request) {
	res, err := scan.client.Do(target.GetHTTPRequestByRequest())
	if err != nil {
		log.Println(err)
		return
	}
	response := aggregate.NewHTTPResponseByResponse(res)
	fmt.Println(target.Info.URL.String(), response.Info.Status)
}

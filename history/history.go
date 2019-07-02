package history

import (
	"net/http"

	"github.com/a-zara-n/Matchlock/datastore"
	"github.com/jinzhu/gorm"
)

//History はhttp.Request/Responseの送受信履歴を保存するためのもの
type History struct {
	gorm.Model
	Identifier string
	IsEdit     bool
}

//SetIdentifier はIdentifierをHistoryにセットする
func (h *History) SetIdentifier(id string) {
	h.Identifier = id
}

//MemoryRequest はリクエストの保存をする
//リクエスのと保存は
//1. SetRequestでリクエストのheaderとbodyを除くものを保存する
//2. SetHeaderでリクエストのheaderを保存する
//3. SetDataでリクエストのdataを保存する
func (h *History) MemoryRequest(r *http.Request, isEdit bool, bstr string) {
	if bstr == "" {
		bstr = r.URL.RawQuery
	}
	var req = Request{Identifier: h.Identifier, IsEdit: isEdit}
	go req.SetRequest(r)
	go req.SetHeader(r.Header)
	go req.SetData(bstr, r.ContentLength, r.TransferEncoding)
	datastore.DB.Insert(h)
}

//MemoryResponse はレスポンスの保存をする
//レスポンスのと保存は
//1. SetResponseでレスポンスのheaderとbodyを除くものを保存する
//2. SetResponseHeaderでレスポンスのheaderを保存する
//3. SetResponseBodyでレスポンスのdataを保存する
func (h *History) MemoryResponse(r *http.Response, bstr string) {
	var res = Response{Identifier: h.Identifier}
	go res.SetResponse(r)
	go res.SetResponseHeader(r.Header)
	go res.SetResponseBody(bstr, r.ContentLength, r.TransferEncoding)
}

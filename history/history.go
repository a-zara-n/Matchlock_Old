package history

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

type History struct {
	gorm.Model
	Identifier string
}

func (h *History) SetIdentifier(id string) {
	h.Identifier = id
}

func (h *History) MemoryRequest(r *http.Request, isEdit bool, bstr string) {
	fmt.Println(bstr)
	var req = Request{Identifier: h.Identifier, IsEdit: isEdit}
	go req.SetRequest(r, bstr)
	go req.SetHeader(r.Header)
	go req.SetData(bstr, r.ContentLength, r.TransferEncoding)
	db.Table = History{}
	db.Insert(h)
}

func (h *History) MemoryResponse(r *http.Response, bstr string) {
	var res = Response{Identifier: h.Identifier}
	go res.SetResponse(r)
	go res.SetResponseHeader(r.Header)
	go res.SetResponseBody(bstr, r.ContentLength, r.TransferEncoding)
}

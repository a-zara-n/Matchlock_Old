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
	req.SetRequest(r, bstr)
	db.Table = History{}
	db.Insert(h)
}

func (h *History) MemoryResponse(r *http.Response, bstr string) {
	var res = Response{Identifier: h.Identifier}
	res.SetResponse(r)
}

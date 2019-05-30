package history

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type Response struct {
	gorm.Model
	Identifier string
}

func (r *Response) SetResponse(res *http.Response) {

}
func (r *Response) GetResponse() {

}

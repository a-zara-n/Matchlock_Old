package aggregate

import (
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//Response は
type Response struct {
	Info   entity.ResponseInfo
	Header entity.HTTPHeader
	Body   entity.Body
}

package aggregate

import "github.com/a-zara-n/Matchlock/src/domain/value"

//HTTPPair はHTTPリクエストとレスポンスをまとめたものになります
type HTTPPair struct {
	value.Identifier
	IsEdit      bool
	Request     Request
	EditRequest Request
	Response    Response
}

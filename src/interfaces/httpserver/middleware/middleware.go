package middleware

import (
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//Middleware はMiddlewareの動作を定義したものになります
type Middleware interface {
}
type middleware struct {
	channel *entity.Channel
}

// NewMiddleware はMiddlewareを新規で作成します
func NewMiddleware(c *entity.Channel) Middleware {
	return middleware{c}
}

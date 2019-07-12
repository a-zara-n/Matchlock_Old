package middleware

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
)

//Middleware はMiddlewareの動作を定義したものになります
type Middleware interface {
	Run()
}
type middleware struct {
	usecase.ManagerUsecase
}

// NewMiddleware はMiddlewareを新規で作成します
func NewMiddleware(m usecase.ManagerUsecase) Middleware {
	return &middleware{m}
}
func (m *middleware) Run() {
	m.InternalCommunication()
}

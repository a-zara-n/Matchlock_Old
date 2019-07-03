package htmlhandler

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase/html"
)

type HTMLHandler interface {
}

type htmlHandler struct {
	u html.HTMLUseCase
}

func NewWarmupHandler(h html.HTMLUseCase) HTMLHandler {
	return &htmlHandler{h}
}

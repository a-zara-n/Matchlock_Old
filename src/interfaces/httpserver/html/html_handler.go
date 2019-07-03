package htmlhandler

import (
	"context"
	"net/http"
)

type HTMLHandler interface {
	Warmup(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

type htmlHandler struct {
	htmlHandler htmlusecase.HTMLUseCase
}

func NewWarmupHandler(h htmlusecase.HTMLUseCase) HTMLHandler {
	return &htmlHandler{h}
}

func (i *IndexHandler) Warmup(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// some code
}

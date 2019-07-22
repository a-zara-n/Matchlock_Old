package api

import (
	"github.com/a-zara-n/Matchlock/src/domain/repository"
)

//HistoryInterface は履歴取得のために必要なmethodを定義したインターフェースです
type HistoryInterface interface {
	FetchAll() interface{}
}

//History は履歴を取得するための各種情報を保持するための構造体です
type History struct {
	repository.HistoryRepository
}

//NewHistory は新規にHistoryinterfaceを提供します
func NewHistory(hh repository.HistoryRepository) HistoryInterface {
	return &History{hh}
}

func (h *History) FetchAll() interface{} {
	return h.Fetch(1)
}

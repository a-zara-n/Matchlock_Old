package api

import (
	"github.com/a-zara-n/Matchlock/src/domain/repository"
)

//MessageInterface はメッセージを取得するmethodを定義します
type MessageInterface interface {
	FetchMessage(identifier string) interface{}
}

//Message は必要な情報を保持します
type Message struct {
	repository.HTTPMessageRepository
}

//NewMessage は
func NewMessage(hm repository.HTTPMessageRepository) MessageInterface {
	return &Message{hm}
}

func (m *Message) FetchMessage(identifier string) interface{} {
	return m.Fetch(identifier)
}

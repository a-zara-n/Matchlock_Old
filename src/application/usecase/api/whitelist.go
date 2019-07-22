package api

import (
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//WhiteListInterface はプロキシを通すドメインのはんだんを行うentityの操作を行います
type WhiteListInterface interface {
	FetchWhiteList(i int) []string
	AddWhiteList(regexstring string) bool
	UpdateWhiteList(i int, regexstring string) bool
	DelWhiteList(i int) bool
}

//WhiteList は情報を保持します
type WhiteList struct {
	list *entity.WhiteList
}

//NewWhiteList は
func NewWhiteList(w *entity.WhiteList) WhiteListInterface {
	return &WhiteList{w}
}

func (w *WhiteList) AddWhiteList(regexstring string) bool {
	w.list.Add(regexstring)
	return true
}

func (w *WhiteList) DelWhiteList(i int) bool {
	w.list.Del(i)
	return true
}

func (w *WhiteList) UpdateWhiteList(i int, regexstring string) bool {
	w.list.Set(i, regexstring)
	return true
}
func (w *WhiteList) FetchWhiteList(i int) []string {
	return w.list.Fetch(i)
}

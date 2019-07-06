package entity

import "regexp"

//WhiteList はHostのWhiteListを管理するmodel
type WhiteList struct {
	List []*regexp.Regexp
}

//Add はWhiteListの追加をするための関数です
func (w *WhiteList) Add(reg string) {
	w.List = append(w.List, regexp.MustCompile(reg))
}

//Del はWhiteListの削除をするための関数です
func (w *WhiteList) Del(i int) {
	if i > len(w.List) {
		return
	}
	w.List = append(w.List[:i-1], w.List[i:]...)
}

//Check はhostがWhiteListに含まれているかを確認する
func (w *WhiteList) Check(host string) bool {
	for _, r := range w.List {
		if r.MatchString(host) {
			return true
		}
	}
	return false
}

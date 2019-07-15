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
func (w *WhiteList) Del(i int) bool {
	if i > len(w.List) {
		return false
	}
	if len(w.List) < 2 {
		w.List = []*regexp.Regexp{}
		return true
	}
	switch i {
	case 0:
		w.List = w.List[i+1:]
	case len(w.List) - 1:
		w.List = w.List[:len(w.List)-2]
	default:
		w.List = append(w.List[:i], w.List[i+1:]...)
	}
	return true
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

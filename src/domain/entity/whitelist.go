package entity

import "regexp"

//WhiteList はHostのWhiteListを管理するmodel
type WhiteList struct {
	Strings []string
	List    []*regexp.Regexp
}

//Fetch は引数から後ろのregex stringを返します
func (w *WhiteList) Fetch(i int) []string {
	return w.Strings[i:]
}

//Set はregexを更新します
func (w *WhiteList) Set(key int, value string) {
	w.Strings[key] = value
	w.List[key] = regexp.MustCompile(value)
}

//Add はWhiteListの追加をするための関数です
func (w *WhiteList) Add(reg string) {
	w.Strings = append(w.Strings, reg)
	w.List = append(w.List, regexp.MustCompile(reg))
}

//Del はWhiteListの削除をするための関数です
func (w *WhiteList) Del(i int) bool {
	if i > len(w.List) {
		return false
	}
	if len(w.List) < 2 {
		w.Strings = []string{}
		w.List = []*regexp.Regexp{}
		return true
	}
	switch i {
	case 0:
		w.Strings = w.Strings[i+1:]
		w.List = w.List[i+1:]
	case len(w.List) - 1:
		w.Strings = w.Strings[:len(w.Strings)-2]
		w.List = w.List[:len(w.List)-2]
	default:
		w.Strings = append(w.Strings[:i], w.Strings[i+1:]...)
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

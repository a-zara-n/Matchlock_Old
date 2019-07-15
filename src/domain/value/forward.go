package value

//Forward はProxyの受け止め可否の情報を保持します
type Forward struct {
	flag bool
}

//Get は可否情報をboolで返します
func (f *Forward) Get() bool {
	return f.flag
}

//Set は受け止めの可否情報を設定します
func (f *Forward) Set(t bool) {
	f.flag = t
}

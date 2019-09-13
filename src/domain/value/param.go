package value

//Param はURLやQueryのParameterを保存するための構造体です。
//連結リストで
type Param struct {
	Name         string
	Type         string
	Format       string
	Param        string
	Data         string
	DefaultParam string
	Required     bool
	thisPath     string
	childs       []*Param
	parents      *Param
}

//GetKeys はParamのkeysを取得するための関数です
//childeの有る限り取得を行います
func (para *Param) GetKeys(parents ...string) []string {
	var path = ""
	if len(parents) >= 0 {
		path = parents[0]
	}
	if len(para.childs) <= 0 {
		if para.thisPath != "" {
			return []string{para.thisPath}
		}
		para.thisPath = path + para.Name
		return []string{para.thisPath}
	}
	var ret = []string{}
	for _, child := range para.childs {
		ret = append(ret, child.GetKeys(path+para.Name+"/")...)
	}
	return ret
}

//Fetch はParamのtypeとformatに沿ってデータを取得できる関数です
func (para *Param) Fetch() interface{} {
	var ret func() interface{}
	switch para.Type {
	case "object":
	case "array":
	default:
		ret = para.GetParam
	}
	return ret()
}

//GetParam はdefaultで設定されたParameterを設定して取得します。
//もし設定されていない場合はそれに近しい値を返します。
func (para *Param) GetParam() interface{} {
	var ret interface{}
	if para.Data != "" {
		return para.Data
	}
	if para.DefaultParam != "" {
		return para.DefaultParam
	}
	switch para.Type {
	case "integer":
		ret = 1
	case "string":
		ret = ""
	case "boolean":
		ret = true
	}
	return ret
}

//GetArray

package attacker

import "github.com/a-zara-n/Matchlock/shared"

type ParamData struct {
	Name     string
	TypeOf   string
	Type     string
	DefaultV string
}

/*
 Attack はattacker.goの関数を動かす仮の関数
*/

func setParamData(pdata []ParamData) ([][]string, []string, map[string]string) {
	var body [][]string
	body = [][]string{{pdata[0].Name, "{{." + pdata[0].Name + "}}"}}
	name, defvlue :=
		[]string{pdata[0].Name},
		map[string]string{pdata[0].Name: pdata[0].DefaultV}
	if len(pdata) > 1 {
		bodys, names, defvalues := setParamData(pdata[1:])
		return append(body, bodys...), append(name, names...), shared.Merge(defvlue, defvalues)
	}
	return body, name, defvlue
}

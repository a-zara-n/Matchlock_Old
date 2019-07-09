package entity

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

//Data はHTTPのデータを管理するためのentityです
type Data struct {
	Keys []string
	Type string
	Data map[string]interface{}
}

//GetKeys はDataのname一覧を取得する
func (d *Data) GetKeys() []string {
	return d.Keys
}

//AddData は指定した値を設定します
func (d *Data) AddData(key string, value interface{}) {
	d.Data[key] = value
	d.Keys = append(d.Keys, key)
}

//RemoveData は指定した値を設定します
func (d *Data) RemoveData(key string) {
	delete(d.Data, key)
	d.Keys = getKeys(d.Data)
}

//FetchData はTypeに合った形式でデータを出力します
func (d *Data) FetchData() string {
	var retdata string
	switch d.Type {
	case "JSON":
		j, _ := json.Marshal(d.Data)
		retdata = string(j)
	case "FORM":
		var tmp []string
		for name, value := range d.Data {
			tmp = append(tmp, name+"="+value.(string))
		}
		sort.Strings(tmp)
		retdata = strings.Join(tmp, "&")
	}
	return retdata
}

//SetDataByHTTPBody はDataエンティティにDataを設定するためのmethod
func (d *Data) SetDataByHTTPBody(body io.ReadCloser) io.ReadCloser {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(body)
	data := bufbody.String()
	d.SetData(data)
	return ioutil.NopCloser(strings.NewReader(data))
}

//SetData はDataエンティティにDataを設定するためのmethod
func (d *Data) SetData(rawdata string) {
	d.Data, d.Type = checkDataType(rawdata)
	d.Keys = getKeys(d.Data)
}

//JSONであるかの検査時に利用する
var typeJSONRegexp = regexp.MustCompile(`^{(\".*\":\"?.*\"?,?)+[^,]}$`)

func checkDataType(rawdata string) (map[string]interface{}, string) {
	if typeJSONRegexp.MatchString(rawdata) {
		return parseJSON(rawdata), "JSON"
	}
	return parseFORM(rawdata), "FORM"
}

func parseJSON(rawdata string) map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(rawdata), &data); err != nil {
		log.Println("not JSON schema")
		return map[string]interface{}{}
	}

	return data
}

func parseFORM(rawdata string) map[string]interface{} {
	var retdata = map[string]interface{}{}
	splitdata := strings.Split(rawdata, "&")
	for _, data := range splitdata {
		v := strings.Split(data, "=")
		retdata[v[0]] = v[1]
	}
	return retdata
}

func getKeys(maps map[string]interface{}) []string {
	ret := []string{}
	for key := range maps {
		ret = append(ret, key)
	}
	sort.Strings(ret)
	return ret
}

package entity

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

var (
	testingDataEntity = &Data{Data: map[string]interface{}{}}
	testData          = map[string]interface{}{
		"FORM": map[string]interface{}{
			"Raw":    `name=hoge&age=20&like=mikan`,
			"Keys":   []string{"name", "age", "like"},
			"Result": map[string]interface{}{"name": "hoge", "age": "20", "like": "mikan"},
			"Fetch":  `age=20&like=mikan&name=hoge`,
		},
		"JSON_Success": map[string]interface{}{
			"Raw":    `{"age":20,"lang":["jp","en","fr"],"like":"mikan","name":"hoge"}`,
			"Keys":   []string{"name", "age", "like", "lang"},
			"Result": map[string]interface{}{"name": "hoge", "age": float64(20), "like": "mikan", "lang": []interface{}{"jp", "en", "fr"}},
			"Fetch":  `{"age":20,"lang":["jp","en","fr"],"like":"mikan","name":"hoge"}`,
		},
		"JSON_Fail": map[string]interface{}{
			"Raw":    `{"name":hoge}`,
			"Keys":   []string{},
			"Result": map[string]interface{}{},
			"Fetch":  `{}`,
		},
	}
)

func TestGetKeysOfDataEntity(t *testing.T) {
	NothingIsIncluded := testingDataEntity.GetKeys()
	if len(NothingIsIncluded) > 0 {
		t.Error("Keyが含まれています。")
	}
	testingDataEntity.Keys = []string{"hoge", "fuga", "piyo"}
	if !reflect.DeepEqual(testingDataEntity.GetKeys(), []string{"hoge", "fuga", "piyo"}) {
		t.Error("正しく返却されていません")
	}
	testingDataEntity.Keys = []string{}
}

func TestAddDataAndRemoveDataOfDataEntity(t *testing.T) {
	testingDataEntity.AddData("name", "hoge")
	if testingDataEntity.Data["name"] != "hoge" {
		t.Error("正しい値が格納されていません")
	}
	testingDataEntity.RemoveData("name")
	_, ok := testingDataEntity.Data["name"]
	if ok {
		t.Error("正しく削除されていません")
	}
}

func TestSetDataAndFetchDataOfDataEntity(t *testing.T) {
	for casestring, data := range testData {
		trance := data.(map[string]interface{})
		testingDataEntity.SetData(trance["Raw"].(string))
		res := trance["Result"].(map[string]interface{})
		for n, v := range res {
			if !reflect.DeepEqual(testingDataEntity.Data[n], v) {
				t.Errorf("%v :値が異なっています", casestring)
			}
		}
		if testingDataEntity.FetchData() != trance["Fetch"] {
			t.Errorf("%v :出力が異なります %v != %v", casestring, testingDataEntity.FetchData(), trance["Fetch"])
		}
	}
}

//Test SetDataByHTTPBody And FetchData Of DataEntity
func TestSetDataByHTTPBodyAndFetchDataOfDataEntity(t *testing.T) {
	for casestring, data := range testData {
		trance := data.(map[string]interface{})
		testingDataEntity.SetDataByHTTPBody(ioutil.NopCloser(strings.NewReader(trance["Raw"].(string))))
		res := trance["Result"].(map[string]interface{})
		for n, v := range res {
			if !reflect.DeepEqual(testingDataEntity.Data[n], v) {
				t.Errorf("%v :値が異なっています", casestring)
			}
		}
		if testingDataEntity.FetchData() != trance["Fetch"] {
			t.Errorf("%v :出力が異なります %v != %v", casestring, testingDataEntity.FetchData(), trance["Fetch"])
		}
	}
}

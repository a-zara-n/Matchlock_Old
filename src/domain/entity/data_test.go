package entity

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/a-zara-n/Matchlock/src/testdata"
)

var testingDataEntity = &Data{Data: map[string]interface{}{}}

func TestGetKeysOfDataEntity(t *testing.T) {
	NothingIsIncluded := testingDataEntity.GetKeys()
	if len(NothingIsIncluded) > 0 {
		t.Error("Keyが含まれています。")
	}
	testingDataEntity.Keys = test.FetchTestRequest(0).Query.Keys
	if !reflect.DeepEqual(testingDataEntity.GetKeys(), test.FetchTestRequest(0).Query.Keys) {
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
	for casestring, data := range testdata.TestQuery {
		testingDataEntity.SetData(data.Raw)
		result := data.Result
		for key, v := range result {
			if !reflect.DeepEqual(testingDataEntity.Data[key], v) {
				t.Errorf("%v :値が異なっています", casestring)
			}
		}
		if testingDataEntity.FetchData() != data.Fetch {
			t.Errorf("%v :出力が異なります %v != %v", casestring, testingDataEntity.FetchData(), test.FetchTestRequest(0).Query.Fetch)
		}
	}
}

//Test SetDataByHTTPBody And FetchData Of DataEntity
func TestSetDataByHTTPBodyAndFetchDataOfDataEntity(t *testing.T) {
	for casestring, data := range testdata.TestQuery {
		testingDataEntity.SetDataByHTTPBody(ioutil.NopCloser(strings.NewReader(data.Raw)))
		res := data.Result
		for n, v := range res {
			if !reflect.DeepEqual(testingDataEntity.Data[n], v) {
				t.Errorf("%v :値が異なっています", casestring)
			}
		}
		if testingDataEntity.FetchData() != data.Fetch {
			t.Errorf("%v :出力が異なります %v != %v", casestring, testingDataEntity.FetchData(), data.Fetch)
		}
	}
}

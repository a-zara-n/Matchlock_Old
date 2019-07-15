package entity

import "testing"

var testrequestinfo = &ResponseInfo{}

func TestResponseInfo(t *testing.T) {
	for i := 0; i < test.GetResponseCount(); i++ {
		testresp := test.FetchTestResponse(i).HTTP
		testrequestinfo.Set(testresp)
		if testrequestinfo.Proto != testresp.Proto || testrequestinfo.Status != testresp.Status {
			t.Error("適正な代入が行われていません")
		}
	}
}

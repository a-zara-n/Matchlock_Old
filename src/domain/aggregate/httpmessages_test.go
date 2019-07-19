package aggregate

import (
	"testing"

	"github.com/a-zara-n/Matchlock/src/domain/value"
)

func TestHTTPMessageSuccess(t *testing.T) {
	testhttpmessage := HTTPMessages{
		Identifier:  value.Identifier{},
		IsEdit:      false,
		Request:     &Request{},
		EditRequest: &Request{},
		Response:    &Response{},
	}
	testhttpmessage.SetRequest(test.FetchTestRequest(0).HTTP)
	testhttpmessage.SetEditedRequest(test.FetchTestRequest(0).HTTP)
	testhttpmessage.SetResponse(test.FetchTestResponse(0).HTTP)
	if testhttpmessage.IsEdited() {
		t.Error("正確な判定がなされていません")
	}
}

func TestHTTPMessageFail(t *testing.T) {
	testhttpmessage := HTTPMessages{
		Identifier:  value.Identifier{},
		IsEdit:      false,
		Request:     &Request{},
		EditRequest: &Request{},
		Response:    &Response{},
	}
	testhttpmessage.SetRequest(test.FetchTestRequest(0).HTTP)
	testhttpmessage.SetEditedRequest(test.FetchTestRequest(1).HTTP)
	testhttpmessage.SetResponse(test.FetchTestResponse(0).HTTP)
	if !testhttpmessage.IsEdited() {
		t.Error("正確な判定がなされていません")
	}
}

func TestHTTPMessageIsIdentifireSuccess(t *testing.T) {
	testhttpmessage := HTTPMessages{
		Identifier:  value.Identifier{},
		IsEdit:      false,
		Request:     &Request{},
		EditRequest: &Request{},
		Response:    &Response{},
	}
	testhttpmessage.SetRequest(test.FetchTestRequest(0).HTTP)
	if testhttpmessage.Get() == "" {
		t.Error("識別子の設定が行われていません")
	}
}

func TestHTTPMessageIsIdentifireFail(t *testing.T) {
	testhttpmessage := HTTPMessages{
		Identifier:  value.Identifier{},
		IsEdit:      false,
		Request:     &Request{},
		EditRequest: &Request{},
		Response:    &Response{},
	}
	if testhttpmessage.Get() != "" {
		t.Error("識別子の設定が行われていません")
	}
}

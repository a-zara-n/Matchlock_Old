package attacker

import (
	"bytes"
	"html"
	"io"
	"net/http"

	"github.com/a-zara-n/Matchlock/extractor"
	"github.com/a-zara-n/Matchlock/scanner/attacker/decid"
	"github.com/a-zara-n/Matchlock/scanner/attacker/payload"
)

func (a attacker) scanClientRun(submitValues map[string]string, payloadData payload.Payload) {
	var buf bytes.Buffer
	a.paramtmplate.Execute(&buf, submitValues)
	a.setSubmitValue(html.UnescapeString(buf.String()))
	resp := a.sender()
	a.decider(resp.Body, payloadData, buf.String())
	resp.Body.Close()
}

func (a attacker) sender() *http.Response {
	resp, err := a.client.Do(a.Request)
	if err != nil {
		panic(err)
	}
	return resp
}

func (a attacker) setSubmitValue(submitValue string) {
	if a.Request.Method == "POST" {
		a.Request.Body = extractor.GetIOReadCloser(submitValue)
	} else {
		a.Request.URL.RawQuery = submitValue
	}
}

func (a attacker) decider(resp io.ReadCloser, payloadData payload.Payload, input string) {
	go decid.Decider(
		lineDiff(a.ResponseBody, extractor.GetStringBody(resp)), payloadData, *a.Request, input,
	)
}

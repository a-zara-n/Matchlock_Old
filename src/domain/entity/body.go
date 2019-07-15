package entity

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
)

//Body は
type Body struct {
	Body       string
	Encodetype string
	Length     int64
}

func (b *Body) Set(body io.ReadCloser) io.ReadCloser {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(body)
	bodystring := bufbody.String()
	if bodystring != "" {
		b.Body = bodystring
		b.Length = int64(len(bodystring))
	}
	return ioutil.NopCloser(strings.NewReader(bodystring))
}

func (b *Body) Get() io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(b.Body))
}

func (b *Body) GetLength() int64 {
	return b.Length
}

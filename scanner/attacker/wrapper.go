package attacker

import (
	"io"

	"../../extractor"
)

func GetStringBody(b io.ReadCloser) string {
	return extractor.GetStringBody(b)
}

func GetIOReadCloser(b string) io.ReadCloser {
	return extractor.GetIOReadCloser(b)
}

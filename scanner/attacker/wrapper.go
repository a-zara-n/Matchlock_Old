package attacker

import (
	"io"

	"github.com/WestEast1st/Matchlock/extractor"
)

func GetStringBody(b io.ReadCloser) string {
	return extractor.GetStringBody(b)
}

func GetIOReadCloser(b string) io.ReadCloser {
	return extractor.GetIOReadCloser(b)
}

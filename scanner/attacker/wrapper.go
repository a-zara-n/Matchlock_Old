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

func merge(m1, m2 map[string]string) map[string]string {
	ans := map[string]string{}
	for k, v := range m1 {
		ans[k] = v
	}
	for k, v := range m2 {
		ans[k] = v
	}
	return (ans)
}

package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func randomString(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), letterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		idx := int(cache & letterIdxMask)
		if idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func GetSha1(s string) string {
	s += randomString(5)
	h := sha1.New()
	h.Write([]byte(s))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return *(*string)(unsafe.Pointer(&hash))
}

package value

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

//Identifier はリクエストのIDの識別子として利用できます
type Identifier struct {
	Value string
}

var randSrc = rand.NewSource(time.Now().UnixNano())

//ランダムな文字列を生成するための要素
const (
	letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

//Get は設定されている識別子を取り出します
func (ident *Identifier) Get() string {
	return ident.Value
}

//Set は引数を元に識別子を設定します
func (ident *Identifier) Set(str string) {
	str += randomString(5)
	h := sha1.New()
	h.Write([]byte(str))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	ident.Value = *(*string)(unsafe.Pointer(&hash))

}

//randomString はSetIdentifiで利用するランダムな文字列を生成する
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

package api

import (
	"github.com/a-zara-n/Matchlock/src/domain/value"
)

//Forward は止めるかとどうか考える
type Forward interface {
	ChangeForward() (bool, error)
}

type forward struct {
	flag *value.Forward
}

func (f *forward) ChangeForward() (bool, error) {
	if f.flag.Get() {
		f.flag.Set(false)
	} else {
		f.flag.Set(true)
	}
	return f.flag.Get(), nil
}

//NewForward は
func NewForward(f *value.Forward) Forward {
	return &forward{f}
}

package api

import (
	"github.com/a-zara-n/Matchlock/src/domain/value"
)

//ForwordInterface は
type ForwordInterface interface {
	ChangeForward() (bool, error)
}

//Forword は
type Forword struct {
	flag *value.Forward
}

//NewForword は
func NewForword(f *value.Forward) ForwordInterface {
	return &Forword{f}
}

func (f *Forword) ChangeForward() (bool, error) {
	if f.flag.Get() {
		f.flag.Set(false)
	} else {
		f.flag.Set(true)
	}
	return f.flag.Get(), nil
}

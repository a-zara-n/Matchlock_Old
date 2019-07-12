package api

import (
	"net/http"

	"github.com/a-zara-n/Matchlock/src/domain/value"

	"github.com/labstack/echo"
)

//Forward は止めるかとどうか考える
type Forward interface {
	ChangeForward(c echo.Context) (bool, error)
}

type forward struct {
	flag *value.Forward
}

func (f *forward) ChangeForward(c echo.Context) (bool, error) {
	if f.flag.Get() {
		f.flag.Set(false)
	} else {
		f.flag.Set(true)
	}
	return f.flag.Get(), c.JSON(http.StatusOK, "{status: \"OK\"}")
}

//NewForward は
func NewForward(f *value.Forward) Forward {
	return &forward{f}
}

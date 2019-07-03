package apiusecase

import (
	"net/http"

	"github.com/a-zara-n/MatchlockDDD/Matchlock/src/domain/entity"

	"github.com/labstack/echo"
)

//Forward は止めるかとどうか考える
type Forward interface {
	ChangeForward(c echo.Context, forward entity.Channel) (bool, error)
}

type forward struct{}

func (f *forward) ChangeForward(c echo.Context, forward entity.Channel) (bool, error) {
	if forward.IsForward {
		forward.IsForward = false
	} else {
		forward.IsForward = true
	}
	return forward.IsForward, c.JSON(http.StatusOK, "{status: \"OK\"}")
}

//NewForward は
func NewForward() Forward {
	return &forward{}
}

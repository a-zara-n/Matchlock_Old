package registry

import (
	"github.com/a-zara-n/Matchlock/src/application/usecase"
	"github.com/a-zara-n/Matchlock/src/domain/entity"
)

//NewLogic はusecase.ProxyLogicを取得
func NewLogic(white *entity.WhiteList) usecase.ProxyLogic {
	return usecase.NewLogic(white)
}

//NewHTMLUseCase はHTMLのレンダリングを行う
func NewHTMLUseCase() usecase.HTMLUseCase {
	return usecase.NewHTMLUseCase()
}

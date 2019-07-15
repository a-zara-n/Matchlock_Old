package registry

import (
	"github.com/a-zara-n/Matchlock/src/domain/entity"
	"github.com/a-zara-n/Matchlock/src/domain/value"
)

//Entity は必要なEntityを取得します
type Entity interface {
	NewForward() *value.Forward
	NewWhiteList() *entity.WhiteList
}

//NewForward はvalue.Forwardを取得する
func (r *registry) NewForward() *value.Forward {
	return &value.Forward{}
}

//NewWhiteList はentity.WhiteListを取得
func (r *registry) NewWhiteList() *entity.WhiteList {
	return &entity.WhiteList{}
}

package registry

import "github.com/a-zara-n/Matchlock/src/domain/entity"

//Entity は必要なEntityを取得します
type Entity interface {
	NewChannel() *entity.Channel
	NewWhiteList() *entity.WhiteList
}

//NewChannel はentity.Channelを取得
func NewChannel() *entity.Channel {
	return entity.NewMatchChannel()
}

//NewWhiteList はentity.WhiteListを取得
func NewWhiteList() *entity.WhiteList {
	return &entity.WhiteList{}
}

package registry

import "github.com/a-zara-n/Matchlock/src/domain/entity"

//NewChannel はentity.Channelを取得
func NewChannel() *entity.Channel {
	return entity.NewMatchChannel()
}

//NewWhiteList はentity.WhiteListを取得
func NewWhiteList() *entity.WhiteList {
	return &entity.WhiteList{}
}

package registry

import "github.com/a-zara-n/MatchlockDDD/Matchlock/src/domain/entity"

//Registry はNewで生成されるものを定義しています
type Registry interface {
	NewChannel() *entity.Channel
}

//NewChannel を取得
func NewChannel() *entity.Channel {
	return entity.NewMatchChannel()
}

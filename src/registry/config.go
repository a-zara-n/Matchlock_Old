package registry

import "github.com/a-zara-n/Matchlock/src/config"

//Config は設定の読み込みに利用されます
type Config interface {
	NewDatabaseConfig() config.DatabaseConfig
	NewMatchlockChannel() config.Channel
}

//NewDatabaseConfig はDBのコンフィグを取得します
func (r *registry) NewDatabaseConfig() config.DatabaseConfig {
	return config.NewDatabaseConfig()
}

//NewMatchlockChannel は新規で通信用チャンネルを生成します
func (r *registry) NewMatchlockChannel() config.Channel {
	return config.NewMatchlockChannel()
}

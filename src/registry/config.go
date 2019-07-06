package registry

import "github.com/a-zara-n/Matchlock/src/config"

//Config は設定の読み込みに利用されます
type Config interface {
	NewDatabaseConfig() config.DatabaseConfig
}

//NewDatabaseConfig はDBのコンフィグを取得します
func NewDatabaseConfig() config.DatabaseConfig {
	return config.NewDatabaseConfig()
}

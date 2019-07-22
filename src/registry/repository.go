package registry

import (
	"github.com/a-zara-n/Matchlock/src/config"
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/a-zara-n/Matchlock/src/infrastructure/persistence/datastore"
)

type Repository interface {
	NewRequestRepositry(dbconfig config.DatabaseConfig) repository.RequestRepositry
	NewResponseRepositry(dbconfig config.DatabaseConfig) repository.ResponseRepositry
	NewHTTPMessageRepositry(dbconfig config.DatabaseConfig) repository.HTTPMessageRepository
	NewHistoryRepositry(dbconfig config.DatabaseConfig) repository.HistoryRepository
}

func (r *registry) NewRequestRepositry(dbconfig config.DatabaseConfig) repository.RequestRepositry {
	return datastore.NewRequest(dbconfig)
}
func (r *registry) NewResponseRepositry(dbconfig config.DatabaseConfig) repository.ResponseRepositry {
	return datastore.NewResponse(dbconfig)
}
func (r *registry) NewHTTPMessageRepositry(dbconfig config.DatabaseConfig) repository.HTTPMessageRepository {
	return datastore.NewHTTPMessage(dbconfig)
}
func (r *registry) NewHistoryRepositry(dbconfig config.DatabaseConfig) repository.HistoryRepository {
	return datastore.NewHTTPHistory(dbconfig)
}

package registry

import (
	"github.com/a-zara-n/Matchlock/src/domain/repository"
	"github.com/a-zara-n/Matchlock/src/infrastructure/persistence/datastore"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	NewRequestRepositry(db *gorm.DB) repository.RequestRepositry
	NewResponseRepositry(db *gorm.DB) repository.ResponseRepositry
}

func (r *registry) NewRequestRepositry(db *gorm.DB) repository.RequestRepositry {
	return datastore.NewRequestRepositry(db)
}
func (r *registry) NewResponseRepositry(db *gorm.DB) repository.ResponseRepositry {
	return datastore.NewResponseRepositry(db)
}

package repository

import (
	"github.com/a-zara-n/Matchlock/src/domain/aggregate"
)

type HTTPMessageRepository interface {
	Fetch(Identifier string, IsEdit bool) *aggregate.HTTPMessages
}

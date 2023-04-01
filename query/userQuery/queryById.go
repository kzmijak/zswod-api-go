package userQuery

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/user"
)

func (uq UserQuery) QueryById(userId uuid.UUID) UserQuery {
	query := uq.Where(user.ID(userId))
	return FromQuery(query)
}
package userQuery

import "github.com/kzmijak/zswod_api_go/ent/user"

func (uq UserQuery) QueryByEmail(email string) UserQuery {
	query := uq.Where(user.Email(email))
	return FromQuery(query)
}
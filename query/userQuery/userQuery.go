package userQuery

import "github.com/kzmijak/zswod_api_go/ent"

type UserQuery struct {
	*ent.UserQuery
}

func FromQuery(userQuery *ent.UserQuery) UserQuery {
	return UserQuery{
		userQuery,
	}
}

func FromTx(tx *ent.Tx) UserQuery {
	return FromQuery(tx.User.Query())
}
package customPageQuery

import "github.com/kzmijak/zswod_api_go/ent"

type CustomPageQuery struct {
	*ent.CustomPageQuery
}

func FromQuery (query *ent.CustomPageQuery) CustomPageQuery {
	return CustomPageQuery{
		query,
	}
}

func FromTx(tx *ent.Tx) CustomPageQuery {
	return FromQuery(tx.CustomPage.Query())
}
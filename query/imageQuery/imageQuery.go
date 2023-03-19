package imageQuery

import "github.com/kzmijak/zswod_api_go/ent"

type ImageQuery struct {
	*ent.ImageQuery
}

func FromQuery (query *ent.ImageQuery) ImageQuery {
	return ImageQuery{
		query,
	}
}

func FromTx (tx *ent.Tx) ImageQuery {
	return FromQuery(tx.Image.Query())
}
package imageQuery

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/image"
)

func (query ImageQuery) QueryOrdered() ImageQuery {
	query.Order(ent.Asc(image.FieldOrder)).Clone()

	return query
}

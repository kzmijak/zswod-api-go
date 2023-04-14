package customPageQuery

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/custompage"
)

func (query CustomPageQuery) QueryById(id uuid.UUID) CustomPageQuery {
	customPageQuery := query.Where(custompage.ID(id))
	return FromQuery(customPageQuery)
}
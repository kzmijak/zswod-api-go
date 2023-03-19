package articleQuery

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/article"
)

func (query ArticleQuery) QueryById(id uuid.UUID) ArticleQuery {
	articleQuery := query.Where(article.ID(id))
	return FromQuery(articleQuery)
}
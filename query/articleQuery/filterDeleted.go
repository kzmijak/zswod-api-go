package articleQuery

import "github.com/kzmijak/zswod_api_go/ent/article"

func (query ArticleQuery) FilterDeleted() ArticleQuery {
	articleQuery := query.Where(article.StatusNotIn(article.StatusRemoved))
	return FromQuery(articleQuery)
}
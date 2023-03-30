package articleQuery

import "github.com/kzmijak/zswod_api_go/ent/article"

func (query ArticleQuery) QueryPublished() ArticleQuery {
	articleQuery := query.Where(article.StatusIn(article.StatusPublished))
	return FromQuery(articleQuery)
}
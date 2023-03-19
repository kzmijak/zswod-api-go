package articleQuery

import "github.com/kzmijak/zswod_api_go/ent/article"

func (query ArticleQuery) QueryByTitleNormalized(titleNormalized string) ArticleQuery {
	articleQuery := query.Where(article.TitleNormalized(titleNormalized))
	return FromQuery(articleQuery)
}
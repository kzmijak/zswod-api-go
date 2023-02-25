package articleRepo

import (
	"github.com/kzmijak/zswod_api_go/ent/articletitleguid"
)

func (query ArticleTitleGuidQuery) QueryArticleEntityByTitle(titleNormalized string) ArticleQuery {
	articleQuery := query.Where(articletitleguid.TitleNormalized(titleNormalized)).
		QueryArticle()

	return QueryArticle(articleQuery)
}
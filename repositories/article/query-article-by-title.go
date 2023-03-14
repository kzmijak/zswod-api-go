package articleRepo

import "github.com/kzmijak/zswod_api_go/ent/article"

func (query ArticleQuery) QueryArticleEntityByTitle(titleNormalized string) ArticleQuery {
	articleQuery := query.Where(article.TitleNormalized(titleNormalized))

	return QueryArticle(articleQuery)
}
package articleRepo

import "github.com/kzmijak/zswod_api_go/ent"

type ArticleQuery struct {
	*ent.ArticleQuery
}

func QueryArticle (query *ent.ArticleQuery) ArticleQuery {
	return ArticleQuery{
		query,
	}
}

func ArticleTx(tx *ent.Tx) ArticleQuery {
	return QueryArticle(tx.Article.Query())
} 
package articleRepo

import "github.com/kzmijak/zswod_api_go/ent"

type ArticleTitleGuidQuery struct {
	ent.ArticleTitleGuidQuery
}

func QueryArticleTitle(query *ent.ArticleTitleGuidQuery) *ArticleTitleGuidQuery {
	return &ArticleTitleGuidQuery{
		*query,
	}
}

func ArticleTitleTx(tx *ent.Tx) *ArticleTitleGuidQuery {
	return QueryArticleTitle(tx.ArticleTitleGuid.Query())
}
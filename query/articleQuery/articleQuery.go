package articleQuery

import "github.com/kzmijak/zswod_api_go/ent"

type ArticleQuery struct {
	*ent.ArticleQuery
}


func FromQuery (query *ent.ArticleQuery) ArticleQuery {
	aq :=  ArticleQuery{
		query,
	}
	
	return aq;
}

func FromTx(tx *ent.Tx) ArticleQuery {
	return FromQuery(tx.Article.Query())
} 
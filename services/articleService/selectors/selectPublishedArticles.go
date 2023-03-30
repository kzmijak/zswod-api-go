package selectors

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/articleQuery"
)

const (
	ErrQueryArticlesFail = "ErrQueryArticlesFail: Failed to query for articles with specified offset and amount"
)

func (s ArticleSelector) SelectPublishedArticles(tx *ent.Tx, offset int, amount int) ([]*ent.Article, error) {
	articleEntities, err := articleQuery.FromTx(tx).
		JoinAllImagesToArticle().
		QueryPublished().
		Order(ent.Desc(article.FieldCreateTime)).
		Limit(amount).
		Offset(offset).
		All(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrQueryArticlesFail)
	}

	return articleEntities, nil
}
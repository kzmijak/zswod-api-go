package selectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/articleQuery"
)

const (
	ErrArticleIdNotFound = "ErrArticleIdNotFound: Could not find article by given ID"
)

func (s ArticleSelector) SelectArticleById(tx *ent.Tx, articleId uuid.UUID) (*ent.Article, error) {
	articleEntity, err := articleQuery.FromTx(tx).
		JoinAllImagesToArticle().
		QueryById(articleId).
		Only(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrArticleIdNotFound)
	}

	return articleEntity, nil
}
package selectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrArticleIdNotFound = "ErrArticleIdNotFound: Could not find article by given ID"
)

func (s ArticleSelector) SelectArticleById(tx *ent.Tx, articleId uuid.UUID) (*ent.Article, error) {
	articleEntity, err := tx.Article.Get(s.ctx, articleId)
	if err != nil {
		return nil, errors.Error(ErrArticleIdNotFound)
	}

	return articleEntity, nil
}
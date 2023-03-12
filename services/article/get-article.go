package article

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrArticleIdNotFound = "ErrArticleIdNotFound: Could not find article by given ID"
)

func (s ArticleService) GetArticle(articleId uuid.UUID, tx *ent.Tx) (*ent.Article, error) {
	article, err := tx.Article.Get(s.ctx, articleId)
	if err != nil {
		return nil, errors.Error(ErrArticleIdNotFound)
	}

	return article, nil
}
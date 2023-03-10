package article

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotParseArticleId = "ErrCouldNotParseArticleId: Could not parse specified string to UUID"
	ErrArticleIdNotFound = "ErrArticleIdNotFound: Could not find article by given ID"
)

func (s ArticleService) GetArticle(articleId string, tx *ent.Tx) (*ent.Article, error) {
	articleIdParsed, err := uuid.Parse(articleId)
	if err != nil {
		return nil, errors.Error(ErrCouldNotParseArticleId)
	}

	article, err := tx.Article.Get(s.ctx, articleIdParsed)
	if err != nil {
		return nil, errors.Error(ErrArticleIdNotFound)
	}

	return article, nil
}
package articleService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/articleModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotUpdateArticle = "ErrCouldNotUpdateArticle: Failed to update the article"
)

func (s ArticleService) UpdateArticle(articleId uuid.UUID, payload articleModel.UpdateArticlePayload, tx *ent.Tx) (*ent.Article, error) {
	article, err := s.selectors.SelectArticleById(tx, articleId)
	if err != nil {
		return nil, err
	}

	articleUpdated, err := article.
		Update().
		SetTitle(payload.Title).
		SetShort(payload.Short).
		SetContent(payload.Content).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrCouldNotUpdateArticle)
	}

	return articleUpdated, nil
}
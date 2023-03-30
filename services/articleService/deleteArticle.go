package articleService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrArticleDeleteFail = "ErrArticleDeleteFail: Failed to delete article by given ID"
)

func (s ArticleService) DeleteArticle(articleId uuid.UUID, tx *ent.Tx) error {
	err := tx.Article.
		UpdateOneID(articleId).
		SetStatus(article.StatusRemoved).
		Exec(s.ctx)
	
	if err != nil {
		return errors.Error(ErrArticleDeleteFail)
	}
	
	return nil
}
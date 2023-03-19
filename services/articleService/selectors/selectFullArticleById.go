package selectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrArticleWithGalleryAndImagesQueryFailed = "ErrArticleWithGalleryAndImagesQueryFailed: Failed to retrieve article with gallery and images"
)

func (s ArticleSelector) SelectFullArticleById(tx *ent.Tx, articleId uuid.UUID) (*ent.Article, error) {
	articleEntity, err := tx.Article.Query().WithGallery(func(gq *ent.GalleryQuery) {
		gq.WithArticle().All(s.ctx)
	}).Where(article.ID(articleId)).Only(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrArticleWithGalleryAndImagesQueryFailed)
	}

	return articleEntity, nil
}
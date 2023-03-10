package article

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotUpdateArticle = "ErrCouldNotUpdateArticle: Failed to update the article"
)

type UpdateArticlePayload struct {
	Title    string `json:"title"`
	Short    string `json:"short"`
	Content  string `json:"content"`
	GalleryId uuid.UUID `json:"galleryId"`
}

func (s ArticleService) UpdateArticle(articleId string, payload UpdateArticlePayload, tx *ent.Tx) (*ent.Article, error) {
	article, err := s.GetArticle(articleId, tx)
	if err != nil {
		return nil, err
	}

	articleUpdated, err := article.
		Update().
		SetTitle(payload.Title).
		SetShort(payload.Short).
		SetContent(payload.Content).
		SetGalleryID(payload.GalleryId).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrCouldNotUpdateArticle)
	}

	return articleUpdated, nil
}
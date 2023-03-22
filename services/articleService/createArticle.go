package articleService

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/articleModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/utils/sanitizer"
)

const (
	ErrFailedCreatingArticle = "ErrFailedCreatingArticle: Failed to create article"
	ErrFailedCreatingArticleTitle = "ErrFailedCreatingArticleTitle: Failed to create article title"
	ErrCouldNotParseUuid = "ErrCouldNotParseUuid: Failed to parse image guid"
)


func (s ArticleService) CreateArticle(req articleModel.CreateArticlePayload, tx *ent.Tx) (uuid.UUID, error) {
	titleSanitized := sanitizer.SanitizeString(req.Title)
	
	articleId := uuid.New()
	err := tx.Article.Create().
		SetID(articleId).
		SetCreateTime(time.Now()).
		SetTitle(req.Title).
		SetTitleNormalized(titleSanitized).
		SetShort(req.Short).
		SetContent(req.Content).
		SetAuthorID(req.AuthorId).
		SetGalleryID(req.GalleryId).
		Exec(s.ctx)

	if err != nil {
		return uuid.Nil, errors.Error(ErrFailedCreatingArticle) 
	}

	return articleId, nil
}

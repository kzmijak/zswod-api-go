package articleService

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/utils/sanitizer"
)

const (
	ErrFailedCreatingArticle = "ErrFailedCreatingArticle: Failed to create article"
	ErrFailedCreatingArticleTitle = "ErrFailedCreatingArticleTitle: Failed to create article title"
	ErrCouldNotParseUuid = "ErrCouldNotParseUuid: Failed to parse image guid"
)

type CreateArticlePayload struct {
	Title    string `json:"title"`
	Short    string `json:"short"`
	Content  string `json:"content"`
}

const MAX_TITLE_LENGTH = 64	


func (s ArticleService) CreateArticle(req CreateArticlePayload, galleryId uuid.UUID, tx *ent.Tx) (uuid.UUID, error) {

	titleSanitized := sanitizer.SanitizeString(req.Title)
	if len(titleSanitized) > MAX_TITLE_LENGTH {
		titleSanitized = titleSanitized[:MAX_TITLE_LENGTH - 1]
	}

	articleId := uuid.New()
	err := tx.Article.Create().
		SetID(articleId).
		SetTitle(req.Title).
		SetTitleNormalized(titleSanitized).
		SetShort(req.Short).
		SetContent(req.Content).
		SetUploadDate(time.Now()).
		SetGalleryID(galleryId).
		Exec(s.ctx)

	if err != nil {
		return uuid.Nil, errors.Error(ErrFailedCreatingArticle) 
	}

	return articleId, nil
}

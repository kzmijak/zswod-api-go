package article

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/services/image"
)

const (
	ErrFailedCreatingArticle = "ErrFailedCreatingArticle: Failed to create article"
	ErrFailedCreatingArticleTitle = "ErrFailedCreatingArticleTitle: Failed to create article title"
	ErrCouldNotParseUuid = "ErrCouldNotParseUuid: Failed to parse image guid"
)

type CreateArticleRequest struct {
	Article BaseArticlePayload `json:"article"`
	Images []BaseImagePayload `json:"images" binding:"min=1"`
}

func (s ArticleService) CreateArticle(req CreateArticleRequest, tx *ent.Tx) (*ent.Article, error) {
	
	article, err := s.createArticle(req.Article, tx)
	if err != nil {
		return nil, err
	}
	
	if err = s.createArticleTitle(article, tx); err != nil {
		return nil, err
	}

	for _, img := range req.Images {
		if err = s.createImage(img, article.ID, tx); err != nil {
			return nil, err
		}
	}

	return article, nil
}

func (s ArticleService) createArticle(req BaseArticlePayload, tx *ent.Tx) (article *ent.Article, err error) {
	articleId := uuid.New()

	article, err = tx.Article.Create().
		SetID(articleId).
		SetTitle(req.Title).
		SetShort(req.Short).
		SetContent(req.Content).
		SetUploadDate(time.Now()).
		Save(s.ctx)

	if err != nil {
		err = errors.Error(ErrFailedCreatingArticle)
		return
	}

	return
}

func (s ArticleService) createArticleTitle(article *ent.Article, tx *ent.Tx) (err error) {
	const MAX_LENGTH = 64	
	articleTitleId := uuid.New()

	titleSanitized := SanitizeTitle(article.Title)
	if len(titleSanitized) > MAX_LENGTH {
		titleSanitized = titleSanitized[:MAX_LENGTH - 1]
	}

	err = tx.ArticleTitleGuid.
		Create().
		SetID(articleTitleId).
		SetTitleNormalized(titleSanitized).
		SetArticleID(article.ID).
		Exec(s.ctx)

	if err != nil {
		err = errors.Error(ErrFailedCreatingArticleTitle)
		return
	}

	return 
}

func (s ArticleService) createImage(img BaseImagePayload, articleId uuid.UUID, tx *ent.Tx) (err error) {
	_, err = s.imageService.CreateImage(image.CreateImageRequest{
		Title: img.Title,
		Alt: img.Alt,
		BlobId: img.BlobId,
		ArticleId: articleId,
		IsPreview: img.IsPreview,
	}, tx)

	return
}
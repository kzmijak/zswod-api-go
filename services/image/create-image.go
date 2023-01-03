package image

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrFailedCreatingImage = "ErrFailedCreatingImage: Failed to create image"
)

type CreateImageRequest struct {
	Title string `json:"title"`
	Alt string `json:"alt"`
	Url string `json:"url"`
	ArticleId uuid.UUID `json:"articleId"`
}

func (s ImageService) CreateImage(req CreateImageRequest) (*ent.Image, error) {
	image, err := database.Client.Image.Create().
		SetID(uuid.New()).
		SetTitle(req.Title).
		SetUploadDate(time.Now()).
		SetAlt(req.Alt).
		SetURL(req.Url).
		SetArticleID(req.ArticleId).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrFailedCreatingImage)
	}

	return image, nil
}
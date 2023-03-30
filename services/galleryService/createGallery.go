package galleryService

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/galleryModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrFailedToCreateGallery = "ErrFailedToCreateGallery: Could not create gallery with given input"
)

func (s GalleryService) CreateGallery(payload galleryModel.CreateGalleryPayload, tx *ent.Tx) (galleryModel.GalleryModel, error) {
	galleryId := uuid.New()
	now := time.Now()

	galleryEntity, err := tx.Gallery.Create().
		SetID(galleryId).
		SetTitle(payload.Title).
		SetAuthorID(payload.AuthorId).
		SetCreateTime(now).
		Save(s.ctx)

	if err != nil {
		return galleryModel.Nil, errors.Error(ErrFailedToCreateGallery)
	}

	return galleryModel.FromEntity(galleryEntity)
}

package galleryService

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrFailedToCreateGallery = "ErrFailedToCreateGallery: Could not create gallery with given input"
)

func (s GalleryService) CreateGallery(title string, tx *ent.Tx) (*ent.Gallery, error) {
	galleryId := uuid.New()
	now := time.Now()

	gallery, err := tx.Gallery.Create().
		SetID(galleryId).
		SetTitle(title).
		SetCreatedAt(now).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrFailedToCreateGallery)
	}

	return gallery, nil
}

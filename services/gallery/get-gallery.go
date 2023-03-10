package gallery

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotParseGalleryId = "ErrCouldNotParseGalleryId: Could not parse specified string to UUID"
	ErrGalleryIdNotFound = "ErrGalleryIdNotFound: Could not find gallery by given ID"
)

func (s GalleryService) GetGallery(galleryId string, tx *ent.Tx) (*ent.Gallery, error) {
	galleryIdParsed, err := uuid.Parse(galleryId)

	if err != nil {
		return nil, errors.Error(ErrCouldNotParseGalleryId)
	}

	gallery, err := tx.Gallery.Get(s.ctx, galleryIdParsed)

	if err != nil {
		return nil, errors.Error(ErrCouldNotParseGalleryId)
	}

	return gallery, nil
}
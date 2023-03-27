package selectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/galleryQuery"
)

const (
	ErrGalleryIdNotFound = "ErrGalleryIdNotFound: Could not find gallery by given ID"
)

func (s GallerySelector) GetGalleryById(tx *ent.Tx, galleryId uuid.UUID) (*ent.Gallery, error) {
	galleryEntity, err := galleryQuery.FromTx(tx).
		JoinAllImagesToGallery().
		QueryById(galleryId).
		Only(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrGalleryIdNotFound)
	}

	return galleryEntity, nil
}
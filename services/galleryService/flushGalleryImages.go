package galleryService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotFlushGallery = "ErrCouldNotFlushGallery: Failed to flush gallery from images" 
)

func (s GalleryService) FlushGalleryImages(galleryId uuid.UUID, tx *ent.Tx) (error) {
	galleryEntity, err := s.selectors.GetGalleryById(tx, galleryId)
	if err != nil {
		return err
	}

	_, err = galleryEntity.Update().
		ClearImages().
		Save(s.ctx)

	if err != nil {
		return errors.Error(ErrCouldNotFlushGallery)
	}

	return nil
}
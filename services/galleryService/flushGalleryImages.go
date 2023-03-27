package galleryService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/galleryModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotFlushGallery = "ErrCouldNotFlushGallery: Failed to flush gallery from images" 
)

func (s GalleryService) FlushGalleryImages(galleryId uuid.UUID, tx *ent.Tx) (galleryModel.GalleryModel, error) {
	galleryEntity, err := s.selectors.GetGalleryById(tx, galleryId)
	if err != nil {
		return galleryModel.Nil, err
	}


	galleryFlushed, err := galleryEntity.Update().
		ClearImages().
		Save(s.ctx)

	if err != nil {
		return galleryModel.Nil, errors.Error(ErrCouldNotFlushGallery)
	}

	return galleryModel.FromEntity(galleryFlushed)
}
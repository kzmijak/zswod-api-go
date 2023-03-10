package gallery

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotFlushGallery = "ErrCouldNotFlushGallery: Failed to flush gallery from images" 
)

func (s GalleryService) FlushGalleryImages(galleryId string, tx *ent.Tx) (*ent.Gallery, error) {
	gallery, err := s.GetGallery(galleryId, tx)

	if err != nil {
		return nil, err
	}

	galleryFlushed, err := gallery.Update().
		ClearImages().
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrCouldNotFlushGallery)
	}

	return galleryFlushed, nil
}
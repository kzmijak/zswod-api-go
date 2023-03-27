package galleryService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/galleryModel"
)

const (
	ErrCouldNotUpdateGallery = "ErrCouldNotUpdateGallery: Failed to update gallery"
)

type UpdateGalleryPayload struct {
	Title string `json:"title"`
}

func (s GalleryService) UpdateGallery(galleryId uuid.UUID, payload UpdateGalleryPayload, tx *ent.Tx) (galleryModel.GalleryModel, error) {
	gallery, err := s.selectors.GetGalleryById(tx, galleryId)
	if err != nil {
		return galleryModel.Nil, err
	}

	galleryUpdated, err := gallery.Update().
		SetTitle(payload.Title).
		Save(s.ctx)
	if err != nil {
		return galleryModel.Nil, err
	}

	return galleryModel.FromEntity(galleryUpdated)
}

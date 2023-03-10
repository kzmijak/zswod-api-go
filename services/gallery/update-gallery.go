package gallery

import (
	"github.com/kzmijak/zswod_api_go/ent"
)

const (
	ErrCouldNotUpdateGallery = "ErrCouldNotUpdateGallery: Failed to update gallery"
)

type UpdateGalleryPayload struct {
	Title string `json:"title"`
}

func (s GalleryService) UpdateGallery(galleryId string, payload UpdateGalleryPayload, tx *ent.Tx) (*ent.Gallery, error) {
	gallery, err := s.GetGallery(galleryId, tx)

	if err != nil {
		return nil, err
	}

	galleryUpdated, err := gallery.Update().
		SetTitle(payload.Title).
		Save(s.ctx)

	if err != nil {
		return nil, err
	}

	return galleryUpdated, nil
}

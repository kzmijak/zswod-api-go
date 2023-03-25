package imageService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrImageUpdateFail = "ErrImageUpdateFail: Could not update image with given ID"
)

type UpdateImagePayload struct {
	Alt       string    `json:"alt"`
	GalleryId uuid.UUID `json:"galleryId"`
}

func (s ImageService) UpdateImage(imageId string, payload UpdateImagePayload, tx *ent.Tx) (*ent.Image, error) {
	image, err := s.GetImage(imageId, tx)
	if err != nil {
		return nil, err
	}

	img, err := image.
		Update().
		SetAlt(payload.Alt).
		SetGalleryID(payload.GalleryId).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrImageUpdateFail)
	}

	return img, nil
}
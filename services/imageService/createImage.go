package imageService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/imageModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrFailedCreatingImage = "ErrFailedCreatingImage: Failed to create image"
)

func (s ImageService) CreateImage(req imageModel.CreateImagePayload, galleryId uuid.UUID, tx *ent.Tx) (*ent.Image, error) {
	image, err := tx.Image.Create().
		SetID(uuid.New()).
		SetAlt(req.Alt).
		SetGalleryID(galleryId).
		SetOrder(req.Order).
		SetSrc(req.Src).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrFailedCreatingImage)
	}

	return image, nil
}